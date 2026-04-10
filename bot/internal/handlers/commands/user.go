package commands

import (
	"bot/internal/http"
	"bot/internal/utils"

	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

func UserStats(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildID := i.GuildID
	requestedUser := i.ApplicationCommandData().Options[0].UserValue(s)
	userID := ""

	if requestedUser == nil {
		userID = i.Member.User.ID
	} else {
		userID = requestedUser.ID
	}

	user, err := http.Interaction.GetUser(guildID, userID)

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
			"**🌟 Server Level %d**\n%s\n**%d / %d XP**\n\n🎙️ **Voice Time** — %s",
			level,
			utils.ProgressBlocks(int(curXP), int(nextXP), 10),
			curXP, nextXP,
			utils.FormatVoiceTime(user.SecondsInVoice),
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