package server

import (
	"context"
	"fmt"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	pb "gachas-service/proto"
	"strings"
)


type GenshinServer struct {
	pb.UnimplementedGenshinServiceServer
	service interfaces.ServiceInterface[dto.GenshinCharacter, dto.GenshinCharacterBrief, dto.GenshinBuild]
}

func (s *GenshinServer) GetAllCharacters(ctx context.Context, req *pb.Empty) (*pb.CharacterListResponse, error) {
	characters, err := s.service.GetCharacters()

	if err != nil {
		return nil, err
	}

	response := &pb.CharacterListResponse{
		Characters: make([]*pb.CharacterShort, len(*characters)),
	}

	for i, character := range *characters {
		response.Characters[i] = &pb.CharacterShort{
			Id:          uint64(character.ID),
			Name:        character.Name,
			Region:      character.Region,
			Element:    character.Element,
		}
	}
	
	return response, nil
}

func (s *GenshinServer) GetCharacterBuild(ctx context.Context, req *pb.CharacterRequest) (*pb.BuildResponse, error) {
	build, err := s.service.GetCharacterBuild(fmt.Sprint(req.Id))
	if err != nil {
		return nil, err
	}

	protoBuild := &pb.Build{
		Name: build.Character.Name,
		Artifacts: func() []*pb.Artifact {
			var arts []*pb.Artifact
			for _, a := range build.Artifacts {
				arts = append(arts, &pb.Artifact{
					Name:           a.Name,
					Set:            a.Set,
					TwoPieceBonus:  a.TwoPieceBonus,
					FourPieceBonus: a.FourPieceBonus,
				})
			}
			return arts
		}(),
		Weapons: func() []*pb.Weapon {
			var ws []*pb.Weapon
			for _, w := range build.Weapons {
				ws = append(ws, &pb.Weapon{
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
		Teams: func() []*pb.CharacterShort {
			var ts []*pb.CharacterShort
			for _, team := range build.Teams {
				for _, c := range team.Characters {
					ts = append(ts, &pb.CharacterShort{
						Id:      uint64(c.ID),
						Name:    c.Name,
						Element: c.Element,
						Region:  c.Region,
					})
				}
			}
			return ts
		}(),
		Stats: &pb.Stats{
			Sands: build.Stats.Sands,
			Goblet: build.Stats.Goblet,
			Circlet: build.Stats.Circlet,
			BestStats: func() []string {
				if build.Stats.BestStats == "" {
					return nil
				}
				return strings.Split(build.Stats.BestStats, ",")
			}(),
			SubStatsPriority: func() []string {
				if build.Stats.SubStatsPriority == "" {
					return nil
				}
				return strings.Split(build.Stats.SubStatsPriority, ",")
			}(),
		},
	}

	return &pb.BuildResponse{
		Build: protoBuild,
	}, nil
}

func (s *GenshinServer) GetCharacterById(ctx context.Context, req *pb.CharacterRequest) (*pb.CharacterDetailResponse, error) {
	character, err := s.service.GetCharacterByID(fmt.Sprint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.CharacterDetailResponse{
		Character: &pb.Character{
			Id:         uint64(character.ID),
			Name:       character.Name,
			Element:    character.Element,
			WeaponType: character.WeaponType,
			BaseStat:   character.BaseStat,
			Region:     character.Region,
			Rarity:     int32(character.Rarity),
			Ascension: &pb.AscensionMaterials{
				LocalSpecialty: character.Ascension.LocalSpecialty,
				BossDrops:      character.Ascension.BossDrops,
			},
			Talents: &pb.TalentMaterials{
				BossDrops: character.Talents.BossDrops,
				Books: &pb.Books{
					Common:   character.Talents.Books.Common,
					Uncommon: character.Talents.Books.Uncommon,
					Rare:     character.Talents.Books.Rare,
				},
				Weekdays:        character.Talents.Weekdays,
				TalentPriority:  character.Talents.TalentPriority,
			},
			CommonMaterials: &pb.CommonMaterials{
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
