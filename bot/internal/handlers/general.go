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

var (
	userHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{
		"play":   PlayCommandHandler,
		"skip":   SkipCommandHandler,
		"stop":   StopCommandHandler,
		"help":   HelpHandler,
		"resume": ResumeHandler,
		
	}
	adminHandlers = map[string]func(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate){
		"add-role-reactions":   settings.AddRole,
		"remove-role-reactions": settings.RemoveRole,
		"set-message-id":        settings.SetMessageId,
		"set-welcome-channel":   settings.SetChannelId,
		"view_reaction_roles":   settings.ShowAllRoles,
		"settings": SettingsHandler,
	}
	buttons = map[string]func(guildKeeper interfaces.GuildKeeperInterface, s *discordgo.Session, i *discordgo.InteractionCreate){
		"view_reaction_roles": settings.ShowAllRoles,
	}

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
	formattedMessage := fmt.Sprintf(config.HelloMessages[randomIndex], u.Member.Mention())

	discord.SendChannelMessage(settings.Settings.Welcome.ChannelId, formattedMessage)
}

func (cd *CommandsDispatcher) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	slog.Info("Bot is ready")
	s.UpdateGameStatus(0, "Listening Music")
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

	if i.Type == discordgo.InteractionMessageComponent {
		if handler, ok := buttons[i.MessageComponentData().CustomID]; ok {
			handler(cd.guildKeeper, s, i)
		}
		return
	}

	if isAdmin {
		if handler, ok := adminHandlers[i.ApplicationCommandData().Name]; ok {
			handler(cd.guildKeeper, s, i)
		}
	} else {
		if handler, ok := userHandlers[i.ApplicationCommandData().Name]; ok {
			if err := handler(s, i); err != nil {
				slog.Error("Error handling command", "err", err)
			}
		}
	}
}
