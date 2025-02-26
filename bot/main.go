package main

import (
	"bot/config"
	"bot/internal/discord"
	"bot/internal/handlers"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Load()
	discord.InitBot()
	discord.InitConnection()
	discord.InitLavalink()

	addHandlers()
	registerCommands()

	defer discord.Bot.Session.Close()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func registerCommands() { // todo: Убрать, на первую необходимость
	for _, command := range discord.CommandsList {
		_, err := discord.Bot.Session.ApplicationCommandCreate(discord.Bot.Session.State.User.ID, "", command)
		if err != nil {
			log.Fatalf("Failed to create command %s: %v", command.Name, err)
		}
		log.Printf("Command %s registered successfully.", command.Name)
	}
}

func addHandlers() {
	discord.Bot.Session.AddHandler(handlers.ReadyHandler)
	discord.Bot.Session.AddHandler(handlers.OnMessageReactionAdd)
	discord.Bot.Session.AddHandler(handlers.OnMessageReactionRemove)
	discord.Bot.Session.AddHandler(handlers.OnNewMemberJoin)
	discord.Bot.Session.AddHandler(handlers.CommandHandler)
	discord.Bot.Session.AddHandler(handlers.OnVoiceServerUpdate)
	discord.Bot.Session.AddHandler(handlers.OnVoiceStateUpdate)
}
