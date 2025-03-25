package settings

import (
	"bot/internal/discord"
	"bot/internal/dto"
	"bot/internal/interfaces"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func SetChannelId(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	channelId := i.ApplicationCommandData().Options[0].ChannelValue(nil).ID
	guildId := i.GuildID

	err := guildKeeper.UpdateWelcomeSetting(guildId, dto.WelcomeSettings{ChannelId: channelId})
	
	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel set",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
