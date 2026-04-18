package repos

import (
	"errors"
	"fmt"
	"user-service/internal/models"
	pb "user-service/proto"
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
				VoiceTime: 0,
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

func (r *UserRepo) GetUsers(guildID string, page, size int, orderBy pb.UserOrderBy, isDescSort bool) ([]models.User, error) {
	var users []models.User

	sortDir := "ASC"
	
	if isDescSort {
		sortDir = "DESC"
	}

	var orderColumn string

	switch orderBy {
	case pb.UserOrderBy_VOICE_TIME:
		orderColumn = "voice_time"
	case pb.UserOrderBy_EXPERIENCE:
		orderColumn = "level"
	default:
		orderColumn = "level"
	}

	err := r.db.Where("guild_id = ?", guildID).
		Offset(page * size).
		Limit(size).
		Order(fmt.Sprintf("%s %s", orderColumn, sortDir)).
		Find(&users).Error

	return users, err
}

func (r *UserRepo) GetCountUsers(guildID string) int {
	var count int64
	r.db.Model(&models.User{}).Where("guild_id = ?", guildID).Count(&count)
	return int(count)
}
