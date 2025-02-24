package handlers

import (
	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)


type SettingsHandlers struct {
	client pb.SettingsServiceClient
	redis interfaces.RedisI
}

func NewClient(client pb.SettingsServiceClient, redis interfaces.RedisI) *SettingsHandlers {
	return &SettingsHandlers{client: client, redis: redis}
}


func (s *SettingsHandlers) GetGuildSettings(c *gin.Context) {
	guildID := c.Param("guild_id")

	key := fmt.Sprintf("guild-settings-%v", guildID)

	exists, _ := s.redis.Exists(key)

	if exists {
		resp, err := s.redis.Get(key)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := s.client.GetSettingsByGuild(context.Background(), &pb.GetSettingsByGuildRequest{GuildId: guildID})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = s.redis.Set(key, resp.String(), 60 * 5)

	if err != nil {
		slog.Error("Error while set data in redis", "error", err)
	}

	c.JSON(http.StatusOK, resp)
}

func (s *SettingsHandlers) UpdateGuildSettings(c *gin.Context) {
	var req pb.UpdateGuildSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := s.client.UpdateGuildSettings(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s *SettingsHandlers) CreateGuildSetting (c *gin.Context) {
	guildID := c.Param("guild_id")

	resp, err := s.client.CreateGuildSettings(context.Background(), &pb.CreateGuildSettingsRequest{GuildId: guildID})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, resp)
}
