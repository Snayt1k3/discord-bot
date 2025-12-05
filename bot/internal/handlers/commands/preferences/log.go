package preferences

import (
	"bot/internal/http"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

func ToggleLogging(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	err = http.Log.Toggle(i.GuildID, !settings.Log.Enabled)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Logging has been %s successfully!", map[bool]string{true: "enabled", false: "disabled"}[!settings.Log.Enabled]),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func AddLoggingChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	channelID := i.ApplicationCommandData().Options[0].ChannelValue(s).ID
	err := http.Log.AddChannel(i.GuildID, channelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Logging channel has been set successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}

func RemoveLoggingChnl(http *http.Container, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if !utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
		utils.SendNoPermissionMessage(s, i)
		return nil
	}

	settings, err := http.Settings.Get(i.GuildID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	err = http.Log.RemoveChannel(i.GuildID, settings.Log.ChannelID)

	if err != nil {
		utils.SendErrorMessage(s, i)
		return err
	}

	utils.Respond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Logging channel has been removed successfully!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	return nil
}
