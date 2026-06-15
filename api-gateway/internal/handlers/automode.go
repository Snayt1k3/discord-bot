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

type AutoMode struct {
	client pb.AutoModServiceClient
}

func NewAutoModeHandlers(cc grpc.ClientConnInterface) *AutoMode {
	return &AutoMode{client: pb.NewAutoModServiceClient(cc)}
}

// ToggleAutoMod godoc
// @Summary      Toggle auto-moderation
// @Description  Enables or disables auto-moderation for the whole guild.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.ToggleAutoModRequest true "Guild ID and enabled flag"
// @Success      200 {object} pb.ToggleAutoModResponse "Updated auto-moderation state"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/toggle [post]
func (s *AutoMode) ToggleAutoMod(c *gin.Context) {
	var req pb.ToggleAutoModRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.ToggleAutoMod(context.Background(), &req)

	if err != nil {
		slog.Error("Error while toggling automod", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddBannedWord godoc
// @Summary      Add banned word
// @Description  Adds a word to the guild's banned-words filter.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.AddBannedWordRequest true "Guild ID and word"
// @Success      200 {object} pb.AddBannedWordResponse "Added banned word"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      429 {object} dto.APIResponse "Quota exceeded for banned words"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/bannedword [post]
func (s *AutoMode) AddBannedWord(c *gin.Context) {
	var req pb.AddBannedWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.AddBannedWord(context.Background(), &req)

	if err != nil {

		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quota exceeded for banned words"})
			return
		}

		slog.Error("Error while adding banned word", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RemoveBannedWord godoc
// @Summary      Remove banned word
// @Description  Removes a word from the guild's banned-words filter.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.RemoveBannedWordRequest true "Guild ID and word"
// @Success      200 {object} pb.RemoveBannedWordResponse "Removed banned word"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/bannedword [delete]
func (s *AutoMode) RemoveBannedWord(c *gin.Context) {
	var req pb.RemoveBannedWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.RemoveBannedWord(context.Background(), &req)

	if err != nil {
		slog.Error("Error while removing banned word", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RemoveAntiLink godoc
// @Summary      Remove anti-link channel
// @Description  Disables the anti-link filter in the given channel.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.RemoveAntiLinkChannelRequest true "Guild ID and channel ID"
// @Success      200 {object} pb.RemoveAntiLinkChannelResponse "Disabled anti-link channel"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      429 {object} dto.APIResponse "Quota exceeded for anti-link"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/antilink [delete]
func (s *AutoMode) RemoveAntiLink(c *gin.Context) {
	var req pb.RemoveAntiLinkChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.RemoveAntiLinkChannel(context.Background(), &req)

	if err != nil {

		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quota exceeded for antilink"})
			return
		}

		slog.Error("Error while removing antilink channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddAntiLink godoc
// @Summary      Add anti-link channel
// @Description  Enables the anti-link filter in the given channel.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.AddAntiLinkChannelRequest true "Guild ID and channel ID"
// @Success      200 {object} pb.AddAntiLinkChannelResponse "Enabled anti-link channel"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/antilink [post]
func (s *AutoMode) AddAntiLink(c *gin.Context) {
	var req pb.AddAntiLinkChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.AddAntiLinkChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding antilink channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddCapsLock godoc
// @Summary      Add caps-lock channel
// @Description  Enables the caps-lock filter in the given channel.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.AddCapsLockChannelRequest true "Guild ID and channel ID"
// @Success      200 {object} pb.AddCapsLockChannelResponse "Enabled caps-lock channel"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      429 {object} dto.APIResponse "Quota exceeded for caps-lock"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/capslock [post]
func (s *AutoMode) AddCapsLock(c *gin.Context) {
	var req pb.AddCapsLockChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.AddCapsLockChannel(context.Background(), &req)

	if err != nil {

		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quota exceeded for capslock"})
			return
		}

		slog.Error("Error while adding capslock channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RemoveCapsLock godoc
// @Summary      Remove caps-lock channel
// @Description  Disables the caps-lock filter in the given channel.
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        request body pb.RemoveCapsLockChannelRequest true "Guild ID and channel ID"
// @Success      200 {object} pb.RemoveCapsLockChannelResponse "Disabled caps-lock channel"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/guild/automode/capslock [delete]
func (s *AutoMode) RemoveCapsLock(c *gin.Context) {
	var req pb.RemoveCapsLockChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.RemoveCapsLockChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding capslock channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
