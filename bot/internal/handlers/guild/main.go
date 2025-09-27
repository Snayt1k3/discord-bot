package guild

import (
	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
	"bot/internal/handlers/guild/settings"
)

type Handlers struct {
	guildService guild.GuildAdapter

	settings settings.GuildPreferencesHandlers
}

func NewHandlers(guildService guild.GuildAdapter) *Handlers {
	return &Handlers{
		guildService: guildService,
		settings:     *settings.NewSettingsHandlers(guildService),
	}
}
func (h *Handlers) menu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return menu(s, i)
}

func (h *Handlers) backToMenu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return backToMenu(s, i)
}

func (h *Handlers) AddHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	handlers["settings"] = h.menu
	handlers["MainMenuSettings"] = h.backToMenu

	h.settings.AddHandlers(handlers)
}
