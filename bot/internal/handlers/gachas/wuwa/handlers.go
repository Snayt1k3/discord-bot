package wuwa

import (
	"bot/internal/discord"
	"bot/internal/interfaces"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	rowSize       = 2
	buttonsPerRow = 4
)

type WuwaHandlers struct {
	adapter interfaces.GachasAdapter
}

func NewWuwaHandlers(adapter interfaces.GachasAdapter) *WuwaHandlers {
	return &WuwaHandlers{adapter: adapter}
}

func (wh *WuwaHandlers) showCharacters(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	var components []discordgo.MessageComponent

	totalButtons := rowSize * buttonsPerRow
	characters, err := wh.adapter.GetWuwaCharacters()

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
				CustomID: fmt.Sprintf("WuwaCharacter_%v", displayedCharacters[j].ID),
			})
		}

		components = append(components, discordgo.ActionsRow{Components: row})
	}

	paginationRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "⬅ Previous",
				Style:    discordgo.SecondaryButton,
				CustomID: "WuwaPagination_0",
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
				CustomID: "WuwaPagination_1",
			},
		},
	}
	components = append(components, paginationRow)

	embed := &discordgo.MessageEmbed{
		Title:       "🌊 Wuthering Waves Characters",
		Description: "Explore all resonators in Wuthering Waves, including their rarity, attributes, and factions. Tap on a resonator to see detailed info, echo setups, weapon builds, and progression materials.",
		Color:       0x6A5ACD,
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

func (wh *WuwaHandlers) showCharacterDetail(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	character, err := wh.adapter.GetWuwaCharacter(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	embed := discordgo.MessageEmbed{
	Title: character.Name + " ⭐" + strconv.Itoa(character.Rarity),
	Color: 0x6A5ACD, 
	Thumbnail: &discordgo.MessageEmbedThumbnail{
		URL: "https://wutheringwaves.gg/wp-content/uploads/sites/8/2024/05/Wuthering-Waves-Aalto-Build-Guide.png", // сюда можно подставить URL изображения персонажа
	},
	Fields: []*discordgo.MessageEmbedField{
		{
			Name:   "🌪️ Element",
			Value:  character.Element,
			Inline: true,
		},
		{
			Name:   "🗡️ Weapon Type",
			Value:  "Pistols" , // todo: Add more information about chars
			Inline: true,
		},
		{
			Name:   "📈 Rarity",
			Value:  strconv.Itoa(character.Rarity) + "-Star",
			Inline: true,
		},
	},
	Description: "A mysterious Resonator from the world of **Wuthering Waves**.\nPrepare to unlock their true potential!",
	Footer: &discordgo.MessageEmbedFooter{
		Text: character.Name + " • Wuthering Waves",
	},
	Image: &discordgo.MessageEmbedImage{
		URL: "https://wuwamerch.com/wp-content/uploads/wuthering-waves-merch-aalto-badge-1.webp",
	},
}

	components := WuwaButtons(int(character.ID))

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

func (wh *WuwaHandlers) showCharacterAscension(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	character, err := wh.adapter.GetWuwaCharacter(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "⬆️ Ascension Materials — " + character.Name,
		Description: fmt.Sprintf("All required materials to fully ascend **%s**.", character.Name),
		Color:       0x6A5ACD,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf("https://your.repo/wuwa/thumbs/%s.png", strings.ToLower(character.Name)),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "🔹 Mob Drops",
				Value: fmt.Sprintf(
					"- %s ×4\n- %s ×12\n- %s ×12\n- %s ×4",
					character.Ascension.MobMaterial.UncommonName,
					character.Ascension.MobMaterial.RareName,
					character.Ascension.MobMaterial.EpicName,
					character.Ascension.MobMaterial.LegendaryName,
				),
				Inline: false,
			},
			{
				Name:   "🧠 Boss Drop",
				Value:  fmt.Sprintf("%s ×46", character.Ascension.BossMaterial),
				Inline: true,
			},
			{
				Name:   "🌿 Local Specialty",
				Value:  fmt.Sprintf("%s ×60", character.Ascension.LocalSpecialty),
				Inline: true,
			},
			{
				Name:   "💰 Shell Credits",
				Value:  "170,000",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: character.Name + " • Wuthering Waves",
		},
	}

	components := WuwaButtons(int(character.ID))

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

func (wh *WuwaHandlers) showCharacterWeapons(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	build, err := wh.adapter.GetWuwaBuild(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	var fields []*discordgo.MessageEmbedField

	for _, weapon := range build.Weapons {
		fieldValue := fmt.Sprintf(
			"🗡️ **Base ATK:** %d\n"+
			"📈 **Substat:** %s (+%.1f%%)\n"+
			"✨ **Passive:** %s",
			weapon.BaseATK,
			weapon.SubStat,
			weapon.SubValue,
			weapon.Passive,
		)

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  weapon.Name + fmt.Sprintf(" %v⭐", weapon.Rarity),
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

	components := WuwaButtons(int(build.Character.ID))

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

func (wh *WuwaHandlers) showCharacterTalents(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	character, err := wh.adapter.GetWuwaCharacter(uint(parsedID))

	if err != nil {
		slog.Error("Error fetching character info", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📘 Talent Materials — " + character.Name,
		Description: fmt.Sprintf("Resources required to max all five talents of **%s**.", character.Name),
		Color:       0xad44d9,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf("https://your.repo/wuwa/thumbs/%s.png", strings.ToLower(character.Name)),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "👤 Mob Drops — Howler Cores",
				Value: fmt.Sprintf(
					"- %s ×25\n- %s ×28\n- %s ×40\n- %s ×57",
					character.Talents.MobMaterial.UncommonName,
					character.Talents.MobMaterial.RareName,
					character.Talents.MobMaterial.EpicName,
					character.Talents.MobMaterial.LegendaryName,
				),
				Inline: false,
			},
			{
				Name: "📘 Dungeon Materials",
				Value: fmt.Sprintf(
					"- %s ×25\n- %s ×28\n- %s ×55\n- %s ×67",
					character.Talents.DungeonMaterial.UncommonName,
					character.Talents.DungeonMaterial.RareName,
					character.Talents.DungeonMaterial.EpicName,
					character.Talents.DungeonMaterial.LegendaryName,
				),
				Inline: false,
			},
			{
				Name:   "🔥 Weekly Boss Material",
				Value:  fmt.Sprintf("- %s ×26", character.Talents.BossMaterial),
				Inline: true,
			},
			{
				Name:   "💰 Shell Credits",
				Value:  "2,030,000",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: fmt.Sprintf("https://your.repo/wuwa/full/%s_talents.png", strings.ToLower(character.Name)),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: character.Name + " • Talent Level-Up Materials",
		},
	}

	components := WuwaButtons(int(character.ID))

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

func (wh *WuwaHandlers) showCharacterTeams(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return nil
}

func (wh *WuwaHandlers) showCharacterEchoes(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, id, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	parsedID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		slog.Error("Error parsing character ID", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	build, err := wh.adapter.GetWuwaBuild(uint(parsedID))
	if err != nil {
		slog.Error("Error fetching character build", "err", err)
		return discord.SendErrorMessage(s, i)
	}

	var fields []*discordgo.MessageEmbedField

	for _, echo := range build.Echoes {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   echo.Name,
			Value:  fmt.Sprintf("2-piece bonus: %s\n5-piece bonus: %s", echo.TwoPieceBonus, echo.FullSetBonus),
			Inline: false,
		})
	}

	stats := build.Stats
	fields = append(fields, &discordgo.MessageEmbedField{
		Name: "📊 Main Stat Suggestions",
		Value: fmt.Sprintf(
			"4-Cost Echo: **%s**\n3-Cost Echo: **%s**\n1-Cost Echo: **%s**",
			stats.FourCostEchoStat,
			stats.ThreeCostEchoStat,
			stats.OneCostEchoStat,
		),
		Inline: false,
	})

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "🔻 Substat Priority",
		Value:  stats.SubStatsPriority,
		Inline: false,
	})

	embed := &discordgo.MessageEmbed{
		Title:       build.Character.Name + " — Echo & Stat Guide",
		Description: fmt.Sprintf("Recommended Echo sets and stat priorities for **%s**.", build.Character.Name),
		Color:       0x3498DB,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "",
		},
		Fields: fields,
		Image: &discordgo.MessageEmbedImage{
			URL: "",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: build.Character.Name + " • Echo Build",
		},
	}

	components := WuwaButtons(int(build.Character.ID))

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

func (wh *WuwaHandlers) pagination(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, pageStr, _ := strings.Cut(i.MessageComponentData().CustomID, "_")
	page, _ := strconv.Atoi(pageStr)

	characters, err := wh.adapter.GetWuwaCharacters()

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
				CustomID: fmt.Sprintf("WuwaCharacter_%v", pageCharacters[j].ID),
			})
		}
		components = append(components, discordgo.ActionsRow{Components: row})
	}

	paginationRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "⬅ Previous",
				Style:    discordgo.SecondaryButton,
				CustomID: fmt.Sprintf("WuwaPagination_%v", page-1),
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
				CustomID: fmt.Sprintf("WuwaPagination_%v", page+1),
				Disabled: end >= len(characters),
			},
		},
	}
	components = append(components, paginationRow)

	embed := &discordgo.MessageEmbed{
		Title:       "🌊 Wuthering Waves Characters",
		Description: "Explore all resonators in Wuthering Waves, including their rarity, attributes, and factions. Tap on a resonator to see detailed info, echo setups, weapon builds, and progression materials.",
		Color:       0x6A5ACD,
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

func (wh *WuwaHandlers) AddWuwaHandlers(handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	handlers["WuwaCharacters"] = wh.showCharacters
	handlers["WuwaPagination"] = wh.pagination
	handlers["WuwaCharacter"] = wh.showCharacterDetail
	handlers["WuwaAscension"] = wh.showCharacterAscension
	handlers["WuwaEchoes"] = wh.showCharacterEchoes
	handlers["WuwaTalents"] = wh.showCharacterTalents
	handlers["WuwaWeapons"] = wh.showCharacterWeapons
	handlers["WuwaTeams"] = wh.showCharacterTeams
}
