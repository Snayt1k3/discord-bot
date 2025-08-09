package routes

import (
		"api-gateway/internal/handlers"
	"api-gateway/internal/routes/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter(SettingsHandlers *handlers.SettingsHandlers) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	v1.SettingsRoutes(r, SettingsHandlers)


	return r
}
