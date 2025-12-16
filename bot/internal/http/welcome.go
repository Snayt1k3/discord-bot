package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type Welcome struct {
	http interfaces.HttpClient
}

func NewWelcome() *Welcome {
	return &Welcome{http: NewDefaultHttpClient()}
}

func (s *Welcome) SetChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Put(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/welcome/channel", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when setting welcome channel", "err", err)
		return err
	}

	return nil
}

func (s *Welcome) AddMessage(guildID, message string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id": guildID,
		"message":  message,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/welcome/message", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding welcome message", "err", err)
		return err
	}

	return nil
}

func (s *Welcome) DeleteMessage(guildID, message string) error {
	bodyBytes, err := json.Marshal(map[string]string{
		"guild_id": guildID,
		"message":  message,
	})

	if err != nil {
		return err
	}

	_, err = s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/welcome/message", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when deleting welcome message", "err", err)
		return err
	}

	return nil
}
