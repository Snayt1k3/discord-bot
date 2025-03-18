package handlers

import (
	"bot/config"
	"bot/internal/adapters"
	"bot/internal/discord"
	"bot/internal/handlers/settings"
	"bot/internal/interfaces"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)


type CommandsDispatcher struct {
	guildKeeper interfaces.GuildKeeperInterface
}

func NewCommandsDispatcher() *CommandsDispatcher {
	return &CommandsDispatcher{guildKeeper: adapters.NewServiceSettingsClient()}
}

func (cd *CommandsDispatcher) OnMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	settings, _ := cd.guildKeeper.GetGuildSettings(u.GuildID)
	
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomIndex := randGen.Intn(len(config.HelloMessages))
	formattedMessage := fmt.Sprintf(config.HelloMessages[randomIndex], u.Member.Nick)

	discord.SendChannelMessage(settings.Settings.Welcome.ChannelId, formattedMessage)
}

func (cd *CommandsDispatcher) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	slog.Info("Bot is ready")
}

func (cd *CommandsDispatcher) OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	settings.OnMessageReactionAdd(cd.guildKeeper, s, r)
}

func (cd *CommandsDispatcher) OnMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	settings.OnMessageReactionRemove(cd.guildKeeper, s, r)
}

func (cd *CommandsDispatcher) Dispatch(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Info("Handling interaction", "type", i.Type)

	isAdmin, err := discord.IsAdmin(s, i.GuildID, i.Member.User.ID)

	if err != nil {
		slog.Error("Error checking admin status", "err", err)
		return
	}

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		slog.Info("Handling command", "command", i.ApplicationCommandData().Name)
		switch i.ApplicationCommandData().Name {

		case "play":
			PlayCommandHandler(s, i)
		case "skip":
			SkipCommandHandler(s, i)
		case "stop":
			StopCommandHandler(s, i)
		case "help":
			HelpHandler(s, i)
		case "resume":
			ResumeHandler(s, i)
		case "settings":
			SettingsHandler(cd.guildKeeper, s, i)
		case "setup_reaction_roles", "add-role-reactions", "remove-role-reactions", "set-message-id", "set-welcome-channel":
			if !isAdmin {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "You do not have administrator rights to perform this action!",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				return
			}

			switch i.ApplicationCommandData().Name {

			case "add-role-reactions":
				settings.AddRole(cd.guildKeeper, s, i)
			case "remove-role-reactions":
				settings.RemoveRole(cd.guildKeeper, s, i)
			case "set-message-id":
				settings.SetMessageId(cd.guildKeeper, s, i)
			case "set-welcome-channel":
				settings.SetChannelId(cd.guildKeeper, s, i)
			default:
				slog.Warn("Unknown command", "command", i.ApplicationCommandData().Name)
			}
		}

	case discordgo.InteractionMessageComponent:
		slog.Info("Handling button", "custom_id", i.MessageComponentData().CustomID)

		switch i.MessageComponentData().CustomID {

		case "view_reaction_roles":
			settings.ShowAllRoles(cd.guildKeeper, s, i)
		default:
			slog.Warn("Unknown button interaction", "custom_id", i.MessageComponentData().CustomID)
		}

	default:
		slog.Warn("Unhandled interaction type", "type", i.Type)
	}
}
