package guild

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)


type WelcomeAdapter struct {
	http interfaces.HttpClient
}

func NewWelcomeAdapter(http interfaces.HttpClient) *WelcomeAdapter {
	return &WelcomeAdapter{http: http}
}


func (s *WelcomeAdapter) SetChannel(guildID, channelID string) (error) {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Put(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/welcome/channel", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when setting welcome channel", "err", err)
		return err	
	}

	return nil
}

func (s *WelcomeAdapter) AddMessage(guildID, message string) (error) {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id": guildID,
		"message":  message,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/welcome/message", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding welcome message", "err", err)
		return err
	}
	
	return nil
}

func (s *WelcomeAdapter) DeleteMessage(guildID, message string) (error) {
	bodyBytes, err := json.Marshal(map[string]string{
		"guild_id": guildID,
		"message":  message,
	})

	if err != nil {
		return err
	}

	_, err = s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/welcome/message", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when deleting welcome message", "err", err)
		return err 
	}

	return nil
}