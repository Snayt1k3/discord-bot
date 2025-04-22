package v1

import (
	"api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func GachaRoutes(r *gin.Engine, hands *handlers.GachaHandlers) {
	gacha := r.Group("/api/v1/gacha")

	// Genshin Impact
	gacha.GET("/genshin/character", hands.GetGenshinCharacters)
	gacha.GET("/genshin/character/:character_id", hands.GetGenshinCharacterByID)
	gacha.GET("/genshin/character/build/:character_id", hands.GetGenshinCharacterBuild)

}
