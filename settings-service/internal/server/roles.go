package server

import (
	"context"
	"settings-service/internal/interfaces"
	pb "settings-service/proto"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)


type RolesReactionServer struct {
	Repo interfaces.ReactionRolesRepository
	pb.UnimplementedRolesServiceServer
}

func (s *RolesReactionServer) AddRole(ctx context.Context, req *pb.AddRoleRequest) (*pb.AddRoleResponse, error) {
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