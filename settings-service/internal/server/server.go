package server

import (
	"context"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
	"settings-service/internal/dto"

)

type SettingsServer struct {
	SettingsService interfaces.SettingsInterface
	pb.UnimplementedSettingsServiceServer
}


func (s *SettingsServer) GetByGuildID(ctx context.Context, req *pb.GetSettingsByGuildRequest) (*pb.GetSettingsByGuildResponse, error) {
	guildSettings, err := s.SettingsService.GetByGuildID(req.GuildId)
	if err != nil {
		return nil, err
	}

	// Преобразуем настройки в формат gRPC
	response := &pb.GetSettingsByGuildResponse{
		Settings: &pb.GuildSettings{
			Id:      string(guildSettings.ID),
			GuildId: guildSettings.GuildID,
			Roles: &pb.RolesSettings{
				MessageId: guildSettings.Roles.MesssageId,
				Matching:  guildSettings.Roles.Matching,

			},
		},
	}
	return response, nil
}


func (s *SettingsServer) GetAllGuildSettings(ctx context.Context, req *pb.GetAllGuildSettingsRequest) (*pb.GetAllGuildSettingsResponse, error) {
	guildSettingsList, err := s.SettingsService.GetAllGuildSettings()
	if err != nil {
		return nil, err
	}

	// Преобразуем в формат gRPC
	var settingsList []*pb.GuildSettings
	for _, setting := range guildSettingsList {
		settingsList = append(settingsList, &pb.GuildSettings{
			Id:      string(setting.ID),
			GuildId: setting.GuildID,
			Roles: &pb.RolesSettings{
				MessageId: setting.Roles.MesssageId,
				Matching:  setting.Roles.Matching,
			},
			},
		)
	}

	response := &pb.GetAllGuildSettingsResponse{
		Settings: settingsList,
	}
	return response, nil
}


func (s *SettingsServer) UpdateGuildSettings(ctx context.Context, req *pb.UpdateGuildSettingsRequest) (*pb.UpdateGuildSettingsResponse, error) {
	updateData := dto.GuildSettingsUpdateDTO{
		Roles: dto.RolesSettings{
			MesssageId: req.Roles.MessageId,
			Matching:   req.Roles.Matching,
		},
	}

	updatedGuildSettings, err := s.SettingsService.UpdateGuildSettings(req.Id, updateData)
	
	if err != nil {
		return nil, err
	}

	// Преобразуем обновленные данные в формат gRPC
	response := &pb.UpdateGuildSettingsResponse{
		GuildSettings: &pb.GuildSettings{
			Id:      string(updatedGuildSettings.ID),
			GuildId: updatedGuildSettings.GuildID,

				Roles: &pb.RolesSettings{
					MessageId: updatedGuildSettings.Roles.MesssageId,
					Matching:  updatedGuildSettings.Roles.Matching,

				},
		},
	}
	return response, nil
}

func (s *SettingsServer) CreateGuildSettings(ctx context.Context, req *pb.CreateGuildSettingsRequest) (*pb.CreateGuildSettingsResponse, error) {
	data := dto.GuildSettingsCreateDTO{GuildId: req.GuildId}
	
	err := s.SettingsService.CreateGuildSetting(data)

	if err != nil {
		return nil, err
	}

	return &pb.CreateGuildSettingsResponse{GuildId: data.GuildId}, nil
}