package repos

import (
	"errors"
	"interaction-service/internal/models"
	"time"

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

func (r *UserRepo) GetOrCreateUser(userID, guildID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("user_id = ? AND guild_id = ?", userID, guildID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = models.User{
				UserID:        userID,
				GuildID:       guildID,
				Experience:    0,
				Level:         1,
				LastMessageAt: time.Now(),
				NextLevelXP:   50,
			}

			err = r.CreateUser(&user)

			if err != nil {
				return nil, err
			}
			return &user, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}
