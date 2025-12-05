package commands

import (
	"bot/internal/http"

	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Rank(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildID := i.GuildID

	user, err := http.Interaction.GetUser(guildID, i.Member.User.ID)
	if err != nil {
		slog.Error("Failed to fetch user profile", "err", err)
		return err
	}

	username := i.Member.User.Username
	avatar := discordgo.EndpointUserAvatar(i.Member.User.ID, i.Member.User.Avatar)

	level := user.Level
	curXP := user.Experience
	nextXP := user.NextLevelXP

	embed := &discordgo.MessageEmbed{
		Color: 0x5865F2,
		Title: fmt.Sprintf("%s's Profile", username),

		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatar,
		},

		Description: fmt.Sprintf(
			"**🌟 Level %d**\n%s\n**%d / %d XP**",
			level,
			progressBlocks(int(curXP), int(nextXP), 10),
			curXP, nextXP,
		),

		Timestamp: time.Now().Format(time.RFC3339),
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

// ▰▱ progress bar builder
func progressBlocks(current, max, length int) string {
	if max <= 0 {
		max = 1
	}

	ratio := float64(current) / float64(max)
	filled := int(ratio * float64(length))

	if filled > length {
		filled = length
	}

	return strings.Repeat("▰", filled) + strings.Repeat("▱", length-filled)
}

func ShowLeaderBoard(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{})
}
