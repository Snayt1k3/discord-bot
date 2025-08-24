package guild

import "github.com/bwmarrin/discordgo"



func menu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "Roles/Reactions.",
			Style:    discordgo.PrimaryButton,
			CustomID: "RolesReactionsSettings",
			Emoji: &discordgo.ComponentEmoji{
				Name: "⚙️",
			},
		},
		discordgo.Button{
			Label:    "Welcome.",
			Style:    discordgo.PrimaryButton,
			CustomID: "WelcomeSettings",
			Emoji: &discordgo.ComponentEmoji{
				Name: "👋",
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

	s.InteractionRespond(i.Interaction, message)
	
	
	return nil
}


func backToMenu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "Roles/Reactions.",
			Style:    discordgo.PrimaryButton,
			CustomID: "RolesReactionsSettings",
			Emoji: &discordgo.ComponentEmoji{
				Name: "⚙️",
			},
		},
		discordgo.Button{
			Label:    "Welcome.",
			Style:    discordgo.PrimaryButton,
			CustomID: "WelcomeSettings",
			Emoji: &discordgo.ComponentEmoji{
				Name: "👋",
			},
		},
	}

	content := "**⚙️ Server Settings**\n\n" +
    "Welcome to the settings panel! Here you can manage various aspects of your server.\n\n" +
    "🔹 *Click the button below to see all roles configured for this server!*\n\n" +
    "**🔧 Admin Commands:**\n" +
    "- `/add-role-reactions <role> <emoji>` – Add a role reaction.\n" +
    "- `/remove-role-reactions <role>` – Remove a role reaction.\n" +
    "- `/set-roles-message-id <message_id>` – Set the message ID for role reactions.\n" +
    "- `/set-welcome-channel <channel_id>` – Set the channel ID for new users.\n\n" +
    "*(Only administrators can use these commands.)*"

	message := &discordgo.WebhookEdit{
		Content: &content,
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: buttons,
			},
		},
	}

	_, err := s.InteractionResponseEdit(i.Interaction, message)
	return err
}