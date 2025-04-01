package discord

import (
	"bot/internal/interfaces"
	"github.com/bwmarrin/discordgo"
)

type HandlerData struct {
	Gk      interfaces.GuildKeeperInterface
	Session *discordgo.Session
	Event   *discordgo.InteractionCreate
}
