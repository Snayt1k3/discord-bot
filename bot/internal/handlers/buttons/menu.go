package buttons

import "github.com/bwmarrin/discordgo"

// Menu
// Available ids [RolesReactionsSettings, WelcomeSettings, AutoModeSettings, LogSettings]
func Menu(disabledIDs ...string) []discordgo.MessageComponent {
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
		discordgo.Button{
			Label:    "Logging",
			Style:    discordgo.PrimaryButton,
			CustomID: "LogSettings",
			Disabled: isDisabled("LogSettings"),
			Emoji: &discordgo.ComponentEmoji{
				Name: "📜",
			},
		},
	}
	return buttons
}
