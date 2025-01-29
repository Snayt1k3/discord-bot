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

// Метод для получения настроек бота
func (s *SettingsServer) GetBotSettings(ctx context.Context, req *pb.GetBotSettingsRequest) (*pb.GetBotSettingsResponse, error) {
	// Получаем настройки бота через SettingsService
	botSettings, err := s.SettingsService.GetBotSettings()
	if err != nil {
		return nil, err
	}

	// Преобразуем в формат gRPC
	response := &pb.GetBotSettingsResponse{
		BotSettings: &pb.BotSettings{
			Id:            string(botSettings.ID),
			BotStatus:     botSettings.BotStatus,
			Description:   botSettings.Description,
			HelpMessage:   botSettings.HelpMessage,
			HelloMessages: botSettings.HelloMessages,
		},
	}
	return response, nil
}

// Метод для обновления настроек бота
func (s *SettingsServer) UpdateBotSettings(ctx context.Context, req *pb.UpdateBotSettingsRequest) (*pb.UpdateBotSettingsResponse, error) {
	// Создаем объект обновлений на основе запроса
	updateData := dto.BotSettingsUpdate{
		BotStatus:   req.BotStatus,
		Description: req.Description,
		HelpMessage: req.HelpMessage,
		HelloMessages: req.HelloMessages,
	}

	// Вызываем сервис для обновления настроек
	updatedBotSettings, err := s.SettingsService.UpdateBotSettings(updateData)
	if err != nil {
		return nil, err
	}

	// Преобразуем обновленные данные в формат gRPC
	response := &pb.UpdateBotSettingsResponse{
		BotSettings: &pb.BotSettings{
			Id:            string(updatedBotSettings.ID),
			BotStatus:     updatedBotSettings.BotStatus,
			Description:   updatedBotSettings.Description,
			HelpMessage:   updatedBotSettings.HelpMessage,
			HelloMessages: updatedBotSettings.HelloMessages,
		},
	}
	return response, nil
}

// Получить настройки по Guild ID
func (s *SettingsServer) GetByGuildID(ctx context.Context, req *pb.GetSettingsByGuildRequest) (*pb.GetSettingsByGuildResponse, error) {
	// Получаем настройки гильдии через SettingsService
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
				IsDisabled: guildSettings.Roles.IsDisabled,
			},
		},
	}
	return response, nil
}

// Получить все настройки для гильдий
func (s *SettingsServer) GetAllGuildSettings(ctx context.Context, req *pb.GetAllGuildSettingsRequest) (*pb.GetAllGuildSettingsResponse, error) {
	// Получаем все настройки гильдий
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
				IsDisabled: setting.Roles.IsDisabled,
			},
			},
		)
	}

	response := &pb.GetAllGuildSettingsResponse{
		Settings: settingsList,
	}
	return response, nil
}

// Обновить настройки гильдии
func (s *SettingsServer) UpdateGuildSettings(ctx context.Context, req *pb.UpdateGuildSettingsRequest) (*pb.UpdateGuildSettingsResponse, error) {
	// Создаем объект обновлений на основе запроса
	updateData := dto.GuildSettingsUpdateDTO{
		Roles: dto.RolesSettings{
			MesssageId: req.Roles.MessageId,
			Matching:   req.Roles.Matching,
			IsDisabled: req.Roles.IsDisabled,
		},
	}

	// Вызываем сервис для обновления настроек
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
					IsDisabled: updatedGuildSettings.Roles.IsDisabled,
				},
		},
	}
	return response, nil
}
