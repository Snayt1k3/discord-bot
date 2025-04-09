package main

import (
	"bot/config"
	"bot/internal/adapters"
	"bot/internal/discord"
	"bot/internal/handlers"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initLogging()
	config.Load()
	discord.InitDiscordBot()

	guildKeeper := adapters.NewServiceSettingsClient()
	dispatcher := handlers.NewCommandsDispatcher(guildKeeper)
	dispatcher.InitHandlers()

	eventHandlers := handlers.NewEventHandlers(guildKeeper, discord.CommandsList)

	addHandlers(dispatcher, eventHandlers)

	defer discord.Bot.Session.Close()

	slog.Info("Bot is running. Press Ctrl + C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func addHandlers(cd *handlers.CommandsDispatcher, eh *handlers.EventHandlers) {
	discord.Bot.Session.AddHandler(cd.Dispatch)

	discord.Bot.Session.AddHandler(eh.OnVoiceServerUpdate)
	discord.Bot.Session.AddHandler(eh.OnVoiceStateUpdate)
	discord.Bot.Session.AddHandler(eh.OnMemberJoin)
	discord.Bot.Session.AddHandler(eh.OnMessageReactionAdd)
	discord.Bot.Session.AddHandler(eh.OnMessageReactionRemove)
	discord.Bot.Session.AddHandler(eh.OnBotReady)

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
