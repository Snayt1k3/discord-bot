package lavalink

import (
	"bot/config"
	"bot/internal/discord"
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type Lavalink struct {
	Client disgolink.Client
	Queue QueueManager
}


func (lv *Lavalink) onPlayerPause(player disgolink.Player, event lavalink.PlayerPauseEvent) {
	fmt.Printf("onPlayerPause: %v\n", event)
}

func (lv *Lavalink) onPlayerResume(player disgolink.Player, event lavalink.PlayerResumeEvent) {
	fmt.Printf("onPlayerResume: %v\n", event)
}

func (lv *Lavalink) onTrackStart(player disgolink.Player, event lavalink.TrackStartEvent) {
	fmt.Printf("onTrackStart: %v\n", event)
}

func (lv *Lavalink) onTrackEnd(player disgolink.Player, event lavalink.TrackEndEvent) {
	fmt.Printf("onTrackEnd: %v\n", event)

	if !event.Reason.MayStartNext() {
		return
	}

	queue := LavalinkClient.Queue.Get(event.GuildID().String())
	var (
		nextTrack lavalink.Track
		ok        bool
	)
	
	nextTrack, ok = queue.Next()

	if !ok {
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithTrack(nextTrack)); err != nil {
		slog.Error(fmt.Sprintf("Failed to play next track: %v", err))
	}
}

func (lv *Lavalink) onTrackException(player disgolink.Player, event lavalink.TrackExceptionEvent) {
	fmt.Printf("onTrackException: %v\n", event)
}

func (lv *Lavalink) onTrackStuck(player disgolink.Player, event lavalink.TrackStuckEvent) {
	fmt.Printf("onTrackStuck: %v\n", event)
}

func (lv *Lavalink) onWebSocketClosed(player disgolink.Player, event lavalink.WebSocketClosedEvent) {
	fmt.Printf("onWebSocketClosed: %v\n", event)
}

var LavalinkClient *Lavalink


func InitLavalink(){
	LavalinkClient = &Lavalink{Queue: QueueManager{queues: make(map[string]*Queue)}}

	Lclient := disgolink.New(
		snowflake.MustParse(discord.Session.State.User.ID),
		disgolink.WithListenerFunc(LavalinkClient.onPlayerPause),
		disgolink.WithListenerFunc(LavalinkClient.onPlayerResume),
		disgolink.WithListenerFunc(LavalinkClient.onTrackStart),
		disgolink.WithListenerFunc(LavalinkClient.onTrackEnd),
		disgolink.WithListenerFunc(LavalinkClient.onTrackException),
		disgolink.WithListenerFunc(LavalinkClient.onTrackStuck),
		disgolink.WithListenerFunc(LavalinkClient.onWebSocketClosed),
		
	)
	LavalinkClient.Client = Lclient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	node, err := LavalinkClient.Client.AddNode(ctx, disgolink.NodeConfig{
		Name:     config.GetLavalinkNodeName(),
		Address:  config.GetLavalinkAddr(),
		Password: config.GetLavalinkPass(),
		Secure:   config.GetLavalinkSecure(),
	})

	if err != nil {
		log.Fatal(err)
	}

	version, err := node.Version(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("node version: %s", version)

}