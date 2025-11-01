package server

import (
	"interaction-service/internal/adapters/repos"
	"interaction-service/internal/interfaces"

	"gorm.io/gorm"
)

type UserServer struct {
	repo interfaces.UserRepo
}

func NewUserServer(db *gorm.DB) *UserServer {
	return &UserServer{
		repo: repos.NewUserRepo(db),
	}
}
