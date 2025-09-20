package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "api-gateway/proto"
)

type RolesHandlers struct {
	clients Clients
}

func NewRolesHandlers(cc grpc.ClientConnInterface) *RolesHandlers {
	return &RolesHandlers{clients: *NewClients(cc)}
}

// SetRoleMessageId godoc
// @Summary      Set role message ID
// @Description  Устанавливает ID сообщения, к которому будут привязаны реакции для выдачи ролей в указанной гильдии
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.SetMessageRequest true "Set role message ID request"
// @Success      200 {object} pb.SetMessageResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/roles/message [put]
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

// AddRole godoc
// @Summary      Add role
// @Description  Добавляет новую реакционную роль в указанной гильдии
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.AddRoleRequest true "Add role request"
// @Success      200 {object} pb.AddRoleResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/roles/role [post]
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

// DeleteRole godoc
// @Summary      Delete role
// @Description  Удаляет реакционную роль из указанной гильдии
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.RemoveRoleRequest true "Delete role request"
// @Success      200 {object} pb.RemoveRoleResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/{guild_id}/roles/role [delete]
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
