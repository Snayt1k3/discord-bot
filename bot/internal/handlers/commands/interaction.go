package commands

import (
	"bot/internal/dto"
	"bot/internal/handlers/buttons"
	"bot/internal/http"
	"bot/internal/utils"

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
			"**🌟 Server Level %d**\n%s\n**%d / %d XP**",
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

func ShowLeaderBoard(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) {
	page := 0
	users, err := http.Interaction.GetUsers(i.GuildID, fmt.Sprintf("%d", page), "10")

	if err != nil {
		slog.Error("Failed to fetch users for leaderboard", "err", err)
		utils.SendErrorMessage(s, i)

	}
	embed := createLeaderboardEmbed(users, page)
	keyboard := buttons.LeaderboardButtons(page, users.TotalCount%10)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: keyboard,
				},
			},
		},
	})
}

func ShowLeaderBoardPaginate(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate, page int) {
	users, err := http.Interaction.GetUsers(i.GuildID, fmt.Sprintf("%d", page), "10")

	if err != nil {
		slog.Error("Failed to fetch users for leaderboard", "err", err)
		utils.SendErrorMessage(s, i)

	}
	embed := createLeaderboardEmbed(users, page)
	keyboard := buttons.LeaderboardButtons(page, users.TotalCount%10)

	msg := &discordgo.MessageEdit{
		ID:      i.Message.ID,
		Channel: i.Message.ChannelID,
		Embeds:  &[]*discordgo.MessageEmbed{embed},
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: keyboard,
			},
		},
	}

	_, err = s.ChannelMessageEditComplex(msg)

	if err != nil {
		slog.Error("Failed to edit leaderboard message", "err", err)
	}
}

func createLeaderboardEmbed(users dto.UsersResponse, page int) *discordgo.MessageEmbed {
	entries := make([]string, 0, len(users.Users))

	for idx, u := range users.Users {
		progress := progressBlocks(int(u.Experience), int(u.NextLevelXP), 10)

		entry := fmt.Sprintf(
			"**%d. Level %d • <@%s>**%s`%d / %d XP`",
			(idx+1)+(page*10),
			u.Level,
			u.UserID,
			progress,
			u.Experience, u.NextLevelXP,
		)

		entries = append(entries, entry)
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🏆 Leaderboard — Page %d", page),
		Color:       0x87CEEB,
		Description: strings.Join(entries, "\n\n"),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	return embed
}

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
