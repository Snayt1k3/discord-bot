package v1

import (
	"api-gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, handler *handlers.User) {
	router := r.Group("/api/v1/")

	router.GET("/user", handler.GetUser)
	router.GET("/users", handler.GetUsers)
	router.POST("/user/addxp", handler.AddXp)
	router.POST("/user/addvoicetime", handler.AddVoiceTime)
}
