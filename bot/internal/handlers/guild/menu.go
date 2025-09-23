package guild

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func menu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	buttons := menuButtons("MainMenuSettings")

	message := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "** Friren – Server Control Panel**\n\n" +
				"Hello, I'm Friren! I'll help you set up some useful features for your server ✨\n\n" +
				"🔹 **Reaction/Roles** – configure adding or removing roles when users react to a specific message.\n" +
				"🔹 **Welcome** – choose a channel and customize the message to greet new members.\n\n" +
				"Use the buttons below to open the settings 👇",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: buttons,
				},
			},
		},
	}

	err := s.InteractionRespond(i.Interaction, message)

	if err != nil {
		slog.Error("Failed to respond to interaction", "err", err)
		return err
	}

	return nil
}

func backToMenu(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	buttons := menuButtons("MainMenuSettings")

	content := "** Friren – Server Control Panel**\n\n" +
		"Hello, I'm Friren! I'll help you set up some useful features for your server ✨\n\n" +
		"🔹 **Reaction/Roles** – configure adding or removing roles when users react to a specific message.\n" +
		"🔹 **Welcome** – choose a channel and customize the message to greet new members.\n\n" +
		"Use the buttons below to open the settings 👇"

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

func menuButtons(disabledIDs ...string) []discordgo.MessageComponent {
	isDisabled := func(id string) bool {
		for _, d := range disabledIDs {
			if d == id {
				return true
			}
		}
		return false
	}
	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "Roles/Reactions",
			Style:    discordgo.PrimaryButton,
			CustomID: "RolesReactionsSettings",
			Disabled: isDisabled("RolesReactionsSettings"),
			Emoji: &discordgo.ComponentEmoji{
				Name: "⚙️",
			},
		},
		discordgo.Button{
			Label:    "Welcome",
			Style:    discordgo.PrimaryButton,
			CustomID: "WelcomeSettings",
			Disabled: isDisabled("WelcomeSettings"),
			Emoji: &discordgo.ComponentEmoji{
				Name: "👋",
			},
		},
		discordgo.Button{
			Label:    "AutoMode",
			Style:    discordgo.PrimaryButton,
			CustomID: "AutoModeSettings",
			Disabled: isDisabled("AutoModeSettings"),
			Emoji: &discordgo.ComponentEmoji{
				Name: "🤖",
			},
		},
	}
	return buttons
}
