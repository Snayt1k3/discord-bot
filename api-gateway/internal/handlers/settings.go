package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
)

type Settings struct {
	client pb.SettingsServiceClient
	redis  interfaces.RedisInterface
}

func NewSettingsHandlers(cc grpc.ClientConnInterface, redis interfaces.RedisInterface) *Settings {
	return &Settings{client: pb.NewSettingsServiceClient(cc), redis: redis}
}

// GetGuildSettings godoc
// @Summary      Get guild settings
// @Description  Получить настройки гильдии. Сначала проверяет Redis-кэш, если нет — тянет данные из gRPC.
// @Tags         settings
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Success      200 {object} pb.GetSettingsResponse
// @Failure      404 {object} dto.APIResponse "Guild settings not found"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id} [get]
func (s *Settings) GetGuildSettings(c *gin.Context) {
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

	resp, err := s.client.GetSettings(context.Background(), &pb.GetSettingsRequest{GuildId: guildID})

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

// CreateSettings godoc
// @Summary      Create guild settings
// @Description  Создаёт настройки гильдии, если их ещё нет
// @Tags         settings
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Success      200 {object} pb.CreateSettingsResponse
// @Failure      409 {object} dto.APIResponse "Guild settings already exist"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id} [post]
func (s *Settings) CreateSettings(c *gin.Context) {
	guildID := c.Param("guild_id")

	resp, err := s.client.CreateSettings(context.Background(), &pb.CreateSettingsRequest{GuildId: guildID})

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
