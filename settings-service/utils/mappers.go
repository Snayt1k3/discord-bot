package utils

import (
	"settings-service/internal/models"
	pb "settings-service/proto"
	"strconv"
)

// MapGuildSettings | Mapping model to a pb.GetSettingsResponse
func MapGuildSettings(guildSettings *models.Settings) (*pb.GetSettingsResponse, error) {
	matching := make(map[string]string)

	var welcomeMessages []string
	var bannedWords []*pb.BannedWord
	var capsLocks []*pb.CapsLock
	var antiLinks []*pb.AntiLink

	for _, role := range guildSettings.Role.Roles {
		matching[role.Emoji] = role.RoleID
	}

	for _, msg := range guildSettings.Welcome.Messages {
		welcomeMessages = append(welcomeMessages, msg.Message)
	}

	for _, bw := range guildSettings.AutoMode.BannedWords {
		bannedWords = append(bannedWords, &pb.BannedWord{
			Id:   strconv.FormatUint(uint64(bw.ID), 10),
			Word: bw.Word,
		})
	}

	for _, c := range guildSettings.AutoMode.CapsLocks {
		capsLocks = append(capsLocks, &pb.CapsLock{
			Id:        strconv.FormatUint(uint64(c.ID), 10),
			ChannelId: c.ChannelID,
		})
	}

	for _, a := range guildSettings.AutoMode.AntiLinks {
		antiLinks = append(antiLinks, &pb.AntiLink{
			Id:        strconv.FormatUint(uint64(a.ID), 10),
			ChannelId: a.ChannelID,
		})
	}

	return &pb.GetSettingsResponse{
		GuildId: guildSettings.GuildID,
		Roles: &pb.RolesSettings{
			MessageId: guildSettings.Role.MessageID,
			Matching:  matching,
		},

		Welcome: &pb.WelcomeSettings{
			ChannelId: guildSettings.Welcome.ChannelId,
			Messages:  welcomeMessages,
		},
		Log: &pb.LogSettings{
			ChannelId: guildSettings.Log.ChannelID,
			Id:        strconv.FormatUint(uint64(guildSettings.Log.ID), 10),
			GuildId:   guildSettings.GuildID,
			Enabled:   guildSettings.Log.Enabled,
		},
		Automode: &pb.AutoModSettings{
			Id:          strconv.FormatUint(uint64(guildSettings.AutoMode.ID), 10),
			GuildId:     guildSettings.GuildID,
			Enabled:     guildSettings.AutoMode.Enabled,
			BannedWords: bannedWords,
			CapsLock:    capsLocks,
			AntiLink:    antiLinks,
		},
	}, nil
}

// MapCreateGuildSettings | Mapping model to a pb.CreateSettingsResponse
func MapCreateGuildSettings(guildSettings *models.Settings) (*pb.CreateSettingsResponse, error) {
	matching := make(map[string]string)

	var welcomeMessages []string
	var bannedWords []*pb.BannedWord
	var capsLocks []*pb.CapsLock
	var antiLinks []*pb.AntiLink

	for _, role := range guildSettings.Role.Roles {
		matching[role.Emoji] = role.RoleID
	}

	for _, msg := range guildSettings.Welcome.Messages {
		welcomeMessages = append(welcomeMessages, msg.Message)
	}

	for _, bw := range guildSettings.AutoMode.BannedWords {
		bannedWords = append(bannedWords, &pb.BannedWord{
			Id:   strconv.FormatUint(uint64(bw.ID), 10),
			Word: bw.Word,
		})
	}

	for _, c := range guildSettings.AutoMode.CapsLocks {
		capsLocks = append(capsLocks, &pb.CapsLock{
			Id:        strconv.FormatUint(uint64(c.ID), 10),
			ChannelId: c.ChannelID,
		})
	}

	for _, a := range guildSettings.AutoMode.AntiLinks {
		antiLinks = append(antiLinks, &pb.AntiLink{
			Id:        strconv.FormatUint(uint64(a.ID), 10),
			ChannelId: a.ChannelID,
		})
	}

	return &pb.CreateSettingsResponse{
		GuildId: guildSettings.GuildID,
		Roles: &pb.RolesSettings{
			MessageId: guildSettings.Role.MessageID,
			Matching:  matching,
		},

		Welcome: &pb.WelcomeSettings{
			ChannelId: guildSettings.Welcome.ChannelId,
			Messages:  welcomeMessages,
		},
		Log: &pb.LogSettings{
			ChannelId: guildSettings.Log.ChannelID,
			Id:        strconv.FormatUint(uint64(guildSettings.Log.ID), 10),
			GuildId:   guildSettings.GuildID,
			Enabled:   guildSettings.Log.Enabled,
		},
		Automode: &pb.AutoModSettings{
			Id:          strconv.FormatUint(uint64(guildSettings.AutoMode.ID), 10),
			GuildId:     guildSettings.GuildID,
			Enabled:     guildSettings.AutoMode.Enabled,
			BannedWords: bannedWords,
			CapsLock:    capsLocks,
			AntiLink:    antiLinks,
		},
	}, nil
}
