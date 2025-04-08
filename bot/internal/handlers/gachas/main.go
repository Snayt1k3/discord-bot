package gachas

import (
	dtoDiscord "bot/internal/dto/discord"
	"bot/internal/handlers/gachas/genshin"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

func showGachas(data dtoDiscord.HandlerData) error {
	buttons := []*discordgo.Button{
		{
			Label:    "Honkai: Star Rail 🌌",
			Style:    discordgo.SecondaryButton,
			CustomID: "HsrCharacters",
		},
		{
			Label:    "Genshin Impact 🔥",
			Style:    discordgo.SecondaryButton,
			CustomID: "GenshinCharacters",
		},
		{
			Label:    "Zenless Zone Zero 🌀",
			Style:    discordgo.SecondaryButton,
			CustomID: "ZenlessCharacters",
		},
		{
			Label:    "Wuthering Waves 🌪️",
			Style:    discordgo.SecondaryButton,
			CustomID: "WuwaCharacters",
		},
	}


	embeds := []*discordgo.MessageEmbed{
		{
			Title:       "📜 **Welcome to the Gacha Hub!**",
			Description: "Ready to dive into the worlds of your favorite gacha games? ✨\n\nFrom builds and ascensions to full team comps — everything you need is just a click away.\n\nChoose a game below to begin exploring characters, guides, and more!",
			Color:       0x9b59b6,
			Image: &discordgo.MessageEmbedImage{
				URL: "https://wotpack.ru/wp-content/uploads/2024/08/ZZZ-potesnila-Genshin-Impact-i-HSR-po-dohodu-za-ijul.jpg",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "**🌌 Honkai: Star Rail**",
					Value: "🚀 Enter the stars and uncover top-tier builds, relics, and path strategies.",
					Inline: false,
				},
				{
					Name: "**🔥 Genshin Impact**",
					Value: "🌍 Explore Teyvat through optimized builds, ascension paths, and full team guides.",
					Inline: false,
				},
				{
					Name: "**🌀 Zenless Zone Zero**",
					Value: "⚙️ Discover W-Engine combos, agent synergies, and stylish squad setups.",
					Inline: false,
				},
				{
					Name: "**🌪️ Wuthering Waves**",
					Value: "🎼 Build your Resonators, learn Echo strategies, and master the combat rhythm of Solaris-3.",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Select a game to start your journey 🌟",
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

func AddHandlers(handlers map[string]func(dtoDiscord.HandlerData) error) {
	genshin.AddGenshinHandlers(handlers)

	handlers["gachas"] = showGachas
}



