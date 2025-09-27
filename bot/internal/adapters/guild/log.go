package guild

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type LogAdapter struct {
	http interfaces.HttpClient
}

func NewLogAdapter(http interfaces.HttpClient) *LogAdapter {
	return &LogAdapter{http: http}
}

func (s *LogAdapter) Toggle(guildID string, enable bool) error {

	bytes, _ := json.Marshal(map[string]interface{}{
		"guild_id": guildID,
		"enable":   enable,
	})

	_, err := s.http.Patch(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/logging/toggle", config.GetApiGatewayAddr(), guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when toggling logging", "err", err)
		return err
	}
	return nil
}

func (s *LogAdapter) AddChannel(guildID, channelID string) error {
	bytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/logging/channel", config.GetApiGatewayAddr(), guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding log channel", "err", err)
		return err
	}
	return nil
}

func (s *LogAdapter) RemoveChannel(guildID, channelID string) error {
	bytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/logging/channel", config.GetApiGatewayAddr(), guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing log channel", "err", err)
		return err
	}
	return nil
}
