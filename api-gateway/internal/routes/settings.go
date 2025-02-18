package routes

import (
	"github.com/gin-gonic/gin"
	"api-gateway/internal/handlers"
)

func SettingsRoutes(r *gin.Engine, handler *handlers.SettingsHandlers) {
	settings := r.Group("/settings")
	{
		settings.GET("/guild/:guild_id", handler.GetGuildSettings)
		settings.GET("/guilds", handler.GetAllGuildsSettings)
		settings.PATCH("/guild", handler.UpdateGuildSettings)
	}
}
