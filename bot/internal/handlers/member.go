package handlers

import (
	"bot/internal/discord"
	"fmt"

	"github.com/bwmarrin/discordgo"
)


var (
	channelId = "" // адрес канала для сообщений
	messageTemplate = "Добро пожаловать, %v!" // шаблон сообшения
)

func OnNewMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	formattedMessage := fmt.Sprintf(messageTemplate, u.Member.User.Mention())
	go discord.SendChannelMessage(channelId, formattedMessage)

}