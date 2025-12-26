package main

import (
	"fmt"
	"log/slog"
	"os"

	"api-gateway/config"
	"api-gateway/internal/adapters"
	"api-gateway/internal/handlers"
	"api-gateway/internal/routes"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	initLogging()
	cfg, err := config.LoadConfig()

	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	settingsServiceConn, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.GrpcSettingsHost, cfg.GrpcSettingsPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Failed to connect to settings service: " + err.Error())
	}

	interactionServiceConn, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.GrpcInteractionHost, cfg.GrpcInteractionPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Failed to connect to interaction service: " + err.Error())
	}

	defer settingsServiceConn.Close()
	defer interactionServiceConn.Close()

	redisClient := adapters.NewRedisAdapter(fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort), cfg.RedisPass, cfg.RedisDB)

	// Инициализация handlers
	settings := handlers.NewSettingsHandlers(settingsServiceConn, redisClient)
	roles := handlers.NewRolesHandlers(settingsServiceConn)
	welcome := handlers.NewWelcomeHandlers(settingsServiceConn)
	log := handlers.NewLogHandlers(settingsServiceConn)
	automode := handlers.NewAutoModeHandlers(settingsServiceConn)
	interaction := handlers.NewInteraction(interactionServiceConn)

	r := routes.SetupRouter(settings, roles, welcome, automode, log, interaction)

	port := ":" + cfg.Port

	slog.Info("Starting server", "port", port)
	slog.Info("Docs available at", "url", "http://localhost"+port+"/swagger/index.html")

	if err := r.Run(port); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}

func initLogging() {
	//file, err := os.OpenFile("api-gateway.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	panic("Failed to open log file: " + err.Error())
	//}
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	//logger := slog.New(slog.NewTextHandler(file, opts))
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}

