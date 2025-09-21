package settings

import (
	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
)

type GuildPreferencesHandlers struct {
	guildService guild.GuildAdapter
}

func NewSettingsHandlers(guildService guild.GuildAdapter) *GuildPreferencesHandlers {
	return &GuildPreferencesHandlers{
		guildService: guildService,
	}
}

func (g *GuildPreferencesHandlers) addRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return addRole(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return removeRole(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) setMessageId(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return setRolesMessage(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) showAllRoles(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return showAllRoles(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) setWelcomeChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return setWelcomeChannel(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) addWelcomeMessage(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return AddWelcomeMessage(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeWelcomeMessage(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return DeleteWelcomeMessage(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) showWelcomeSettings(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return showWelcomeSettings(g.guildService, s, i)
}

func (gp *GuildPreferencesHandlers) AddHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {

	// Настройка Roles/Reactions
	handlers["rr-add"] = gp.addRole
	handlers["rr-remove"] = gp.removeRole
	handlers["rr-message"] = gp.setMessageId

	// Настройка Welcome
	handlers["welcome-chnl"] = gp.setWelcomeChannel
	handlers["welcomemsg-add"] = gp.addWelcomeMessage
	handlers["welcomemsg-remove"] = gp.removeWelcomeMessage

	// Меню, Кнопки связанные с меню
	handlers["RolesReactionsSettings"] = gp.showAllRoles
	handlers["WelcomeSettings"] = gp.showWelcomeSettings
}
