package adapters

import (
	"bot/config"
	dtoGachas "bot/internal/dto/gachas"
	"bot/internal/interfaces"

	"context"
	"encoding/json"
	"fmt"
)

type GachasAdapter struct {
	client interfaces.HttpClient
}

func NewGachasAdapter(client interfaces.HttpClient) *GachasAdapter {
	return &GachasAdapter{
		client: client,
	}
}

func (ga *GachasAdapter) GetGenshinCharacters() ([]dtoGachas.GenshinCharacterBrief, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/genshin/character", config.GetApiGatewayAddr()),
		nil,
	)

	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	var characters []dtoGachas.GenshinCharacterBrief

	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, err
	}

	return characters, nil
}

func (ga *GachasAdapter) GetGenshinCharacter(id uint) (dtoGachas.GenshinCharacter, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/genshin/character/%d", config.GetApiGatewayAddr(), id),
		nil,
	)

	if err != nil {
		return dtoGachas.GenshinCharacter{}, err
	}

	defer resp.Body.Close()

	var character dtoGachas.GenshinCharacter

	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return dtoGachas.GenshinCharacter{}, err
	}

	return character, nil
}

func (ga *GachasAdapter) GetGenshinBuild(id uint) (dtoGachas.GenshinBuild, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/genshin/build/%d", config.GetApiGatewayAddr(), id),
		nil,
	)

	if err != nil {
		return dtoGachas.GenshinBuild{}, err
	}

	defer resp.Body.Close()

	var build dtoGachas.GenshinBuild

	if err := json.NewDecoder(resp.Body).Decode(&build); err != nil {
		return dtoGachas.GenshinBuild{}, err
	}

	return build, nil
}