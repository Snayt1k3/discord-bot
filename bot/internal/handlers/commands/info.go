package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func UserInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	var user *discordgo.User
	var member *discordgo.Member

	// get user from option
	if len(data.Options) > 0 {
		user = data.Options[0].UserValue(s)
		member, _ = s.GuildMember(i.GuildID, user.ID)
	} else {
		user = i.Member.User
		member = i.Member
	}

	created, _ := discordgo.SnowflakeTimestamp(user.ID)
	joined := member.JoinedAt

	// find highest role
	var topRole *discordgo.Role
	for _, roleID := range member.Roles {
		role, _ := s.State.Role(i.GuildID, roleID)

		if topRole == nil || role.Position > topRole.Position {
			topRole = role
		}
	}

	roleCount := len(member.Roles)

	// status
	status := "Offline"
	if pres, err := s.State.Presence(i.GuildID, user.ID); err == nil {
		switch pres.Status {
		case "online":
			status = "Online"
		case "idle":
			status = "Idle"
		case "dnd":
			status = "Do Not Disturb"
		}
	}

	bot := "No"
	if user.Bot {
		bot = "Yes"
	}

	embed := &discordgo.MessageEmbed{
		Title: "User Info",
		Color: 0x5865F2,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL("256"),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "👤 Username",
				Value:  fmt.Sprintf("%s#%s", user.Username, user.Discriminator),
				Inline: true,
			},
			{
				Name:   "🆔 User ID",
				Value:  user.ID,
				Inline: true,
			},
			{
				Name:   "🤖 Bot",
				Value:  bot,
				Inline: true,
			},
			{
				Name:   "📅 Account Created",
				Value:  created.Format("2 January 2006"),
				Inline: true,
			},
			{
				Name:   "📥 Joined Server",
				Value:  joined.Format("2 January 2006"),
				Inline: true,
			},
			{
				Name:   "🎭 Highest Role",
				Value:  topRole.Name,
				Inline: true,
			},
			{
				Name:   "🏷 Roles",
				Value:  fmt.Sprintf("%d", roleCount),
				Inline: true,
			},
			{
				Name:   "🟢 Status",
				Value:  status,
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Requested by %s", i.Member.User.Username),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func ServerInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, err := s.State.Guild(i.GuildID)

	if err != nil {
		guild, _ = s.Guild(i.GuildID)
	}

	// owner
	owner, _ := s.User(guild.OwnerID)

	// channels
	channels, _ := s.GuildChannels(guild.ID)

	textChannels := 0
	voiceChannels := 0

	for _, ch := range channels {
		switch ch.Type {
		case discordgo.ChannelTypeGuildText:
			textChannels++
		case discordgo.ChannelTypeGuildVoice:
			voiceChannels++
		}
	}

	// members online
	online := 0
	for _, p := range guild.Presences {
		if p.Status != discordgo.StatusOffline {
			online++
		}
	}

	// verification level
	verification := map[discordgo.VerificationLevel]string{
		discordgo.VerificationLevelNone:     "None",
		discordgo.VerificationLevelLow:      "Low",
		discordgo.VerificationLevelMedium:   "Medium",
		discordgo.VerificationLevelHigh:     "High",
		discordgo.VerificationLevelVeryHigh: "Very High",
	}[guild.VerificationLevel]

	created, _ := discordgo.SnowflakeTimestamp(guild.ID)

	embed := &discordgo.MessageEmbed{
		Title: guild.Name + " — Server Info",
		Color: 0x5865F2,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL("256"),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📛 Name",
				Value:  guild.Name,
				Inline: true,
			},
			{
				Name:   "🆔 ID",
				Value:  guild.ID,
				Inline: true,
			},
			{
				Name:   "👑 Owner",
				Value:  "<@" + owner.ID + ">",
				Inline: true,
			},
			{
				Name:   "👥 Members",
				Value:  fmt.Sprintf("%d", guild.MemberCount),
				Inline: true,
			},
			{
				Name:   "🟢 Online",
				Value:  fmt.Sprintf("%d", online),
				Inline: true,
			},
			{
				Name:   "📅 Created",
				Value:  created.Format("2 January 2006"),
				Inline: true,
			},
			{
				Name:   "💬 Text Channels",
				Value:  fmt.Sprintf("%d", textChannels),
				Inline: true,
			},
			{
				Name:   "🔊 Voice Channels",
				Value:  fmt.Sprintf("%d", voiceChannels),
				Inline: true,
			},
			{
				Name:   "🏷 Roles",
				Value:  fmt.Sprintf("%d", len(guild.Roles)),
				Inline: true,
			},
			{
				Name:   "😀 Emojis",
				Value:  fmt.Sprintf("%d", len(guild.Emojis)),
				Inline: true,
			},
			{
				Name:   "🔐 Verification Level",
				Value:  verification,
				Inline: true,
			},
			{
				Name:   "🛡 Boost Level",
				Value:  fmt.Sprintf("%d", guild.PremiumTier),
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
