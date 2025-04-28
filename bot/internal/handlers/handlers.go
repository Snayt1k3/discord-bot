package handlers

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func HelpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "🌿 Frieren Bot - Traces of Music 🌿",
		Description: "Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:",
		Color:       0x2ECC71,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/2f/eb/71/2feb71b7fb35c35886b87324b6cef144.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "🎼 Commands to Guide the Melody:",
				Value: "`/play <song_name/link>` – Let the music flow, one song at a time.\n" +
					"`/pause` – Even melodies need a moment of rest.\n" +
					"`/resume` – Continue where you left off, like an old journey resumed.\n" +
					"`/stop` – Bring the music to a quiet end, clearing all that remains.\n" +
					"`/skip` – Move past this tune, towards the next story in sound.",
			},
			{
				Name:  "📖 Knowledge in the Wind:",
				Value: "`/help` – If you have forgotten, let this guide you once more.",
			},
			{
				Name: "🌾 A Few Words of Caution:",
				Value: "- A melody can only be heard if you are present—join a voice channel first.\n" +
					"- If questions linger, seek wisdom from those who lead this place.",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Music drifts like memories in the wind. Enjoy it while it lasts. 🎧",
		},
	}

	s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{helpEmbed},
			},
		},
	)
	return nil
}

func ShowSupportedGachas(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	buttons := []*discordgo.Button{
		{
			Label:    "Honkai: Star Rail 🌌",
			Style:    discordgo.SecondaryButton,
			CustomID: "HsrCharacters",
			Disabled: true,
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
			Disabled: true,
		},
		{
			Label:    "Wuthering Waves 🌪️",
			Style:    discordgo.SecondaryButton,
			CustomID: "WuwaCharacters",
			Disabled: true,
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
					Name:   "**🌌 Honkai: Star Rail**",
					Value:  "🚀 Enter the stars and uncover top-tier builds, relics, and path strategies.",
					Inline: false,
				},
				{
					Name:   "**🔥 Genshin Impact**",
					Value:  "🌍 Explore Teyvat through optimized builds, ascension paths, and full team guides.",
					Inline: false,
				},
				{
					Name:   "**🌀 Zenless Zone Zero**",
					Value:  "⚙️ Discover W-Engine combos, agent synergies, and stylish squad setups.",
					Inline: false,
				},
				{
					Name:   "**🌪️ Wuthering Waves**",
					Value:  "🎼 Build your Resonators, learn Echo strategies, and master the combat rhythm of Solaris-3.",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Select a game to start your journey 🌟",
			},
		},
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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
