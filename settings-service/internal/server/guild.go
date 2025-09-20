package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"settings-service/internal/interfaces"
	pb "settings-service/proto"
)

type GuildServer struct {
	Repo interfaces.GuildSettingsRepository
	pb.UnimplementedSettingsServiceServer
}

func (s *GuildServer) GetSettings(ctx context.Context, req *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	guildSettings, err := s.Repo.GetGuildSettings(req.GuildId)

	if err != nil {
		return nil, err
	}

	if guildSettings.ID == 0 {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

	matching := make(map[string]string)
	var welcomeMessages []string

	for _, role := range guildSettings.Role.Roles {
		matching[role.Emoji] = role.RoleID
	}

	for _, msg := range guildSettings.Welcome.Messages {
		welcomeMessages = append(welcomeMessages, msg.Message)
	}

	response := &pb.GetSettingsResponse{
		GuildId: guildSettings.GuildID,
		Roles: &pb.RolesSettings{
			MessageId: guildSettings.Role.MessageID,
			Matching:  matching,
		},

		Welcome: &pb.WelcomeSettings{
			ChannelId: guildSettings.Welcome.ChannelId,
			Messages:  welcomeMessages,
		}, // todo: Добавить остальные настройки
	}

	return response, nil
}

func (s *GuildServer) CreateSettings(ctx context.Context, req *pb.CreateSettingsRequest) (*pb.CreateSettingsResponse, error) {

	guildSettings, _ := s.Repo.GetGuildSettings(req.GuildId)

	if guildSettings != nil && guildSettings.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, "Guild settings already exist")
	}

	settings, err := s.Repo.CreateGuildSetting(req.GuildId)

	if err != nil {
		return nil, err
	}

	return &pb.CreateSettingsResponse{GuildId: settings.GuildID}, nil // todo: return full settings
}
