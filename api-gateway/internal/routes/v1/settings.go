package v1

import (
	"api-gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SettingsRoutes(r *gin.Engine, setting *handlers.Settings, roles *handlers.Roles, welcome *handlers.Welcome, automode *handlers.AutoMode, log *handlers.Log) {
	settings := r.Group("/api/v1/settings")

	settings.GET("/guild/:guild_id", setting.GetGuildSettings)
	settings.POST("/guild/:guild_id", setting.CreateSettings)

	settings.PUT("/guild/:guild_id/roles/message", roles.SetRoleMessageId)
	settings.POST("/guild/:guild_id/roles/role", roles.AddRole)
	settings.DELETE("/guild/:guild_id/roles/role", roles.DeleteRole)

	settings.PUT("/guild/:guild_id/welcome/channel", welcome.SetWelcomeChannel)
	settings.POST("/guild/:guild_id/welcome/message", welcome.AddWelcomeMessage)
	settings.DELETE("/guild/:guild_id/welcome/message", welcome.DeleteWelcomeMessage)

	settings.POST("/guild/:guild_id/automode/toggle", automode.ToggleAutoMod)
	settings.POST("/guild/:guild_id/automode/bannedword", automode.AddBannedWord)
	settings.DELETE("/guild/:guild_id/automode/bannedword", automode.RemoveBannedWord)
	settings.POST("/guild/:guild_id/automode/antilink", automode.AddAntiLink)
	settings.DELETE("/guild/:guild_id/automode/antilink", automode.RemoveAntiLink)
	settings.POST("/guild/:guild_id/automode/capslock", automode.AddCapsLock)
	settings.DELETE("/guild/:guild_id/automode/capslock", automode.RemoveCapsLock)

	settings.POST("/guild/:guild_id/logging/toggle", log.ToggleLog)
	settings.POST("/guild/:guild_id/logging/channel", log.AddLogChannel)
	settings.DELETE("/guild/:guild_id/logging/channel", log.RemoveLogChannel)
}
