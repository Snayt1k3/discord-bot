package main

import (
	"bot/internal/http"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"bot/config"
	"bot/internal/discord"
	"bot/internal/handlers"

	"github.com/bwmarrin/discordgo"
)

func main() {
	initLogging()
	config.Load()
	discord.InitDiscordBot()
	ReRegisterCommands(discord.Bot.Session, config.GetApplicationId(), discord.CommandsList)

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

func initLogging() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}

func ReRegisterCommands(s *discordgo.Session, appID string, commands []*discordgo.ApplicationCommand) error {
    // 1. Получаем список всех текущих команд
    // existing, err := s.ApplicationCommands(appID, "609869875053199366")
    // if err != nil {
    //     return fmt.Errorf("failed to get existing commands: %w", err)
    // }

    // 2. Удаляем все команды
    // for _, cmd := range existing {
    //     err := s.ApplicationCommandDelete(appID, "609869875053199366", cmd.ID)
    //     if err != nil {
    //         return fmt.Errorf("failed to delete command %s: %w", cmd.Name, err)
    //     }
    // }

    // 3. Регистрируем заново из списка (commands)
    for _, cmd := range commands {
        _, err := s.ApplicationCommandCreate(appID, "609869875053199366", cmd)
        if err != nil {
            slog.Error("failed to create command", "command", cmd.Name, "err", err)
        }
    }

    return nil
}