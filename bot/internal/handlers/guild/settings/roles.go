package settings

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
	"bot/internal/utils"
)

func showAllRoles(gk guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := gk.Settings.Get(i.GuildID)
	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	var roleList strings.Builder
	for emoji, roleID := range settings.Roles.Matching {
		emojiStr := emoji
		if _, err := strconv.ParseInt(emoji, 10, 64); err == nil {
			emojiStr = fmt.Sprintf("<:emoji:%s>", emoji)
		}
		roleList.WriteString(fmt.Sprintf("%s - <@&%s>\n", emojiStr, roleID))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📜 Roles configured for this server:",
		Description: roleList.String(),
		Color:       0x3498DB,
	}

	// отвечаем на интеракцию
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Flags:  discordgo.MessageFlagsEphemeral, // приватное сообщение
		},
	})

	if err != nil {
		slog.Error("Failed to respond to interaction", "err", err)
		return err
	}

	return nil
}

func addRole(gk guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

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

	err := gk.RolesReactions.Add(roleID, emojiKey, i.GuildID)

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func removeRole(gk guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	roleID := i.ApplicationCommandData().Options[0].RoleValue(s, i.GuildID).ID

	err := gk.RolesReactions.Delete(roleID, "", i.GuildID)

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been deleted successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func setRolesMessage(gk guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	messageId := i.ApplicationCommandData().Options[0].StringValue()

	err := gk.RolesReactions.SetMessageID(messageId, i.GuildID)

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message ID has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}
