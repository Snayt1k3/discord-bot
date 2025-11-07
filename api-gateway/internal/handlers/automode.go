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
// @Summary      Toggle automoderation
// @Description  Включает или выключает систему автомодерации для указанной гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.ToggleAutoModRequest true "Toggle automod request"
// @Success      200 {object} pb.ToggleAutoModResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/toggle [post]
func (s *AutoMode) ToggleAutoMod(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.ToggleAutoModRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
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
// @Description  Добавляет новое запрещённое слово для фильтрации сообщений в гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.AddBannedWordRequest true "Add banned word request"
// @Success      200 {object} pb.AddBannedWordResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      429 {object} dto.APIResponse "Quota exceeded"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/bannedword [post]
func (s *AutoMode) AddBannedWord(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddBannedWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.AddBannedWord(context.Background(), &req)

	if err != nil {

		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			slog.Warn("Quota exceeded for banned words", "guild_id", guildID)
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
// @Description  Удаляет запрещённое слово из фильтра сообщений в указанной гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.RemoveBannedWordRequest true "Remove banned word request"
// @Success      200 {object} pb.RemoveBannedWordResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/bannedword [delete]
func (s *AutoMode) RemoveBannedWord(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveBannedWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
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
// @Description  Удаляет канал из списка фильтра ссылок для указанной гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.RemoveAntiLinkChannelRequest true "Remove anti-link channel request"
// @Success      200 {object} pb.RemoveAntiLinkChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      429 {object} dto.APIResponse "Quota exceeded"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/antilink [delete]
func (s *AutoMode) RemoveAntiLink(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveAntiLinkChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.RemoveAntiLinkChannel(context.Background(), &req)

	if err != nil {

		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			slog.Warn("Quota exceeded for antilink", "guild_id", guildID)
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
// @Description  Добавляет канал в список фильтра ссылок для указанной гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.AddAntiLinkChannelRequest true "Add anti-link channel request"
// @Success      200 {object} pb.AddAntiLinkChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/antilink [post]
func (s *AutoMode) AddAntiLink(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddAntiLinkChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.AddAntiLinkChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding antilink channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddCapsLock godoc
// @Summary      Add capslock channel
// @Description  Добавляет канал в список фильтра капс-лока для указанной гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.AddCapsLockChannelRequest true "Add capslock channel request"
// @Success      200 {object} pb.AddCapsLockChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      429 {object} dto.APIResponse "Quota exceeded"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/capslock [post]
func (s *AutoMode) AddCapsLock(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.AddCapsLockChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.AddCapsLockChannel(context.Background(), &req)

	if err != nil {

		st, ok := status.FromError(err)

		if ok && st.Code() == codes.ResourceExhausted {
			slog.Warn("Quota exceeded for capslock", "guild_id", guildID)
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
// @Summary      Remove capslock channel
// @Description  Удаляет канал из списка фильтра капс-лока для указанной гильдии
// @Tags         automode
// @Accept       json
// @Produce      json
// @Param        guild_id path string true "Guild ID"
// @Param        request body pb.RemoveCapsLockChannelRequest true "Remove capslock channel request"
// @Success      200 {object} pb.RemoveCapsLockChannelResponse
// @Failure      400 {object} dto.APIResponse "Bad request"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/settings/http/{guild_id}/automode/capslock [delete]
func (s *AutoMode) RemoveCapsLock(c *gin.Context) {
	guildID := c.Param("guild_id")
	var req pb.RemoveCapsLockChannelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.GuildId = guildID
	resp, err := s.client.RemoveCapsLockChannel(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding capslock channel", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
