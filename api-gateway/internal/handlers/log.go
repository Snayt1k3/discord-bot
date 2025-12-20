package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "api-gateway/proto"
)

type Log struct {
	client pb.LogServiceClient
}

func NewLogHandlers(cc grpc.ClientConnInterface) *Log {
	return &Log{client: pb.NewLogServiceClient(cc)}
}

// ToggleLog godoc
// @Summary      Toggle logging
// @Description  Включает или выключает систему логирования для указанной гильдии
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        request body pb.ToggleLogRequest true "Toggle logging request"
// @Success      200 {object} pb.ToggleLogResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/logging/toggle [post]
func (s *Log) ToggleLog(c *gin.Context) {
	var req pb.ToggleLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.ToggleLog(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling log", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddLogChannel godoc
// @Summary      Add log channel
// @Description  Добавляет настройку для логирования действий в указанной гильдии
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        request body pb.UpdateLogChannelRequest true "Add logs setting"
// @Success      200 {object} pb.UpdateLogChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/logging/channel [post]
func (s *Log) AddLogChannel(c *gin.Context) {
	var req pb.UpdateLogChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.AddLogs(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling log", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RemoveLogChannel godoc
// @Summary      Remove log channel
// @Description  Удаляет настройку для логирования действий в указанной гильдии
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        request body pb.UpdateLogChannelRequest true "Remove logs settings"
// @Success      200 {object} pb.UpdateLogChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/logging/channel [delete]
func (s *Log) RemoveLogChannel(c *gin.Context) {
	var req pb.UpdateLogChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.RemoveLogs(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling log", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
