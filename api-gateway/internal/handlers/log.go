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
// @Description  Enables or disables event logging for the whole guild.
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        request body pb.ToggleLogRequest true "Guild ID and enabled flag"
// @Success      200 {object} pb.ToggleLogResponse "Updated logging state"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
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
// @Description  Routes the given event types to a channel and returns the resulting event-to-channel mapping for the guild.
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        request body pb.UpdateLogChannelRequest true "Guild ID, event types and channel ID"
// @Success      200 {object} pb.UpdateLogChannelResponse "Updated event-to-channel mapping"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
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
// @Description  Stops logging the given event types for the guild.
// @Tags         logging
// @Accept       json
// @Produce      json
// @Param        request body pb.UpdateLogChannelRequest true "Guild ID and event types to stop logging"
// @Success      200 {object} pb.RemoveLogChannelResponse "Updated logging configuration"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
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
