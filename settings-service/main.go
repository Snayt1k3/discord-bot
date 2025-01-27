package main

import (
	"log"
	"net"
	"settings-service/config"
	"settings-service/internal/adapters"
	"settings-service/internal/server"
	pb "settings-service/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Создание репозиториев
	guildRepo := adapters.NewGuildSettingRepository(db)
	botRepo := adapters.NewBotSettingRepository(db)

	// Создание сервиса с указателями на репозитории
	settingsService := &adapters.SettingsService{
		GuildRepo: guildRepo,
		BotRepo:   botRepo,
	}

	// Создаём слушатель на порту 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Создаём gRPC-сервер
	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис
	pb.RegisterSettingsServiceServer(grpcServer, &server.SettingsServer{
		SettingsService: settingsService,
	})

	log.Println("gRPC server is running on port :50051")

	// Запускаем сервер
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}