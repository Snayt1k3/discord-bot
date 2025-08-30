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

	"github.com/bwmarrin/discordgo"
)

func main() {
	initLogging()
	config.Load()
	discord.InitDiscordBot()

	// init deps
	guildKeeper := adapters.NewServiceSettingsClient()


	// init handlers/commands
	dispatcher := handlers.NewCommandsDispatcher(guildKeeper)
	eventHandlers := handlers.NewEventHandlers(guildKeeper, discord.CommandsList)
	dispatcher.InitHandlers()

	addHandlers(dispatcher, eventHandlers)

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

func addHandlers(cd *handlers.CommandsDispatcher, eh *handlers.EventHandlers) {
	discord.Bot.Session.AddHandler(cd.Dispatch)

	discord.Bot.Session.AddHandler(eh.OnVoiceServerUpdate)
	discord.Bot.Session.AddHandler(eh.OnVoiceStateUpdate)
	discord.Bot.Session.AddHandler(eh.OnMemberJoin)
	discord.Bot.Session.AddHandler(eh.OnMessageReactionAdd)
	discord.Bot.Session.AddHandler(eh.OnMessageReactionRemove)
	discord.Bot.Session.AddHandler(eh.OnGuildCreate)

}

func initBot(s *discordgo.Session, cmds []*discordgo.ApplicationCommand) {
	appID := s.State.User.ID

	for _, guild := range s.State.Guilds {
		slog.Info("Syncing commands for server",
			"server_name", guild.Name,
			"server_id", guild.ID,
		)

		// BulkOverwrite заменяет все команды сразу
		commands, err := s.ApplicationCommandBulkOverwrite(appID, guild.ID, cmds)
		if err != nil {
			slog.Error("failed to overwrite commands",
				"server_name", guild.Name,
				"error", err,
			)
			continue
		}

		slog.Info("Commands synced",
			"server_name", guild.Name,
			"count", len(commands),
		)
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
