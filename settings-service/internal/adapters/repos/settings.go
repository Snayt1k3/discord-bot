package repos

import (
	"gorm.io/gorm"
	"settings-service/internal/models"
)


type GuildSettingsRepoImpl struct {
	db *gorm.DB
}

func NewGuildSettingsRepo(db *gorm.DB) *GuildSettingsRepoImpl {
	return &GuildSettingsRepoImpl{
		db: db,
	}
}

func (r *GuildSettingsRepoImpl) CreateGuildSetting(guildId string) (models.Settings, error) {
	obj := &models.Settings{
		GuildID: guildId,
		Role: models.RolesSettings{
			GuildID: guildId,
		},
		Welcome: models.WelcomeSettings{
			GuildID: guildId,
		},
		AutoMode: models.AutoModeSettings{
			GuildID: guildId,
			Enabled: false,
		},
		Log: models.LogSettings{
			GuildID: guildId,
			Enabled: false,
		},

	}
	if err := r.db.Create(obj).Error; err != nil {
		return models.Settings{}, err
	}

	return *obj, nil
}

func (r *GuildSettingsRepoImpl) GetGuildSettings(guildID string) (*models.Settings, error) {
	var settings models.Settings
	if err := r.db.
	Preload("Role.Roles").
	Preload("Welcome.Messages").
	Preload("AutoMode.BannedWords").
	Preload("AutoMode.AntiCapsChannels").
	Preload("AutoMode.AntiLinkChannels").
	Where("guild_id = ?", guildID).First(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}