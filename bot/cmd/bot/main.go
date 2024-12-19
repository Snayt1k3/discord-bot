package main

import (
	"bot/config"
	"bot/internal/discord"
	"bot/internal/handlers"
	"fmt"

	"os"
	"os/signal"
	"syscall"
)


func main() {
	config.Load()

	discord.InitSession()
	addHandlers()
	discord.InitConnection()

	defer discord.Session.Close()

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
}