package handlers

import (
	"bot/config"
	"bot/internal/discord"
	"bot/internal/interfaces"
	er "bot/internal/errors"
	"errors"
	"context"
	"log/slog"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/snowflake/v2"
)

func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateCustomStatus(config.GetBotStatus())
	if err != nil {
		slog.Warn("failed to update custom status", "error", err)
	}
}

func OnVoiceStateUpdate(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	if event.UserID != session.State.User.ID {
		return
	}

	var channelID *snowflake.ID
	if event.ChannelID != "" {
		id := snowflake.MustParse(event.ChannelID)
		channelID = &id
	}
	discord.Bot.Lavalink.OnVoiceStateUpdate(context.TODO(), snowflake.MustParse(event.GuildID), channelID, event.SessionID)

	if event.ChannelID == "" {
		discord.Bot.Queues.Delete(event.GuildID)
	}
}

func OnVoiceServerUpdate(session *discordgo.Session, event *discordgo.VoiceServerUpdate) {
	discord.Bot.Lavalink.OnVoiceServerUpdate(context.TODO(), snowflake.MustParse(event.GuildID), event.Token, event.Endpoint)
}

func HelpHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
	helpMessage := "**🎵 Frieren Bot Help Menu 🎵**\n" +
		"Hello! Here are the commands you can use:\n\n" +
		"**Main Commands:**\n" +
		"- `/play <song_name/link>` – Add a song to the queue and start playing.\n" +
		"- `/pause` – Pause the music.\n" +
		"- `/resume` – Resume playing the music.\n" +
		"- `/stop` – Stop the music and clear the queue.\n" +
		"- `/skip` – Skip the current song.\n\n" +

		"**Information:**\n" +
		"- `/help` – Show this help menu.\n\n" +

		"**Notes:**\n" +
		"- Make sure you're in a voice channel before using music commands.\n" +
		"- For questions or suggestions, contact the server administrator.\n\n" +
		"**Thank you for using me!** 🎧"

	discord.Bot.Session.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: helpMessage,
			},
		},
	)
}

func SettingsHandler(guildKeeper interfaces.GuildKeeperInterface, session *discordgo.Session, i *discordgo.InteractionCreate) {
	
	_, err := guildKeeper.GetGuildSettings(i.GuildID)

	if errors.Is(err, er.ErrGuildSettingsNotFound){
		guildKeeper.CreateSettings(i.GuildID)
	}

	isAdmin, err := discord.IsAdmin(session, i.GuildID, i.Member.User.ID)

	if err != nil {
		slog.Error("Error checking admin status", "err", err)
		return
	}

	if !isAdmin {
		session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You do not have administrator rights to perform this action!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "⚙️ Configure Reaction Roles",
			Style:    discordgo.SuccessButton,
			CustomID: "setup_reaction_roles",
			Emoji: &discordgo.ComponentEmoji{
				Name: "🔧",
			},
		},
	}

	message := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "**⚙️ Server Settings**\n\n" +
				"Welcome to the settings panel! Here you can configure various aspects of your server.\n\n" +
				"🔹 *Click the button below to set up reaction roles!*",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: buttons,
				},
			},
		},
	}

	err = session.InteractionRespond(i.Interaction, message)
	if err != nil {
		slog.Error("Error sending settings message", "err", err)
	}
}
