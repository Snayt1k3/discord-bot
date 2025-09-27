package settings

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
	Word    string `json:"word"`
	Id      string `json:"id"`
}

type LogSettings struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Enabled   bool   `json:"enabled"`
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
