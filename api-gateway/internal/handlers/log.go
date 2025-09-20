package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "api-gateway/proto"
)

type LogHandlers struct {
	clients Clients
}

func NewLogHandlers(cc grpc.ClientConnInterface) *LogHandlers {
	return &LogHandlers{clients: *NewClients(cc)}
}

// ToggleLog godoc
// @Summary      Toggle logging
// @Description  Включает или выключает систему логирования для указанной гильдии
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.ToggleLogRequest true "Toggle logging request"
// @Success      200 {object} pb.ToggleLogResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/logging/toggle [patch]
func (s *LogHandlers) ToggleLog(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.ToggleLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Log.ToggleLog(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling log", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}


// AddLogChannel godoc
// @Summary      Add log channel
// @Description  Добавляет новый канал для логирования действий в указанной гильдии
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.UpdateLogChannelRequest true "Add log channel request"
// @Success      200 {object} pb.UpdateLogChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/logging/channel [post]
func (s *LogHandlers) AddLogChannel(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.UpdateLogChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Log.AddLogChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling log", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddLogChannel godoc
// @Summary      Add log channel
// @Description  Добавляет новый канал для логирования действий в указанной гильдии
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.UpdateLogChannelRequest true "Add log channel request"
// @Success      200 {object} pb.UpdateLogChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/logging/channel [post]
func (s *LogHandlers) RemoveLogChannel(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.UpdateLogChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Log.RemoveLogChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling log", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
