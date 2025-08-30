package server

import (
	"context"

	"settings-service/internal/interfaces"
	pb "settings-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SettingsServer struct {
	GuildRepo interfaces.GuildRepository
	pb.UnimplementedGuildServiceServer
}

func (s *SettingsServer) GetSettings(ctx context.Context, req *pb.GetSettingsByGuildRequest) (*pb.GetSettingsByGuildResponse, error) {
	guildSettings, err := s.GuildRepo.GetGuildSettings(req.GuildId)

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



	response := &pb.GetSettingsByGuildResponse{
		Settings: &pb.GuildSettings{
			GuildId:       guildSettings.GuildID,
			Roles: &pb.RolesSettings{
				MessageId: guildSettings.Role.MessageID,
				Matching: matching,
			},

			Welcome: &pb.WelcomeSettings{
				ChannelId: guildSettings.Welcome.ChannelId,
				Messages:  welcomeMessages,
			},
		},
	}
	return response, nil
}

func (s *SettingsServer) CreateSettings(ctx context.Context, req *pb.CreateGuildSettingsRequest) (*pb.CreateGuildSettingsResponse, error) {

	err := s.GuildRepo.CreateGuildSetting(req.GuildId)

	if err != nil {
		return nil, err
	}

	return &pb.CreateGuildSettingsResponse{GuildId: req.GuildId}, nil
}

func (s *SettingsServer) AddRole(ctx context.Context, req *pb.Role) (*pb.Role, error) {
	err := s.GuildRepo.AddRole(req.GuildId, req.RoleId, req.Emoji)

	if err != nil {
		return nil, err
	}

	response := &pb.Role{
		GuildId: req.GuildId,
		RoleId:  req.RoleId,
		Emoji:   req.Emoji,
	}
	return response, nil
}

func (s *SettingsServer) DeleteRole(ctx context.Context, req *pb.RoleDelete) (*pb.Role, error) {
	err := s.GuildRepo.DeleteRole(req.GuildId, req.RoleId)

	if err != nil {
		return nil, err
	}

	response := &pb.Role{
		GuildId: req.GuildId,
		RoleId:  req.RoleId,
		Emoji:   "",
	}
	return response, nil
}

func (s *SettingsServer) SetRoleMessageId(ctx context.Context, req *pb.RoleMessage) (*pb.RoleMessage, error) {
	err := s.GuildRepo.SetRoleMessageId(req.GuildId, req.MessageId)

	if err != nil {
		return nil, err
	}

	response := &pb.RoleMessage{
		GuildId:   req.GuildId,
		MessageId: req.MessageId,
	}
	return response, nil
}

func (s *SettingsServer) SetWelcomeChannel(ctx context.Context, req *pb.SetWelcomeChannelRequest) (*pb.SetWelcomeChannelResponse, error) {
	err := s.GuildRepo.SetWelcomeChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.SetWelcomeChannelResponse{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}
	return response, nil
}

func (s *SettingsServer) AddWelcomeMessage(ctx context.Context, req *pb.WelcomeMessageRequest) (*pb.WelcomeMessageResponse, error) {
	err := s.GuildRepo.AddWelcomeMessage(req.GuildId, req.Message)

	if err != nil {
		return nil, err
	}

	response := &pb.WelcomeMessageResponse{
		GuildId: req.GuildId,
		Message: req.Message,
	}
	return response, nil
}

func (s *SettingsServer) DeleteWelcomeMessage(ctx context.Context, req *pb.WelcomeMessageRequest) (*pb.WelcomeMessageResponse, error) {
	err := s.GuildRepo.DeleteWelcomeMessage(req.GuildId, req.Message)

	if err != nil {
		return nil, err
	}

	response := &pb.WelcomeMessageResponse{
		GuildId: req.GuildId,
		Message: req.Message,
	}
	return response, nil
}
