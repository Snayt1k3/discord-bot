package v1

import (
	"api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(r *gin.Engine, handler *handlers.SettingsHandlers) {
	settings := r.Group("/api/v1/settings")

	settings.GET("/guild/:guild_id", handler.GetGuildSettings)
	settings.POST("/guild/:guild_id", handler.CreateGuildSetting)
	settings.PATCH("/guild/:guild_id/roles", handler.UpdateRoles)
	settings.PATCH("/guild/:guild_id/welcome", handler.UpdateWelcome)

}
