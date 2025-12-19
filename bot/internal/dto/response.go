package dto

type SetWelcomeChannelResponse struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}

type WelcomeMessageResponse struct {
	GuildID string `json:"guild_id"`
	Message string `json:"message"`
}

type RoleMessage struct {
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
}

type Role struct {
	GuildID string `json:"guild_id"`
	RoleID  string `json:"role_id"`
	Emoji   string `json:"emoji"`
}

type RolesSettings struct {
	MessageID string            `json:"message_id"`
	Matching  map[string]string `json:"matching"`
}

type WelcomeSettings struct {
	ChannelID string   `json:"channel_id"`
	Messages  []string `json:"messages"`
}

type AutoModeSettings struct {
	Enabled     bool         `json:"enabled"`
	GuildID     string       `json:"guild_id"`
	BannedWords []BannedWord `json:"banned_words"`
	AntiLink    []AntiLink   `json:"anti_link"`
	CapsLock    []CapsLock   `json:"caps_lock"`
}

type AntiLink struct {
	ChannelId string `json:"channel_id"`
	Id        string `json:"id"`
}

type CapsLock struct {
	ChannelId string `json:"channel_id"`
	Id        string `json:"id"`
}

type BannedWord struct {
	Word string `json:"word"`
	Id   string `json:"id"`
}

type LogSettings struct {
	GuildID   string `json:"guild_id"`
	Events    []EventSettings `json:"event"`
	Enabled   bool   `json:"enabled"`
}

type EventSettings struct {
	EventType int32 `json:"event_type"`
	ChannelID string `json:"channel_id"`
}

type GuildSettings struct {
	ID       string           `json:"id"`
	GuildID  string           `json:"guild_id"`
	Roles    RolesSettings    `json:"roles"`
	Welcome  WelcomeSettings  `json:"welcome"`
	AutoMode AutoModeSettings `json:"automode"`
	Log      LogSettings      `json:"log"`
}

type GuildSettingsResponse struct {
	Settings GuildSettings `json:"settings"`
}

type User struct {
	UserID        string `json:"user_id"`
	GuildID       string `json:"guild_id"`
	Experience    int32  `json:"experience"`
	Level         int32  `json:"level"`
	NextLevelXP   int32  `json:"next_level_xp"`
	LastMessageAt string `json:"last_message_at"` // ISO timestamp string
}

type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
}

type AddXpResponse struct {
	AddedXp int  `json:"added_xp"`
	LevelUp bool `json:"level_up"`
	User    User `json:"user"`
}
