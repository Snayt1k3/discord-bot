package settings

import (
	dtoDiscord "bot/internal/dto/discord"
	er "bot/internal/errors"
	"errors"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

func settingsHandler(data dtoDiscord.HandlerData) error {

	_, err := data.Gk.GetGuildSettings(data.Event.GuildID)

	if errors.Is(err, er.ErrGuildSettingsNotFound) {
		data.Gk.CreateSettings(data.Event.GuildID)
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

	err = data.Session.InteractionRespond(data.Event.Interaction, message)
	if err != nil {
		slog.Error("Error sending settings message", "err", err)
		return err
	}

	return nil
}

func AddSettingsHandlers(handlers map[string]func(data dtoDiscord.HandlerData) error) {
	handlers["add-role-reactions"] = addRole
	handlers["remove-role-reactions"] = removeRole
	handlers["set-roles-message-id"] = setMessageId
	handlers["set-welcome-channel"] = setChannelId
	handlers["settings"] = settingsHandler
	handlers["view_reaction_roles"] = showAllRoles
	// todo добавить проверку на админа. И поменять сообщение: Улучшить оформление
}
