package routes

import (

	"github.com/gin-gonic/gin"
	pb "api-gateway/proto" 
)

func SetupRouter(settingsClient pb.SettingsServiceClient) *gin.Engine {
	r := gin.Default()
	SettingsRoutes(r, settingsClient)
	return r
}
