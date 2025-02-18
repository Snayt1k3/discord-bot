package main

import (
	"api-gateway/config"
	"api-gateway/internal/adapters"
	"api-gateway/internal/handlers"
	"api-gateway/internal/routes"
	pb "api-gateway/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.LoadConfig()
	
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}
	
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.GrpcHost, cfg.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC-серверу: %v", err)
	}
	
	defer conn.Close()

	protoClient := pb.NewSettingsServiceClient(conn)
	redisClient := adapters.NewRedisAdapter(fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort), cfg.RedisPass, cfg.RedisDB)
	settingsClient := handlers.NewClient(protoClient, redisClient)
	r := routes.SetupRouter(settingsClient)

	port := ":8080"
	log.Printf("Сервер запущен на %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
