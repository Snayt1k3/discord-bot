package server

import (
	"context"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"
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
	eventTypes := make([]models.EventType, len(req.EventType))
	for i, et := range req.EventType {
		eventTypes[i] = models.EventType(et)
	}
	res, err := s.Repo.AddLogs(req.GuildId, req.ChannelId, eventTypes)

	if err != nil {
		return nil, err
	}

	var eventSettings []*pb.EventSettings
	for _, ev := range res.Events {
		eventSettings = append(eventSettings, &pb.EventSettings{ChannelId: ev.ChannelID, Event: pb.EventType(ev.EventType)})
	}

	response := &pb.UpdateLogChannelResponse{
		GuildId:   req.GuildId,
		EventSettings: eventSettings,
	}

	return response, nil
}

func (s *LogServer) RemoveLogChannel(ctx context.Context, req *pb.UpdateLogChannelRequest) (*pb.RemoveLogChannelResponse, error) {
	eventTypes := make([]models.EventType, len(req.EventType))
	for i, et := range req.EventType {
		eventTypes[i] = models.EventType(et)
	}

	err := s.Repo.RemoveLogs(req.GuildId, req.ChannelId, eventTypes)

	if err != nil {
		return nil, err
	}

	response := &pb.RemoveLogChannelResponse{
		GuildId:   req.GuildId,
	}

	return response, nil
}
