package wuwa

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func WuwaButtons(id int) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Back",
					CustomID: "WuwaCharacters",
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "↩️",
					},
				},
				discordgo.Button{
					Label:    "Ascension",
					CustomID: fmt.Sprintf("WuwaAscension_%v", id),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "⬆️",
					},
				},
				discordgo.Button{
					Label:    "Talents",
					CustomID: fmt.Sprintf("WuwaTalents_%v", id),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "📘",
					},
				},
				discordgo.Button{
					Label:    "Echoes",
					CustomID: fmt.Sprintf("WuwaEchoes_%v", id),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "🌀",
					},
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Weapons",
					CustomID: fmt.Sprintf("WuwaWeapons_%v", id),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "⚔️",
					},
				},
				discordgo.Button{
					Label:    "Team Synergy",
					CustomID: fmt.Sprintf("WuwaTeams_%v", id),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "👥",
					},
				},
				discordgo.Button{
					Label:    "Overview",
					CustomID: fmt.Sprintf("WuwaOverview_%v", id),
					Style:    discordgo.SecondaryButton,
					Emoji: &discordgo.ComponentEmoji{
						Name: "📖",
					},
				},
			},
		},
	}
}
