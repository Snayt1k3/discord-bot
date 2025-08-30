package guild

import (
	"bot/internal/interfaces"
	"github.com/bwmarrin/discordgo"
)

type GuildPreferencesHandlers struct {
	guildService interfaces.GuildServiceInterface
}

func NewSettingsHandlers(guildService interfaces.GuildServiceInterface) *GuildPreferencesHandlers {
	return &GuildPreferencesHandlers{
		guildService: guildService,
	}
}
func (g *GuildPreferencesHandlers) menu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return menu(s, i)
}

func (g *GuildPreferencesHandlers) backToMenu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return backToMenu(s, i)
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

func (gp *GuildPreferencesHandlers) AddSettingsHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	
	// Настройка Roles/Reactions
	handlers["rr-add"] = gp.addRole
	handlers["rr-remove"] = gp.removeRole
	handlers["rr-message"] = gp.setMessageId

	// Настройка Welcome
	handlers["set-welcome-channel"] = gp.setWelcomeChannel
	handlers["add-welcome-msg"] = gp.setWelcomeChannel

	// Меню, Кнопки связанные с меню
	handlers["settings"] = gp.menu
	handlers["MainMenuSettings"] = gp.backToMenu
	handlers["RolesReactionsSettings"] = gp.showAllRoles
	// todo добавить проверку на админа. И поменять сообщение: Улучшить оформление
}
