package v1

import (
	"api-gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InteractionRoutes(r *gin.Engine, handler *handlers.Interaction) {
	router := r.Group("/api/v1/interaction")

	router.GET("/user", handler.GetUser)
	router.POST("/user/addxp", handler.AddXp)

}
