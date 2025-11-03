package guild

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"bot/config"
	dtoGuild "bot/internal/dto/guild"
	"bot/internal/interfaces"
)

type SettingsAdapter struct {
	http interfaces.HttpClient
}

func NewSettingsAdapter(http interfaces.HttpClient) *SettingsAdapter {
	return &SettingsAdapter{http: http}
}

func (s *SettingsAdapter) Get(guild_id string) (dtoGuild.GuildSettings, error) {
	resp, err := s.http.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v", config.GetApiGatewayAddr(), guild_id),
		nil,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when getting settings", "err", err)
		return dtoGuild.GuildSettings{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dtoGuild.GuildSettings{}, err
	}

	var settings dtoGuild.GuildSettings

	err = json.Unmarshal(body, &settings)

	if err != nil {
		slog.Error("Bad response when getting settings", "err", err)
		return dtoGuild.GuildSettings{}, nil
	}

	return settings, nil

}

func (s *SettingsAdapter) Create(guild_id string) error {
	resp, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v", config.GetApiGatewayAddr(), guild_id),
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
