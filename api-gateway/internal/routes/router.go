package routes

import (
	"api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(SettingsHandlers *handlers.SettingsHandlers) *gin.Engine {
	r := gin.Default()
	SettingsRoutes(r, SettingsHandlers)
	return r
}
