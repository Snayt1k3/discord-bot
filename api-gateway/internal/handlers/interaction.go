package handlers

import (
	pb "api-gateway/proto"
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Interaction struct {
	client pb.InteractionServiceClient
}

func NewInteraction(cc grpc.ClientConnInterface) *Interaction {
	return &Interaction{client: pb.NewInteractionServiceClient(cc)}
}

// GetUser godoc
// @Summary      Get user
// @Description  Получает профиль пользователя из interaction-сервиса
// @Tags         interaction
// @Accept       json
// @Produce      json
// @Param        request body pb.GetUserRequest
// @Success      200 {object} pb.GetUserResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/interaction/user
func (i *Interaction) GetUser(c *gin.Context) {
	var req pb.GetUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := i.client.GetUser(context.Background(), &req)

	if err != nil {
		slog.Error("Error while getting user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddXp godoc
// @Summary      Add Experience to user
// @Description  Добавляет опыт пользователю
// @Tags         interaction
// @Accept       json
// @Produce      json
// @Param        request body pb.AddXPRequest
// @Success      200 {object} pb.AddXPResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/interaction/user/addxp
func (i *Interaction) AddXp(c *gin.Context) {
	var req pb.AddXPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := i.client.AddXP(context.Background(), &req)

	if err != nil {
		slog.Error("Error while getting user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)

}
