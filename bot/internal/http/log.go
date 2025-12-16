package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type Log struct {
	http interfaces.HttpClient
}

func NewLog() *Log {
	return &Log{http: NewDefaultHttpClient()}
}

func (s *Log) Toggle(guildID string, enable bool) error {

	bytes, _ := json.Marshal(map[string]interface{}{
		"guild_id": guildID,
		"enabled":  enable,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/logging/toggle", config.GetApiGatewayAddr()),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when toggling logging", "err", err)
		return err
	}
	return nil
}

func (s *Log) AddChannel(guildID, channelID string) error {
	bytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/logging/channel", config.GetApiGatewayAddr()),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding log channel", "err", err)
		return err
	}
	return nil
}

func (s *Log) RemoveChannel(guildID, channelID string) error {
	bytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/logging/channel", config.GetApiGatewayAddr()),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing log channel", "err", err)
		return err
	}
	return nil
}
