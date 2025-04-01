package settings

import (
	"bot/internal/discord"
	dtoDiscord "bot/internal/dto/discord"
	dtoGuild "bot/internal/dto/settings"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func SetChannelId(data dtoDiscord.HandlerData) {
	channelId := data.Event.ApplicationCommandData().Options[0].ChannelValue(nil).ID
	guildId := data.Event.GuildID

	err := data.Gk.UpdateWelcomeSetting(guildId, dtoGuild.WelcomeSettings{ChannelId: channelId})

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel set",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
