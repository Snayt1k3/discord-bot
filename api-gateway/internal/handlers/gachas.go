package handlers

import (
	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GachaHandlers struct {
	genshinClient pb.GenshinServiceClient
	wuwaClient    pb.WuwaServiceClient
	honkaiClient  pb.HsrServiceClient
	zenlessClient pb.ZenlessServiceClient
	redis         interfaces.RedisInterface
}

func (g *GachaHandlers) GetGenshinCharacters(c *gin.Context) {
	cacheKey := "genshin_characters"
	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	resp, err := g.genshinClient.GetAllCharacters(c, &pb.Empty{})
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

	resp, err := g.genshinClient.GetCharacterById(c, &pb.CharacterRequest{Id: id})
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

	resp, err := g.genshinClient.GetCharacterBuild(c, &pb.CharacterRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetWuwaCharacters(c *gin.Context) {
	cacheKey := "wuwa_characters"
	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	resp, err := g.wuwaClient.GetAllCharacters(c, &pb.Empty{})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetWuwaCharacterByID(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "wuwa_character_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.wuwaClient.GetCharacterByID(c, &pb.CharacterRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)

}

func (g *GachaHandlers) GetWuwaCharacterBuild(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "wuwa_character_build_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.wuwaClient.GetCharacterBuild(c, &pb.CharacterRequest{Id: id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetZenlessCharacters(c *gin.Context) {
	cacheKey := "zenless_characters"
	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	resp, err := g.zenlessClient.GetAllCharacters(c, &pb.Empty{})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetZenlessCharacterByID(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "zenless_character_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.zenlessClient.GetCharacterByID(c, &pb.CharacterRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)

}

func (g *GachaHandlers) GetZenlessCharacterBuild(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "zenless_character_build_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.zenlessClient.GetCharacterBuild(c, &pb.CharacterRequest{Id: id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetHsrCharacters(c *gin.Context) {
	cacheKey := "hsr_characters"
	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	resp, err := g.honkaiClient.GetAllCharacters(c, &pb.Empty{})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func (g *GachaHandlers) GetHsrCharacterByID(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "hsr_character_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.honkaiClient.GetCharacterByID(c, &pb.CharacterRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)

}

func (g *GachaHandlers) GetHsrCharacterBuild(c *gin.Context) {
	characterID := c.Param("character_id")
	cacheKey := "hsr_character_build_" + characterID

	if cachedData, err := g.redis.Get(cacheKey); err == nil {
		c.JSON(200, gin.H{"data": cachedData})
		return
	}

	id, err := strconv.ParseUint(characterID, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid character ID"})
		return
	}

	resp, err := g.honkaiClient.GetCharacterBuild(c, &pb.CharacterRequest{Id: id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = g.redis.Set(cacheKey, resp.String(), 3600)
	c.JSON(200, resp)
}

func NewGachaHandlers(
	genshinClient pb.GenshinServiceClient,
	wuwaClient pb.WuwaServiceClient,
	zenlessClient pb.ZenlessServiceClient,
	honkaiClient pb.HsrServiceClient,
	redis interfaces.RedisInterface,
) *GachaHandlers {
	return &GachaHandlers{
		genshinClient: genshinClient,
		wuwaClient:    wuwaClient,
		zenlessClient: zenlessClient,
		honkaiClient:  honkaiClient,
		redis:         redis,
	}
}
