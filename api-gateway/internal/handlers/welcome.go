package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "api-gateway/proto"
)

type Welcome struct {
	client pb.WelcomeServiceClient
}

func NewWelcomeHandlers(cc grpc.ClientConnInterface) *Welcome {
	return &Welcome{client: pb.NewWelcomeServiceClient(cc)}
}

// SetWelcomeChannel godoc
// @Summary      Set welcome channel
// @Description  Устанавливает канал, в который бот будет отправлять приветственные сообщения.
// @Tags         welcome
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.SetWelcomeChannelRequest true "Welcome channel data"
// @Success      200 {object} pb.SetWelcomeChannelResponse
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/welcome/channel [put]
func (s *Welcome) SetWelcomeChannel(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.SetWelcomeChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.SetWelcomeChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddWelcomeMessage godoc
// @Summary      Add welcome message
// @Description  Добавляет приветственное сообщение для гильдии.
// @Tags         welcome
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.WelcomeMessageRequest true "Welcome message data"
// @Success      200 {object} pb.WelcomeMessageResponse
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      429 {object} dto.APIResponse "Quota exceeded"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/welcome/message [post]
func (s *Welcome) AddWelcomeMessage(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.WelcomeMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.AddWelcomeMessage(context.Background(), &req)

	if err != nil {
		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			slog.Warn("Quota exceeded for welcome messages", "guild_id", guildID)
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quota exceeded for welcome messages"})
			return
		}

		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteWelcomeMessage godoc
// @Summary      Delete welcome message
// @Description  Удаляет одно из приветственных сообщений гильдии.
// @Tags         welcome
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.WelcomeMessageRequest true "Welcome message data"
// @Success      200 {object} pb.WelcomeMessageResponse
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/welcome/message [delete]
func (s *Welcome) DeleteWelcomeMessage(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.WelcomeMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.DeleteWelcomeMessage(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
