package preferences

import (
	"bot/internal/http"
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

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
