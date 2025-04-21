package handlers

import (
	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GachaHandlers struct {
	client pb.GenshinServiceClient
	redis  interfaces.RedisInterface
}

func (g *GachaHandlers) GetGenshinCharacters(c *gin.Context) {
	cacheKey := "genshin_characters"
	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	resp, err := g.client.GetAllCharacters(c, &pb.Empty{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600) 
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetGenshinCharacterByID(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "genshin_character_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.client.GetCharacterById(c, &pb.CharacterRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600) 
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetGenshinCharacterBuild(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "genshin_character_build_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.client.GetCharacterBuild(c, &pb.CharacterRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func NewGachaHandlers(client pb.GenshinServiceClient, redis interfaces.RedisInterface) *GachaHandlers {
	return &GachaHandlers{
		client: client,
		redis:  redis,
	}
}