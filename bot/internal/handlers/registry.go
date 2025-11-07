package handlers

import (
	"bot/internal/handlers/commands"
	"bot/internal/handlers/commands/preferences"
	"bot/internal/http"
	"log/slog"

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

func (cd *Container) Menu(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.Menu(s, i)
}

func (cd *Container) InteractionProfile(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := commands.InteractionProfile(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) WelcomePreferences(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.WelcomeSettings(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) AddWelcomeMsg(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AddWelcomeMessage(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RemoveWelcomeMsg(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.DeleteWelcomeMessage(cd.Http, s, i)
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

func (cd *Container) LoggingSettings(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.LogSettings(cd.Http, s, i)
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

func (cd *Container) AddLogChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AddLoggingChnl(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) ToggleLog(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.ToggleLogging(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) ToggleModeration(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.ToggleModeration(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) AddBannedWord(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AddBannedWord(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RemoveBannedWord(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.RemoveBannedWord(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) AddAntiLinkChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AddAntiLinkChnl(cd.Http, s, i)
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

func (cd *Container) AddCapsLockChnl(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AddCapsLockChnl(cd.Http, s, i)
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

func (cd *Container) ModerationSettings(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.AutoModeSettings(cd.Http, s, i)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (cd *Container) RolesSettings(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := preferences.RolesSettings(cd.Http, s, i)
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
