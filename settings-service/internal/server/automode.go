package server

import (
	"context"
	"fmt"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
)

type AutomodeServer struct {
	Repo interfaces.AutoModeRepository
	pb.UnimplementedAutoModServiceServer
}

func (s *AutomodeServer) ToggleAutoMode(ctx context.Context, req *pb.ToggleAutoModRequest) (*pb.ToggleAutoModResponse, error) {
	_, err := s.Repo.ToggleAutoMode(req.GuildId, req.Enabled)

	if err != nil {
		return nil, err
	}

	response := &pb.ToggleAutoModResponse{
		GuildId: req.GuildId,
		Enabled: req.Enabled,
	}
	return response, nil
}

func (s *AutomodeServer) AddBannedWord(ctx context.Context, req *pb.AddBannedWordRequest) (*pb.AddBannedWordResponse, error) {
	res, err := s.Repo.AddBannedWord(req.GuildId, req.Word)

	if err != nil {
		return nil, err
	}

	response := &pb.AddBannedWordResponse{
		Id:      fmt.Sprintf("%d", res.ID),
		GuildId: req.GuildId,
		Word:    req.Word,
	}

	return response, nil
}

func (s *AutomodeServer) RemoveBannedWord(ctx context.Context, req *pb.RemoveBannedWordRequest) (*pb.RemoveBannedWordResponse, error) {
	err := s.Repo.DeleteBannedWord(req.GuildId, req.Word)

	if err != nil {
		return nil, err
	}

	response := &pb.RemoveBannedWordResponse{
		GuildId: req.GuildId,
		Word:    req.Word,
	}

	return response, nil
}

func (s *AutomodeServer) AddCapsLockChannel(ctx context.Context, req *pb.AddCapsLockChannelRequest) (*pb.AddCapsLockChannelResponse, error) {
	res, err := s.Repo.AddCapsLockChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.AddCapsLockChannelResponse{
		Id:        fmt.Sprintf("%d", res.ID),
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}

	return response, nil
}

func (s *AutomodeServer) RemoveCapsLockChannel(ctx context.Context, req *pb.RemoveCapsLockChannelRequest) (*pb.RemoveCapsLockChannelResponse, error) {
	err := s.Repo.DeleteCapsLockChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.RemoveCapsLockChannelResponse{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}

	return response, nil
}

func (s *AutomodeServer) AddAntiLinkChannel(ctx context.Context, req *pb.AddAntiLinkChannelRequest) (*pb.AddAntiLinkChannelResponse, error) {
	res, err := s.Repo.AddAntiLinkChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.AddAntiLinkChannelResponse{
		Id:        fmt.Sprintf("%d", res.ID),
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}

	return response, nil
}

func (s *AutomodeServer) RemoveAntiLinkChannel(ctx context.Context, req *pb.RemoveAntiLinkChannelRequest) (*pb.RemoveAntiLinkChannelResponse, error) {
	err := s.Repo.DeleteAntiLinkChannel(req.GuildId, req.ChannelId)

	if err != nil {
		return nil, err
	}

	response := &pb.RemoveAntiLinkChannelResponse{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	}

	return response, nil
}
