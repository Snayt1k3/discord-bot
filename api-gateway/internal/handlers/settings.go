package handlers

import (
	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SettingsHandlers struct {
	clients Clients
	redis  interfaces.RedisInterface
}

func NewSettingsHandlers(cc grpc.ClientConnInterface, redis interfaces.RedisInterface) *SettingsHandlers {
	return &SettingsHandlers{clients: *NewClients(cc), redis: redis}
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

	resp, err := s.clients.Settings.GetSettings(context.Background(), &pb.GetSettingsRequest{GuildId: guildID})

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

func (s *SettingsHandlers) CreateSettings(c *gin.Context) {
	guildID := c.Param("guild_id")

	resp, err := s.clients.Settings.CreateSettings(context.Background(), &pb.CreateSettingsRequest{GuildId: guildID})

	if err != nil {
		st, ok := status.FromError(err)

		if ok && st.Code() == codes.AlreadyExists {
			slog.Warn("Guild settings already exist", "guild_id", guildID)
			c.JSON(http.StatusConflict, gin.H{"error": "Guild settings already exist"})
			return
		}

		slog.Error("Error while creating guild settings", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

