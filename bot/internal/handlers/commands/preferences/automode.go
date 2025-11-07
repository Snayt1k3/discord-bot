package preferences

import (
	"bot/internal/http"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

func ToggleModeration(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	guildSettings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	autoMode := guildSettings.AutoMode

	newStatus := !autoMode.Enabled
	err = http.Moderation.Toggle(i.GuildID, newStatus)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	statusText := "enabled"

	if !newStatus {
		statusText = "disabled"
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "AutoMode has been " + statusText + " successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func AddBannedWord(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	word := i.ApplicationCommandData().Options[0].StringValue()

	err := http.Moderation.AddBannedWord(i.GuildID, word)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "The word has been added to the banned list successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil

}

func RemoveBannedWord(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	word := i.ApplicationCommandData().Options[0].StringValue()

	err := http.Moderation.RemoveBannedWord(i.GuildID, word)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "The word has been removed from banned list successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddAntiLinkChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := http.Moderation.AddAntiLinkChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been added to Anti-Link successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func RemoveAntiLinkChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := http.Moderation.RemoveAntiLinkChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been remove from Anti-Link successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddCapsLockChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := http.Moderation.AddCapsLockChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been added to CapsLock successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func RemoveCapsLockChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	err := http.Moderation.RemoveCapsLockChannel(i.GuildID, i.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This channel has been removed from CapsLock successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AutoModeSettings(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}
	guildSettings, err := http.Settings.Get(i.GuildID)

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

	joinOrDash := func(items []string) string {
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
		Color:       0x57F287, // Discord green
		Description: "Here are the current automatic moderation settings:",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🚫 Banned Words",
				Value:  fmt.Sprintf("%s\n**Quota:** %d / 50", joinOrDash(bannedWords), len(bannedWords)),
				Inline: false,
			},
			{
				Name:   "🔗 Anti-Link (Channels)",
				Value:  fmt.Sprintf("%s\n**Quota:** %d / 5", joinOrDash(mentionChannels(antiLinkChannels)), len(antiLinkChannels)),
				Inline: false,
			},
			{
				Name:   "🔠 CapsLock (Channels)",
				Value:  fmt.Sprintf("%s\n**Quota:** %d / 5", joinOrDash(mentionChannels(capsLockChannels)), len(capsLockChannels)),
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
