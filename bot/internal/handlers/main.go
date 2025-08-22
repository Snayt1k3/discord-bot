package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func HelpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "🌿 Frieren Bot - Traces of Music 🌿",
		Description: "Time passes, but music stays with us. If you wish to fill the silence, here’s what you can do:",
		Color:       0x2ECC71,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/2f/eb/71/2feb71b7fb35c35886b87324b6cef144.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "🎼 Commands to Guide the Melody:",
				Value: "`/play <song_name/link>` – Let the music flow, one song at a time.\n" +
					"`/pause` – Even melodies need a moment of rest.\n" +
					"`/resume` – Continue where you left off, like an old journey resumed.\n" +
					"`/stop` – Bring the music to a quiet end, clearing all that remains.\n" +
					"`/skip` – Move past this tune, towards the next story in sound.",
			},
			{
				Name:  "📖 Knowledge in the Wind:",
				Value: "`/help` – If you have forgotten, let this guide you once more.",
			},
		},
	}

	s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{helpEmbed},
			},
		},
	)
	return nil
}

func SettingsMenuHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// Создаем Embed для меню настроек
	settingsEmbed := &discordgo.MessageEmbed{
		Title:       "⚙ Bot Settings ⚙",
		Description: "Выберите действие ниже:",
		Color:       0x2ECC71, // зеленый цвет
		Image: &discordgo.MessageEmbedImage{
			URL: "", // если нужно, можно убрать
		},
	}

	// Создаем кнопки
	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Настройка ролей",
					Style:    discordgo.PrimaryButton,
					CustomID: "settings_roles_reactions",
				},
				discordgo.Button{
					Label:    "Настройка приветствия",
					Style:    discordgo.PrimaryButton,
					CustomID: "settings_welcome",
				},
				
			},
		},
	}

	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{settingsEmbed},
				Components: components,
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}