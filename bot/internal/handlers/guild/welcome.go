package guild

import (
	"bot/internal/utils"
	"bot/internal/interfaces"
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func showWelcomeSettings(gk interfaces.GuildServiceInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	settings, err := gk.GetGuildSettings(i.GuildID)
	if err != nil {
		slog.Error("Error while fetching welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	// Проверяем, есть ли сообщения
	if len(settings.Welcome.Messages) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "📜 Messages configured for this server:",
			Description: "⚠️ No welcome messages configured.",
			Color:       0xFFFFFF, // белый
		}
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		})
	}

	// Собираем список сообщений
	var messageList strings.Builder
	for _, message := range settings.Welcome.Messages {
		messageList.WriteString("• ")
		messageList.WriteString(message)
		messageList.WriteString("\n")
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📜 Messages configured for this server:",
		Description: messageList.String(),
		Color:       0xFFFFFF, // белый
	}

	// Отвечаем на интеракцию
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



func setWelcomeChannel(gk interfaces.GuildServiceInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	channelId := i.ApplicationCommandData().Options[0].ChannelValue(nil).ID

	_, err := gk.SetWelcomeChannel(i.GuildID, channelId)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel set successfully! New members will be welcomed in this channel.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddWelcomeMessage(gk interfaces.GuildServiceInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	msg := i.ApplicationCommandData().Options[0].StringValue()

	_, err := gk.AddWelcomeMessage(i.GuildID, msg)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func DeleteWelcomeMessage(gk interfaces.GuildServiceInterface, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	msg := i.ApplicationCommandData().Options[0].StringValue()

	_, err := gk.DeleteWelcomeMessage(i.GuildID, msg)

	if err != nil {
		slog.Error("Error while updating welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}