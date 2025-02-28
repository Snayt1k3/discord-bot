package adapters

import (
	"bot/config"
	"bot/internal/dto"
	"bot/internal/interfaces"
	"context"
	"log/slog"
	"encoding/json"
	"fmt"
	"io"
)

type ServiceSettingsClient struct {
	client interfaces.HttpClient
}

func (s *ServiceSettingsClient) CreateSettings(guild_id string) error {
	resp, err := s.client.Post(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", config.GetApiGatewayAddr(), guild_id),
		nil,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when creating settings", "err", err)
		return err
	}

	defer resp.Body.Close()

	return err

}

func (s *ServiceSettingsClient) UpdateGuildSettings(guildId string, roles dto.RolesSettings) (dto.GuildSettingsDTO, error) {
	body, _ := json.Marshal(roles)

	resp, err := s.client.Patch(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", config.GetApiGatewayAddr(), guildId),
		body,
		nil,
	)

	if err != nil {
		
		return dto.GuildSettingsDTO{}, err
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	var settings dto.GuildSettingsDTO

	err = json.Unmarshal(body, &settings)

	if err != nil {
		return dto.GuildSettingsDTO{}, nil
	}
	slog.Warn("Bad response when updating settings", "err", err)
	return settings, nil
}

func (s *ServiceSettingsClient) GetGuildSettings(guildId string) (dto.GuildSettingsDTO, error) {
	resp, err := s.client.Get(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", config.GetApiGatewayAddr(), guildId),
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when getting settings", "err", err)
		return dto.GuildSettingsDTO{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	var settings dto.GuildSettingsDTO

	err = json.Unmarshal(body, &settings)

	if err != nil {
		return dto.GuildSettingsDTO{}, nil
	}

	return settings, nil
}

func NewServiceSettingsClient() *ServiceSettingsClient {
	return &ServiceSettingsClient{client: NewDefaultHttpClient()}
}
