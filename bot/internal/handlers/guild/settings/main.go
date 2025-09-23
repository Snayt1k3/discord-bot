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
	return ShowWelcomeSettings(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) showAutoModeSettings(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return ShowAutoModeSettings(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) addBannedWord(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return AddBannedWord(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeBannedWord(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return RemoveBannedWord(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) addAntiLinkChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return AddAntiLinkChannel(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeAntiLinkChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return RemoveAntiLinkChannel(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) addAntiCapsLock(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return AddCapsLockChannel(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeAntiCapsLock(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return RemoveCapsLockChannel(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) toggleLogging(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return ToggleLogging(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) addLoggingChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return AddLoggingChannel(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeLoggingChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return RemoveLoggingChannel(g.guildService, s, i)
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
	handlers["AutoModeSettings"] = gp.showAutoModeSettings

	// Automode
	handlers["automod-bw-add"] = gp.addBannedWord
	handlers["automod-bw-rm"] = gp.removeBannedWord
	handlers["automod-al-enable"] = gp.addAntiLinkChannel
	handlers["automod-al-disable"] = gp.removeAntiLinkChannel
	handlers["automod-ac-enable"] = gp.addAntiCapsLock
	handlers["automod-ac-disable"] = gp.removeAntiCapsLock

	// Logging
	handlers["log-toggle"] = gp.toggleLogging
	handlers["log-chnl"] = gp.addLoggingChannel
	handlers["log-chnl-rm"] = gp.removeLoggingChannel
}
