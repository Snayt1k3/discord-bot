package settings

import (
	"bot/internal/discord"
	"bot/internal/dto"
	"bot/internal/interfaces"
	"fmt"
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ShowAllRoles(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	settings, err := guildKeeper.GetGuildSettings(i.GuildID)
	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	roles := settings.Settings.Roles

	var roleList strings.Builder
	roleList.WriteString("📜 **Roles configured for this server:**\n\n")

	for key, roleID := range roles.Matching {
		role, err := s.State.Role(i.GuildID, roleID)
		if err != nil {
			slog.Warn("Could not fetch role", "roleID", roleID, "error", err)
			continue
		}
		roleList.WriteString(fmt.Sprintf("`%s.` **%s** (<@&%s>)\n", key, role.Name, roleID))
	}

	roleListStr := roleList.String()


	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	if err != nil {
		slog.Error("Failed to defer interaction response", "err", err)
		return
	}

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content:    &roleListStr,
		Components: nil,
	})

	if err != nil {
		slog.Error("Failed to edit interaction response", "err", err)
	}
}

func AddRole(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := strings.Split(i.ApplicationCommandData().Options[1].StringValue(), ":")
	emojiName := options[len(options)-1]
	roleID := i.ApplicationCommandData().Options[0].RoleValue(s, i.GuildID).ID

	guildSetting, err := guildKeeper.GetGuildSettings(i.GuildID)
	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	if guildSetting.Settings.Roles.Matching == nil {
		guildSetting.Settings.Roles.Matching = make(map[string]string)
	}

	guildSetting.Settings.Roles.Matching[emojiName] = roleID

	_, err = guildKeeper.UpdateGuildSettings(i.GuildID, dto.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:   guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}
	
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Role has been added successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func RemoveRole(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	emojiName := i.ApplicationCommandData().Options[0].StringValue()

	guildSetting, err := guildKeeper.GetGuildSettings(i.GuildID)
	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}
	guildSetting.Settings.Roles.Matching[emojiName] = ""

	_, err = guildKeeper.UpdateGuildSettings(i.GuildID, dto.RolesSettings{
		MessageId: guildSetting.Settings.Roles.MessageId,
		Matching:   guildSetting.Settings.Roles.Matching,
	})
	
	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}
	
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message ID has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func SetMessageId(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	messageId := i.ApplicationCommandData().Options[0].StringValue()
	guildSetting, err := guildKeeper.GetGuildSettings(i.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}

	_, err = guildKeeper.UpdateGuildSettings(i.GuildID, dto.RolesSettings{
		MessageId: messageId,
		Matching:   guildSetting.Settings.Roles.Matching,
	})

	if err != nil {
		slog.Error("Error while updating guild settings", "err", err)
		discord.SendErrorMessage(s, i)
		return
	}
	
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message ID has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

}


func OnMessageReactionAdd(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	slog.Info("%v reacted with %v", r.UserID, r.Emoji.Name)
	
	guildSetting, err := guildKeeper.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Settings.Roles.MessageId {
		return
	}

	roleId, exists := guildSetting.Settings.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error adding role", "error", err)
	}

}

func OnMessageReactionRemove(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	slog.Info("%v remove reaction %v", r.UserID, r.Emoji.Name)

	guildSetting, err := guildKeeper.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Settings.Roles.MessageId {
		return
	}

	roleId, exists := guildSetting.Settings.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleRemove(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error removing role", "error", err)
	}
}