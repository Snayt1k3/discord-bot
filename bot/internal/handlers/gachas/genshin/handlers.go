package genshin

import (
	"bot/internal/discord"
	"bot/internal/interfaces"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"strconv"
	"strings"
)

const (
	rowSize       = 2
	buttonsPerRow = 4
)

type GenshinHandlers struct {
	adapter interfaces.GachasAdapter
}

func NewGenshinHandlers(adapter interfaces.GachasAdapter) *GenshinHandlers {
	return &GenshinHandlers{
		adapter: adapter,
	}
}

func (gh *GenshinHandlers) showCharacters(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	var components []discordgo.MessageComponent

	totalButtons := rowSize * buttonsPerRow
	characters, err := gh.adapter.GetGenshinCharacters()

	if err != nil {
		slog.Error("Error fetching Genshin characters", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	if totalButtons > len(characters) {
		totalButtons = len(characters)
	}

	displayedCharacters := characters[:totalButtons]

	for i := 0; i < totalButtons; i += buttonsPerRow {
		var row []discordgo.MessageComponent

		for j := i; j < i+buttonsPerRow && j < totalButtons; j++ {
			row = append(row, discordgo.Button{
				Label:    displayedCharacters[j].Name,
				Style:    discordgo.PrimaryButton,
				CustomID: fmt.Sprintf("GenshinCharacter_%v", displayedCharacters[j].ID),
			})
		}

		components = append(components, discordgo.ActionsRow{Components: row})
	}

	paginationRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "⬅ Previous",
				Style:    discordgo.SecondaryButton,
				CustomID: "genshinPagination_0",
				Disabled: true,
			},
			discordgo.Button{
				Label:    "Main menu",
				Style:    discordgo.SecondaryButton,
				CustomID: "GachasMenu",
			},
			discordgo.Button{
				Label:    "Next ➡",
				Style:    discordgo.SecondaryButton,
				CustomID: "genshinPagination_1",
			},
		},
	}
	components = append(components, paginationRow)

	embed := &discordgo.MessageEmbed{
		Title:       "🎮 Genshin Impact Characters",
		Description: "Discover all available characters in Genshin Impact, including their rarity, elements, and regions. Tap on a character to view detailed info, ascension materials, and builds.",
		Color:       0x3498DB,
		// todo: add page number
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) showCharacterInfo(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	character, err := gh.adapter.GetGenshinCharacter(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	embed := discordgo.MessageEmbed{
		Title: character.Name,
		Color: 0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "⭐ Rarity",
				Value:  strconv.Itoa(character.Rarity) + " Star",
				Inline: true,
			},
			{
				Name:   "⚔️ Weapon",
				Value:  character.WeaponType,
				Inline: true,
			},
			{
				Name:   "🌩️ Element",
				Value:  character.Element,
				Inline: false,
			},
			// {
			// 	Name:   "🎙️ Voice Actors",
			// 	Value:  "**EN:** Anne Yatco\n**JP:** Miyuki Sawashiro",
			// 	Inline: true,
			// },
			{
				Name:   "🗺️ Region",
				Value:  character.Region,
				Inline: false,
			},
			// {
			// 	Name:   "📅 Released",
			// 	Value:  "Version 2.1 – September 1, 2021",
			// 	Inline: true,
			// },
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: character.Name + " • Genshin Impact",
		},
	}

	components := genshinButtons(character.ID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{&embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) showCharacterAscension(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	character, err := gh.adapter.GetGenshinCharacter(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "⬆️ Ascension Materials — " + character.Name,
		Description: fmt.Sprintf("Materials required to fully ascend %v to Lv. 90, including Talent Level-Up materials.", character.Name),
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   character.Ascension.Gem,
				Value:  "- Sliver ×1\n- Fragment ×9\n- Chunk ×9\n- Gemstone ×6",
				Inline: true,
			},
			{
				Name:   character.Ascension.BossDrops,
				Value:  "- Total: ×46",
				Inline: true,
			},
			{
				Name:   character.Ascension.LocalSpecialty,
				Value:  "- Total: ×168",
				Inline: true,
			},
			{
				Name:   "Ascension Materials",
				Value:  fmt.Sprintf("- %v ×18\n- %v ×30\n- %v ×36", character.CommonMaterials.Common, character.CommonMaterials.Uncommon, character.CommonMaterials.Rare),
				Inline: true,
			},
			{
				Name:   "💰 Mora",
				Value:  "- Total: 420,000",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg", // todo: Заменить
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: character.Name + " • Full Ascension & Talent Materials",
		},
	}

	components := genshinButtons(character.ID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) showCharacterTalents(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	character, err := gh.adapter.GetGenshinCharacter(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📘 Talent Materials — " + character.Name,
		Description: fmt.Sprintf("Resources required to level up all three of %v talents to Lv. 10.", character.Name),
		Color:       0xad44d9,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📚 Talent Books",
				Value:  fmt.Sprintf("- %v ×9\n- %v ×63\n- %v ×114", character.Talents.Books.Common, character.Talents.Books.Uncommon, character.Talents.Books.Rare),
				Inline: true,
			},
			{
				Name:   "Materials",
				Value:  fmt.Sprintf("- %v ×18\n- %v ×66\n- %v ×93", character.CommonMaterials.Common, character.CommonMaterials.Uncommon, character.CommonMaterials.Rare),
				Inline: true,
			},
			{
				Name:   "🔥 Weekly Boss Material",
				Value:  fmt.Sprintf("- %v ×18", character.Talents.BossDrops),
				Inline: true,
			},
			{
				Name:   "👑 Crown of Insight",
				Value:  "- Total: ×3 (for maxing all 3 talents)",
				Inline: true,
			},
			{
				Name:   "💰 Mora",
				Value:  "- Total: 4,950,000",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/0b/18/e8/0b18e8acbf645b7b227689f33785d5c3.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: character.Name + " • Talent Level-Up Costs",
		},
	}

	components := genshinButtons(character.ID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) showCharacterWeapons(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	build, err := gh.adapter.GetGenshinBuild(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	var fields []*discordgo.MessageEmbedField

	for _, weapon := range build.Weapons {
		fieldValue := fmt.Sprintf(
			"⭐ **Rarity:** %d★\n"+
				"🗡️ **Base ATK:** %d\n"+
				"📈 **Substat:** %s (+%.1f%%)\n"+
				"✨ **Passive:** %s",
			weapon.Rarity,
			weapon.BaseATK,
			weapon.SubStat,
			weapon.SubValue,
			weapon.Passive,
		)

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  weapon.Name,
			Value: fieldValue,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:       build.Character.Name + " — Weapon Guide",
		Description: fmt.Sprintf("Recommended weapons for %v depending on your build and availability.", build.Character.Name),
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: fields,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: build.Character.Name + " • Weapon Recommendations",
		},
	}

	components := genshinButtons(build.Character.ID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) showCharacterTeams(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	build, err := gh.adapter.GetGenshinBuild(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	var Fields []*discordgo.MessageEmbedField

	for _, team := range build.Teams {
		Fields = append(Fields, &discordgo.MessageEmbedField{
			Name: team.Characters[0].Name + " + " + team.Characters[1].Name + " + " + team.Characters[2].Name + " + " + team.Characters[3].Name,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:       build.Character.Name + " — Best Team Compositions",
		Description: fmt.Sprintf("Top team compositions for %v, including elemental reactions and synergy.", build.Character.Name),
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: Fields,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: build.Character.Name + " • Team Composition Guide",
		},
	}

	components := genshinButtons(build.Character.ID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) showCharacterArtifacts(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	build, err := gh.adapter.GetGenshinBuild(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	var fields []*discordgo.MessageEmbedField

	for _, artifact := range build.Artifacts {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  artifact.Name,
			Value: fmt.Sprintf("2-piece: %v \n 4-piece: %v", artifact.TwoPieceBonus, artifact.FourPieceBonus),
		})
	}

	stats := build.Stats
	fields = append(fields, &discordgo.MessageEmbedField{
		Name: "Main Stats",
		Value: fmt.Sprintf(
			"🏺 Sands: **%s**\n🍷 Goblet: **%s**\n👑 Circlet: **%s**",
			stats.Sands,
			stats.Goblet,
			stats.Circlet,
		),
	})

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "Substat Priority",
		Value: stats.SubStatsPriority,
	})

	embed := &discordgo.MessageEmbed{
		Title:       build.Character.Name + " — Artifact Guide",
		Description: fmt.Sprintf("Top artifact sets for different %v builds.\nChoose based on your team and role preferences.", build.Character.Name),
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: fields,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: build.Character.Name + " • Artifact Sets Overview",
		},
	}

	components := genshinButtons(build.Character.ID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil

}

func (gh *GenshinHandlers) genshinPagination(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, pageStr, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	page, _ := strconv.Atoi(pageStr)

	characters, err := gh.adapter.GetGenshinCharacters()

	if err != nil {
		slog.Error("Error fetching Genshin characters", "err", err)
		return err
	}

	pageSize := rowSize * buttonsPerRow
	start := page * pageSize
	end := start + pageSize

	if end > len(characters) {
		end = len(characters)
	}
	pageCharacters := characters[start:end]

	var components []discordgo.MessageComponent

	for i := 0; i < len(pageCharacters); i += buttonsPerRow {
		var row []discordgo.MessageComponent
		for j := i; j < i+buttonsPerRow && j < len(pageCharacters); j++ {
			row = append(row, discordgo.Button{
				Label:    pageCharacters[j].Name,
				Style:    discordgo.PrimaryButton,
				CustomID: fmt.Sprintf("GenshinCharacter_%v", pageCharacters[j].ID),
			})
		}
		components = append(components, discordgo.ActionsRow{Components: row})
	}

	// Добавляем пагинацию
	paginationRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "⬅ Previous",
				Style:    discordgo.SecondaryButton,
				CustomID: fmt.Sprintf("genshinPagination_%v", page-1),
				Disabled: page == 0,
			},
			discordgo.Button{
				Label:    "Main menu",
				Style:    discordgo.SecondaryButton,
				CustomID: "GachasMenu",
			},
			discordgo.Button{
				Label:    "Next ➡",
				Style:    discordgo.SecondaryButton,
				CustomID: fmt.Sprintf("genshinPagination_%v", page+1),
				Disabled: end >= len(characters),
			},
		},
	}
	components = append(components, paginationRow)

	embed := &discordgo.MessageEmbed{
		Title:       "🎮 Genshin Impact Characters",
		Description: "Discover all available characters in Genshin Impact, including their rarity, elements, and regions. Tap on a character to view detailed info, ascension materials, and builds.",
		Color:       0x3498DB,
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(s, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    i.ChannelID,
		ID:         i.Message.ID,
	})

	return nil
}

func (gh *GenshinHandlers) AddGenshinHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	handlers["GenshinCharacters"] = gh.showCharacters
	handlers["genshinPagination"] = gh.genshinPagination
	handlers["GenshinCharacter"] = gh.showCharacterInfo
	handlers["GenshinAscension"] = gh.showCharacterAscension
	handlers["GenshinArtifacts"] = gh.showCharacterArtifacts
	handlers["GenshinTalents"] = gh.showCharacterTalents
	handlers["GenshinWeapons"] = gh.showCharacterWeapons
	handlers["GenshinTeams"] = gh.showCharacterTeams
}
