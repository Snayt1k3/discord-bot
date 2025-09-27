package settings

import (
	"fmt"
	"log/slog"

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

func ShowLogSettings(gk guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// Check admin permissions
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	// Get guild settings
	settings, err := gk.Settings.Get(i.GuildID)

	if err != nil {
		slog.Error("Error while fetching welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	if !settings.Log.Enabled {
		// Если AutoMode полностью выключен
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "🚨 **Logging is disabled on this server.**",
			},
		})
	}

	// Format channel mention
	channelMention := "—"
	color := 0xED4245 // red

	if settings.Log.ChannelID != "" {
		channelMention = "<#" + settings.Welcome.ChannelID + ">"
		color = 0x57F287 // green
	}

	embed := &discordgo.MessageEmbed{
		Title: "📜 Logging Events configuration",
		Color: color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "📍 Channel",
				Value: channelMention,
			},
		},
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
