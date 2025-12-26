package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type AutoMode struct {
	http interfaces.HttpClient
}

func NewAutoMode() *AutoMode {
	return &AutoMode{http: NewDefaultHttpClient()}
}

func (s *AutoMode) Toggle(guildID string, enable bool) error {
	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"guild_id": guildID,
		"enabled":  enable,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/toggle", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when toggling automode", "err", err)
		return err
	}
	return nil
}

func (s *AutoMode) AddCapsLockChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/capslock", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding capslock channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoMode) RemoveCapsLockChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/capslock", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing capslock channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoMode) AddAntiLinkChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/antilink", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding AntiLink channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoMode) RemoveAntiLinkChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/antilink", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing capslock channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoMode) AddBannedWord(guildID, word string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id": guildID,
		"word":     word,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/bannedword", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding banned word", "err", err)
		return err
	}
	return nil
}

func (s *AutoMode) RemoveBannedWord(guildID, word string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id": guildID,
		"word":     word,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/automode/bannedword", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing banned word", "err", err)
		return err
	}
	return nil
}
