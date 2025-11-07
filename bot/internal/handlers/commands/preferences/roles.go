package preferences

import (
	"bot/internal/http"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

func RolesSettings(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)
	if err != nil {
		slog.Error("Error while getting http settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	// Check if there are roles configured
	if len(settings.Roles.Matching) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "📜 Roles configured for this server:",
			Description: "⚠️ No roles configured.",
			Color:       0xED4245, // red
		}
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}

	// Build role list
	var roleList strings.Builder
	for emoji, roleID := range settings.Roles.Matching {
		emojiStr := emoji
		// If the emoji is numeric, assume it's a custom emoji ID
		if _, err := strconv.ParseInt(emoji, 10, 64); err == nil {
			emojiStr = fmt.Sprintf("<:emoji:%s>", emoji)
		}
		roleList.WriteString(fmt.Sprintf("%s - <@&%s>\n", emojiStr, roleID))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📜 Roles configured for this server",
		Description: roleList.String(),
		Color:       0x57F287, // green
		Footer: &discordgo.MessageEmbedFooter{
			Text: "React with the corresponding emoji to get the role",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Respond to the interaction
	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	return nil
}

func AddRole(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
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

	err := http.Roles.Add(roleID, emojiKey, i.GuildID)

	if err != nil {
		slog.Error("Error while updating http settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func RemoveRole(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	roleID := i.ApplicationCommandData().Options[0].RoleValue(s, i.GuildID).ID

	err := http.Roles.Delete(roleID, "", i.GuildID)

	if err != nil {
		slog.Error("Error while updating http settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been deleted successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func SetRolesMsg(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	messageId := i.ApplicationCommandData().Options[0].StringValue()

	err := http.Roles.SetMessageID(messageId, i.GuildID)

	if err != nil {
		slog.Error("Error while updating http settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message ID has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}
