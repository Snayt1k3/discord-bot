package handlers

import (
	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SettingsHandlers struct {
	client pb.SettingsServiceClient
	redis  interfaces.RedisInterface
}

func NewClient(client pb.SettingsServiceClient, redis interfaces.RedisInterface) *SettingsHandlers {
	return &SettingsHandlers{client: client, redis: redis}
}

func (s *SettingsHandlers) GetGuildSettings(c *gin.Context) {
	guildID := c.Param("guild_id")

	key := fmt.Sprintf("guild-settings-%v", guildID)

	exists, _ := s.redis.Exists(key)

	if exists {
		resp, err := s.redis.Get(key)

		if err != nil {
			slog.Error("Error while get data from redis", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := s.client.GetSettingsByGuild(context.Background(), &pb.GetSettingsByGuildRequest{GuildId: guildID})

	if err != nil {
		st, ok := status.FromError(err)

		if ok && st.Code() == codes.NotFound {
			slog.Warn("Guild settings not found", "guild_id", guildID)
			c.JSON(http.StatusNotFound, gin.H{"error": "Guild settings not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	err = s.redis.Set(key, resp.String(), 60*5)

	if err != nil {
		slog.Error("Error while set data in redis", "error", err)
	}

	c.JSON(http.StatusOK, resp)
}

func (s *SettingsHandlers) UpdateRoles(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.UpdateRolesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := fmt.Sprintf("guild-settings-%v", guildID)
	s.redis.Delete(key)

	req.GuildId = guildID
	resp, err := s.client.UpdateRoles(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating roles", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *SettingsHandlers) CreateGuildSetting(c *gin.Context) {
	guildID := c.Param("guild_id")

	resp, err := s.client.CreateGuildSettings(context.Background(), &pb.CreateGuildSettingsRequest{GuildId: guildID})

	if err != nil {
		slog.Error("Error while creating guild settings", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *SettingsHandlers) UpdateWelcome(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.UpdateWelcomeChannelIdRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.UpdateWelcomeChannelId(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
