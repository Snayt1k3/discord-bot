package commands

import (
	"bot/internal/handlers/buttons"
	"bot/internal/utils"

	"github.com/bwmarrin/discordgo"
)

func Menu(s *discordgo.Session, i *discordgo.InteractionCreate) {
	btns := buttons.Menu("MainMenuSettings")

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
					Components: btns,
				},
			},
		},
	}

	utils.Respond(s, i, message)
}
