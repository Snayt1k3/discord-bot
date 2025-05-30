package server

import (
	"context"
	"fmt"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	pb "gachas-service/proto"
)

type HonkaiServer struct {
	pb.UnimplementedHsrServiceServer
	service interfaces.ServiceInterface[dto.HonkaiCharacterDTO, dto.HonkaiCharacterBriefDTO, dto.HonkaiBuildDTO]
}

func (s *HonkaiServer) GetAllCharacters(ctx context.Context, req *pb.Empty) (*pb.HonkaiCharacterList, error) {
	characters, err := s.service.GetCharacters()
	if err != nil {
		return nil, err
	}

	response := &pb.HonkaiCharacterList{
		Characters: make([]*pb.HonkaiCharacterShort, len(characters)),
	}

	for i, character := range characters {
		response.Characters[i] = &pb.HonkaiCharacterShort{
			Id:     uint64(character.ID),
			Name:   character.Name,
			Rarity: int32(character.Rarity),
		}
	}

	return response, nil
}

func (s *HonkaiServer) GetCharacterById(ctx context.Context, req *pb.CharacterRequest) (*pb.HonkaiCharacter, error) {
	character, err := s.service.GetCharacterByID(fmt.Sprint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.HonkaiCharacter{
		Id:        uint64(character.ID),
		Name:      character.Name,
		Rarity:    int32(character.Rarity),
		Path:      character.Path,
		Attribute: character.Attribute,
		Ascension: &pb.HonkaiAscension{
			BossMaterial: character.Ascension.BossMaterial,
			Resource: &pb.HonkaiResource{
				Uncommon: character.Ascension.Resource.Uncommon,
				Rare:     character.Ascension.Resource.Rare,
				Epic:     character.Ascension.Resource.Epic,
			},
		},
		Traces: &pb.HonkaiTraces{
			BossMaterial: character.Traces.BossMaterial,
			DungeonResource: &pb.HonkaiResource{
				Uncommon: character.Traces.DungeonResource.Uncommon,
				Rare:     character.Traces.DungeonResource.Rare,
				Epic:     character.Traces.DungeonResource.Epic,
			},
			MobResource: &pb.HonkaiResource{
				Uncommon: character.Traces.MobResource.Uncommon,
				Rare:     character.Traces.MobResource.Rare,
				Epic:     character.Traces.MobResource.Epic,
			},
		},
	}, nil
}

func (s *HonkaiServer) GetCharacterBuild(ctx context.Context, req *pb.CharacterRequest) (*pb.HonkaiBuild, error) {
	build, err := s.service.GetCharacterBuild(fmt.Sprint(req.Id))
	
	if err != nil {
		return nil, err
	}

	character := &pb.HonkaiCharacter{
		Id:        uint64(build.Character.ID),
		Name:      build.Character.Name,
		Rarity:    int32(build.Character.Rarity),
		Path:      build.Character.Path,
		Attribute: build.Character.Attribute,
		Ascension: &pb.HonkaiAscension{
			BossMaterial: build.Character.Ascension.BossMaterial,
			Resource: &pb.HonkaiResource{
				Uncommon: build.Character.Ascension.Resource.Uncommon,
				Rare:     build.Character.Ascension.Resource.Rare,
				Epic:     build.Character.Ascension.Resource.Epic,
			},
		},
		Traces: &pb.HonkaiTraces{
			BossMaterial: build.Character.Traces.BossMaterial,
			DungeonResource: &pb.HonkaiResource{
				Uncommon: build.Character.Traces.DungeonResource.Uncommon,
				Rare:     build.Character.Traces.DungeonResource.Rare,
				Epic:     build.Character.Traces.DungeonResource.Epic,
			},
			MobResource: &pb.HonkaiResource{
				Uncommon: build.Character.Traces.MobResource.Uncommon,
				Rare:     build.Character.Traces.MobResource.Rare,
				Epic:     build.Character.Traces.MobResource.Epic,
			},
		},
	}


	var cones []*pb.HonkaiLightCone
	for _, cone := range build.Cones {
		cones = append(cones, &pb.HonkaiLightCone{
			Name:    cone.Name,
			Rarity:  int32(cone.Rarity),
			Path:    cone.Path,
			BaseDef: cone.BaseDEF,
			BaseHp:  cone.BaseHP,
			BaseAtk: cone.BaseATK,
			Passive: cone.Passive,
		})
	}

	// 3. Artifacts
	var artifacts []*pb.HonkaiArtifactsPreset
	for _, art := range build.Artifacts {
		artifact := &pb.HonkaiArtifactsPreset{}

		artifact.Relics = &pb.HonkaiRelics{
			Name:           art.Relics.Name,
			TwoPieceBonus:  art.Relics.TwoPieceBonus,
			FourPieceBonus: art.Relics.FourPieceBonus,
		}

		artifact.Planar = &pb.HonkaiPlanar{
			Name:     art.Planar.Name,
			SetBonus: art.Planar.SetBonus,
		}


		artifacts = append(artifacts, artifact)
	}

	stats := &pb.HonkaiStats{
		Body:         build.Stats.Body,
		Feet:         build.Stats.Feet,
		PlanarSphere: build.Stats.PlanarSphere,
		LinkRope:     build.Stats.LinkRope,
		SubStats:     build.Stats.SubStats,
	}

	return &pb.HonkaiBuild{
		Character: character,
		Cones:     cones,
		Artifacts: artifacts,
		Stats:     stats,
	}, nil

	
}

func NewHonkaiServer(service interfaces.ServiceInterface[dto.HonkaiCharacterDTO, dto.HonkaiCharacterBriefDTO, dto.HonkaiBuildDTO]) *HonkaiServer {
	return &HonkaiServer{
		service: service,
	}
}