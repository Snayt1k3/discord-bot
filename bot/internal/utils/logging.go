package utils

import (
	"bot/internal/dto"
	"bot/internal/http"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)


func SendLoggingMessage(
	http *http.Container,
	s *discordgo.Session,
	event dto.EventType,
	guildId string,
	embed *discordgo.MessageEmbed,
) error {
	settings, err := http.Settings.Get(guildId)

	if err != nil {
		slog.Error("Error while getting settings", "error", err)
		return err
	}

	if !settings.Log.Enabled {
		return nil
	}

	for _, ev := range settings.Log.Events {
		if int(event) == int(ev.EventType) {
			_, err = s.ChannelMessageSendEmbed(ev.ChannelID, embed)
			if err != nil {
				slog.Error("Error while sending log message", "error", err)
			}
		}

	}

	return nil
}