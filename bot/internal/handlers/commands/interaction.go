package commands

import (
	"bot/internal/http"
	"bot/internal/utils"
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

func InteractionProfile(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildID := i.GuildID
	user, err := http.Interaction.GetUser(guildID, i.User.ID)

	if err != nil {
		slog.Error("Failed to fetch user profile", "err", err)
		return err
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("📜 Profile of %s", i.User.Username),
		Color:       0x5865F2, // Discord blurple
		Description: fmt.Sprintf("Here’s your current progress in **%s**!", guildID),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🏅 Level",
				Value:  fmt.Sprintf("%d", user.Level),
				Inline: true,
			},
			{
				Name:   "⭐ Experience",
				Value:  fmt.Sprintf("%d / %d", user.Experience, user.NextLevelXP),
				Inline: true,
			},
			{
				Name:   "🕒 Last Message",
				Value:  utils.FormatLastMessage(user.LastMessageAt),
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Username - %s", i.User.Username),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})
}

func ShowLeaderBoard(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{})
}
