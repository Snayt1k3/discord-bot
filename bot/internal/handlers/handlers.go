package handlers

import (
	dtoDiscord "bot/internal/dto/discord"
	"github.com/bwmarrin/discordgo"
)

func HelpHandler(data dtoDiscord.HandlerData) error {
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "🌿 Frieren Bot - Traces of Music 🌿",
		Description: "Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:",
		Color:       0x2ECC71, // Soft green color
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/2f/eb/71/2feb71b7fb35c35886b87324b6cef144.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "🎼 Commands to Guide the Melody:",
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
				Name:  "🌾 A Few Words of Caution:",
				Value: "- A melody can only be heard if you are present—join a voice channel first.\n" +
					"- If questions linger, seek wisdom from those who lead this place.",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Music drifts like memories in the wind. Enjoy it while it lasts. 🎧",
		},
	}

	data.Session.InteractionRespond(
		data.Event.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{helpEmbed},
			},
		},
	)
	return nil
}




