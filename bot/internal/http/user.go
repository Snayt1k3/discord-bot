package http

import (
	"bot/config"
	dtoGuild "bot/internal/dto"
	"bot/internal/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
)

type UserService struct {
	http interfaces.HttpClient
}

func NewUserService() *UserService {
	return &UserService{http: NewDefaultHttpClient()}
}

func (i *UserService) GetUser(guildId, userId string) (dtoGuild.User, error) {
	params := map[string]string{
		"user_id":  userId,
		"guild_id": guildId,
	}

	response, err := i.http.Get(
		context.Background(),
			fmt.Sprintf("%v/api/v1/user", config.GetApiGatewayAddr()),
		params,
		nil,
	)
	if err != nil {
		slog.Warn("Bad response when adding xp", "err", err)
		return dtoGuild.User{}, err
	}

	defer response.Body.Close()

	var user dtoGuild.UserResponse
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		slog.Warn("Failed to decode user response", "err", err)
		return dtoGuild.User{}, err
	}

	return user.User, nil
}

func (i *UserService) GetUsers(guildId, page, size, order_by string, is_desc_sort bool) (dtoGuild.UsersResponse, error) {
	params := map[string]string{
		"page":     page,
		"size":     size,
		"guild_id": guildId,
		"order_by": order_by,
		"is_desc_sort": strconv.FormatBool(is_desc_sort),
	}

	response, err := i.http.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/users", config.GetApiGatewayAddr()),
		params,
		nil,
	)
	if err != nil {
		slog.Warn("Bad response when getting users", "err", err)
		return dtoGuild.UsersResponse{}, err
	}

	defer response.Body.Close()

	var users dtoGuild.UsersResponse
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		slog.Warn("Failed to decode user response", "err", err)
		return dtoGuild.UsersResponse{}, err
	}

	return users, nil
}

func (i *UserService) AddXP(guildId, userId string, xp int32) (dtoGuild.AddXpResponse, error) {
	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"guild_id": guildId,
		"user_id":  userId,
		"xp":       xp,
	})

	response, err := i.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/user/addxp", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding xp", "err", err)
		return dtoGuild.AddXpResponse{}, err
	}

	var resp dtoGuild.AddXpResponse
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		slog.Warn("Failed to decode user response", "err", err)
		return dtoGuild.AddXpResponse{}, err
	}

	return resp, nil
}

func (i *UserService) AddVoiceTime(guildId, userId string, seconds int64) (dtoGuild.AddVoiceTimeResponse, error) {
	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"guild_id": guildId,
		"user_id":  userId,
		"seconds":  seconds,
	})

	response, err := i.http.Post(
		context.Background(),
		fmt.Sprintf("%v/api/v1/user/addvoicetime", config.GetApiGatewayAddr()),
		bodyBytes,
		nil,
	)

	if err != nil {
		slog.Warn("Bad response when adding voice time", "err", err)
		return dtoGuild.AddVoiceTimeResponse{}, err
	}

	var resp dtoGuild.AddVoiceTimeResponse

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		slog.Warn("Failed to decode user response", "err", err)
		return dtoGuild.AddVoiceTimeResponse{}, err
	}

	return resp, nil
}
