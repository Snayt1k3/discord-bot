package interfaces

import (
	"user-service/internal/models"
	pb "user-service/proto"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	GetOrCreateUser(userID, guildID string) (*models.User, error)
	UpdateUser(user *models.User) error
	GetUsers(guildID string, page, size int, orderBy pb.UserOrderBy, isDescSort bool) ([]models.User, error)
	GetCountUsers(guildID string) int
}
