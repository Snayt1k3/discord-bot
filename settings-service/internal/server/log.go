package server

import (
	"context"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

type LogServer struct {
	Repo interfaces.LogRepository
	pb.UnimplementedLogServiceServer
}

func (s *LogServer) ToggleLog(ctx context.Context, req *pb.ToggleLogRequest) (*pb.ToggleLogResponse, error) {
	err := s.Repo.ToggleLog(req.GuildId, req.Enabled)

	if err != nil {
		return nil, err
	}

	response := &pb.ToggleLogResponse{
		GuildId: req.GuildId,
		Enabled: req.Enabled,
	}
	return response, nil
}

func (s *LogServer) AddLogChannel(ctx context.Context, req *pb.UpdateLogChannelRequest) (*pb.UpdateLogChannelResponse, error) {
	_, err := s.Repo.AddLogChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.UpdateLogChannelResponse{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}

	return response, nil
}

func (s *LogServer) RemoveLogChannel(ctx context.Context, req *pb.UpdateLogChannelRequest) (*pb.UpdateLogChannelResponse, error) {
	err := s.Repo.RemoveLogChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.UpdateLogChannelResponse{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}

	return response, nil
}
