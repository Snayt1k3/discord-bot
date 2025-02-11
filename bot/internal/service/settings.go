package service

import (
	"bot/internal/interfaces"
	"bot/internal/adapters"
)


type ServiceSettingsClient struct {
	client interfaces.HttpClient
}

func NewServiceSettingsClient() *ServiceSettingsClient {
	return &ServiceSettingsClient{client: adapters.NewDefaultHttpClient()}
}