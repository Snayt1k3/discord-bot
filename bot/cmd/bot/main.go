package main

import (
	"bot/config"

	"bot/internal/commands"
	"bot/internal/discord"
	"bot/internal/handlers"
	"bot/internal/lavalink"
	"fmt"
	"log/slog"

	"os"
	"os/signal"
	"syscall"
)


func main() {
	config.Load()

	discord.InitSession()
	addHandlers()
	discord.InitConnection()

	lavalink.InitLavalink()
	defer discord.Session.Close()
	defer lavalink.LavalinkClient.Client.Close()

	// TODO: Добавить handler для GuildAdd ивентов
	for _, cmd := range commands.CommandsList {
		_, err := discord.Session.ApplicationCommandCreate(config.GetApplicationId(), config.GetGuildId(), cmd)
		if err != nil {
			slog.Info(fmt.Sprintf("Error creating command '%s' for guild '%s': %v", cmd.Name, config.GetGuildId(), err))
		} else {
			slog.Info(fmt.Sprintf("Command '%s' registered for guild '%s'", cmd.Name, config.GetGuildId()))
		}
	}


	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<- sc
}




func addHandlers(){
	discord.Session.AddHandler(handlers.ReadyHandler)
	discord.Session.AddHandler(handlers.MessageCreateHandler)
	discord.Session.AddHandler(handlers.OnMessageReactionAdd)
	discord.Session.AddHandler(handlers.OnMessageReactionRemove)
	discord.Session.AddHandler(handlers.OnNewMemberJoin)
	discord.Session.AddHandler(commands.CommandHandler)
	discord.Session.AddHandler(handlers.OnVoiceServerUpdate)
	discord.Session.AddHandler(handlers.OnVoiceStateUpdate)
}