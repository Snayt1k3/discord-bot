package v1

import (
	"api-gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SettingsRoutes(r *gin.Engine, setting *handlers.Settings, roles *handlers.Roles, welcome *handlers.Welcome, automode *handlers.AutoMode, log *handlers.Log) {
	settings := r.Group("/api/v1/settings")

	settings.GET("/guild", setting.GetGuildSettings)
	settings.POST("/guild", setting.CreateSettings)

	settings.PUT("/guild/roles/message", roles.SetRoleMessageId)
	settings.POST("/guild/roles/role", roles.AddRole)
	settings.DELETE("/guild/roles/role", roles.DeleteRole)

	settings.PUT("/guild/welcome/channel", welcome.SetWelcomeChannel)
	settings.POST("/guild/welcome/message", welcome.AddWelcomeMessage)
	settings.DELETE("/guild/welcome/message", welcome.DeleteWelcomeMessage)

	settings.POST("/guild/automode/toggle", automode.ToggleAutoMod)
	settings.POST("/guild/automode/bannedword", automode.AddBannedWord)
	settings.DELETE("/guild/automode/bannedword", automode.RemoveBannedWord)
	settings.POST("/guild/automode/antilink", automode.AddAntiLink)
	settings.DELETE("/guild/automode/antilink", automode.RemoveAntiLink)
	settings.POST("/guild/automode/capslock", automode.AddCapsLock)
	settings.DELETE("/guild/automode/capslock", automode.RemoveCapsLock)

	settings.POST("/guild/logging/toggle", log.ToggleLog)
	settings.POST("/guild/logging/channel", log.AddLogChannel)
	settings.DELETE("/guild/logging/channel", log.RemoveLogChannel)
}
