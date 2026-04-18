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


func ShowLeaderBoard(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) {
	page := 0
	users, err := http.User.GetUsers(i.GuildID, fmt.Sprintf("%d", page), "10", "experience", true)

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
	users, err := http.User.GetUsers(i.GuildID, fmt.Sprintf("%d", page), "10", "experience", true)

	if err != nil {
		slog.Error("Failed to fetch users for leaderboard", "err", err)
		utils.SendErrorMessage(s, i)

	}
	embed := createLeaderboardEmbed(users, page)
	keyboard := buttons.LeaderboardButtons(page, users.TotalCount%10)
	utils.Acknowledge(s, i)

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
		progress := utils.ProgressBlocks(int(u.Experience), int(u.NextLevelXP), 10)

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
		Title:       "🏆 Leaderboard",
		Color:       0x87CEEB,
		Description: strings.Join(entries, "\n\n"),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	return embed
}


