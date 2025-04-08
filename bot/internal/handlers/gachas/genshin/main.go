package genshin

import (
	dtoDiscord "bot/internal/dto/discord"
)

func AddGenshinHandlers(handlers map[string]func(dtoDiscord.HandlerData) error) {
	handlers["GenshinCharacters"] = showGenshinCharacters
	// handlers["genshinChPagination"] = func
	handlers["GenshinCharacter"] = showCharacterInfo
	handlers["GenshinAsc"] = showCharacterAscension
	handlers["GenshinArtifacts"] = showCharacterArtifacts
	handlers["GenshinWeapons"] = showCharacterWeapons
	handlers["GenshinTeams"] = showCharacterComps
}