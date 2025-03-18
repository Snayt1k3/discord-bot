package settings

import (
	"bot/internal/discord"
	"bot/internal/dto"
	"bot/internal/interfaces"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ShowAllRoles(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	settings, err := guildKeeper.GetGuildSettings(i.GuildID)
	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
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

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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

func AddRole(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	emojiRaw := i.ApplicationCommandData().Options[1].StringValue()
	var emojiKey string
	if strings.HasPrefix(emojiRaw, "<:") && strings.HasSuffix(emojiRaw, ">") {
		parts := strings.Split(emojiRaw, ":")
		if len(parts) == 3 {
			emojiKey = strings.TrimSuffix(parts[2], ">")
		}
	} else {
		emojiKey = emojiRaw
	}

	roleID := i.ApplicationCommandData().Options[0].RoleValue(s, i.GuildID).ID
	guildSetting, err := guildKeeper.GetGuildSettings(i.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	if guildSetting.Settings.Roles.Matching == nil {
		guildSetting.Settings.Roles.Matching = make(map[string]string)
	}

	guildSetting.Settings.Roles.Matching[emojiKey] = roleID

	_, err = guildKeeper.UpdateRolesSetting(i.GuildID, dto.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func RemoveRole(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	emojiRaw := i.ApplicationCommandData().Options[1].StringValue()
	var emojiKey string

	if strings.HasPrefix(emojiRaw, "<:") && strings.HasSuffix(emojiRaw, ">") {
		parts := strings.Split(emojiRaw, ":")
		if len(parts) == 3 {
			emojiKey = strings.TrimSuffix(parts[2], ">")
		}
	} else {
		emojiKey = emojiRaw
	}

	guildSetting, err := guildKeeper.GetGuildSettings(i.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}
	guildSetting.Settings.Roles.Matching[emojiKey] = ""

	_, err = guildKeeper.UpdateRolesSetting(i.GuildID, dto.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been deleted successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func SetMessageId(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	messageId := i.ApplicationCommandData().Options[0].StringValue()

	_, err := guildKeeper.UpdateRolesSetting(i.GuildID, dto.RolesSettings{
		MessageId: messageId,
		Matching:  map[string]string{},
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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
