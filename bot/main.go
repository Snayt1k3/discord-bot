package bot

import (
	"bot/config"

	"bot/internal/commands"
	"bot/internal/handlers"
	"fmt"

	"os"
	"os/signal"
	"syscall"
	"bot/internal/discord"

)

func RunBot() {
	config.Load()
	discord.InitBot()
	discord.InitLavalink()
	discord.InitConnection()
	addHandlers()
	
	defer discord.Bot.Session.Close()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<- sc
}




func addHandlers(){
	discord.Bot.Session.AddHandler(handlers.ReadyHandler)
	discord.Bot.Session.AddHandler(handlers.OnMessageReactionAdd)
	discord.Bot.Session.AddHandler(handlers.OnMessageReactionRemove)
	discord.Bot.Session.AddHandler(handlers.OnNewMemberJoin)
	discord.Bot.Session.AddHandler(commands.CommandHandler)
	discord.Bot.Session.AddHandler(handlers.OnVoiceServerUpdate)
	discord.Bot.Session.AddHandler(handlers.OnVoiceStateUpdate)
	
}