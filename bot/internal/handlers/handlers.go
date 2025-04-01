package handlers

import (
	"bot/config"
	"bot/internal/discord"
	dto "bot/internal/dto/discord"
	er "bot/internal/errors"
	"context"
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/snowflake/v2"
	"log/slog"
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

func HelpHandler(data dto.HandlerData) error {
	helpMessage := "**🌿 Frieren Bot - Traces of Music 🌿**\n" +
		"Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:\n\n" +
		"**🎼 Commands to Guide the Melody:**\n" +
		"- `/play <song_name/link>` – Let the music flow, one song at a time.\n" +
		"- `/pause` – Even melodies need a moment of rest.\n" +
		"- `/resume` – Continue where you left off, like an old journey resumed.\n" +
		"- `/stop` – Bring the music to a quiet end, clearing all that remains.\n" +
		"- `/skip` – Move past this tune, towards the next story in sound.\n\n" +

		"**📖 Knowledge in the Wind:**\n" +
		"- `/help` – If you have forgotten, let this guide you once more.\n\n" +

		"**🌾 A Few Words of Caution:**\n" +
		"- A melody can only be heard if you are present—join a voice channel first.\n" +
		"- If questions linger, seek wisdom from those who lead this place.\n\n" +
		"Music drifts like memories in the wind. Enjoy it while it lasts. 🎧"

	data.Session.InteractionRespond(
		data.Event.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: helpMessage,
			},
		},
	)
	return nil
}

func SettingsHandler(data dto.HandlerData) {

	_, err := data.Gk.GetGuildSettings(data.Event.GuildID)

	if errors.Is(err, er.ErrGuildSettingsNotFound) {
		data.Gk.CreateSettings(data.Event.GuildID)
	}

	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "⚙️ Show all roles.",
			Style:    discordgo.SuccessButton,
			CustomID: "view_reaction_roles",
			Emoji: &discordgo.ComponentEmoji{
				Name: "🔧",
			},
		},
	}

	message := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "**⚙️ Server Settings**\n\n" +
				"Welcome to the settings panel! Here you can manage various aspects of your server.\n\n" +
				"🔹 *Click the button below to see all roles configured for this server!*\n\n" +
				"**🔧 Admin Commands:**\n" +
				"- `/add-role-reactions <role> <emoji>` – Add a role reaction.\n" +
				"- `/remove-role-reactions <role>` – Remove a role reaction.\n" +
				"- `/set-message-id <message_id>` – Set the message ID for role reactions.\n\n" +
				"*(Only administrators can use these commands.)*",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: buttons,
				},
			},
		},
	}

	err = data.Session.InteractionRespond(data.Event.Interaction, message)
	if err != nil {
		slog.Error("Error sending settings message", "err", err)
	}
}
