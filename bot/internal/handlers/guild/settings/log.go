package settings

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
	"bot/internal/utils"
)

func ToggleLogging(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := guildService.Settings.Get(i.GuildID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	err = guildService.Log.Toggle(i.GuildID, !settings.Log.Enabled)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(("Logging has been %s successfully!"), map[bool]string{true: "enabled", false: "disabled"}[!settings.Log.Enabled]),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddLoggingChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	channelID := i.ApplicationCommandData().Options[0].ChannelValue(s).ID
	err := guildService.Log.AddChannel(i.GuildID, channelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Logging channel has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func RemoveLoggingChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := guildService.Settings.Get(i.GuildID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	err = guildService.Log.RemoveChannel(i.GuildID, settings.Log.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Logging channel has been removed successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}
