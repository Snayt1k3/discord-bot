package settings

import (
	"bot/internal/discord"
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

	roles := settings.Roles
	if len(roles.Matching) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ **No roles are configured for this server.**",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

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

	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "➕ Add Role",
				CustomID: "role_add",
				Style:    discordgo.PrimaryButton,
			},
			discordgo.Button{
				Label:    "➖ Remove Role",
				CustomID: "role_remove",
				Style:    discordgo.DangerButton,
			},
			discordgo.Button{
				Label:    "📝 Set Message ID",
				CustomID: "set_message_id",
				Style:    discordgo.SecondaryButton,
			},
		},
	}

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content:    &roleListStr,
		Components: &[]discordgo.MessageComponent{actionRow},
	})

	if err != nil {
		slog.Error("Failed to edit interaction response", "err", err)
	}
}

func AddRole(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	return
}

func RemoveRole(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	return
}

func SetMessageId(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate) {
	return
}
