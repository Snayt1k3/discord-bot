package handlers

import (
	pb "api-gateway/proto"
	"context"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

)

type RolesHandlers struct {
	clients Clients
}

func NewRolesHandlers(cc grpc.ClientConnInterface) *RolesHandlers {
	return &RolesHandlers{clients: *NewClients(cc)}
}


func (s *RolesHandlers) SetRoleMessageId(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.SetMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Roles.SetRoleMessageId(context.Background(), &req)

	if err != nil {
		slog.Error("Error while deleting role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *RolesHandlers) AddRole(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Roles.AddRole(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *RolesHandlers) DeleteRole(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.clients.Roles.RemoveRole(context.Background(), &req)

	if err != nil {
		slog.Error("Error while deleting role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}