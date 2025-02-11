package main

import (
	"api-gateway/config"
	"api-gateway/internal/routes"
	pb "api-gateway/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}
	
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.GrpcHost, cfg.GrpcPort))
	
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC-серверу: %v", err)
	}
	
	defer conn.Close()

	settingsClient := pb.NewSettingsServiceClient(conn)

	r := routes.SetupRouter(settingsClient)

	port := ":8080"
	log.Printf("Сервер запущен на %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
