package guild

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type AutoModeAdapter struct {
	http interfaces.HttpClient
}

func NewAutoMode(http interfaces.HttpClient) *AutoModeAdapter {
	return &AutoModeAdapter{http: http}
}

func (s *AutoModeAdapter) Toggle(guildID string, enable bool) error {
	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"guild_id": guildID,
		"enabled":  enable,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/toggle", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when toggling automode", "err", err)
		return err
	}
	return nil
}

func (s *AutoModeAdapter) AddCapsLockChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/capslock", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding capslock channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoModeAdapter) RemoveCapsLockChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/capslock", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing capslock channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoModeAdapter) AddAntiLinkChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/antilink", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding AntiLink channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoModeAdapter) RemoveAntiLinkChannel(guildID, channelID string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/antilink", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing capslock channel", "err", err)
		return err
	}
	return nil
}

func (s *AutoModeAdapter) AddBannedWord(guildID, word string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id": guildID,
		"word":     word,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/bannedword", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding banned word", "err", err)
		return err
	}
	return nil
}

func (s *AutoModeAdapter) RemoveBannedWord(guildID, word string) error {
	bodyBytes, _ := json.Marshal(map[string]string{
		"guild_id": guildID,
		"word":     word,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/automode/bannedword", config.GetApiGatewayAddr(), guildID),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when removing banned word", "err", err)
		return err
	}
	return nil
}
