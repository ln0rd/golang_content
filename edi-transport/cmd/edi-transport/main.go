package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/InVisionApp/go-health/v2"
	hc "github.com/????lab/pkg/healthcheck"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"net/http"

	istioproxy "github.com/allisson/go-istio-proxy-wait"
	"github.com/????lab/edi-transport/client"
	"github.com/????lab/edi-transport/server"
	"github.com/????lab/pkg/env"
	"github.com/????lab/pkg/log"
)

var (
	logger     *zap.Logger
	istioProxy istioproxy.Proxy
)

func mustGetenv(k string) string {
	// nolint:errcheck
	defer logger.Sync()
	v := env.GetString(k, "")
	if v == "" {
		logger.Fatal("environment-variable-not-set.", zap.String("variable", k))
	}
	return v
}

func init() {
	// nolint:errcheck
	logger, _ = log.NewLogger()

	// Wait until the istio-proxy is ready
	istioProxy = istioproxy.New(
		time.Duration(env.GetInt("ISTIO_PROXY_TIMEOUT", 5))*time.Second,
		time.Duration(env.GetInt("ISTIO_PROXY_RETRY_DELAY", 5))*time.Second,
		env.GetInt("ISTIO_PROXY_MAX_RETRIES", 10),
	)
	if err := istioProxy.Wait(); err != nil {
		closeIstioProxy()
		logger.Fatal("istio-proxy-wait", zap.Error(err))
	}
}

func healthCheckServer() {
	port := env.GetInt("HEALTH_CHECK_PORT", 8001)

	// Create health configs
	healthConfigs := []*health.Config{}

	// Create mux
	mux, err := hc.NewMux(healthConfigs, healthConfigs)
	if err != nil {
		zap.L().Error("healthcheck-serve-mux", zap.Error(err))
		return
	}

	// Start server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		zap.L().Error("healthcheck-server-failed-to-start", zap.Error(err))
	}
}

func closeIstioProxy() {
	if err := istioProxy.Close(); err != nil {
		logger.Error("istio-proxy-close", zap.Error(err))
	}
}

func main() {
	// nolint:errcheck
	defer logger.Sync()
	defer closeIstioProxy()

	app := cli.NewApp()
	app.Name = "SFTP Service"
	app.Usage = "CLI"
	app.Authors = []*cli.Author{
		{
			Name:  "???? Developers",
			Email: "devs@????.com.br",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Starts SFTP Server",
			Action: func(c *cli.Context) error {
				var serverConfig = server.Config{
					Username:           mustGetenv("SFTP_USERNAME"),
					Password:           mustGetenv("SFTP_PASSWORD"),
					Port:               mustGetenv("SFTP_PORT"),
					KeyFile:            mustGetenv("SFTP_SERVER_KEY_FILE"),
					AuthorizedKeysFile: env.GetString("SFTP_AUTHORIZED_KEYS_FILE", ""),
					GCSBucket:          mustGetenv("GCS_BUCKET"),
					GCSCredentialsFile: env.GetString("GCS_CREDENTIALS_FILE", ""),
				}

				// Start health check
				go healthCheckServer()

				sftpServer, err := server.NewServer(serverConfig)
				if err != nil {
					logger.Fatal("failed-create-sftp-sftpServer", zap.Error(err))
				}

				idleConnsClosed := make(chan struct{})
				go func() {
					sigint := make(chan os.Signal, 1)

					// interrupt signal sent from terminal
					signal.Notify(sigint, os.Interrupt)
					// sigterm signal sent from kubernetes
					signal.Notify(sigint, syscall.SIGTERM)

					<-sigint

					// We received an interrupt signal, shut down.
					logger.Info("sftp-sftpServer-shutdown-started")
					if err = sftpServer.Close(); err != nil {
						logger.Error("sftp-server-close-error", zap.Error(err))
					}
					close(idleConnsClosed)
					logger.Info("sftp-sftpServer-shutdown-finished")
				}()

				if err = sftpServer.Run(); err != nil {
					logger.Fatal("failed-run-sftp-sftpServer", zap.Error(err))
				}

				<-idleConnsClosed
				return nil
			},
		},
		{
			Name:    "client",
			Aliases: []string{"c"},
			Usage:   "Starts SFTP Client",
			Action: func(c *cli.Context) error {
				var clientConfig = client.Config{
					Addresses:          strings.Fields(mustGetenv("TIVIT_ADDRESSES")),
					Port:               env.GetString("TIVIT_PORT", "22"),
					Username:           mustGetenv("TIVIT_USERNAME"),
					Password:           mustGetenv("TIVIT_PASSWORD"),
					KnownHostsFile:     env.GetString("KNOWN_HOSTS_FILE", "/tmp/known_hosts"),
					BasePath:           env.GetString("TIVIT_BASE_PATH", ""),
					KeyFile:            mustGetenv("SFTP_SERVER_KEY_FILE"),
					GCSBucket:          mustGetenv("GCS_BUCKET"),
					GCSCredentialsFile: env.GetString("GCS_CREDENTIALS_FILE", ""),
					Timeout:            env.GetInt("CONNECTION_TIMEOUT", 30),
				}

				sftpClient, err := client.NewClient(clientConfig)
				if err != nil {
					logger.Fatal("failed-create-client", zap.Error(err))
				}

				if err = sftpClient.Run(); err != nil {
					logger.Fatal("failed-running-sftp-client", zap.Error(err))
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal("app", zap.Error(err))
	}
}
