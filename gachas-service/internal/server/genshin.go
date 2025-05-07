package server

import (
	"context"
	"fmt"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	pb "gachas-service/proto"
)

type GenshinServer struct {
	pb.UnimplementedGenshinServiceServer
	service interfaces.ServiceInterface[dto.GenshinCharacter, dto.GenshinCharacterBrief, dto.GenshinBuild]
}

func (s *GenshinServer) GetAllCharacters(ctx context.Context, req *pb.Empty) (*pb.GenshinCharactersResponse, error) {
	characters, err := s.service.GetCharacters()

	if err != nil {
		return nil, err
	}

	response := &pb.GenshinCharactersResponse{
		Characters: make([]*pb.GenshinCharacterShort, len(characters)),
	}

	for i, character := range characters {
		response.Characters[i] = &pb.GenshinCharacterShort{
			Id:      uint64(character.ID),
			Name:    character.Name,
			Region:  character.Region,
			Element: character.Element,
		}
	}

	return response, nil
}

func (s *GenshinServer) GetCharacterBuild(ctx context.Context, req *pb.CharacterRequest) (*pb.GenshinBuildResponse, error) {
	build, err := s.service.GetCharacterBuild(fmt.Sprint(req.Id))

	if err != nil {
		return nil, err
	}

	protoBuild := &pb.GenshinBuild{
		Name: build.Character.Name,
		Artifacts: func() []*pb.GenshinArtifact {
			var arts []*pb.GenshinArtifact
			for _, a := range build.Artifacts {
				arts = append(arts, &pb.GenshinArtifact{
					Name:           a.Name,
					TwoPieceBonus:  a.TwoPieceBonus,
					FourPieceBonus: a.FourPieceBonus,
				})
			}
			return arts
		}(),
		Weapons: func() []*pb.GenshinWeapon {
			var ws []*pb.GenshinWeapon
			for _, w := range build.Weapons {
				ws = append(ws, &pb.GenshinWeapon{
					Name:     w.Name,
					Type:     w.Type,
					Rarity:   int32(w.Rarity),
					BaseAtk:  int32(w.BaseATK),
					SubStat:  w.SubStat,
					SubValue: float32(w.SubValue),
					Passive:  w.Passive,
				})
			}
			return ws
		}(),
		Teams: func() []*pb.GenshinCharacterShort {
			var ts []*pb.GenshinCharacterShort
			for _, team := range build.Teams {
				for _, c := range team.Characters {
					ts = append(ts, &pb.GenshinCharacterShort{
						Id:      uint64(c.ID),
						Name:    c.Name,
						Element: c.Element,
						Region:  c.Region,
					})
				}
			}
			return ts
		}(),
		Stats: &pb.GenshinStats{
			Sands:            build.Stats.Sands,
			Goblet:           build.Stats.Goblet,
			Circlet:          build.Stats.Circlet,
			BestStats:        build.Stats.BestStats,
			SubStatsPriority: build.Stats.SubStatsPriority,
		},
	}

	return &pb.GenshinBuildResponse{
		Build: protoBuild,
	}, nil
}

func (s *GenshinServer) GetCharacterById(ctx context.Context, req *pb.CharacterRequest) (*pb.GenshinCharacterDetailResponse, error) {
	character, err := s.service.GetCharacterByID(fmt.Sprint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GenshinCharacterDetailResponse{
		Character: &pb.GenshinCharacter{
			Id:         uint64(character.ID),
			Name:       character.Name,
			Element:    character.Element,
			WeaponType: character.WeaponType,
			BaseStat:   character.BaseStat,
			Region:     character.Region,
			Rarity:     int32(character.Rarity),
			Ascension: &pb.GenshinAscensionMaterials{
				LocalSpecialty: character.Ascension.LocalSpecialty,
				BossDrops:      character.Ascension.BossDrops,
			},
			Talents: &pb.GenshinTalentMaterials{
				BossDrops: character.Talents.BossDrops,
				Books: &pb.GenshinBooks{
					Common:   character.Talents.Books.Common,
					Uncommon: character.Talents.Books.Uncommon,
					Rare:     character.Talents.Books.Rare,
					Weekdays: character.Talents.Books.Weekdays,
				},

				TalentPriority: character.Talents.TalentPriority,
			},
			CommonMaterials: &pb.GenshinCommonMaterials{
				Common:   character.CommonMaterials.Common,
				Uncommon: character.CommonMaterials.Uncommon,
				Rare:     character.CommonMaterials.Rare,
			},
		},
	}, nil
}

func NewGenshinServer(service interfaces.ServiceInterface[dto.GenshinCharacter, dto.GenshinCharacterBrief, dto.GenshinBuild]) *GenshinServer {
	return &GenshinServer{
		service: service,
	}
}
