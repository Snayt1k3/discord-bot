package routes

import (
	"api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(r *gin.Engine, handler *handlers.SettingsHandlers) {
	settings := r.Group("/settings")
	{
		settings.GET("/guild/:guild_id", handler.GetGuildSettings)
		settings.POST("/guild/:guild_id", handler.CreateGuildSetting)
		settings.PATCH("/guild/:guild_id", handler.UpdateGuildSettings)
	}
}
