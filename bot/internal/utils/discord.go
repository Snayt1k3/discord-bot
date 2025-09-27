package utils

import (
	"bot/internal/discord"
	"log/slog"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SearchGuildByChannelID(textChannelID string) (guildID string) {
	channel, _ := discord.Bot.Session.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

func SendChannelMessage(channelID string, message string) {
	_, err := discord.Bot.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		slog.Error("failed to send message to channel", "channelId", channelID, "message", message, "error", err)
	}
}

func SendChannelFile(channelID string, filepath string, filename string) {
	reader, err := os.Open(filepath)

	if err != nil {
		slog.Warn("failed to open file", "filepath", filepath, "error", err)
		return
	}

	_, err = discord.Bot.Session.ChannelFileSend(channelID, filename, reader)

	if err != nil {
		slog.Warn("failed to send file to channel", "channelId", channelID, "filepath", filepath, "error", err)
	}
}

func IsAdmin(session *discordgo.Session, guildID, userID string) bool {
	member, err := session.GuildMember(guildID, userID)

	if err != nil {
		return false
	}

	guild, err := session.State.Guild(guildID)
	if err != nil {
		return false
	}

	for _, roleID := range member.Roles {
		for _, role := range guild.Roles {
			if role.ID == roleID && (role.Permissions&discordgo.PermissionAdministrator) != 0 {
				return true
			}
		}
	}

	return false
}

func SendErrorMessage(session *discordgo.Session, i *discordgo.InteractionCreate) error {
	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "⚠️ Oops! Something went wrong. Please try again. ⚠️",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func SendNoPermissionMessage(session *discordgo.Session, i *discordgo.InteractionCreate) error {
	return session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "⛔ You do not have permission to use this command/button.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func EditMessage(session *discordgo.Session, message *discordgo.MessageEdit) {
	_, err := session.ChannelMessageEditComplex(message)
	if err != nil {
		slog.Error("failed to edit message", "error", err)
	}
}

// SendTempMessage sends a message and deletes it after 5 seconds
func SendTempMessage(s *discordgo.Session, channelID, content string) {
	msg, err := s.ChannelMessageSend(channelID, content)
	if err != nil {
		slog.Error("failed to send temp message", "error", err)
		return
	}

	// Удаляем сообщение через 5 секунд (не блокируем основной поток)
	time.AfterFunc(5*time.Second, func() {
		err := s.ChannelMessageDelete(channelID, msg.ID)
		if err != nil {
			slog.Error("failed to delete temp message", "error", err)
		}
	})
}
