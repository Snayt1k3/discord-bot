package server

import (
	"context"
	"interaction-service/internal/adapters/repos"
	"interaction-service/internal/interfaces"
	pb "interaction-service/proto"
	"time"

	"gorm.io/gorm"
)

type UserServer struct {
	repo interfaces.UserRepo
	pb.UnimplementedInteractionServiceServer
}

func NewUserServer(db *gorm.DB) *UserServer {
	return &UserServer{
		repo: repos.NewUserRepo(db),
	}
}

func (s *UserServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetOrCreateUser(request.GetUserId(), request.GetGuildId())

	if err != nil { // if user doesn't exist, we create a new one
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			UserId:        user.UserID,
			GuildId:       user.GuildID,
			Experience:    int32(user.Experience),
			Level:         int32(user.Level),
			NextLevelXp:   int32(user.NextLevelXP),
			LastMessageAt: user.LastMessageAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *UserServer) AddXP(ctx context.Context, req *pb.AddXPRequest) (*pb.AddXPResponse, error) {
	user, err := s.repo.GetOrCreateUser(req.GetUserId(), req.GetGuildId())

	if err != nil {
		return nil, err
	}

	currentExperience := int32(user.Experience) + req.Xp
	nextLevelXP := int32(user.NextLevelXP)
	leveledUp := false
	user.LastMessageAt = time.Now()

	if currentExperience >= int32(user.NextLevelXP) {
		user.Experience = int(currentExperience - int32(user.NextLevelXP))
		user.NextLevelXP = int(float32(nextLevelXP) * 1.5)
		user.Level += 1
		leveledUp = true
	} else {
		user.Experience = int(currentExperience)
	}

	err = s.repo.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	return &pb.AddXPResponse{
		User: &pb.User{
			UserId:        user.UserID,
			GuildId:       user.GuildID,
			Experience:    int32(user.Experience),
			Level:         int32(user.Level),
			NextLevelXp:   int32(user.NextLevelXP),
			LastMessageAt: user.LastMessageAt.Format(time.RFC3339),
		},
		LevelUp: leveledUp,
		AddedXp: req.Xp,
	}, nil
}
