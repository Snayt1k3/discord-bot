package routes

import (
	_ "api-gateway/docs"
	"api-gateway/internal/handlers"
	"api-gateway/internal/metrics"
	v1 "api-gateway/internal/routes/v1"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(handlers *handlers.Handlers) *gin.Engine {
	r := gin.Default()
	r.Use(metrics.MetricsMiddleware())

	// Swagger and Prometheus endpoints
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Setup routes
	v1.SettingsRoutes(r, handlers)

	return r
}
