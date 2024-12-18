package discord

import (
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
)

// SearchGuildByChannelID search the guild ID.
func SearchGuildByChannelID(textChannelID string) (guildID string) {
	channel, _ := Session.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// SearchVoiceChannelByUserID search the voice channel id into from guild.
func SearchVoiceChannelByUserID(userID string) (voiceChannelID string) {
	for _, g := range Session.State.Guilds {
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
	_, err := Session.ChannelMessageSend(channelID, message)
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

	_, err = Session.ChannelFileSend(channelID, filename, reader)
	if err != nil {
		slog.Warn("failed to send file to channel", "channelId", channelID, "filepath", filepath, "error", err)
	}
}

func JoinVoiceChannel(guildID string, voiceChannelID string, mute bool, deafen bool) (*discordgo.VoiceConnection, error) {
	voiceConnection, err := Session.ChannelVoiceJoin(guildID, voiceChannelID, mute, deafen)
	if err != nil {
		slog.Warn("failed to join voice channel", "error", err)
	}

	return voiceConnection, err
}