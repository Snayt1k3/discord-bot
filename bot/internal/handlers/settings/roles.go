package settings

import (
	"bot/internal/discord"
	dtoDiscord "bot/internal/dto/discord"
	dtoGuild "bot/internal/dto/settings"
	"bot/internal/interfaces"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ShowAllRoles(data dtoDiscord.HandlerData) {
	settings, err := data.Gk.GetGuildSettings(data.Event.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return
	}

	roles := settings.Settings.Roles

	var roleList strings.Builder
	roleList.WriteString("📜 **Roles configured for this server:**\n\n")

	for emoji, roleID := range roles.Matching {
		emojiStr := emoji

		if _, err := strconv.ParseInt(emoji, 10, 64); err == nil {
			emojiStr = fmt.Sprintf("<:emoji:%s>", emoji)
		}

		roleList.WriteString(fmt.Sprintf("%s - (<@&%s>)\n", emojiStr, roleID))
	}

	roleListStr := roleList.String()

	err = data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: roleListStr,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		slog.Error("Failed to respond to interaction", "err", err)
	}
}

func AddRole(data dtoDiscord.HandlerData) {
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
		return
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
		return
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func RemoveRole(data dtoDiscord.HandlerData) {
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
		return
	}
	guildSetting.Settings.Roles.Matching[emojiKey] = ""

	_, err = data.Gk.UpdateRolesSetting(data.Event.GuildID, dtoGuild.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been deleted successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func SetMessageId(data dtoDiscord.HandlerData) {
	messageId := data.Event.ApplicationCommandData().Options[0].StringValue()

	_, err := data.Gk.UpdateRolesSetting(data.Event.GuildID, dtoGuild.RolesSettings{
		MessageId: messageId,
		Matching:  map[string]string{},
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(data.Session, data.Event)
		return
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message ID has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

}

func OnMessageReactionAdd(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	slog.Info("%v reacted with %v", r.UserID, r.Emoji.Name)

	guildSetting, err := guildKeeper.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Settings.Roles.MessageId {
		return
	}

	roleId, exists := guildSetting.Settings.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error adding role", "error", err)
	}

}

func OnMessageReactionRemove(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	slog.Info("%v remove reaction %v", r.UserID, r.Emoji.Name)

	guildSetting, err := guildKeeper.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Settings.Roles.MessageId {
		return
	}

	roleId, exists := guildSetting.Settings.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleRemove(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error removing role", "error", err)
	}
}
