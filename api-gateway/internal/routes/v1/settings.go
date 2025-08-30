package v1

import (
	"api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(r *gin.Engine, handler *handlers.SettingsHandlers) {
	settings := r.Group("/api/v1/settings")

	settings.GET("/guild/:guild_id", handler.GetGuildSettings)
	settings.POST("/guild/:guild_id", handler.CreateSettings)

	settings.PUT("/guild/:guild_id/roles/message", handler.SetRoleMessageId)
	settings.POST("/guild/:guild_id/roles/role", handler.AddRole)
	settings.DELETE("/guild/:guild_id/roles/role", handler.DeleteRole)

	settings.PUT("/guild/:guild_id/welcome/channel", handler.SetWelcomeChannel)
	settings.POST("/guild/:guild_id/welcome/message", handler.AddWelcomeMessage)
	settings.DELETE("/guild/:guild_id/welcome/message", handler.DeleteWelcomeMessage)

}
