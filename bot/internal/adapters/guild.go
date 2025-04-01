package adapters

import (
	"bot/config"
	dtoGuild "bot/internal/dto/settings"
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
	slog.Info("Creating settings", "guild_id", guild_id)

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

func (s *GuildKeeper) UpdateRolesSetting(guildId string, roles dtoGuild.RolesSettings) (dtoGuild.GuildSettingsResponse, error) {
	body, _ := json.Marshal(map[string]any{"roles": roles.Matching, "message_id": roles.MessageId})

	slog.Info("Updating roles settings", "body", string(body))

	resp, err := s.client.Patch(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/roles", config.GetApiGatewayAddr(), guildId),
		body,
		nil,
	)

	if err != nil {
		slog.Error("Bad response when updating settings", "err", err)
		return dtoGuild.GuildSettingsResponse{}, err
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return dtoGuild.GuildSettingsResponse{}, err
	}

	var settings dtoGuild.GuildSettingsResponse

	err = json.Unmarshal(body, &settings)

	if err != nil {
		slog.Error("Bad response when updating settings", "err", err)
		return dtoGuild.GuildSettingsResponse{}, nil
	}

	return settings, nil
}

func (s *GuildKeeper) GetGuildSettings(guildId string) (dtoGuild.GuildSettingsResponse, error) {
	slog.Info("Getting guild settings", "guild_id", guildId)

	resp, err := s.client.Get(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", config.GetApiGatewayAddr(), guildId),
		nil,
	)

	if err != nil {
		slog.Error("Failed to get guild settings", "error", err)
		return dtoGuild.GuildSettingsResponse{}, err
	}

	if resp.StatusCode == 404 {
		slog.Warn("Guild settings not found", "guild_id", guildId)
		return dtoGuild.GuildSettingsResponse{}, errors.ErrGuildSettingsNotFound
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dtoGuild.GuildSettingsResponse{}, err
	}

	var settings dtoGuild.GuildSettingsResponse

	err = json.Unmarshal(body, &settings)

	if err != nil {
		slog.Error("Failed to unmarshal guild settings", "error", err)
		return dtoGuild.GuildSettingsResponse{}, err
	}

	return settings, nil
}

func (s *GuildKeeper) UpdateWelcomeSetting(guildId string, welcome dtoGuild.WelcomeSettings) error {
	body, _ := json.Marshal(map[string]any{"channel_id": welcome.ChannelId})

	slog.Info("Updating welcome settings", "body", string(body))

	resp, err := s.client.Patch(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/welcome", config.GetApiGatewayAddr(), guildId),
		body,
		nil,
	)

	if err != nil {
		slog.Error("Bad response when updating settings", "err", err)
		return err
	}

	defer resp.Body.Close()

	return nil
}

func NewServiceSettingsClient() *GuildKeeper {
	return &GuildKeeper{client: NewDefaultHttpClient()}
}
