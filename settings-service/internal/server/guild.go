package server

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"settings-service/internal/interfaces"
	pb "settings-service/proto"
)

type GuildServer struct {
	Repo interfaces.GuildSettingsRepository
	pb.UnimplementedSettingsServiceServer
}

func (s *GuildServer) GetSettings(ctx context.Context, req *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	guildSettings, err := s.Repo.GetGuildSettings(req.GuildId)

	if err != nil {
		return nil, err
	}

	if guildSettings.ID == 0 {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

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
			Id:      strconv.FormatUint(uint64(bw.ID), 10),
			Word:    bw.Word,
			GuildId: bw.GuildID,
		})
	}

	for _, c := range guildSettings.AutoMode.CapsLocks {
		capsLocks = append(capsLocks, &pb.CapsLock{
			Id:        strconv.FormatUint(uint64(c.ID), 10),
			ChannelId: c.ChannelID,
			GuildId:   c.GuildID,
		})
	}

	for _, a := range guildSettings.AutoMode.AntiLinks {
		antiLinks = append(antiLinks, &pb.AntiLink{
			Id:        strconv.FormatUint(uint64(a.ID), 10),
			ChannelId: a.ChannelID,
			GuildId:   a.GuildID,
		})
	}

	response := &pb.GetSettingsResponse{
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
	}

	return response, nil
}

func (s *GuildServer) CreateSettings(ctx context.Context, req *pb.CreateSettingsRequest) (*pb.CreateSettingsResponse, error) {

	guildSettings, _ := s.Repo.GetGuildSettings(req.GuildId)

	if guildSettings != nil && guildSettings.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, "Guild settings already exist")
	}

	settings, err := s.Repo.CreateGuildSetting(req.GuildId)

	if err != nil {
		return nil, err
	}
	
	matching := make(map[string]string)
	var welcomeMessages []string
	var bannedWords []*pb.BannedWord
	var capsLocks []*pb.CapsLock
	var antiLinks []*pb.AntiLink

	for _, role := range settings.Role.Roles {
		matching[role.Emoji] = role.RoleID
	}

	for _, msg := range settings.Welcome.Messages {
		welcomeMessages = append(welcomeMessages, msg.Message)
	}

	for _, bw := range settings.AutoMode.BannedWords {
		bannedWords = append(bannedWords, &pb.BannedWord{
			Id:      strconv.FormatUint(uint64(bw.ID), 10),
			Word:    bw.Word,
			GuildId: bw.GuildID,
		})
	}

	for _, c := range settings.AutoMode.CapsLocks {
		capsLocks = append(capsLocks, &pb.CapsLock{
			Id:        strconv.FormatUint(uint64(c.ID), 10),
			ChannelId: c.ChannelID,
			GuildId:   c.GuildID,
		})
	}

	for _, a := range settings.AutoMode.AntiLinks {
		antiLinks = append(antiLinks, &pb.AntiLink{
			Id:        strconv.FormatUint(uint64(a.ID), 10),
			ChannelId: a.ChannelID,
			GuildId:   a.GuildID,
		})
	}

	return &pb.CreateSettingsResponse{
		GuildId: settings.GuildID,
		Roles: &pb.RolesSettings{
			MessageId: settings.Role.MessageID,
			Matching:  matching,
		},

		Welcome: &pb.WelcomeSettings{
			ChannelId: settings.Welcome.ChannelId,
			Messages:  welcomeMessages,
		},
		Log: &pb.LogSettings{
			ChannelId: settings.Log.ChannelID,
			Id:        strconv.FormatUint(uint64(settings.Log.ID), 10),
			GuildId:   settings.GuildID,
			Enabled:   settings.Log.Enabled,
		},
		Automod: &pb.AutoModSettings{
			Id:          strconv.FormatUint(uint64(settings.AutoMode.ID), 10),
			GuildId:     settings.GuildID,
			Enabled:     settings.AutoMode.Enabled,
			BannedWords: bannedWords,
			CapsLock:    capsLocks,
			AntiLink:    antiLinks,
		},
	}, nil 
}
