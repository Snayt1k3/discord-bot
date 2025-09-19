package handlers

import (
	pb "api-gateway/proto"
	"context"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

)

type WelcomeHandlers struct {
	clients Clients
}

func NewWelcomeHandlers(cc grpc.ClientConnInterface) *WelcomeHandlers {
	return &WelcomeHandlers{clients: *NewClients(cc)}
}

func (s *WelcomeHandlers) SetWelcomeChannel(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.SetWelcomeChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Welcome.SetWelcomeChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *WelcomeHandlers) AddWelcomeMessage(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.WelcomeMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Welcome.AddWelcomeMessage(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (s *WelcomeHandlers) DeleteWelcomeMessage(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.WelcomeMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Welcome.DeleteWelcomeMessage(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}