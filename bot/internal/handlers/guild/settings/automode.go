package settings

import (
	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
)

func AddBannedWord(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}

func RemoveBannedWord(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}

func AddAntiLinkChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}

func RemoveAntiLinkChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}

func AddCapsLockChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}

func RemoveCapsLockChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}
