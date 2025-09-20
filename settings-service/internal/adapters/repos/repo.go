package repos

import (
	"gorm.io/gorm"
	"settings-service/internal/interfaces"
)

type Repos struct {
	AutoMode      interfaces.AutoModeRepository
	Log           interfaces.LogRepository
	ReactionRoles interfaces.ReactionRolesRepository
	Settings      interfaces.GuildSettingsRepository
	Welcome       interfaces.WelcomeRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		AutoMode:      NewAutoModeRepository(db),
		Log:           NewLogRepository(db),
		ReactionRoles: NewReactionRolesRepo(db),
		Settings:      NewGuildSettingsRepo(db),
		Welcome:       NewWelcomeRepository(db),
	}
}
