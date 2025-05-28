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

type ZenlessServer struct {
	pb.UnimplementedZenlessServiceServer
	service interfaces.ServiceInterface[dto.ZenlessCharacterFull, dto.ZenlessCharacterDTO, dto.ZenlessBuildDTO]
}

func NewZenlessServer(service interfaces.ServiceInterface[dto.ZenlessCharacterFull, dto.ZenlessCharacterDTO, dto.ZenlessBuildDTO]) *ZenlessServer {
	return &ZenlessServer{service: service}
}

func (s *ZenlessServer) GetAllCharacters(ctx context.Context, _ *pb.Empty) (*pb.ZenlessCharacterList, error) {
	characters, err := s.service.GetCharacters()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	resp := &pb.ZenlessCharacterList{
		Characters: make([]*pb.ZenlessCharacterShort, len(characters)),
	}

	for i, c := range characters {
		resp.Characters[i] = &pb.ZenlessCharacterShort{
			Id:        uint64(c.ID),
			Name:      c.Name,
			Rank:      c.Rank,
			Attribute: c.Attribute,
			Specialty: c.Specialty,
		}
	}

	return resp, nil
}

func (s *ZenlessServer) GetCharacterByID(ctx context.Context, req *pb.CharacterRequest) (*pb.ZenlessCharacterFull, error) {
	character, err := s.service.GetCharacterByID(strconv.FormatUint(req.Id, 10))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pb.ZenlessCharacterFull{
		Id:        uint64(character.ID),
		Name:      character.Name,
		Rank:      character.Rank,
		Attribute: character.Attribute,
		Specialty: character.Specialty,
		Faction:   character.Faction,
	}, nil
}

func (s *ZenlessServer) GetCharacterBuild(ctx context.Context, req *pb.CharacterRequest) (*pb.ZenlessBuildResponse, error) {
	build, err := s.service.GetCharacterBuild(strconv.FormatUint(req.Id, 10))
	
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	var weapons []*pb.ZenlessWeapon
	for _, w := range build.Weapons {
		weapons = append(weapons, &pb.ZenlessWeapon{
			Id:       uint64(w.ID),
			Name:     w.Name,
			BaseAtk:  w.BaseATK,
			Rarity:   w.Rarity,
			Type:     w.Type,
			Passive:  w.Passive,
			SubStat:  w.SubStat,
			SubValue: w.SubValue,
			Priority: int32(w.Priority),
		})
	}

	var discs []*pb.ZenlessDiscPreset
	for _, d := range build.Discs {
		discs = append(discs, &pb.ZenlessDiscPreset{
			TwoPieceSet: &pb.ZenlessDiscSet{
				Name:            d.TwoPiece.Name,
				TwoPieceBonus:   d.TwoPiece.TwoPieceBonus,
				FourPieceBonus:  d.TwoPiece.FourPieceBonus,
			},
			FourPieceSet: &pb.ZenlessDiscSet{
				Name:            d.FourPiece.Name,
				TwoPieceBonus:   d.FourPiece.TwoPieceBonus,
				FourPieceBonus:  d.FourPiece.FourPieceBonus,
			},
			Priority: int32(d.Priority),
		})
	}

	stats := &pb.ZenlessStats{
		FourthDisc:       build.Stats.FourthDisc,
		FifthDisc:        build.Stats.FifthDisc,
		SixthDisc:        build.Stats.SixthDisc,
		SubStatsPriority: build.Stats.SubStatsPriority,
	}

	character := &pb.ZenlessCharacterFull{
		Id:        uint64(build.Character.ID),
		Name:      build.Character.Name,
		Specialty: build.Character.Specialty,
		Rank:      build.Character.Rank,
		Attribute: build.Character.Attribute,
		Faction:   build.Character.Faction,
		Ascension: &pb.ZenlessMaterial{
			Common: build.Character.Ascension.Common,
			Rare:   build.Character.Ascension.Rare,
			Epic:   build.Character.Ascension.Epic,
		},
		Nodes: &pb.ZenlessNode{
			HuntBossMaterial:   build.Character.Nodes.HuntBossMaterial,
			ExpertBossMaterial: build.Character.Nodes.ExpertBossMaterial,
			Resource: &pb.ZenlessMaterial{
				Common: build.Character.Nodes.Resource.Common,
				Rare:   build.Character.Nodes.Resource.Rare,
				Epic:   build.Character.Nodes.Resource.Epic,
			},
			HamsterCagePass: int32(build.Character.Nodes.HamsterCagePass),
		},
	}

	return &pb.ZenlessBuildResponse{
		Build: &pb.ZenlessBuild{
			Character: character,
			Weapons:   weapons,
			Discs:     discs,
			Stats:     stats,
		},
	}, nil
}