package models

type EventType int32

const (
	EVENT_TYPE_UNSPECIFIED EventType = 0
	MESSAGE_DELETE         EventType = 1
	MESSAGE_EDIT           EventType = 2
	USER_JOIN              EventType = 3
	USER_LEAVE             EventType = 4
	CREATE_INVITE          EventType = 5
	JOIN_CHANNEL           EventType = 6
	LEAVE_CHANNEL          EventType = 7
	MOVE_CHANNEL           EventType = 8
)

type LogSettings struct {
	ID        uint   `gorm:"primaryKey"`
	GuildID   string `gorm:"not null"`
	Events    []LogEvent `gorm:"many2many:log_settings_events;constraint:OnDelete:CASCADE;"`
	Enabled   bool   `gorm:"default:true"`
}

type LogEvent struct {
	ID         uint      `gorm:"primaryKey"`
	EventType  EventType `gorm:"not null;index"`
	ChannelID string `json:"channel_id"`
}