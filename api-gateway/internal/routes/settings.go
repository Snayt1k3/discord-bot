package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api-gateway/proto"
)

func SettingsRoutes(r *gin.Engine, settingsClient pb.SettingsServiceClient) {
	settings := r.Group("/settings")
	{
		settings.GET("/bot", func(c *gin.Context) {
			resp, err := settingsClient.GetBotSettings(context.Background(), &pb.GetBotSettingsRequest{})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		settings.PUT("/bot", func(c *gin.Context) {
			var req pb.UpdateBotSettingsRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := settingsClient.UpdateBotSettings(context.Background(), &req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		settings.GET("/guild/:guild_id", func(c *gin.Context) {
			guildID := c.Param("guild_id")
			resp, err := settingsClient.GetSettingsByGuild(context.Background(), &pb.GetSettingsByGuildRequest{GuildId: guildID})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		settings.GET("/guilds", func(c *gin.Context) {
			resp, err := settingsClient.GetAllGuildSettings(context.Background(), &pb.GetAllGuildSettingsRequest{})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		settings.PUT("/guild", func(c *gin.Context) {
			var req pb.UpdateGuildSettingsRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := settingsClient.UpdateGuildSettings(context.Background(), &req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		})
	}
}
