package routes

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	SettingsRoutes(r)

	return r
}