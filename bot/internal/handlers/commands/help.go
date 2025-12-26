package commands

import (
	"bot/internal/utils"
	"log/slog"

	"bot/internal/discord"
	"bot/internal/handlers/buttons"

	"github.com/bwmarrin/discordgo"
)

var (
	PageSize   = 8
	COMMANDS   = discord.CommandsList
	TotalPages = len(COMMANDS) % PageSize
)

func helpFields(currentPage int) []*discordgo.MessageEmbedField {
	start := currentPage * PageSize
	end := start + PageSize

	if start > len(COMMANDS) {
		return []*discordgo.MessageEmbedField{}
	}

	if end > len(COMMANDS) {
		end = len(COMMANDS)
	}

	fields := make([]*discordgo.MessageEmbedField, 0, end-start)

	for _, cmd := range COMMANDS[start:end] {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  cmd.Name,
			Value: cmd.Description,
		})
	}

	return fields
}

func Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "🌿 Frieren Bot 🌿",
		Description: "Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:",
		Color:       0x2ECC71,
		Fields:      helpFields(0),
	}

	btn := buttons.HelpPaginationButtons(0, TotalPages)

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{helpEmbed},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: btn,
				},
			},
		},
	}

	utils.Respond(s, i, resp)
}

func HelpPaginate(s *discordgo.Session, i *discordgo.InteractionCreate, page int) {
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "🌿 Frieren Bot 🌿",
		Description: "Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:",
		Color:       0x2ECC71,
		Fields:      helpFields(page),
	}

	utils.Acknowledge(s, i)

	btn := buttons.HelpPaginationButtons(page, TotalPages)

	edit := &discordgo.MessageEdit{
		ID:      i.Message.ID,
		Channel: i.Message.ChannelID,
		Embeds:  &[]*discordgo.MessageEmbed{helpEmbed},
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: btn,
			},
		},
	}

	_, err := s.ChannelMessageEditComplex(edit)

	if err != nil {
		slog.Error(err.Error())
	}
}
