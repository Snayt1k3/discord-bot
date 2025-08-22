package guild

import (
	"bot/internal/discord"
	"bot/internal/interfaces"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func setWelcomeChannel(gk interfaces.GuildServiceInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	channelId := i.ApplicationCommandData().Options[0].ChannelValue(nil).ID

	_, err := gk.SetWelcomeChannel(i.GuildID, channelId)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		discord.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel set successfully! New members will be welcomed in this channel.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

// func AddWelcomeMessage(gk interfaces.GuildServiceInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
// 	msg := i.ApplicationCommandData().Options[0].ChannelValue(nil).ID

// 	_, err := gk.AddWelcomeMessage(i.GuildID, msg)

// 	if err != nil {
// 		slog.Error("Error while updating welcome settings", "err", err)
// 		discord.SendErrorMessage(s, i)
// 		return err
// 	}

// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: "Message added successfully!",
// 			Flags:   discordgo.MessageFlagsEphemeral,
// 		},
// 	})
// 	return nil
// }