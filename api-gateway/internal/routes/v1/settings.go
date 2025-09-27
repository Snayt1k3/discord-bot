package v1

import (
	"api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(r *gin.Engine, handler *handlers.Handlers) {
	settings := r.Group("/api/v1/settings")

	settings.GET("/guild/:guild_id", handler.Settings.GetGuildSettings)
	settings.POST("/guild/:guild_id", handler.Settings.CreateSettings)

	settings.PUT("/guild/:guild_id/roles/message", handler.Roles.SetRoleMessageId)
	settings.POST("/guild/:guild_id/roles/role", handler.Roles.AddRole)
	settings.DELETE("/guild/:guild_id/roles/role", handler.Roles.DeleteRole)

	settings.PUT("/guild/:guild_id/welcome/channel", handler.Welcome.SetWelcomeChannel)
	settings.POST("/guild/:guild_id/welcome/message", handler.Welcome.AddWelcomeMessage)
	settings.DELETE("/guild/:guild_id/welcome/message", handler.Welcome.DeleteWelcomeMessage)

	settings.POST("/guild/:guild_id/automode/toggle", handler.Automode.ToggleAutoMod)
	settings.POST("/guild/:guild_id/automode/bannedword", handler.Automode.AddBannedWord)
	settings.DELETE("/guild/:guild_id/automode/bannedword", handler.Automode.RemoveBannedWord)
	settings.POST("/guild/:guild_id/automode/antilink", handler.Automode.AddAntiLink)
	settings.DELETE("/guild/:guild_id/automode/antilink", handler.Automode.RemoveAntiLink)
	settings.POST("/guild/:guild_id/automode/capslock", handler.Automode.AddCapsLock)
	settings.DELETE("/guild/:guild_id/automode/capslock", handler.Automode.RemoveCapsLock)

	settings.POST("/guild/:guild_id/logging/toggle", handler.Log.ToggleLog)
	settings.POST("/guild/:guild_id/logging/channel", handler.Log.AddLogChannel)
	settings.DELETE("/guild/:guild_id/logging/channel", handler.Log.RemoveLogChannel)
}
