package main

import (
	"bot/config"
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

	dispatcher := handlers.NewCommandsDispatcher()
	addHandlers(dispatcher)
	
	registerCommands() // todo: Заменить на событие GuildCreate

	defer discord.Bot.Session.Close()

	slog.Info("Bot is running. Press Ctrl + C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func registerCommands() { // todo: Заменить на событие GuildCreate
	for _, command := range discord.CommandsList {
		_, err := discord.Bot.Session.ApplicationCommandCreate(discord.Bot.Session.State.User.ID, "609869875053199366", command)
		if err != nil {
			slog.Error("Failed to create command", "name", command.Name, "error", err)
			continue
		}
		slog.Info("Command registered successfully", "name", command.Name)
	}
}

func addHandlers(cd *handlers.CommandsDispatcher) {
	discord.Bot.Session.AddHandler(handlers.ReadyHandler)
	discord.Bot.Session.AddHandler(handlers.OnVoiceServerUpdate)
	discord.Bot.Session.AddHandler(handlers.OnVoiceStateUpdate)

	discord.Bot.Session.AddHandler(cd.OnMemberJoin)
	discord.Bot.Session.AddHandler(cd.Dispatch)
	discord.Bot.Session.AddHandler(cd.OnMessageReactionAdd)
	discord.Bot.Session.AddHandler(cd.OnMessageReactionRemove)
}

func initLogging() {
	file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(file, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}
