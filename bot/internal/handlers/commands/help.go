package commands

import (
	"bot/internal/utils"

	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "🌿 Frieren Bot - Traces of Music 🌿",
		Description: "Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:",
		Color:       0x2ECC71,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/2f/eb/71/2feb71b7fb35c35886b87324b6cef144.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "📖 Knowledge in the Wind:",
				Value: "`/help` – If you have forgotten, let this guide you once more.",
			},
		},
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{helpEmbed},
		},
	}

	utils.Respond(s, i, resp)
}
