package handlers

import (
	"bot/internal/handlers/commands"
	"bot/internal/handlers/commands/preferences"
	"bot/internal/http"
	"log/slog"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Container struct {
	Http *http.Container
}

func NewContainer(http *http.Container) *Container {
	return &Container{
		Http: http,
	}
}

func (cd *Container) Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.Help(s, i)
}

func (cd *Container) Rank(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := commands.Rank(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) WelcomeMsg(s *discordgo.Session, i *discordgo.InteractionCreate) {
	action := i.ApplicationCommandData().Options[0].StringValue()

	var err error

	switch action {
	case "Add":
		err = preferences.AddWelcomeMessage(cd.Http, s, i)
	case "Remove":
		err = preferences.DeleteWelcomeMessage(cd.Http, s, i)
	}

	if err != nil {
		slog.Error(err.Error())
	}
}


func (cd *Container) SetWelcomeChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.SetWelcomeChnl(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) Menu(s *discordgo.Session, i *discordgo.InteractionCreate) {
	menuType := i.ApplicationCommandData().Options[0].StringValue()
	var err error

	switch menuType {
		case "Welcome":
			err = commands.WelcomeSettings(cd.Http, s, i)
		case "RolesReactions":
			err = commands.RolesSettings(cd.Http, s, i)
		case "AutoMode":
			err = commands.ModerationSettings(cd.Http, s, i)
		case "Logging":
			err = commands.LoggingSettings(cd.Http, s, i)
	}
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RemoveLogChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.RemoveLoggingChnl(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) LogChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ApplicationCommandData().Options[0]
	var err error

	if channel != nil {
		err = preferences.AddLoggingChnl(cd.Http, s, i)
	} else {
		err = preferences.RemoveLoggingChnl(cd.Http, s, i)
	}

	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) ToggleFeature(s *discordgo.Session, i *discordgo.InteractionCreate) {
	feature := i.ApplicationCommandData().Options[0].StringValue()
	var err error

	switch feature { 
	case "automod":
		err = preferences.ToggleModeration(cd.Http, s, i)
	case "logging":
		err = preferences.ToggleLogging(cd.Http, s, i)
	}

	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) BannedWord(s *discordgo.Session, i *discordgo.InteractionCreate) {
	action := i.ApplicationCommandData().Options[0].StringValue()
	var err error

	switch action {
	case "Add":
		err = preferences.AddBannedWord(cd.Http, s, i)
	case "Remove":
		err = preferences.RemoveBannedWord(cd.Http, s, i)
	}

	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) AntiLink(s *discordgo.Session, i *discordgo.InteractionCreate) {
	action := i.ApplicationCommandData().Options[0].StringValue()
	var err error

	switch action {
	case "Add":
		err = preferences.AddAntiLinkChnl(cd.Http, s, i)
	case "Remove":
		err = preferences.RemoveAntiLinkChnl(cd.Http, s, i)
	}

	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RemoveAntiLinkChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.RemoveAntiLinkChnl(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) AntiCaps(s *discordgo.Session, i *discordgo.InteractionCreate) {
	action := i.ApplicationCommandData().Options[0].StringValue()
	var err error

	switch action {
	case "Add":
		err = preferences.AddCapsLockChnl(cd.Http, s, i)
	case "Remove":
		err = preferences.RemoveCapsLockChnl(cd.Http, s, i)
	}
	
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RemoveCapsLockChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.RemoveCapsLockChnl(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) AddRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AddRole(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RemoveRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.RemoveRole(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) SetRoleMsg(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.SetRolesMsg(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) HelpPagination(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, pageStr, _ := strings.Cut(i.MessageComponentData().CustomID, "_")

	page, _ := strconv.Atoi(pageStr)
	commands.HelpPaginate(s, i, page)
}

func (cd *Container) HelpPaginationFirst(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.HelpPaginate(s, i, 0)
}

func (cd *Container) HelpPaginationLast(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.HelpPaginate(s, i, commands.TotalPages)
}
