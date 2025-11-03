package discord

import (
	"github.com/bwmarrin/discordgo"
)

var CommandsList = []*discordgo.ApplicationCommand{
	{
		Name:        "settings",
		Description: "View server settings.",
	},
	{
		Name:        "rr-add",
		Description: "Add a role.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role",
				Description: "The role to add",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "emoji",
				Description: "The emoji to use",
				Required:    true,
			},
		},
	},
	{
		Name:        "rr-remove",
		Description: "Remove a role.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role",
				Description: "The role to remove",
				Required:    true,
			},
		},
	},
	{
		Name:        "rr-message",
		Description: "Set the message ID for role reactions.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message_id",
				Description: "The message ID",
				Required:    true,
			},
		},
	},
	{
		Name:        "welcome-chnl",
		Description: "Set the welcome channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "The channel to send welcome messages",
				Required:    true,
			},
		},
	},
	{
		Name:        "welcomemsg-add",
		Description: "Add a welcome message. Variables: {username}",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "msg",
				Description: "The welcome message to add",
				Required:    true,
			},
		},
	},
	{
		Name:        "welcomemsg-remove",
		Description: "Remove a welcome message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "msg",
				Description: "The welcome message to remove",
				Required:    true,
			},
		},
	},
	{
		Name:        "automod-bw-add",
		Description: "Add a banned word to automod.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "word",
				Description: "The word to ban",
				Required:    true,
			},
		},
	},
	{
		Name:        "automod-bw-rm",
		Description: "Remove a banned word from automod.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "word",
				Description: "The word to ban",
				Required:    true,
			},
		},
	},
	{
		Name:        "automod-al-enable",
		Description: "Enable Antilink filter in current chat.",
	},
	{
		Name:        "automod-al-disable",
		Description: "Disable Antilink filter in current chat.",
	},
	{
		Name:        "automod-ac-enable",
		Description: "Disable AntiCaps filter in current chat.",
	},
	{
		Name:        "automod-ac-disable",
		Description: "Disable AntiCaps filter in current chat.",
	},
	{
		Name:        "automod-toggle",
		Description: "Toggle automod on/off.",
	},
	{
		Name:        "log-toggle",
		Description: "Toggle logging on/off.",
	},
	{
		Name:        "log-chnl",
		Description: "Set the logging channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "The channel to send logs",
				Required:    true,
			},
		},
	},
	{
		Name:        "log-chnl-rm",
		Description: "Remove the logging channel",
	},
	{
		Name:        "iprofile",
		Description: "View your profile",
	},
}
