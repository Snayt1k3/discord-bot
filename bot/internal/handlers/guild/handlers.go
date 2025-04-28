package guild

import (
	er "bot/internal/errors"
	"bot/internal/interfaces"
	"errors"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type GuildPreferencesHandlers struct {
	Gk interfaces.GuildKeeperInterface
}

func NewSettingsHandlers(gk interfaces.GuildKeeperInterface) *GuildPreferencesHandlers {
	return &GuildPreferencesHandlers{
		Gk: gk,
	}
}
func (gp *GuildPreferencesHandlers) showGuildPreferences(s *discordgo.Session, i *discordgo.InteractionCreate) error {

	_, err := gp.Gk.GetGuildSettings(i.GuildID)

	if errors.Is(err, er.ErrGuildSettingsNotFound) {
		gp.Gk.CreateSettings(i.GuildID) // todo: перенести в обработчик при добаление бота на сервер
	}

	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "⚙️ Show all roles.",
			Style:    discordgo.SuccessButton,
			CustomID: "view_reaction_roles",
			Emoji: &discordgo.ComponentEmoji{
				Name: "🔧",
			},
		},
	}

	message := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "**⚙️ Server Settings**\n\n" +
				"Welcome to the settings panel! Here you can manage various aspects of your server.\n\n" +
				"🔹 *Click the button below to see all roles configured for this server!*\n\n" +
				"**🔧 Admin Commands:**\n" +
				"- `/add-role-reactions <role> <emoji>` – Add a role reaction.\n" +
				"- `/remove-role-reactions <role>` – Remove a role reaction.\n" +
				"- `/set-roles-message-id <message_id>` – Set the message ID for role reactions.\n" +
				"- `/set-welcome-channel <channel_id>` – Set the channel ID for new users.\n\n" +
				"*(Only administrators can use these commands.)*",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: buttons,
				},
			},
		},
	}

	err = s.InteractionRespond(i.Interaction, message)

	if err != nil {
		slog.Error("Error sending settings message", "err", err)
		return err
	}

	return nil
}

func (gp *GuildPreferencesHandlers) addRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return addRole(gp.Gk, s, i)
}

func (gp *GuildPreferencesHandlers) removeRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return removeRole(gp.Gk, s, i)
}

func (gp *GuildPreferencesHandlers) setMessageId(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return setRolesMessage(gp.Gk, s, i)
}

func (gp *GuildPreferencesHandlers) showAddedRoles(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return showAllRoles(gp.Gk, s, i)
}

func (gp *GuildPreferencesHandlers) setWelcomeChannel(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return setWelcomeChannel(gp.Gk, s, i)
}

func (gp *GuildPreferencesHandlers) AddSettingsHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	handlers["add-role-reactions"] = gp.addRole
	handlers["remove-role-reactions"] = gp.removeRole
	handlers["set-roles-message-id"] = gp.setMessageId
	handlers["set-welcome-channel"] = gp.setWelcomeChannel
	handlers["settings"] = gp.showGuildPreferences
	handlers["ViewReactionRoles"] = gp.showAddedRoles
	// todo добавить проверку на админа. И поменять сообщение: Улучшить оформление
}
