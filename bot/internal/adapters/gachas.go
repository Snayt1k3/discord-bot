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

func NewGachasAdapter() *GachasAdapter {
	return &GachasAdapter{
		client: NewDefaultHttpClient(),
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

	var raw map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var characters []dtoGachas.GenshinCharacterBrief
	if err := json.Unmarshal(raw["characters"], &characters); err != nil {
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

	var wrapper map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return dtoGachas.GenshinCharacter{}, err
	}

	var character dtoGachas.GenshinCharacter
	if err := json.Unmarshal(wrapper["character"], &character); err != nil {
		return dtoGachas.GenshinCharacter{}, err
	}

	return character, nil
}

func (ga *GachasAdapter) GetGenshinBuild(id uint) (dtoGachas.GenshinBuild, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/genshin/character/%d/builds", config.GetApiGatewayAddr(), id),
		nil,
	)
	if err != nil {
		return dtoGachas.GenshinBuild{}, err
	}
	defer resp.Body.Close()

	var wrapper map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return dtoGachas.GenshinBuild{}, err
	}

	var build dtoGachas.GenshinBuild
	if err := json.Unmarshal(wrapper["build"], &build); err != nil {
		return dtoGachas.GenshinBuild{}, fmt.Errorf("unmarshal build error: %w", err)
	}

	return build, nil
}

func (ga *GachasAdapter) GetWuwaCharacters() ([]dtoGachas.WuwaCharacterShort, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/wuwa/character", config.GetApiGatewayAddr()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var characters []dtoGachas.WuwaCharacterShort

	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, err
	}

	return characters, nil
}

func (ga *GachasAdapter) GetWuwaCharacter(id uint) (dtoGachas.WuwaCharacterFull, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/wuwa/character/%d", config.GetApiGatewayAddr(), id),
		nil,
	)

	if err != nil {
		return dtoGachas.WuwaCharacterFull{}, err
	}

	defer resp.Body.Close()

	var character dtoGachas.WuwaCharacterFull

	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return character, err
	}

	return character, nil
}

func (ga *GachasAdapter) GetWuwaBuild(id uint) (dtoGachas.WuwaCharacterBuild, error) {
	resp, err := ga.client.Get(
		context.Background(),
		fmt.Sprintf("%v/api/v1/gacha/wuwa/build/%d", config.GetApiGatewayAddr(), id),
		nil,
	)

	var build dtoGachas.WuwaCharacterBuild

	if err != nil {
		return build, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&build); err != nil {
		return dtoGachas.WuwaCharacterBuild{}, err
	}

	return build, nil
}
