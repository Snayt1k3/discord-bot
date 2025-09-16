package server

import (
	"context"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
)


type WelcomeServer struct {
	Repo interfaces.WelcomeRepository
	pb.UnimplementedWelcomeServiceServer
}


func (s *WelcomeServer) SetWelcomeChannel(ctx context.Context, req *pb.SetWelcomeChannelRequest) (*pb.SetWelcomeChannelResponse, error) {
	err := s.Repo.SetWelcomeChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.SetWelcomeChannelResponse{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}
	return response, nil
}

func (s *WelcomeServer) AddWelcomeMessage(ctx context.Context, req *pb.WelcomeMessageRequest) (*pb.WelcomeMessageResponse, error) {
	err := s.Repo.AddWelcomeMessage(req.GuildId, req.Message)

	if err != nil {
		return nil, err
	}

	response := &pb.WelcomeMessageResponse{
		GuildId: req.GuildId,
		Message: req.Message,
	}
	return response, nil
}

func (s *WelcomeServer) DeleteWelcomeMessage(ctx context.Context, req *pb.WelcomeMessageRequest) (*pb.WelcomeMessageResponse, error) {
	err := s.Repo.DeleteWelcomeMessage(req.GuildId, req.Message)

	if err != nil {
		return nil, err
	}

	response := &pb.WelcomeMessageResponse{
		GuildId: req.GuildId,
		Message: req.Message,
	}
	return response, nil
}