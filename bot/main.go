package main

import (
	"bot/internal/http"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"bot/config"
	"bot/internal/discord"
	"bot/internal/handlers"
)

func main() {
	initLogging()
	config.Load()
	discord.InitDiscordBot()

	// init deps
	httpContainer := http.NewContainer()

	// init handlers/commands
	dispatcher := handlers.NewCommandsDispatcher()
	eventHandlers := handlers.NewEventHandlers(httpContainer, discord.CommandsList)
	handlersContainer := handlers.NewContainer(httpContainer)
	dispatcher.InitHandlers(handlersContainer)
	addEventHandlers(dispatcher, eventHandlers)

	if err := discord.Bot.Session.UpdateCustomStatus(config.GetBotStatus()); err != nil {
		slog.Warn("failed to update custom status", "error", err)
	}

	initBot(discord.Bot.Session, discord.CommandsList)

	defer discord.Bot.Session.Close()

	slog.Info("Bot is running. Press Ctrl + C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func addEventHandlers(cd *handlers.CommandsDispatcher, eh *handlers.EventHandlers) {
	discord.Bot.Session.AddHandler(cd.Dispatch)
	discord.Bot.Session.AddHandler(eh.OnMemberJoin)
	discord.Bot.Session.AddHandler(eh.OnMessageReactionAdd)
	discord.Bot.Session.AddHandler(eh.OnMessageReactionRemove)
	discord.Bot.Session.AddHandler(eh.OnGuildCreate)
	discord.Bot.Session.AddHandler(eh.MessageCreate)
	discord.Bot.Session.AddHandler(eh.GuildBanAdd)
	discord.Bot.Session.AddHandler(eh.GuildBanRemove)
	discord.Bot.Session.AddHandler(eh.GuildMemberRemove)
	discord.Bot.Session.AddHandler(eh.MessageDelete)
	discord.Bot.Session.AddHandler(eh.MessageDeleteBulk)
	discord.Bot.Session.AddHandler(eh.OnInviteCreate)

}

func initBot(s *discordgo.Session, cmds []*discordgo.ApplicationCommand) {
	appID := s.State.User.ID

	for _, guild := range s.State.Guilds {
		// Получаем текущие команды
		// oldCommands, _ := s.ApplicationCommands(appID, http.ID)
		// for _, cmd := range oldCommands {
		// 	_ = s.ApplicationCommandDelete(appID, http.ID, cmd.ID)
		// }

		// slog.Info("Old commands deleted, registering new ones...",
		// 	"server_name", http.Name,
		// 	"server_id", http.ID,
		// )

		for _, cmd := range cmds {
			_, err := s.ApplicationCommandCreate(appID, guild.ID, cmd)
			if err != nil {
				slog.Error("Error creating command", "command", cmd.Name, "error", err)
			}
		}
	}
}

func initLogging() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}
