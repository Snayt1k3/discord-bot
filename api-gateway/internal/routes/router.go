package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/docs"
	"api-gateway/internal/handlers"
	v1 "api-gateway/internal/routes/v1"
)

func SetupRouter(handlers *handlers.Handlers) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	v1.SettingsRoutes(r, handlers)

	return r
}
