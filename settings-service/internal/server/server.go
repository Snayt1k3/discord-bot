package server

import (
	"context"

	"settings-service/internal/dto"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SettingsServer struct {
	SettingsService interfaces.SettingsService
	pb.UnimplementedSettingsServiceServer
}

func (s *SettingsServer) GetSettingsByGuild(ctx context.Context, req *pb.GetSettingsByGuildRequest) (*pb.GetSettingsByGuildResponse, error) {
	guildSettings, err := s.SettingsService.GetSettingsByGuildID(req.GuildId)
	
	if err != nil {
		return nil, err
	}

	if guildSettings.ID == 0 {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

	response := &pb.GetSettingsByGuildResponse{
		Settings: &pb.GuildSettings{
			Id:      strconv.Itoa(int((guildSettings.ID))),
			GuildId: guildSettings.GuildID,
			Roles: &pb.RolesSettings{
				MessageId: guildSettings.Roles.MessageId,
				Matching:  guildSettings.Roles.Matching,
			},
		},
	}
	return response, nil
}

func (s *SettingsServer) CreateGuildSettings(ctx context.Context, req *pb.CreateGuildSettingsRequest) (*pb.CreateGuildSettingsResponse, error) {

	err := s.SettingsService.CreateGuildSettings(req.GuildId)

	if err != nil {
		return nil, err
	}

	return &pb.CreateGuildSettingsResponse{GuildId: req.GuildId}, nil
}

func (s *SettingsServer) UpdateRoles(ctx context.Context, req *pb.UpdateRolesRequest) (*pb.UpdateRolesResponse, error) {
	settings, err := s.SettingsService.UpdateRolesSettings(&dto.RolesSettings{
		MessageId: req.MessageId,
		Matching:   req.Roles,
		GuildID:    req.GuildId,
	})

	if err != nil {
		return nil, err
	}

	response := &pb.UpdateRolesResponse{
		GuildSettings: &pb.GuildSettings{
			Id:      strconv.Itoa(int((settings.ID))),
			GuildId: settings.GuildID,
			Roles: &pb.RolesSettings{
				MessageId: settings.Roles.MessageId,
				Matching:  settings.Roles.Matching,
			},
		},
	}

	return response, nil

}
