package genshin

import (
	"bot/internal/discord"
	dtoDiscord "bot/internal/dto/discord"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func genshinPagination(data dtoDiscord.HandlerData) error {
	_, pageStr, _ := strings.Cut(data.Event.MessageComponentData().CustomID, "_")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
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
				Label:    pageCharacters[j],
				Style:    discordgo.PrimaryButton,
				CustomID: fmt.Sprintf("GenshinCharacter_%s", pageCharacters[j]),
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

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil
}
