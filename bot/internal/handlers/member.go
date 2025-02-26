package handlers

import (
	"bot/internal/discord"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	channelId       = ""
	messageTemplate = "Добро пожаловать, %v!"
)

func OnNewMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	formattedMessage := fmt.Sprintf(messageTemplate, u.Member.Nick)

	go discord.SendChannelMessage(channelId, formattedMessage)

}
