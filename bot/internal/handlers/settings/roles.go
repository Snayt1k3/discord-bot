package settings

import (
	"bot/internal/discord"
	dtoDiscord "bot/internal/dto/discord"
	dtoGuild "bot/internal/dto/settings"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func showAllRoles(data dtoDiscord.HandlerData) error {
	settings, err := data.Gk.GetGuildSettings(data.Event.GuildID)
	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return err
	}

	roles := settings.Settings.Roles

	var roleList strings.Builder
	for emoji, roleID := range roles.Matching {
		emojiStr := emoji
		if _, err := strconv.ParseInt(emoji, 10, 64); err == nil {
			emojiStr = fmt.Sprintf("<:emoji:%s>", emoji)
		}
		roleList.WriteString(fmt.Sprintf("%s - (<@&%s>)\n", emojiStr, roleID))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📜 Roles configured for this server:",
		Description: roleList.String(),
		Color:       0x3498DB, // Soft blue color
	}

	err = data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		slog.Error("Failed to respond to interaction", "err", err)
		return err
	}

	return nil
}


func addRole(data dtoDiscord.HandlerData) error {
	emojiRaw := data.Event.ApplicationCommandData().Options[1].StringValue()
	var emojiKey string

	if strings.HasPrefix(emojiRaw, "<:") && strings.HasSuffix(emojiRaw, ">") {
		parts := strings.Split(emojiRaw, ":")
		if len(parts) == 3 {
			emojiKey = strings.TrimSuffix(parts[2], ">")
		}
	} else {
		emojiKey = emojiRaw
	}

	roleID := data.Event.ApplicationCommandData().Options[0].RoleValue(data.Session, data.Event.GuildID).ID
	guildSetting, err := data.Gk.GetGuildSettings(data.Event.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return err
	}

	if guildSetting.Settings.Roles.Matching == nil {
		guildSetting.Settings.Roles.Matching = make(map[string]string)
	}

	guildSetting.Settings.Roles.Matching[emojiKey] = roleID

	_, err = data.Gk.UpdateRolesSetting(data.Event.GuildID, dtoGuild.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return err
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func removeRole(data dtoDiscord.HandlerData) error{
	emojiRaw := data.Event.ApplicationCommandData().Options[1].StringValue()
	var emojiKey string

	if strings.HasPrefix(emojiRaw, "<:") && strings.HasSuffix(emojiRaw, ">") {
		parts := strings.Split(emojiRaw, ":")
		if len(parts) == 3 {
			emojiKey = strings.TrimSuffix(parts[2], ">")
		}
	} else {
		emojiKey = emojiRaw
	}

	guildSetting, err := data.Gk.GetGuildSettings(data.Event.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return err
	}
	guildSetting.Settings.Roles.Matching[emojiKey] = ""

	_, err = data.Gk.UpdateRolesSetting(data.Event.GuildID, dtoGuild.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return err
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been deleted successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func setMessageId(data dtoDiscord.HandlerData) error {
	messageId := data.Event.ApplicationCommandData().Options[0].StringValue()

	_, err := data.Gk.UpdateRolesSetting(data.Event.GuildID, dtoGuild.RolesSettings{
		MessageId: messageId,
		Matching:  map[string]string{},
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return err
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message ID has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

