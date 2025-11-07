package interfaces

import (
	"interaction-service/internal/models"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	GetOrCreateUser(userID, guildID string) (*models.User, error)
	UpdateUser(user *models.User) error
}
