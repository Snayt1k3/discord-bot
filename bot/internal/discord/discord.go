package discord

import (
	"bot/config"
	"fmt"
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
)

// SearchGuildByChannelID search the guild ID.
func SearchGuildByChannelID(textChannelID string) (guildID string) {
	channel, _ := Bot.Session.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// SearchVoiceChannelByUserID search the voice channel id into from guild.
func SearchVoiceChannelByUserID(userID string) (voiceChannelID string) {
	for _, g := range Bot.Session.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == userID {
				return v.ChannelID
			}
		}
	}
	return ""
}

// SendChannelMessage sends a channel message to channel with channel id equal to m.ChannelID.
func SendChannelMessage(channelID string, message string) {
	_, err := Bot.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		slog.Warn("failed to send message to channel", "channelId", channelID, "message", message, "error", err)
	}
}

func SendChannelFile(channelID string, filepath string, filename string) {
	reader, err := os.Open(filepath)
	if err != nil {
		slog.Warn("failed to open file", "filepath", filepath, "error", err)
		return
	}

	_, err = Bot.Session.ChannelFileSend(channelID, filename, reader)
	if err != nil {
		slog.Warn("failed to send file to channel", "channelId", channelID, "filepath", filepath, "error", err)
	}
}

func JoinVoiceChannel(guildID string, voiceChannelID string, mute bool, deafen bool) (*discordgo.VoiceConnection, error) {
	voiceConnection, err := Bot.Session.ChannelVoiceJoin(guildID, voiceChannelID, mute, deafen)
	if err != nil {
		slog.Warn("failed to join voice channel", "error", err)
	}

	return voiceConnection, err
}

func SendMusicEmbedMessage(title string, url string, duration string, thumbnail string) {
	Bot.Session.ChannelMessageSendEmbed(config.MusicChannelId, &discordgo.MessageEmbed{
		Title:       "Now is playing 🎶",
		Description: fmt.Sprintf("[%s](%s)", title, url),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: thumbnail,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Duration",
				Value:  duration,
				Inline: true,
			},
		},
	})
}