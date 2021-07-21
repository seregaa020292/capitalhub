package main

import (
	"log"
	"os"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/server"
	"github.com/seregaa020292/capitalhub/pkg/db/aws"
	"github.com/seregaa020292/capitalhub/pkg/db/postgres"
	"github.com/seregaa020292/capitalhub/pkg/db/redis"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/tracer"
	"github.com/seregaa020292/capitalhub/pkg/webSocket"
)

// @title REST API
// @version 1.0
// @description REST API

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey Auth
// @in header
// @name Authorization
func main() {
	log.Println("Starting api server")

	// Setup Config
	cfg, err := config.NewConfig(os.Getenv("CONFIG"))
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	// Setup Logger
	appLogger := logger.NewApiLogger(cfg)
	closeLogger := appLogger.InitLogger()
	defer closeLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL,
	)

	// Setup Psql
	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	// Setup Redis
	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	// Setup Aws
	awsClient, err := aws.NewAWSClient(cfg.AWS)
	if err != nil {
		appLogger.Errorf("AWS Client init: %s", err)
	}
	appLogger.Info("AWS S3 connected")

	// Setup WebSocket
	webSocketClient := webSocket.NewHub()
	go webSocketClient.Run()

	// Setup Tracer
	closerJaeger := tracer.NewJaeger(cfg.Jaeger)
	defer closerJaeger.Close()
	appLogger.Info("Jaeger connected")
	appLogger.Info("Opentracing connected")

	// Run Server
	newServer := server.NewServer(
		cfg,
		psqlDB,
		webSocketClient,
		redisClient,
		awsClient,
		appLogger,
	)
	if err = newServer.Run(); err != nil {
		log.Fatal(err)
	}
}
