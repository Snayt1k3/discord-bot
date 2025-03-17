package dto

import (
	"encoding/json"
	"fmt"
)

type RolesSettings struct {
	GuildID    string            `json:"guild_id"`
	MessageId string            `json:"message_id"`
	Matching   map[string]string `json:"matching"`
}

func (rs *RolesSettings) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("failed to scan RolesSettings, expected []byte, got %T", value)
	}

	return json.Unmarshal(bytes, rs)
}

func (rs RolesSettings) Value() (interface{}, error) {
	return json.Marshal(rs)
}
