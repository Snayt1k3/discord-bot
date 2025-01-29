package main

import (
	"fmt"
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

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	guildRepo := adapters.NewGuildSettingRepository(db)
	botRepo := adapters.NewBotSettingRepository(db)

	settingsService := &adapters.SettingsService{
		GuildRepo: guildRepo,
		BotRepo:   botRepo,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSettingsServiceServer(grpcServer, &server.SettingsServer{
		SettingsService: settingsService,
	})

	log.Printf("gRPC server is running on port :%v \n", cfg.GrpcPort)

	// Запускаем сервер
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}