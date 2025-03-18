package settings

import (
	"bot/internal/discord"
	"bot/internal/dto"
	"bot/internal/interfaces"

	"github.com/bwmarrin/discordgo"
)

func SetChannelId(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	channelId := i.ApplicationCommandData().Options[0].StringValue()
	guildId := i.GuildID

	err := guildKeeper.UpdateWelcomeSetting(guildId, dto.WelcomeSettings{ChannelId: channelId})
	if err != nil {
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
