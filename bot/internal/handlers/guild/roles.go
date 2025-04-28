package guild

import (
	"bot/internal/discord"
	"bot/internal/interfaces"

	dtoGuild "bot/internal/dto/settings"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func showAllRoles(gk interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	settings, err := gk.GetGuildSettings(i.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
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
		Color:       0x3498DB,
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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

func addRole(gk interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
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
	guildSetting, err := gk.GetGuildSettings(i.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return err
	}

	if guildSetting.Settings.Roles.Matching == nil {
		guildSetting.Settings.Roles.Matching = make(map[string]string)
	}

	guildSetting.Settings.Roles.Matching[emojiKey] = roleID

	_, err = gk.UpdateRolesSetting(i.GuildID, dtoGuild.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
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

func removeRole(gk interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
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

	guildSetting, err := gk.GetGuildSettings(i.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return err
	}

	guildSetting.Settings.Roles.Matching[emojiKey] = ""

	_, err = gk.UpdateRolesSetting(i.GuildID, dtoGuild.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:  guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
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

func setRolesMessage(gk interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	messageId := i.ApplicationCommandData().Options[0].StringValue()

	_, err := gk.UpdateRolesSetting(i.GuildID, dtoGuild.RolesSettings{
		MessageId: messageId,
		Matching:  map[string]string{},
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
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
