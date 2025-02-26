package main

import (
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"settings-service/config"
	"settings-service/internal/adapters"
	"settings-service/internal/models"
	"settings-service/internal/server"
	pb "settings-service/proto"
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
	models.AutoMigrate(db)
	guildRepo := adapters.NewGuildSettingRepository(db)

	settingsService := &adapters.SettingsService{
		GuildRepo: guildRepo,
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
