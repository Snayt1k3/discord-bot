package genshin

import (
	dtoDiscord "bot/internal/dto/discord"
)

func AddGenshinHandlers(handlers map[string]func(dtoDiscord.HandlerData) error) {
	handlers["GenshinCharacters"] = showGenshinCharacters
	handlers["genshinPagination"] = genshinPagination
	handlers["GenshinCharacter"] = showCharacterInfo
	handlers["GenshinAsc"] = showCharacterAscension
	handlers["GenshinArtifacts"] = showCharacterArtifacts
	handlers["GenshinTalents"] = showCharacterTalents
	handlers["GenshinWeapons"] = showCharacterWeapons
	handlers["GenshinTeams"] = showCharacterComps
}