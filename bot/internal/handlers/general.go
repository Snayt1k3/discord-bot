package handlers

import (
	"bot/internal/adapters"
	"bot/internal/discord"
	"bot/internal/handlers/settings"
	"bot/internal/interfaces"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"math/rand"
	"time"
)

var (
	// Temporary solution
	channelId       = ""
	messageTemplate = []string{
		"Еще один путешественник прибыл… %v, время на этом сервере течет иначе, чем в мире снаружи. Надеюсь, твое пребывание здесь будет долгим и спокойным.",
		"Ты присоединился, %v… Люди говорят, что время летит незаметно в хорошей компании. Возможно, ты тоже это почувствуешь.",
		"Еще один человек… %v, люди приходят и уходят, но, возможно, ты задержишься здесь дольше, чем кажется.",
		"Когда-то здесь уже были другие… Но каждое новое знакомство — это возможность узнать что-то новое. Добро пожаловать, %v.",
		"Добро пожаловать, %v. В этом месте нет конца пути, только новые встречи. Оставайся, если хочешь.",
	}
)

type CommandsDispatcher struct {
	guildKeeper interfaces.GuildKeeperInterface
}

func NewCommandsDispatcher() *CommandsDispatcher {
	return &CommandsDispatcher{guildKeeper: adapters.NewServiceSettingsClient()}
}

func (cd *CommandsDispatcher) OnMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomIndex := randGen.Intn(len(messageTemplate))
	formattedMessage := fmt.Sprintf(messageTemplate[randomIndex], u.Member.Nick)

	discord.SendChannelMessage(``, formattedMessage)
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
		case "setup_reaction_roles", "add-role-reactions", "remove-role-reactions", "set-message-id":
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

			case "setup_reaction_roles":
				settings.ShowAllRoles(cd.guildKeeper, s, i)
			case "add-role-reactions":
				settings.AddRole(cd.guildKeeper, s, i)
			case "remove-role-reactions":
				settings.RemoveRole(cd.guildKeeper, s, i)
			case "set-message-id":
				settings.SetMessageId(cd.guildKeeper, s, i)
			default:
				slog.Warn("Unknown command", "command", i.ApplicationCommandData().Name)
			}
		}

	case discordgo.InteractionMessageComponent:
		slog.Info("Handling button", "custom_id", i.MessageComponentData().CustomID)
		switch i.MessageComponentData().CustomID {
		case "setup_reaction_roles":
			settings.ShowAllRoles(cd.guildKeeper, s, i)
		default:
			slog.Warn("Unknown button interaction", "custom_id", i.MessageComponentData().CustomID)
		}

	default:
		slog.Warn("Unhandled interaction type", "type", i.Type)
	}
}
