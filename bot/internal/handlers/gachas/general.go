package gachas

import (
	dto "bot/internal/dto/discord"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

func ShowGachas(data dto.HandlerData) error {
	buttons := []*discordgo.Button{
		{
			Label:    "Honkai: Star Rail 🌌",
			Style:    discordgo.PrimaryButton,
			CustomID: "hsr",
		},
		{
			Label:    "Genshin Impact 🔥",
			Style:    discordgo.PrimaryButton,
			CustomID: "genshin",
		},
		{
			Label:    "Zenless Zone Zero 🌀",
			Style:    discordgo.PrimaryButton,
			CustomID: "zenless",
		},
		{
			Label:    "Wuthering Waves 🌊",
			Style:    discordgo.PrimaryButton,
			CustomID: "wuwa",
		},
	}

	// Create message with buttons and embed
	embeds := []*discordgo.MessageEmbed{
		{
			Title:       "🎮 **Choose Your Gacha!**",
			Description: "Don't miss the chance to learn more about the games! 🤩\n\nEach game offers unique guides, builds, and the opportunity to sign up for pulls. Choose your favorite game below:",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://wotpack.ru/wp-content/uploads/2024/08/ZZZ-potesnila-Genshin-Impact-i-HSR-po-dohodu-za-ijul.jpg", // Image URL
			},
			Color: 0x1F7A1F, // Green color
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "**Honkai: Star Rail 🌌**",
					Value:  "🚀 *Guides, Builds, Pull Recordings* - A guide to the universe of **Honkai: Star Rail**.",
					Inline: false,
				},
				{
					Name:   "**Genshin Impact 🔥**",
					Value:  "🌍 *Guides, Builds, Pull Recordings* - Dive into the world of **Teyvat** with deep mechanics and characters.",
					Inline: false,
				},
				{
					Name:   "**Zenless Zone Zero 🌀**",
					Value:  "⚙️ *Guides, Builds, Pull Recordings* - A new world with exciting gameplay and intriguing mechanics.",
					Inline: false,
				},
				{
					Name:   "**Wuthering Waves 🌊**",
					Value:  "🌌 *Guides, Builds, Pull Recordings* - A unique game with a dynamic world and combat mechanics.",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Gather your team and set off on an adventure! 🚀",
			},
		},
	}

	err := data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					buttons[0], buttons[1], buttons[2],
				}},

				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					buttons[3],
				}},
			},
		},
	})

	if err != nil {
		slog.Error("Failed to send message", "error", err)
		return err
	}

	return nil
}

func GenshinButtonHandler(data dto.HandlerData) error {
	// Handle Genshin button click
	slog.Info("Genshin button clicked")
	return nil
}

func HSRButtonHandler(data dto.HandlerData) error {
	// Handle HSR button click
	slog.Info("HSR button clicked")
	return nil
}

func ZenlessButtonHandler(data dto.HandlerData) error {
	// Handle Zenless button click
	slog.Info("Zenless button clicked")
	return nil
}

func WuwaButtonHandler(data dto.HandlerData) error {
	// Handle Wuwa button click
	slog.Info("Wuwa button clicked")
	return nil
}
