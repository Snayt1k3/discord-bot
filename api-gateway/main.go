package main

import (
	"fmt"
	"log/slog"
	"os"

	"api-gateway/config"
	"api-gateway/internal/adapters"
	"api-gateway/internal/handlers"
	"api-gateway/internal/metrics"
	"api-gateway/internal/routes"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	initLogging()
	InitProm()
	cfg, err := config.LoadConfig()

	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	settingsServiceConn, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.GrpcSettingsHost, cfg.GrpcSettingsPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Failed to connect to settings service: " + err.Error())
	}

	defer settingsServiceConn.Close()

	redisClient := adapters.NewRedisAdapter(fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort), cfg.RedisPass, cfg.RedisDB)
	settingsHandlers := handlers.NewHandlers(settingsServiceConn, redisClient)

	r := routes.SetupRouter(settingsHandlers)

	port := ":" + cfg.Port

	slog.Info("Starting server", "port", port)
	slog.Info("Docs available at", "url", "http://localhost"+port+"/swagger/index.html")

	if err := r.Run(port); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}

func initLogging() {
	file, err := os.OpenFile("api-gateway.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(file, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}

func InitProm() {
	prometheus.MustRegister(metrics.HttpRequestsTotal)
	prometheus.MustRegister(metrics.HttpRequestDuration)
}
