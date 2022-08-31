package server

import (
	"context"
	"fmt"
	gsftp "github.com/????lab/edi-transport/handler"
	"github.com/????lab/pkg/log"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"net"
	"sync"
)

var logger *zap.Logger

// Config has configuration needed to run sftp server
type Config struct {
	Port               string
	Username           string
	Password           string
	KeyFile            string
	AuthorizedKeysFile string
	GCSCredentialsFile string
	GCSBucket          string
}

// Server is an implementation of tivitSftp.SftpServer
type Server struct {
	config    Config
	sshConfig *ssh.ServerConfig
	listener  net.Listener
	wg        sync.WaitGroup
}

func init() {
	logger, _ = log.NewLogger()
}

func processPublicKeyAuth(authorizedKeysFile string, config *ssh.ServerConfig) {
	// nolint:errcheck
	defer logger.Sync()

	if authorizedKeysFile == "" {
		return
	}

	authorizedKeysBytes, err := ioutil.ReadFile(authorizedKeysFile)
	if err != nil {
		logger.Fatal("failed-load-authorized-keys-file", zap.Error(err))
	}

	var authorizedKeysArray []ssh.PublicKey
	for {
		out, _, _, rest, err := ssh.ParseAuthorizedKey(authorizedKeysBytes)
		if err != nil {
			if err.Error() == "ssh: no key found" {
				break
			}
			logger.Fatal("failed-parse-authorized-key", zap.Error(err))
		}
		authorizedKeysArray = append(authorizedKeysArray, out)
		authorizedKeysBytes = rest
	}

	config.PublicKeyCallback = func(conn ssh.ConnMetadata, auth ssh.PublicKey) (*ssh.Permissions, error) {
		for _, pubKey := range authorizedKeysArray {
			if string(pubKey.Marshal()) == string(auth.Marshal()) {
				return &ssh.Permissions{
					// Record the public key used for authentication.
					Extensions: map[string]string{
						"pubkey-fp": ssh.FingerprintSHA256(auth),
					},
				}, nil
			}
		}

		return nil, fmt.Errorf("unknown public key for %q", conn.User())
	}
}

func (s *Server) passwordCallback(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
	if c.User() == s.config.Username && string(pass) == s.config.Password {
		return nil, nil
	}
	return nil, fmt.Errorf("password rejected for %q", c.User())
}

func (s *Server) handleConn(nConn net.Conn) {
	// nolint:errcheck
	defer logger.Sync()

	defer s.wg.Done()
	// Before use, a handshake must be performed on the incoming net.Conn.
	sconn, chans, reqs, err := ssh.NewServerConn(nConn, s.sshConfig)
	if err != nil {
		logger.Error("handshake-failed", zap.Error(err))
		return
	}
	logger.Debug("login-detected", zap.String("user", sconn.User()))

	// The incoming Request channel must be serviced.
	go ssh.DiscardRequests(reqs)

	// Service the incoming Channel channel.
	for newChannel := range chans {
		// Channels have a type, depending on the application level
		// protocol intended. In the case of an SFTP session, this is "subsystem"
		// with a payload string of "<length=4>sftp"
		if newChannel.ChannelType() != "session" {
			if err = newChannel.Reject(ssh.UnknownChannelType, "unknown channel type"); err != nil {
				logger.Error("channel-reject-error", zap.Error(err))
			}
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			logger.Error("could-not-accept-channel", zap.Error(err))
			return
		}

		// Sessions have out-of-band requests such as "shell",
		// "pty-req" and "env".  Here we handle only the
		// "subsystem" request.
		go func(in <-chan *ssh.Request) {
			for req := range in {
				ok := false
				switch req.Type {
				case "subsystem":
					if string(req.Payload[4:]) == "sftp" {
						ok = true
					}
				}
				if err = req.Reply(ok, nil); err != nil {
					logger.Error("req-reply-error", zap.Error(err), zap.Bool("ok", ok))
				}
			}
		}(requests)

		ctx := context.Background()

		var opts []option.ClientOption
		if s.config.GCSCredentialsFile != "" {
			opts = append(opts, option.WithCredentialsFile(s.config.GCSCredentialsFile))
		}

		root, err := gsftp.GoogleCloudStorageHandler(ctx, s.config.GCSBucket, opts...)
		if err != nil {
			logger.Fatal("gcs-init-failed", zap.Error(err))
		}

		server := sftp.NewRequestServer(channel, *root)
		if err := server.Serve(); err == io.EOF {
			server.Close()

			logger.Debug("sftp-client-exited-session.")
		} else if err != nil {
			logger.Error("sftp-server-completed-with-error", zap.Error(err))
		}
	}
}

// Run starts the sftp server
func (s *Server) Run() error {
	// nolint:errcheck
	defer logger.Sync()

	// Once a Server has been configured, connections can be accepted.
	listener, err := net.Listen("tcp", "0.0.0.0:"+s.config.Port)
	if err != nil {
		logger.Error("failed-to-listen-for-connection", zap.Error(err))
		return err
	}
	s.listener = listener
	logger.Info("listening", zap.String("address", s.listener.Addr().String()))

	for {
		nConn, err := s.listener.Accept()
		if err != nil {
			if x, ok := err.(*net.OpError); ok && x.Op == "accept" {
				logger.Warn("listener-closed")
				break
			}

			logger.Error("failed-accept-incoming-connection", zap.Error(err))
			continue
		}

		s.wg.Add(1)
		go s.handleConn(nConn)
	}

	return nil
}

// Close closes the server listener
func (s *Server) Close() error {
	if err := s.listener.Close(); err != nil {
		logger.Error("litener-close-error", zap.Error(err))
	}

	s.wg.Wait()
	return nil
}

// NewServer creates a new sftp server
func NewServer(config Config) (*Server, error) {
	// nolint:errcheck
	defer logger.Sync()

	server := Server{}

	server.config = config

	// An SSH server is represented by a ServerConfig, which holds
	// certificate details and handles authentication of ServerConns.
	server.sshConfig = &ssh.ServerConfig{
		NoClientAuth:  false,
		ServerVersion: "SSH-2.0-GCS-SFTP",
		AuthLogCallback: func(conn ssh.ConnMetadata, method string, err error) {
			if err != nil {
				logger.Error("ssh2-method-failed", zap.String("method", method),
					zap.String("user", conn.User()),
					zap.String("remote-address", conn.RemoteAddr().String()))
			} else {
				logger.Info("ssh2-accepted-connection", zap.String("method", method),
					zap.String("user", conn.User()),
					zap.String("remote-address", conn.RemoteAddr().String()))
			}
		},
		PasswordCallback: server.passwordCallback,
	}

	privateBytes, err := ioutil.ReadFile(config.KeyFile)
	if err != nil {
		logger.Error("failed-load-private-key", zap.Error(err))
		return nil, err
	}

	hostKey, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		logger.Error("failed-parse-server-key", zap.Error(err))
		return nil, err
	}
	server.sshConfig.AddHostKey(hostKey)

	processPublicKeyAuth(config.AuthorizedKeysFile, server.sshConfig)
	return &server, nil
}
