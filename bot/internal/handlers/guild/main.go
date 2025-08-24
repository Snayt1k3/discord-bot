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

func (g *GuildPreferencesHandlers) addRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return addRole(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) removeRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return removeRole(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) setMessageId(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return setRolesMessage(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) showAddedRoles(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return showAllRoles(g.guildService, s, i)
}

func (g *GuildPreferencesHandlers) setWelcomeChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return setWelcomeChannel(g.guildService, s, i)
}

func (gp *GuildPreferencesHandlers) AddSettingsHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	handlers["add-role-reactions"] = gp.addRole
	handlers["remove-role-reactions"] = gp.removeRole
	handlers["set-roles-message-id"] = gp.setMessageId
	handlers["set-welcome-channel"] = gp.setWelcomeChannel
	handlers["menu"] = gp.menu
	handlers["ViewReactionRoles"] = gp.showAddedRoles
	// todo добавить проверку на админа. И поменять сообщение: Улучшить оформление
}
