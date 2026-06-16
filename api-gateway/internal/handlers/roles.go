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

type Roles struct {
	client pb.RolesServiceClient
}

func NewRolesHandlers(cc grpc.ClientConnInterface) *Roles {
	return &Roles{client: pb.NewRolesServiceClient(cc)}
}

// SetRoleMessageId godoc
// @Summary      Set reaction-role message
// @Description  Sets the message whose reactions drive reaction-role assignment for the guild.
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body pb.SetMessageRequest true "Guild ID and target message ID"
// @Success      200 {object} pb.SetMessageResponse "Updated reaction-role message"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/roles/message [put]
func (s *Roles) SetRoleMessageId(c *gin.Context) {
	var req pb.SetMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.SetRoleMessageId(context.Background(), &req)

	if err != nil {
		slog.Error("Error while deleting role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddRole godoc
// @Summary      Add reaction role
// @Description  Maps an emoji to a role so that reacting with it grants the role in the guild.
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body pb.AddRoleRequest true "Guild ID, role ID and emoji"
// @Success      200 {object} pb.AddRoleResponse "Created emoji-to-role mapping"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      429 {object} dto.APIResponse "Quota exceeded for roles/reactions"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/roles/role [post]
func (s *Roles) AddRole(c *gin.Context) {
	var req pb.AddRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.AddRole(context.Background(), &req)

	if err != nil {
		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quota exceeded for Roles/Reactions"})
			return
		}

		slog.Error("Error while adding role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteRole godoc
// @Summary      Delete reaction role
// @Description  Removes the emoji-to-role mapping for the given role in the guild.
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body pb.RemoveRoleRequest true "Guild ID and role ID"
// @Success      200 {object} pb.RemoveRoleResponse "Removed reaction role"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/roles/role [delete]
func (s *Roles) DeleteRole(c *gin.Context) {
	var req pb.RemoveRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.RemoveRole(context.Background(), &req)

	if err != nil {
		slog.Error("Error while deleting role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
