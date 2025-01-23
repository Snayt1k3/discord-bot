package main

import (
	"log"
	"net"
	pb "settings-service/proto"
	"google.golang.org/grpc"
	"settings-service/internal/handlers"
)

func main() {

	// Создаём слушатель на порту 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Создаём gRPC-сервер
	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис
	pb.RegisterSettingsServiceServer(grpcServer, &handlers.SettingsServer{})

	log.Println("gRPC server is running on port :50051")

	// Запускаем сервер
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}


