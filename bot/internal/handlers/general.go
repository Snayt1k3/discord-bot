package handlers

import (
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

	discord.SendChannelMessage(channelId, formattedMessage)
}

func (cd *CommandsDispatcher) Dispatch(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Info("Handling interaction", "type", i.Type)

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
		default:
			slog.Warn("Unknown command", "command", i.ApplicationCommandData().Name)
		}

	case discordgo.InteractionMessageComponent:
		slog.Info("Handling button", "custom_id", i.MessageComponentData().CustomID)
		switch i.MessageComponentData().CustomID {
		case "setup_reaction_roles":
			settings.ShowAllRoles(cd.guildKeeper, s, i)
		case "role_add":
			settings.AddRole(cd.guildKeeper, s, i)
		case "role_remove":
			settings.RemoveRole(cd.guildKeeper, s, i)
		case "set_message_id":
			settings.SetMessageId(cd.guildKeeper, s, i)
		default:
			slog.Warn("Unknown button interaction", "custom_id", i.MessageComponentData().CustomID)
		}

	default:
		slog.Warn("Unhandled interaction type", "type", i.Type)
	}
}
