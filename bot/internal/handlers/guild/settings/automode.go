package settings

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
	"bot/internal/utils"
)

func AddBannedWord(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	word := i.ApplicationCommandData().Options[0].StringValue()

	err := guildService.AutoMode.AddBannedWord(i.GuildID, word)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "The word has been added to the banned list successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil

}

func RemoveBannedWord(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	word := i.ApplicationCommandData().Options[0].StringValue()

	err := guildService.AutoMode.RemoveBannedWord(i.GuildID, word)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "The word has been removed from banned list successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddAntiLinkChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := guildService.AutoMode.AddAntiLinkChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been added to Anti-Link successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func RemoveAntiLinkChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := guildService.AutoMode.RemoveAntiLinkChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been remove from Anti-Link successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddCapsLockChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := guildService.AutoMode.AddCapsLockChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been added to CapsLock successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func RemoveCapsLockChannel(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := guildService.AutoMode.RemoveCapsLockChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been removed from CapsLock successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

// Выводит сообщение в дискорд, о том какие настройки 'Automode' установлены
func ShowAutoModeSettings(guildService guild.GuildAdapter, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	guildSettings, err := guildService.Settings.Get(i.GuildID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	autoMode := guildSettings.AutoMode

	if !autoMode.Enabled {
		// Если AutoMode полностью выключен
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "🚨 **AutoMode is disabled on this server.**",
			},
		})
	}

	channelList := func(items []string) string {
		if len(items) == 0 {
			return "—"
		}
		return strings.Join(items, ", ")
	}

	mentionChannels := func(list []string) []string {
		res := make([]string, 0, len(list))
		for _, id := range list {
			res = append(res, "<#"+id+">")
		}
		return res
	}

	// Собираем ID каналов
	antiLinkChannels := make([]string, 0, len(autoMode.AntiLink))
	for _, v := range autoMode.AntiLink {
		antiLinkChannels = append(antiLinkChannels, v.ChannelId)
	}

	capsLockChannels := make([]string, 0, len(autoMode.CapsLock))
	for _, v := range autoMode.CapsLock {
		capsLockChannels = append(capsLockChannels, v.ChannelId)
	}

	bannedWords := make([]string, 0, len(autoMode.BannedWords))
	for _, v := range autoMode.BannedWords {
		bannedWords = append(bannedWords, v.Word)
	}

	// Формируем Embed
	embed := &discordgo.MessageEmbed{
		Title:       "⚙️ AutoMode Settings",
		Color:       0x2ECC71,
		Description: "Current automatic moderation settings:",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🚫 Banned Words",
				Value:  channelList(bannedWords),
				Inline: false,
			},
			{
				Name:   "🔗 Anti-Link (Channels)",
				Value:  channelList(mentionChannels(antiLinkChannels)),
				Inline: false,
			},
			{
				Name:   "🔠 CapsLock (Channels)",
				Value:  channelList(mentionChannels(capsLockChannels)),
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
