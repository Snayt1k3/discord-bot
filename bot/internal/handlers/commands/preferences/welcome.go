package preferences

import (
	"bot/internal/http"
	"fmt"
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

func WelcomeSettings(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		slog.Error("Error while fetching welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	channelMention := "—"
	if settings.Welcome.ChannelID != "" {
		channelMention = "<#" + settings.Welcome.ChannelID + ">"
	}

	if len(settings.Welcome.Messages) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "📜 Welcome messages configuration",
			Description: "⚠️ No welcome messages have been configured.",
			Color:       0xED4245,
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

	messageList := strings.Join(settings.Welcome.Messages, "\n• ")
	embed := &discordgo.MessageEmbed{
		Title: "📜 Welcome messages for this server",
		Color: 0x57F287,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📍 Channel",
				Value:  channelMention,
				Inline: false,
			},
			{
				Name:  "〽️ Quota:",
				Value: fmt.Sprintf("%d/5 Messages", len(settings.Welcome.Messages)),
			},
			{
				Name:  "✉️ Messages",
				Value: "• " + messageList,
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

func SetWelcomeChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	channelId := i.ApplicationCommandData().Options[0].ChannelValue(nil).ID

	err := http.Welcome.SetChannel(i.GuildID, channelId)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel set successfully! New members will be welcomed in this channel.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddWelcomeMessage(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	msg := i.ApplicationCommandData().Options[0].StringValue()

	err := http.Welcome.AddMessage(i.GuildID, msg)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func DeleteWelcomeMessage(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	msg := i.ApplicationCommandData().Options[0].StringValue()

	err := http.Welcome.DeleteMessage(i.GuildID, msg)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message removed successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}
