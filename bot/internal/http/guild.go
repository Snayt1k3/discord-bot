package http

import (
	dtoGuild "bot/internal/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type Settings struct {
	http interfaces.HttpClient
}

func NewSettings() *Settings {
	return &Settings{http: NewDefaultHttpClient()}
}

func (s *Settings) Get(guildId string) (dtoGuild.GuildSettings, error) {
	resp, err := s.http.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild", config.GetApiGatewayAddr()),
		map[string]string{"guild_id": guildId},
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
