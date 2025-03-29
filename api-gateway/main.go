package main

import (
	"api-gateway/config"
	"api-gateway/internal/adapters"
	"api-gateway/internal/handlers"
	"api-gateway/internal/routes"
	pb "api-gateway/proto"
	"fmt"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.LoadConfig()
	initLogging()

	if err != nil {
		slog.Error("Unable to load config", "error", err)
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.GrpcHost, cfg.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Unable to connect to grpc server", "error", err)
	}

	defer conn.Close()

	protoClient := pb.NewSettingsServiceClient(conn)
	redisClient := adapters.NewRedisAdapter(fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort), cfg.RedisPass, cfg.RedisDB)
	settingsClient := handlers.NewClient(protoClient, redisClient)
	
	r := routes.SetupRouter(settingsClient)

	port := ":8080"
	slog.Info("Starting server", "port", port)

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
