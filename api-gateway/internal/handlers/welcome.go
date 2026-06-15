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

type Welcome struct {
	client pb.WelcomeServiceClient
}

func NewWelcomeHandlers(cc grpc.ClientConnInterface) *Welcome {
	return &Welcome{client: pb.NewWelcomeServiceClient(cc)}
}

// SetWelcomeChannel godoc
// @Summary      Set welcome channel
// @Description  Sets the channel where the bot posts welcome messages for new members.
// @Tags         welcome
// @Accept       json
// @Produce      json
// @Param        request body pb.SetWelcomeChannelRequest true "Guild ID and channel ID"
// @Success      200 {object} pb.SetWelcomeChannelResponse "Updated welcome channel"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/welcome/channel [put]
func (s *Welcome) SetWelcomeChannel(c *gin.Context) {
	var req pb.SetWelcomeChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.SetWelcomeChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddWelcomeMessage godoc
// @Summary      Add welcome message
// @Description  Adds a message to the guild's pool of welcome messages.
// @Tags         welcome
// @Accept       json
// @Produce      json
// @Param        request body pb.WelcomeMessageRequest true "Guild ID and message text"
// @Success      200 {object} pb.WelcomeMessageResponse "Added welcome message"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      429 {object} dto.APIResponse "Quota exceeded for welcome messages"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/welcome/message [post]
func (s *Welcome) AddWelcomeMessage(c *gin.Context) {
	var req pb.WelcomeMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.AddWelcomeMessage(context.Background(), &req)

	if err != nil {
		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quota exceeded for welcome messages"})
			return
		}

		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteWelcomeMessage godoc
// @Summary      Delete welcome message
// @Description  Removes a message from the guild's pool of welcome messages.
// @Tags         welcome
// @Accept       json
// @Produce      json
// @Param        request body pb.WelcomeMessageRequest true "Guild ID and message text"
// @Success      200 {object} pb.WelcomeMessageResponse "Removed welcome message"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/welcome/message [delete]
func (s *Welcome) DeleteWelcomeMessage(c *gin.Context) {
	var req pb.WelcomeMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.DeleteWelcomeMessage(context.Background(), &req)

	if err != nil {
		slog.Error("Error while updating welcome channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
