package dto

import (
	"encoding/json"
	"fmt"
)


type RolesSettings struct {
	MesssageId string `json:"message_id"`
	Matching map[string]string `json:"matching"`
}

type SettingsJson struct {
	Roles RolesSettings `json:"roles"`
}

type GuildSettingsDTO struct {
	ID uint `json:id`
	GuildID string `json:"guild_id"`
	Roles RolesSettings `json:"roles"`
}

type GuildSettingsUpdateDTO struct {
	GuildId string `json:"guild_id"`
	Roles RolesSettings `json:"roles"`
}

type GuildSettingsCreateDTO struct {
	GuildId string `json:"guild_id"`
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