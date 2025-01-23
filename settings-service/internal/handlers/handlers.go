package handlers

import (
	"context"
	pb "settings-service/proto"
)

type SettingsServer struct {
	pb.UnimplementedSettingsServiceServer
}

func (s *SettingsServer) UpdateBotSettings(ctx context.Context, req *pb.UpdateBotSettingsRequest) (*pb.UpdateBotSettingsResponse, error) {
	return nil, nil
}

func (s *SettingsServer) GetBotSettings(ctx context.Context, req *pb.GetBotSettingsRequest) (*pb.GetBotSettingsResponse, error) {
	return nil, nil
}

func (s *SettingsServer) GetByGuildID(ctx context.Context, req *pb.GetSettingsByGuildRequest) (*pb.GetSettingsByGuildResponse, error) {
	return nil, nil
}

func (s *SettingsServer) GetAllGuildSettings(ctx context.Context, req *pb.GetAllGuildSettingsRequest) (*pb.GetAllGuildSettingsResponse, error) {
	return nil, nil
}

func (s *SettingsServer) UpdateGuildSettings(ctx context.Context, req *pb.UpdateGuildSettingsRequest) (*pb.UpdateGuildSettingsResponse, error) {
	return nil, nil
}

func (s *SettingsServer) DeteleSetting(ctx context.Context, req *pb.DeleteGuildSettingRequest) (*pb.DeleteGuildSettingResponse, error) {
	return nil, nil
}