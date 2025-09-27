package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
)

type RolesReactionServer struct {
	Repo      interfaces.ReactionRolesRepository
	GuildRepo interfaces.GuildSettingsRepository
	pb.UnimplementedRolesServiceServer
}

func (s *RolesReactionServer) AddRole(ctx context.Context, req *pb.AddRoleRequest) (*pb.AddRoleResponse, error) {

	settungs, err := s.GuildRepo.GetGuildSettings(req.GuildId)

	if err != nil {
		return nil, err
	}

	if len(settungs.AutoMode.CapsLocks) >= 50 {
		return nil, status.Errorf(
			codes.ResourceExhausted,
			"Reaction roles limit (50) reached",
		)
	}

	role, err := s.Repo.AddRole(req.GuildId, req.RoleId, req.Emoji)

	if err != nil {
		return nil, err
	}

	response := &pb.AddRoleResponse{
		GuildId: req.GuildId,
		RoleId:  role.RoleID,
		Emoji:   role.Emoji,
	}
	return response, nil
}

func (s *RolesReactionServer) DeleteRole(ctx context.Context, req *pb.RemoveRoleResponse) (*pb.RemoveRoleResponse, error) {
	err := s.Repo.DeleteRole(req.GuildId, req.RoleId)

	if err != nil {
		return nil, err
	}

	response := &pb.RemoveRoleResponse{
		GuildId: req.GuildId,
		RoleId:  req.RoleId,
	}
	return response, nil
}

func (s *RolesReactionServer) SetRoleMessageId(ctx context.Context, req *pb.SetMessageRequest) (*pb.SetMessageResponse, error) {
	err := s.Repo.SetRoleMessageId(req.GuildId, req.MessageId)

	if err != nil {
		return nil, err
	}

	response := &pb.SetMessageResponse{
		GuildId:   req.GuildId,
		MessageId: req.MessageId,
	}
	return response, nil
}
