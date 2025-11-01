package repos

import (
	"interaction-service/internal/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetUser(userID, guildID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("user_id = ? AND guild_id = ?", userID, guildID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}
