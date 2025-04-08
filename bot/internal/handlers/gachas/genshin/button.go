package genshin

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func genshinButtons(name string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Back",
					CustomID: "GenshinCharacters",
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "↩️",
					},
				},
				discordgo.Button{
					Label:    "Ascension",
					CustomID: fmt.Sprintf("GenshinAsc_%s", name),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "⬆️",
					},
				},
				discordgo.Button{
					Label:    "Talents",
					CustomID: fmt.Sprintf("GenshinTalents_%s", name),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "📘",
					},
				},
				discordgo.Button{
					Label:    "Artifacts",
					CustomID: fmt.Sprintf("GenshinArtifacts_%s", name),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "🛠️",
					},
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Weapons",
					CustomID: fmt.Sprintf("GenshinWeapons_%s", name),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "⚔️",
					},
				},
				discordgo.Button{
					Label:    "Teams",
					CustomID: fmt.Sprintf("GenshinTeams_%s", name),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "👥",
					},
				},
				discordgo.Button{
					Label:    "Overview",
					CustomID: fmt.Sprintf("GenshinOverview_%s", name),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "📖",
					},
				},
			},
		},
	}
}