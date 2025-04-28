package guild

import (
	"bot/internal/discord"
	dtoGuild "bot/internal/dto/settings"
	"bot/internal/interfaces"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func setWelcomeChannel(gk interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	channelId := i.ApplicationCommandData().Options[0].ChannelValue(nil).ID
	guildId := i.GuildID

	err := gk.UpdateWelcomeSetting(guildId, dtoGuild.WelcomeSettings{ChannelId: channelId})

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
