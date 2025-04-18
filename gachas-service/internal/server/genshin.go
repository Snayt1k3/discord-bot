package server

import (
	"context"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	pb "gachas-service/proto"
)

// todo implement the methods
type GenshinServer struct {
	pb.UnimplementedGenshinServiceServer
	service interfaces.ServiceInterface[dto.GenshinCharacter, dto.GenshinCharacterBrief, dto.GenshinBuild]
}

func (s *GenshinServer) GetCharacterById(ctx context.Context, req *pb.CharacterRequest) (*pb.CharacterDetailResponse, error) {
	return nil, nil
}

func GetCharacterBuild(ctx context.Context, req *pb.CharacterRequest) (*pb.BuildResponse, error) {
	return nil, nil
}

func GetCharacterById(ctx context.Context, req *pb.CharacterRequest) (*pb.CharacterDetailResponse, error) {
	return nil, nil
}

func NewGenshinServer(service interfaces.ServiceInterface[dto.GenshinCharacter, dto.GenshinCharacterBrief, dto.GenshinBuild]) *GenshinServer {
	return &GenshinServer{
		service: service,
	}
}
