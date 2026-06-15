package handlers

import (
	pb "api-gateway/proto"
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type User struct {
	client pb.UserServiceClient
}

func NewUserService(cc grpc.ClientConnInterface) *User {
	return &User{client: pb.NewUserServiceClient(cc)}
}

// GetUser godoc
// @Summary      Get a user
// @Description  Returns the progression profile (experience, level, voice time) of a single user within a guild.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id  query string true  "Discord user ID"
// @Param        guild_id query string true  "Discord guild ID"
// @Success      200 {object} pb.GetUserResponse "User profile"
// @Failure      400 {object} dto.APIResponse "user_id and guild_id are required"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/user [get]
func (i *User) GetUser(c *gin.Context) {
	userID := c.Query("user_id")
	guildID := c.Query("guild_id")

	if userID == "" || guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and guild_id are required"})
		return
	}

	req := &pb.GetUserRequest{
		UserId:  userID,
		GuildId: guildID,
	}

	resp, err := i.client.GetUser(context.Background(), req)
	if err != nil {
		slog.Error("Error while getting user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUsers godoc
// @Summary      List users
// @Description  Returns a paginated, sorted list of user profiles for a guild (e.g. for a leaderboard).
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        guild_id     query string true  "Discord guild ID"
// @Param        page         query int    true  "Page number (zero-based)"
// @Param        size         query int    true  "Items per page (max 50)"
// @Param        order_by     query string false "Sort field: 'experience' or 'voice_time'" Enums(experience, voice_time)
// @Param        is_desc_sort query bool   false "Sort in descending order" default(true)
// @Success      200 {object} pb.GetUsersResponse "Paginated list of users"
// @Failure      400 {object} dto.APIResponse "Missing or invalid query parameters"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/users [get]
func (i *User) GetUsers(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	guildID := c.Query("guild_id")
	orderBy := c.Query("order_by")
	isDescSort := c.Query("is_desc_sort")

	if page == "" || guildID == "" || size == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page, size and guild_id are required"})
		return
	}

	pageInt, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
		return
	}

	sizeInt, err := strconv.ParseInt(size, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid size number"})
		return
	}
	orderByValue, ok := pb.UserOrderBy_value[strings.ToUpper(orderBy)]

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order_by value, must be 'experience' or 'voice_time'"})
		return
	}

	isDescSortBool, err := strconv.ParseBool(isDescSort)
	if err != nil {
		isDescSortBool = false
	}

	req := pb.GetUsersRequest{
		GuildId:    guildID,
		Page:       int32(pageInt),
		Size:       int32(sizeInt),
		OrderBy:    pb.UserOrderBy(orderByValue),
		IsDescSort: isDescSortBool,
	}

	resp, err := i.client.GetUsers(context.Background(), &req)
	if err != nil {
		slog.Error("Error while getting users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}	

// AddXp godoc
// @Summary      Add experience to a user
// @Description  Grants experience to a user and recalculates their level. The response reports whether the user leveled up.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body pb.AddXPRequest true "User and amount of XP to add"
// @Success      200 {object} pb.AddXPResponse "Updated user with level-up flag"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/user/addxp [post]
func (i *User) AddXp(c *gin.Context) {
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

// AddVoiceTime godoc
// @Summary      Add voice time to a user
// @Description  Adds the given number of seconds spent in voice channels to the user's accumulated voice time.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body pb.AddVoiceTimeRequest true "User and number of seconds to add"
// @Success      200 {object} pb.AddVoiceTimeResponse "Updated user"
// @Failure      400 {object} dto.APIResponse "Invalid request body"
// @Failure      500 {object} dto.APIResponse "Internal server error"
// @Router       /api/v1/user/addvoicetime [post]
func (i *User) AddVoiceTime(c *gin.Context) {
	var req pb.AddVoiceTimeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error while binding json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := i.client.AddVoiceTime(context.Background(), &req)

	if err != nil {
		slog.Error("Error while adding voice time", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
