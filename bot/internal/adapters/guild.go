package adapters

import (
	"bot/config"
	dtoGuild "bot/internal/dto/guild"
	"bot/internal/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
)

type GuildService struct {
	client interfaces.HttpClient
	gatewayAddr string
}

func (s *GuildService) GetGuildSettings(guild_id string) (dtoGuild.GuildSettings, error) {
	resp, err := s.client.Get(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", s.gatewayAddr, guild_id),
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

func (s *GuildService) CreateSettings(guild_id string) error {
	resp, err := s.client.Post(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v", s.gatewayAddr, guild_id),
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

func (s *GuildService) AddRole(roleId, emoji, guildID string) (dtoGuild.Role, error) {
	
	bytes, _ := json.Marshal(map[string]string{
		"role_id": roleId,
		"emoji":   emoji,
		"guild_id": guildID,
	})
	
	resp, err := s.client.Post(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/roles/role", s.gatewayAddr, guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding role", "err", err)
		return dtoGuild.Role{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dtoGuild.Role{}, err
	}

	var addedRole dtoGuild.Role

	err = json.Unmarshal(body, &addedRole)

	if err != nil {
		slog.Error("Bad response when adding role", "err", err)
		return dtoGuild.Role{}, err
	}

	return addedRole, nil
}

func (s *GuildService) DeleteRole(roleId, emoji, guildID string) (dtoGuild.Role, error) {
	
	bytes, _ := json.Marshal(map[string]string{
		"role_id": roleId,
		"emoji":   emoji,
		"guild_id": guildID,
	})
	
	resp, err := s.client.Delete(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/roles/role", s.gatewayAddr, guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding role", "err", err)
		return dtoGuild.Role{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dtoGuild.Role{}, err
	}

	var addedRole dtoGuild.Role

	err = json.Unmarshal(body, &addedRole)

	if err != nil {
		slog.Error("Bad response when adding role", "err", err)
		return dtoGuild.Role{}, err
	}

	return addedRole, nil
}

func (s *GuildService) SetRoleMessageID(messageID, guildID string) (dtoGuild.RoleMessage, error) {
	bytes, _ := json.Marshal(map[string]string{
		"message_id": messageID,
		"guild_id":   guildID,
	})

	resp, err := s.client.Put(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/roles/message", s.gatewayAddr, guildID),
		bytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when setting role message ID", "err", err)
		return dtoGuild.RoleMessage{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return dtoGuild.RoleMessage{}, err
	}

	var roleMessage dtoGuild.RoleMessage

	err = json.Unmarshal(body, &roleMessage)

	if err != nil {
		slog.Error("Bad response when setting role message ID", "err", err)
		return dtoGuild.RoleMessage{}, err
	}

	return roleMessage, nil
}

func (s *GuildService) SetWelcomeChannel(guildID, channelID string) (dtoGuild.SetWelcomeChannelResponse, error) {
	bodyBytes, err := json.Marshal(map[string]string{
		"guild_id":   guildID,
		"channel_id": channelID,
	})
	if err != nil {
		return dtoGuild.SetWelcomeChannelResponse{}, err
	}

	resp, err := s.client.Put(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/welcome/channel", s.gatewayAddr, guildID),
		bodyBytes,
		nil,
	)
	if err != nil {
		return dtoGuild.SetWelcomeChannelResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtoGuild.SetWelcomeChannelResponse{}, err
	}

	var result dtoGuild.SetWelcomeChannelResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return dtoGuild.SetWelcomeChannelResponse{}, err
	}

	return result, nil
}

func (s *GuildService) AddWelcomeMessage(guildID, message string) (dtoGuild.WelcomeMessageResponse, error) {
	bodyBytes, err := json.Marshal(map[string]string{
		"guild_id": guildID,
		"message":  message,
	})
	if err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}

	resp, err := s.client.Post(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/welcome/message", s.gatewayAddr, guildID),
		bodyBytes,
		nil,
	)
	if err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}

	var result dtoGuild.WelcomeMessageResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}

	return result, nil
}

func (s *GuildService) DeleteWelcomeMessage(guildID, message string) (dtoGuild.WelcomeMessageResponse, error) {
	bodyBytes, err := json.Marshal(map[string]string{
		"guild_id": guildID,
		"message":  message,
	})
	if err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}

	resp, err := s.client.Delete(
		context.Background(),
		fmt.Sprintf("%v/settings/guild/%v/welcome/message", s.gatewayAddr, guildID),
		bodyBytes,
		nil,
	)
	if err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}

	var result dtoGuild.WelcomeMessageResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return dtoGuild.WelcomeMessageResponse{}, err
	}

	return result, nil
}

func NewServiceSettingsClient() interfaces.GuildServiceInterface {
	return &GuildService{client: NewDefaultHttpClient(), gatewayAddr: config.GetApiGatewayAddr()}
}
