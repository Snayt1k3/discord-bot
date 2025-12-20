package server

import (
	"context"
	"errors"
	"settings-service/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"settings-service/internal/interfaces"
	pb "settings-service/proto"
)

type GuildServer struct {
	Repo interfaces.GuildSettingsRepository
	pb.UnimplementedSettingsServiceServer
}

func (s *GuildServer) GetSettings(ctx context.Context, req *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	guildSettings, err := s.Repo.GetGuildSettings(req.GuildId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if guildSettings == nil {
		guildSettings, err = s.Repo.CreateGuildSetting(req.GuildId)

		if err != nil {
			return nil, err
		}
	}

	return utils.MapGuildSettings(guildSettings)
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

	return utils.MapCreateGuildSettings(settings)
}
