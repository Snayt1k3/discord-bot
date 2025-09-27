package main

import (
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net"
	"os"
	"settings-service/config"
	"settings-service/internal/adapters/repos"
	"settings-service/internal/models"
	"settings-service/internal/server"
	pb "settings-service/proto"
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
	repositories := repos.NewRepos(db)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAutoModServiceServer(grpcServer, &server.AutomodeServer{
		Repo: repositories.AutoMode, GuildRepo: repositories.Settings,
	})
	pb.RegisterLogServiceServer(grpcServer, &server.LogServer{
		Repo: repositories.Log,
	})
	pb.RegisterRolesServiceServer(grpcServer, &server.RolesReactionServer{
		Repo: repositories.ReactionRoles, GuildRepo: repositories.Settings,
	})
	pb.RegisterSettingsServiceServer(grpcServer, &server.GuildServer{
		Repo: repositories.Settings,
	})
	pb.RegisterWelcomeServiceServer(grpcServer, &server.WelcomeServer{
		Repo: repositories.Welcome, GuildRepo: repositories.Settings,
	})

	log.Printf("gRPC server is running on port :%v \n", cfg.GrpcPort)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
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
