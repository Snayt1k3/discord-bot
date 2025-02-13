package routes

import (

	"github.com/gin-gonic/gin"
	"api-gateway/internal/handlers" 
)

func SetupRouter(SettingsHandlers *handlers.SettingsHandlers) *gin.Engine {
	r := gin.Default()
	SettingsRoutes(r, SettingsHandlers)
	return r
}
