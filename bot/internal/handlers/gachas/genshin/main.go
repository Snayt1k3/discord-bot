package genshin

import (
	dtoDiscord "bot/internal/dto/discord"
)

func AddGenshinHandlers(handlers map[string]func(dtoDiscord.HandlerData) error) {}