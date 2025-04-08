package genshin

import (
	dtoDiscord "bot/internal/dto/discord"
	"bot/internal/discord"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	characters = []string{"Varesa", "Iansan", "Diluc"}
)

func showGenshinCharacters(data dtoDiscord.HandlerData) error {
	var buttons []discordgo.MessageComponent

	for _, char := range characters {
		buttons = append(buttons, &discordgo.Button{
			Label:    char,
			Style:    discordgo.PrimaryButton,
			CustomID: fmt.Sprintf("GenshinCharacter_%s", char),
		})
	}

	paginationRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.Button{
				Label:    "⬅ Previous",
				Style:    discordgo.SecondaryButton,
				CustomID: "genshinPaginationBack",

			},
			&discordgo.Button{
				Label:    "Main menu",
				Style:    discordgo.SecondaryButton,
				CustomID: "GachasMenu",

			},
			&discordgo.Button{
				Label:    "Next ➡",
				Style:    discordgo.SecondaryButton,
				CustomID: "genshinPaginationNext",

			},
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🎮 Genshin Impact Characters",
		Description: "Discover all available characters in Genshin Impact, including their rarity, elements, and regions. Tap on a character to view detailed info, ascension materials, and builds.",
		Color:       0x3498DB,
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{Components: buttons},
		paginationRow,
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

func showCharacterInfo(data dtoDiscord.HandlerData) error {
	embed := discordgo.MessageEmbed{
		Title:       "🌩️ Raiden Shogun", 
		Description: "“Eternity is in the eyes of the beholder.”",
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "⭐ Rarity",
				Value:  "5-Star",
				Inline: true,
			},
			{
				Name:   "⚔️ Weapon",
				Value:  "Polearm",
				Inline: true,
			},
			{
				Name:   "🌩️ Element",
				Value:  "Electro",
				Inline: true,
			},
			{
				Name: "🎙️ Voice Actors",
				Value: "**EN:** Anne Yatco\n**JP:** Miyuki Sawashiro",
				Inline: true,
			},
			{
				Name:   "🗺️ Region",
				Value:  "Inazuma",
				Inline: true,
			},
			{
				Name: "📅 Released",
				Value: "Version 2.1 – September 1, 2021",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Raiden Shogun • Genshin Impact",
		},
	}

	components := genshinButtons("shogun")

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{&embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil
}