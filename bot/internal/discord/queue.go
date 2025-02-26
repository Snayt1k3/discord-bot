package discord

import "github.com/disgoorg/disgolink/v3/lavalink"

type Queue struct {
	Tracks []lavalink.Track
}

func (q *Queue) Add(track ...lavalink.Track) {
	q.Tracks = append(q.Tracks, track...)
}

func (q *Queue) Next() (lavalink.Track, bool) {
	if len(q.Tracks) == 0 {
		return lavalink.Track{}, false
	}

	track := q.Tracks[0]
	q.Tracks = q.Tracks[1:]

	return track, true
}

func (q *Queue) Clear() {
	q.Tracks = make([]lavalink.Track, 0)
}

type QueueManager struct {
	Queues map[string]*Queue
}

func (q *QueueManager) Get(guildID string) *Queue {
	queue, ok := q.Queues[guildID]
	if !ok {
		queue = &Queue{
			Tracks: make([]lavalink.Track, 0),
		}
		q.Queues[guildID] = queue
	}
	return queue
}

func (q *QueueManager) Delete(guildID string) {
	delete(q.Queues, guildID)
}
