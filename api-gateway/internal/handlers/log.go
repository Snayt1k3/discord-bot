package handlers

import (
	pb "api-gateway/proto"
	"context"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

)

type LogHandlers struct {
	clients Clients
}

func NewLogHandlers(cc grpc.ClientConnInterface) *LogHandlers {
	return &LogHandlers{clients: *NewClients(cc)}
}

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