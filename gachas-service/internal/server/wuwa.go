package server

import (
	"context"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	pb "gachas-service/proto"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WuwaServer struct {
	pb.UnimplementedWuwaServiceServer
	service interfaces.ServiceInterface[dto.WuwaCharacterFull, dto.WuwaCharacterShort, dto.WuwaCharacterBuild]
}

func NewWuwaServer(service interfaces.ServiceInterface[dto.WuwaCharacterFull, dto.WuwaCharacterShort, dto.WuwaCharacterBuild]) *WuwaServer {
	return &WuwaServer{service: service}
}

func (s *WuwaServer) GetAllCharacters(ctx context.Context, request *pb.Empty) (*pb.WuwaCharacterList, error) {
	characters, err := s.service.GetCharacters()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	response := &pb.WuwaCharacterList{
		Characters: make([]*pb.WuwaCharacterShort, len(characters)),
	}

	for i, character := range characters {
		response.Characters[i] = &pb.WuwaCharacterShort{
			Id:      uint64(character.ID),
			Name:    character.Name,
			Element: character.Element,
			Rarity:  int32(character.Rarity),
		}
	}

	return response, nil

}
func (s *WuwaServer) GetCharacterByID(ctx context.Context, request *pb.CharacterRequest) (*pb.WuwaCharacterFull, error) {
	character, err := s.service.GetCharacterByID(strconv.FormatUint(request.Id, 10))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pb.WuwaCharacterFull{
		Id:         uint64(character.ID),
		Element:    character.Element,
		Name:       character.Name,
		Rarity:     int32(character.Rarity),
		WeaponType: character.WeaponType,
		Ascension: &pb.WuwaAscension{
			LocalSpecialty: character.Ascension.LocalSpecialty,
			BossMaterial:   character.Ascension.BossMaterial,
			MobMaterial: &pb.WuwaResource{
				UncommonName:  character.Ascension.MobMaterial.UncommonName,
				RareName:      character.Ascension.MobMaterial.RareName,
				EpicName:      character.Ascension.MobMaterial.EpicName,
				LegendaryName: character.Ascension.MobMaterial.LegendaryName,
			},
		},
		Talents: &pb.WuwaTalent{
			DungeonMaterial: &pb.WuwaResource{
				UncommonName:  character.Talents.DungeonMaterial.UncommonName,
				RareName:      character.Talents.DungeonMaterial.RareName,
				EpicName:      character.Talents.DungeonMaterial.EpicName,
				LegendaryName: character.Talents.DungeonMaterial.LegendaryName,
			},
			MobMaterial: &pb.WuwaResource{
				UncommonName:  character.Talents.MobMaterial.UncommonName,
				RareName:      character.Talents.MobMaterial.RareName,
				EpicName:      character.Talents.MobMaterial.EpicName,
				LegendaryName: character.Talents.MobMaterial.LegendaryName,
			},
		},
	}, nil

}
func (s *WuwaServer) GetCharacterBuild(ctx context.Context, request *pb.CharacterRequest) (*pb.WuwaCharacterBuild, error) {
	build, err := s.service.GetCharacterBuild(strconv.FormatUint(request.Id, 10))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	var weapons []*pb.WuwaWeapon
	for _, w := range build.Weapons {
		weapons = append(weapons, &pb.WuwaWeapon{
			Name:       w.Name,
			WeaponType: w.WeaponType,
			Rarity:     int32(w.Rarity),
			Passive:    w.Passive,
			BaseAtk:    int32(w.BaseATK),
			SubStat:    w.SubStat,
			SubValue:   w.SubValue,
		})
	}

	var echoes []*pb.WuwaEcho
	for _, e := range build.Echoes {
		echoes = append(echoes, &pb.WuwaEcho{
			Name:          e.Name,
			TwoPieceBonus: e.TwoPieceBonus,
			FullSetBonus:  e.FullSetBonus,
		})
	}

	stats := &pb.WuwaStats{
		FourCostEchoStat:  build.Stats.FourCostEchoStat,
		ThreeCostEchoStat: build.Stats.ThreeCostEchoStat,
		SubStatsPriority:  build.Stats.SubStatsPriority,
		OneCostEchoStat:   build.Stats.OneCostEchoStat,
	}

	character := &pb.WuwaCharacterFull{
		Id:         uint64(build.Character.ID),
		Name:       build.Character.Name,
		Element:    build.Character.Element,
		WeaponType: build.Character.WeaponType,
		Rarity:     int32(build.Character.Rarity),
		Ascension: &pb.WuwaAscension{
			LocalSpecialty: build.Character.Ascension.LocalSpecialty,
			BossMaterial:   build.Character.Ascension.BossMaterial,
			MobMaterial: &pb.WuwaResource{
				UncommonName:  build.Character.Ascension.MobMaterial.UncommonName,
				RareName:      build.Character.Ascension.MobMaterial.RareName,
				EpicName:      build.Character.Ascension.MobMaterial.EpicName,
				LegendaryName: build.Character.Ascension.MobMaterial.LegendaryName,
			},
		},
		Talents: &pb.WuwaTalent{
			DungeonMaterial: &pb.WuwaResource{
				UncommonName:  build.Character.Talents.DungeonMaterial.UncommonName,
				RareName:      build.Character.Talents.DungeonMaterial.RareName,
				EpicName:      build.Character.Talents.DungeonMaterial.EpicName,
				LegendaryName: build.Character.Talents.DungeonMaterial.LegendaryName,
			},
			MobMaterial: &pb.WuwaResource{
				UncommonName:  build.Character.Talents.MobMaterial.UncommonName,
				RareName:      build.Character.Talents.MobMaterial.RareName,
				EpicName:      build.Character.Talents.MobMaterial.EpicName,
				LegendaryName: build.Character.Talents.MobMaterial.LegendaryName,
			},
			BossMaterial: build.Character.Talents.BossMaterial,
		},
	}

	return &pb.WuwaCharacterBuild{
		Character:       character,
		Weapons:         weapons,
		Echoes:          echoes,
		Stats:           stats,
		BestPrimaryEcho: build.BestPrimaryEcho,
	}, nil
}
