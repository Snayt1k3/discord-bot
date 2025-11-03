package guild

import (
	"bot/internal/adapters/guild"
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ViewInteractionProfile(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildID := i.GuildID
	user, err := guildService.Interaction.GetUser(guildID, i.User.ID)
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
				Value:  formatLastMessage(user.LastMessageAt),
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

func formatLastMessage(ts string) string {
	if ts == "" {
		return "No messages yet"
	}
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return ts // fallback — показать как есть
	}
	return t.Format("02 Jan 2006 15:04")
}
