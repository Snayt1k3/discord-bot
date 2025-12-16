package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)

type Roles struct {
	http interfaces.HttpClient
}

func NewRoles() *Roles {
	return &Roles{http: NewDefaultHttpClient()}
}

func (s *Roles) Add(roleId, emoji, guildID string) error {

	bytes, _ := json.Marshal(map[string]string{
		"role_id":  roleId,
		"emoji":    emoji,
		"guild_id": guildID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/roles/role", config.GetApiGatewayAddr()),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding role", "err", err)
		return err
	}

	return nil
}

func (s *Roles) Delete(roleId, emoji, guildID string) error {

	bytes, _ := json.Marshal(map[string]string{
		"role_id":  roleId,
		"emoji":    emoji,
		"guild_id": guildID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/roles/role", config.GetApiGatewayAddr()),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when deleting role", "err", err)
		return err
	}

	return nil
}

func (s *Roles) SetMessageID(messageID, guildID string) error {
	bytes, _ := json.Marshal(map[string]string{
		"message_id": messageID,
		"guild_id":   guildID,
	})

	_, err := s.http.Put(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/roles/message", config.GetApiGatewayAddr()),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when setting role message ID", "err", err)
		return err
	}

	return nil

}
