package handlers

import (
	pb "api-gateway/proto"
	"context"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

)

type AutoModeHandlers struct {
	clients Clients
}

func NewAutoModeHandlers(cc grpc.ClientConnInterface) *AutoModeHandlers {
	return &AutoModeHandlers{clients: *NewClients(cc)}
}


func (s *AutoModeHandlers) ToggleAutoMod(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.ToggleAutoModRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.ToggleAutoMod(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling automod", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *AutoModeHandlers) AddBannedWord(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddBannedWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.AddBannedWord(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding banned word", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *AutoModeHandlers) RemoveBannedWord(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveBannedWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.RemoveBannedWord(context.Background(), &req)

	if err != nil {
		slog.Error("Error while removing banned word", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *AutoModeHandlers) RemoveAntiLink(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveAntiLinkChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.RemoveAntiLinkChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while removing antilink channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *AutoModeHandlers) AddAntiLink(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddAntiLinkChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.AddAntiLinkChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding antilink channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *AutoModeHandlers) AddCapsLock(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddCapsLockChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.AddCapsLockChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding capslock channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *AutoModeHandlers) RemoveCapsLock(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveCapsLockChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.AutoMode.RemoveCapsLockChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding capslock channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
