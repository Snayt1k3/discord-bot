package commands

import (
	"bot/internal/dto"
	"bot/internal/http"
	"bot/internal/utils"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func RolesSettings(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		slog.Error("Error while getting http settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	if len(settings.Roles.Matching) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "📜 Roles configured for this server:",
			Description: "⚠️ No roles configured.",
			Color:       0xED4245, // red
		}
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}

	var roleList strings.Builder

	for emoji, roleID := range settings.Roles.Matching {
		emojiStr := emoji
		if _, err := strconv.ParseInt(emoji, 10, 64); err == nil {
			emojiStr = fmt.Sprintf("<:emoji:%s>", emoji)
		}
		roleList.WriteString(fmt.Sprintf("%s - <@&%s>\n", emojiStr, roleID))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📜 Roles configured for this server",
		Description: roleList.String(),
		Color:       0x57F287, // green
		Footer: &discordgo.MessageEmbedFooter{
			Text: "React with the corresponding emoji to get the role",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	return nil
}

func WelcomeSettings(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		slog.Error("Error while fetching welcome settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	channelMention := "—"
	if settings.Welcome.ChannelID != "" {
		channelMention = "<#" + settings.Welcome.ChannelID + ">"
	}

	if len(settings.Welcome.Messages) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "📜 Welcome messages configuration",
			Description: "⚠️ No welcome messages have been configured.",
			Color:       0xED4245,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "📍 Channel",
					Value: channelMention,
				},
			},
		}
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}

	messageList := strings.Join(settings.Welcome.Messages, "\n• ")
	embed := &discordgo.MessageEmbed{
		Title: "📜 Welcome messages for this server",
		Color: 0x57F287,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📍 Channel",
				Value:  channelMention,
				Inline: false,
			},
			{
				Name:  "〽️ Quota:",
				Value: fmt.Sprintf("%d/5 Messages", len(settings.Welcome.Messages)),
			},
			{
				Name:  "✉️ Messages",
				Value: "• " + messageList,
			},
		},
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func ModerationSettings(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
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

func LoggingSettings(
	http *http.Container,
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) error {

	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		slog.Error("Error while fetching log settings", "err", err)
		utils.SendErrorMessage(s, i)
		return err
	}

	logSettings := settings.Log


	statusText := "🔴 Disabled"
	color := 0xED4245

	if logSettings.Enabled {
		statusText = "🟢 Enabled"
		color = 0x57F287
	}

	channel := "—"
	eventsValue := "—"

	if len(logSettings.Events) > 0 {
		lines := make([]string, 0, len(logSettings.Events))
		groupedEvents := groupEvents(logSettings.Events)
		
		for channel, events := range groupedEvents {
			channel = "<#" + channel + ">"
			lines = append(lines,
				"• **"+strings.Join(events, ", ")+"** → "+channel,
			)
		}

		eventsValue = strings.Join(lines, "\n")
	}

	embed := &discordgo.MessageEmbed{
		Title: "📜 Logging Events Configuration",
		Color: color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📌 Status",
				Value:  statusText,
				Inline: true,
			},
			{
				Name:   "📍 Default Channel",
				Value:  channel,
				Inline: true,
			},
			{
				Name:  "🧾 Events",
				Value: eventsValue,
			},
		},
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	return nil
}

// groupEvents groups event settings by their channel IDs.
func groupEvents(events []dto.EventSettings) map[string][]string {
	grouped := make(map[string][]string)

	for _, event := range events {
		channelID := event.ChannelID
		grouped[channelID] = append(grouped[channelID], dto.EventType(event.EventType).String())
	}
	return grouped
}