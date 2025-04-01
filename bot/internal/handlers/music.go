package handlers

import (
	"bot/internal/discord"
	dto "bot/internal/dto/discord"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
	"log/slog"
	"regexp"
	"time"
)

var (
	urlPattern    = regexp.MustCompile("^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?")
	searchPattern = regexp.MustCompile(`^(.{2})search:(.+)`)
)

func PlayCommandHandler(data dto.HandlerData) error {
	identifier := data.Event.ApplicationCommandData().Options[0].StringValue()

	if !urlPattern.MatchString(identifier) && !searchPattern.MatchString(identifier) {
		identifier = lavalink.SearchTypeYouTube.Apply(identifier)
	}

	// Проверка голосового состояния пользователя
	voiceState, err := data.Session.State.VoiceState(data.Event.GuildID, data.Event.Member.User.ID)
	if err != nil || voiceState == nil {
		return data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "⚠️ You are not in a voice channel! 🎤",
			},
		})
	}

	// Отправка предварительного ответа
	if err := data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}); err != nil {
		return err
	}

	// Проверка Lavalink
	if discord.Bot.Lavalink == nil {
		return fmt.Errorf("Lavalink is not initialized")
	}

	player := discord.Bot.Lavalink.Player(snowflake.MustParse(data.Event.GuildID))
	if player == nil {
		return fmt.Errorf("Player not found for guild: %s", data.Event.GuildID)
	}

	queue := discord.Bot.Queues.Get(data.Event.GuildID)
	if queue == nil {
		return fmt.Errorf("Queue not initialized for guild: %s", data.Event.GuildID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toPlay *lavalink.Track

	discord.Bot.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
		func(track lavalink.Track) {
			_, _ = data.Session.InteractionResponseEdit(data.Event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Loading: [`%s`](<%s>)", track.Info.Title, *track.Info.URI)),
			})
			if player.Track() == nil {
				toPlay = &track
			} else {
				queue.Add(track)
			}
		},
		func(playlist lavalink.Playlist) {
			_, _ = data.Session.InteractionResponseEdit(data.Event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Loading playlist: `%s` with `%d` tracks", playlist.Info.Name, len(playlist.Tracks))),
			})
			if player.Track() == nil {
				toPlay = &playlist.Tracks[0]
				queue.Add(playlist.Tracks[1:]...)
			} else {
				queue.Add(playlist.Tracks...)
			}
		},
		func(tracks []lavalink.Track) {
			_, _ = data.Session.InteractionResponseEdit(data.Event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Loading: [`%s`](<%s>)", tracks[0].Info.Title, *tracks[0].Info.URI)),
			})
			if player.Track() == nil {
				toPlay = &tracks[0]
			} else {
				queue.Add(tracks[0])
			}
		},
		func() {
			_, _ = data.Session.InteractionResponseEdit(data.Event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("No matches: `%s`", identifier)),
			})
		},
		func(err error) {
			_, _ = data.Session.InteractionResponseEdit(data.Event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Error while searching: `%s`", err)),
			})
		},
	))

	if toPlay == nil {
		return nil // Ничего не загружено
	}

	if err := data.Session.ChannelVoiceJoinManual(data.Event.GuildID, voiceState.ChannelID, false, false); err != nil {
		return err
	}

	return player.Update(context.TODO(), lavalink.WithTrack(*toPlay))

}

func StopCommandHandler(data dto.HandlerData) error {
	queue := discord.Bot.Queues.Get(data.Event.GuildID)
	queue.Clear()

	voiceState, _ := data.Session.State.VoiceState(data.Event.GuildID, data.Session.State.User.ID)
	if voiceState == nil || voiceState.ChannelID == "" {
		return nil
	}

	err := data.Session.ChannelVoiceJoinManual(data.Event.GuildID, "", false, false)

	if err != nil {
		slog.Error("Error while disconnecting: ", "error", err)
	}

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Music has been stopped and the queue has been cleared! 🎵",
		},
	})
	return nil

}

func SkipCommandHandler(data dto.HandlerData) error {
	queue := discord.Bot.Queues.Get(data.Event.GuildID)
	track, exists := queue.Next()

	if !exists {
		return data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "The queue is empty. Add a new song using the /play command.",
			},
		})
	}

	player := discord.Bot.Lavalink.Player(snowflake.MustParse(data.Event.GuildID))

	if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
		slog.Error("Failed to play next track: ", "error", err)
	}
	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🎉 Successfully skipped. 🚀",
		},
	})
	discord.SendMusicEmbedMessage(track.Info.Title, *track.Info.URI, track.Info.Length.String(), *track.Info.ArtworkURL)
	return nil

}

func PauseHandler(data dto.HandlerData) error {
	player := discord.Bot.Lavalink.Player(snowflake.MustParse(data.Event.GuildID))

	if player == nil {
		return nil
	}

	if player.Paused() {
		return nil
	}

	if err := player.Update(context.TODO(), lavalink.WithPaused(false)); err != nil {
		return data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error while pausing",
			}})
	}

	return data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Player is now Paused",
		},
	})
}
func ResumeHandler(data dto.HandlerData) error {
	player := discord.Bot.Lavalink.Player(snowflake.MustParse(data.Event.GuildID))

	if player == nil {
		return nil
	}

	if !player.Paused() {
		return nil
	}

	if err := player.Update(context.TODO(), lavalink.WithPaused(true)); err != nil {
		return data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error while resuming",
			}})
	}

	return data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Player is now Resumed",
		},
	})
}
