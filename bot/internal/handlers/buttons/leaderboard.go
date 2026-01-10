package buttons

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func LeaderboardButtons(currentPage, totalPages int) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "⏮",
			Style:    discordgo.PrimaryButton,
			CustomID: "LeaderboardFirst",
			Disabled: currentPage == 0,
		},
		discordgo.Button{
			Label:    "◀",
			Style:    discordgo.SecondaryButton,
			CustomID: fmt.Sprintf("LeaderboardPage_%d", currentPage-1),
			Disabled: currentPage < 1,
		},
		discordgo.Button{
			Label:    fmt.Sprintf("Page %d/%d", currentPage, totalPages),
			Style:    discordgo.SecondaryButton,
			CustomID: "page_current",
			Disabled: true,
		},
		discordgo.Button{
			Label:    "▶",
			Style:    discordgo.SecondaryButton,
			CustomID: fmt.Sprintf("LeaderboardPage_%d", currentPage+1),
			Disabled: currentPage > totalPages,
		},
		discordgo.Button{
			Label:    "⏭",
			Style:    discordgo.PrimaryButton,
			CustomID: fmt.Sprintf("LeaderboardLast_%d", totalPages),
			Disabled: currentPage == totalPages,
		},
	}
}
