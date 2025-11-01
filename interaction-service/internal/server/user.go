package server

import (
	"context"
	"errors"
	"interaction-service/internal/adapters/repos"
	"interaction-service/internal/interfaces"
	"interaction-service/internal/models"
	pb "interaction-service/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	user, err := s.repo.GetUser(request.GetUserId(), request.GetGuildId())

	if err != nil { // if user doesn't exist, we create a new one
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = &models.User{
				UserID:        request.GetUserId(),
				GuildID:       request.GetGuildId(),
				Experience:    0,
				Level:         1,
				LastMessageAt: time.Now(),
				NextLevelXP:   50,
			}

			err = s.repo.CreateUser(user)

			if err != nil {
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
	user, err := s.repo.GetUser(req.GetUserId(), req.GetGuildId())

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "User not found")
		}
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
