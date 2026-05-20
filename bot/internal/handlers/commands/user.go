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
	userID := i.Member.User.ID
	targetUser := i.Member.User

	if options := i.ApplicationCommandData().Options; len(options) > 0 {
		if u := options[0].UserValue(s); u != nil {
			userID = u.ID
			targetUser = u 
		}
	}

	user, err := http.User.GetUser(guildID, userID)
	if err != nil {
		slog.Error("Failed to fetch user profile", "err", err)
		return err
	}

	username := targetUser.Username
	avatar := discordgo.EndpointUserAvatar(targetUser.ID, targetUser.Avatar)

	if err != nil {
		slog.Error("Failed to fetch user profile", "err", err)
		return err
	}

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
			utils.FormatVoiceTime(user.VoiceTime),
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