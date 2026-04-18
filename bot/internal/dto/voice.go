package dto

import (
	"fmt"
	"time"
)

type VoiceSession struct {
	ChannelID string    `json:"channel_id"`
	JoinedAt  time.Time `json:"joined_at"`
}

func VoiceSessionKey(guildID, userID string) string {
	return fmt.Sprintf("voice:session:%s:%s", guildID, userID)
}
