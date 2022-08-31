package client

import (
	"context"
	"fmt"
	"io"
	"net"
	"path"
	"time"

	"github.com/????lab/pkg/log"

	"cloud.google.com/go/storage"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"google.golang.org/api/option"
)

var logger *zap.Logger

// Config is a client configuration structure
type Config struct {
	Addresses          []string
	Port               string
	Username           string
	Password           string
	KnownHostsFile     string
	BasePath           string
	KeyFile            string
	GCSCredentialsFile string
	GCSBucket          string
	Timeout            int
}

// Client implements a tivitSftp.SftpClient Interface
type Client struct {
	config        Config
	storageBucket *storage.BucketHandle
	sshConfig     *ssh.ClientConfig
}

func init() {
	logger, _ = log.NewLogger()
}

func closeRemoteFile(file *sftp.File) {
	if err := file.Close(); err != nil {
		logger.Error("remote-file-close-error", zap.Error(err))
	}
}

func closeWriter(writer *storage.Writer) {
	if err := writer.Close(); err != nil {
		logger.Error("bucket-writer-close-error", zap.Error(err))
	}
}

func (c *Client) writeDestFile(dest string, content []byte) (n int, err error) {
	logger.Debug("writing-file", zap.String("name", dest))

	timeout := time.Duration(c.config.Timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	writer := c.storageBucket.Object(dest).NewWriter(ctx)
	defer closeWriter(writer)

	return writer.Write(content)
}

func (c *Client) readRemoteFile(client *sftp.Client, fileName string) ([]byte, error) {
	remote, err := client.Open(fileName)
	if err != nil {
		logger.Error("failed-open-file", zap.String("name", fileName), zap.Error(err))
		return nil, err
	}
	defer closeRemoteFile(remote)

	fullContent := make([]byte, 0)
	for {
		raw := make([]byte, 4096)
		l, err := remote.Read(raw)
		if err != nil && err == io.EOF {
			if l > 0 {
				fullContent = append(fullContent, raw[:l]...)
			}
			break
		} else if err != nil {
			logger.Error("read-remote-file-fail", zap.String("name", fileName),
				zap.Error(err))
			return nil, err
		}
		fullContent = append(fullContent, raw[:l]...)
	}

	return fullContent, nil
}

func (c *Client) copyRemoteFile(client *sftp.Client, originName string) error {
	logger.Debug("copy-file", zap.String("name", originName))

	content, err := c.readRemoteFile(client, originName)
	if err != nil {
		logger.Error("failed-read-remote-file", zap.String("name", originName),
			zap.Error(err))
		return err
	}

	destName := path.Base(originName)
	l, err := c.writeDestFile(destName, content)
	if err != nil {
		logger.Error("failed-write-file", zap.String("name", destName), zap.Error(err))
		return err
	}

	logger.Debug("wrote-file", zap.String("filename", originName),
		zap.String("destination", destName), zap.Int("bytes", l))

	return nil
}

func (c *Client) downloadTree(client *sftp.Client, dirPath string) error {
	// nolint:errcheck
	defer logger.Sync()
	// List remote files
	logger.Debug("read-dir", zap.String("path", dirPath))

	fileInfos, err := client.ReadDir(dirPath)
	if err != nil {
		logger.Error("failed-read-dir", zap.String("path", dirPath), zap.Error(err))
		return err
	}

	for _, fileInfo := range fileInfos {
		fileFullPath := fmt.Sprintf("%s/%s", dirPath, fileInfo.Name())
		if fileInfo.IsDir() {
			err = c.downloadTree(client, path.Clean(fileFullPath))
			if err != nil {
				logger.Error("failed-read-dir", zap.String("path", fileFullPath), zap.Error(err))
			}
			continue
		}

		if err = c.copyRemoteFile(client, path.Clean(fileFullPath)); err != nil {
			logger.Error("copy-remote-file-error", zap.Error(err))
		}
	}

	return nil
}

// Run try to connect on one of addresses configured and download files available on remote to storage parameterized
func (c *Client) Run() error {
	// nolint:errcheck
	defer logger.Sync()
	var (
		sftpClient *sftp.Client
		sshClient  *ssh.Client
		err        error
	)

	for _, host := range c.config.Addresses {
		strHost := fmt.Sprintf("%s:%s", host, c.config.Port)

		logger.Debug("connecting-to", zap.String("host", strHost))

		sshClient, err = ssh.Dial("tcp", strHost, c.sshConfig)
		if err != nil {
			logger.Error("failed-connect-host", zap.String("host", strHost), zap.Error(err))
			continue
		}

		sftpClient, err = sftp.NewClient(sshClient)
		if err != nil {
			logger.Error("failed-create-sftp-client", zap.String("host", strHost),
				zap.String("port", c.config.Port), zap.Error(err))
			if err = sshClient.Close(); err != nil {
				logger.Error("ssh-client-close-error", zap.Error(err))
			}

			continue
		}

		break
	}

	if err != nil {
		logger.Error("could-not-connect", zap.Error(err))
		return err
	}

	defer sshClient.Close()
	defer sftpClient.Close()

	root, err := sftpClient.Getwd()
	if err != nil {
		logger.Error("failed-get-current-working-directory", zap.Error(err))
		return err
	}

	logger.Debug("root-path", zap.String("cwd", root))

	root = fmt.Sprintf("%s%s", root, c.config.BasePath)

	return c.downloadTree(sftpClient, path.Clean(root))
}

/*
While issue below is not fixed, by pass host validation.

{"severity":"ERROR","time":"2020-08-24T13:58:28.244Z","caller":"client/client.go:187",
"message":"known-hosts-new-error",
"error":"knownhosts: /tmp/known_hosts:1: illegal base64 data at input byte 188", ...
*/
func hostKeyCallback(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return nil
}

// NewClient creates a new sftp client
func NewClient(config Config) (*Client, error) {
	// nolint:errcheck
	defer logger.Sync()
	client := Client{}
	client.config = config

	var opts []option.ClientOption
	if config.GCSCredentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(config.GCSCredentialsFile))
	}

	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("storage client error: %s", err)
	}

	client.storageBucket = storageClient.Bucket(config.GCSBucket)

	client.sshConfig = &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: hostKeyCallback,
		Timeout:         time.Duration(config.Timeout) * time.Second,
	}
	return &client, nil
}
