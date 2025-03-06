package adapters

import (
	"bot/config"
	"bot/internal/dto"
	"bot/internal/errors"
	"bot/internal/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
)

type GuildKeeper struct {
	client interfaces.HttpClient
}

func (s *GuildKeeper) CreateSettings(guild_id string) error {
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

func (s *GuildKeeper) UpdateGuildSettings(guildId string, roles dto.RolesSettings) (dto.GuildSettingsResponse, error) {
	body, _ := json.Marshal(map[string]interface{}{"roles": roles})

	resp, err := s.client.Patch(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", config.GetApiGatewayAddr(), guildId),
		body,
		nil,
	)

	if err != nil {
		return dto.GuildSettingsResponse{}, err
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return dto.GuildSettingsResponse{}, err
	}

	var settings dto.GuildSettingsResponse

	err = json.Unmarshal(body, &settings)

	if err != nil {
		slog.Warn("Bad response when updating settings", "err", err)
		return dto.GuildSettingsResponse{}, nil
	}

	return settings, nil
}

func (s *GuildKeeper) GetGuildSettings(guildId string) (dto.GuildSettingsResponse, error) {
	resp, err := s.client.Get(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", config.GetApiGatewayAddr(), guildId),
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when getting settings", "err", err)
		return dto.GuildSettingsResponse{}, err
	}

	if resp.StatusCode == 404 {
		return dto.GuildSettingsResponse{}, errors.ErrGuildSettingsNotFound
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dto.GuildSettingsResponse{}, err
	}

	var settings dto.GuildSettingsResponse
	
	err = json.Unmarshal(body, &settings)

	if err != nil {
		slog.Error("Failed to unmarshal guild settings", "error", err)
		return dto.GuildSettingsResponse{}, err
	}
	fmt.Println(settings)
	return settings, nil
}

func NewServiceSettingsClient() *GuildKeeper {
	return &GuildKeeper{client: NewDefaultHttpClient()}
}
