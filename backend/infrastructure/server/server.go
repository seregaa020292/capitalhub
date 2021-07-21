package server

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"

	"github.com/seregaa020292/capitalhub/config"
	_ "github.com/seregaa020292/capitalhub/docs"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/webSocket"
)

const (
	certFile       = "docker/common/ssl/server.crt"
	keyFile        = "docker/common/ssl/server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo            *echo.Echo
	webSocketClient *webSocket.Hub
	cfg             *config.Config
	db              *sqlx.DB
	redisClient     *redis.Client
	awsClient       *minio.Client
	logger          logger.Logger
}

// NewServer New Server constructor
func NewServer(
	cfg *config.Config,
	db *sqlx.DB,
	webSocketClient *webSocket.Hub,
	redisClient *redis.Client,
	awsS3Client *minio.Client,
	logger logger.Logger,
) *Server {
	return &Server{
		echo:            echo.New(),
		cfg:             cfg,
		db:              db,
		webSocketClient: webSocketClient,
		redisClient:     redisClient,
		awsClient:       awsS3Client,
		logger:          logger,
	}
}

func (server *Server) Run() error {
	if server.cfg.Server.SSL {
		return server.https()
	}
	return server.http()
}

func (server Server) http() error {
	httpServer := &http.Server{
		Addr:           server.cfg.Server.Port,
		ReadTimeout:    time.Second * server.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * server.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		server.logger.Infof("Server is listening on PORT: %s", server.cfg.Server.Port)
		if err := server.echo.StartServer(httpServer); err != nil {
			server.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	go func() {
		server.logger.Infof("Starting Debug Server on PORT: %s", server.cfg.Server.PprofPort)
		if err := http.ListenAndServe(server.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			server.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	if err := server.MapHandlers(server.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	server.logger.Info("Server Exited Properly")
	return server.echo.Server.Shutdown(ctx)
}

func (server *Server) https() error {
	if err := server.MapHandlers(server.echo); err != nil {
		return err
	}

	server.echo.Server.ReadTimeout = time.Second * server.cfg.Server.ReadTimeout
	server.echo.Server.WriteTimeout = time.Second * server.cfg.Server.WriteTimeout

	go func() {
		server.logger.Infof("Server is listening on PORT: %s", server.cfg.Server.Port)
		server.echo.Server.ReadTimeout = time.Second * server.cfg.Server.ReadTimeout
		server.echo.Server.WriteTimeout = time.Second * server.cfg.Server.WriteTimeout
		server.echo.Server.MaxHeaderBytes = maxHeaderBytes
		if err := server.echo.StartTLS(server.cfg.Server.Port, certFile, keyFile); err != nil {
			server.logger.Fatalf("Error starting TLS Server: ", err)
		}
	}()

	go func() {
		server.logger.Infof("Starting Debug Server on PORT: %s", server.cfg.Server.PprofPort)
		if err := http.ListenAndServe(server.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			server.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	server.logger.Info("Server Exited Properly")
	return server.echo.Server.Shutdown(ctx)
}
