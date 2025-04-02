package handlers

import (
	dtoDiscord "bot/internal/dto/discord"
	"github.com/bwmarrin/discordgo"


)

func HelpHandler(data dtoDiscord.HandlerData) error {
	helpMessage := "**🌿 Frieren Bot - Traces of Music 🌿**\n" +
		"Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:\n\n" +
		"**🎼 Commands to Guide the Melody:**\n" +
		"- `/play <song_name/link>` – Let the music flow, one song at a time.\n" +
		"- `/pause` – Even melodies need a moment of rest.\n" +
		"- `/resume` – Continue where you left off, like an old journey resumed.\n" +
		"- `/stop` – Bring the music to a quiet end, clearing all that remains.\n" +
		"- `/skip` – Move past this tune, towards the next story in sound.\n\n" +

		"**📖 Knowledge in the Wind:**\n" +
		"- `/help` – If you have forgotten, let this guide you once more.\n\n" +

		"**🌾 A Few Words of Caution:**\n" +
		"- A melody can only be heard if you are present—join a voice channel first.\n" +
		"- If questions linger, seek wisdom from those who lead this place.\n\n" +
		"Music drifts like memories in the wind. Enjoy it while it lasts. 🎧"

	data.Session.InteractionRespond(
		data.Event.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: helpMessage,
			},
		},
	)
	return nil
}


