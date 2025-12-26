package main

import (

	"fmt"
	"interaction-service/config"
	"log"
	"log/slog"
	"net"

	"os"
	"interaction-service/internal/models"
	"interaction-service/internal/server"
	pb "interaction-service/proto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

func main() {
	initLogging()
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	models.Migrate(db)
	grpcServer := grpc.NewServer()

	pb.RegisterInteractionServiceServer(grpcServer, server.NewUserServer(db))

	// run grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))
	
	if err != nil {
		log.Fatalf("Error listening on grpc port: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initLogging() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	// Вывод логов в консоль
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}

