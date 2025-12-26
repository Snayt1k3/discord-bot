package routes

import (
	_ "api-gateway/docs"
	"api-gateway/internal/handlers"
	v1 "api-gateway/internal/routes/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(setting *handlers.Settings, roles *handlers.Roles, welcome *handlers.Welcome, automode *handlers.AutoMode, log *handlers.Log, interaction *handlers.Interaction) *gin.Engine {
	r := gin.Default()

	// Swagger and Prometheus endpoints
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	v1.SettingsRoutes(r, setting, roles, welcome, automode, log)
	v1.InteractionRoutes(r, interaction)

	return r
}
