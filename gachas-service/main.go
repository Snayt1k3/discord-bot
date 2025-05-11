package main

import (
	"fmt"
	"gachas-service/config"
	"gachas-service/internal/adapters/storage/models/genshin"
	"gachas-service/internal/adapters/storage/models/wuwa"
	"gachas-service/internal/adapters/storage/repos"
	"gachas-service/internal/server"
	"gachas-service/internal/services"
	pb "gachas-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}
	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	// Perform database migrations
	genshin.Migrate(db)
	wuwa.Migrate(db)

	// Initialize the repositories
	genshinRepo := repos.NewGenshinRepository(db)
	wuwaRepo := repos.NewWuwaRepository(db)

	// Initialize the Services
	genshinService := services.NewGenshinService(genshinRepo)
	wuwaService := services.NewWuwaService(wuwaRepo)

	// Initialize the gRPC server
	genshinServer := server.NewGenshinServer(genshinService)
	wuwaServer := server.NewWuwaServer(wuwaService)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGenshinServiceServer(grpcServer, genshinServer)
	pb.RegisterWuwaServiceServer(grpcServer, wuwaServer)

	log.Printf("gRPC server is running on port :%v \n", cfg.GrpcPort)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
