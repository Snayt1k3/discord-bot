package guild

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"bot/config"
	"bot/internal/interfaces"
)


type RolesAdapter struct {
	http interfaces.HttpClient
}

func NewRolesAdapter(http interfaces.HttpClient) *RolesAdapter {
	return &RolesAdapter{http: http}
}

func (s *RolesAdapter) Add(roleId, emoji, guildID string) (error) {

	bytes, _ := json.Marshal(map[string]string{
		"role_id":  roleId,
		"emoji":    emoji,
		"guild_id": guildID,
	})

	_, err := s.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/roles/role", config.GetApiGatewayAddr(), guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding role", "err", err)
		return err
	}

	return nil
}

func (s *RolesAdapter) Delete(roleId, emoji, guildID string) (error) {

	bytes, _ := json.Marshal(map[string]string{
		"role_id":  roleId,
		"emoji":    emoji,
		"guild_id": guildID,
	})

	_, err := s.http.Delete(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/roles/role", config.GetApiGatewayAddr(), guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when deleting role", "err", err)
		return err
	}

	return nil
}

func (s *RolesAdapter) SetMessageID(messageID, guildID string) (error) {
	bytes, _ := json.Marshal(map[string]string{
		"message_id": messageID,
		"guild_id":   guildID,
	})

	_, err := s.http.Put(
		context.Background(),
		fmt.Sprintf("%v/api/v1/settings/guild/%v/roles/message", config.GetApiGatewayAddr(), guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when setting role message ID", "err", err)
		return err
	}

	return nil
	
}