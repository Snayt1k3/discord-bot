package discord

import (
	"github.com/bwmarrin/discordgo"
)

var CommandsList = []*discordgo.ApplicationCommand{
	{
		Name:        "menu",
		Description: "View server settings.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "type",
				Description: "Menu to view",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Roles/Reactions",
						Value: "RolesReactions",
					},
					{
						Name:  "Welcome",
						Value: "Welcome",
					},
					{
						Name:  "AutoMode",
						Value: "AutoMode",
					},
					{
						Name:  "Logging",
						Value: "Logging",
					},
				},
			},
		},
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
		Name:        "welcome-msg",
		Description: "Add/Remove a welcome message. Available variables: {username}",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Action to perform",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Add",
						Value: "Add",
					},
					{
						Name:  "Remove",
						Value: "Remove",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "msg",
				Description: "The welcome message to add",
				Required:    true,
			},
		},
	},
	{
		Name:        "automod-bannedword",
		Description: "Add/Remove a banned word to automod.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Action to perform",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Add",
						Value: "Add",
					},
					{
						Name:  "Remove",
						Value: "Remove",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "word",
				Description: "The word to ban",
				Required:    true,
			},
		},
	},
	{
		Name:        "automod-antilink",
		Description: "Enable/Disable Antilink filter in current chat.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Action to perform",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Add",
						Value: "Add",
					},
					{
						Name:  "Remove",
						Value: "Remove",
					},
				},
			},
		},
	},
	{
		Name:        "automod-anticaps",
		Description: "Enable/Disable AntiCaps filter in current chat.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Action to perform",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Add",
						Value: "Add",
					},
					{
						Name:  "Remove",
						Value: "Remove",
					},
				},
			},
		},
	},
	{
		Name:        "toggle",
		Description: "Toggle automod or log - on/off.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "feature",
				Description: "Feature to toggle",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Automod",
						Value: "automod",
					},
					{
						Name:  "Logging",
						Value: "logging",
					},
				},
			},
		},
	},
	{
		Name:        "log-edit",
		Description: "Edit the logging sttings",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Action to perform",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Remove",
						Value: "remove",
					},
					{
						Name:  "Add",
						Value: "add",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "The channel to send logs",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "event",
				Description: "The type of log",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Message Delete",
						Value: "1",
					},
					{
						Name:  "Message Edit",
						Value: "2",
					},
					{
						Name:  "User Join",
						Value: "3",
					},
					{
						Name:  "User Leave",
						Value: "4",
					},
					{
						Name:  "Create Invite",
						Value: "5",
					},
					{
						Name:  "Join Channel",
						Value: "6",
					},
					{
						Name:  "Leave Channel",
						Value: "7",
					},
					{
						Name:  "Move Channel",
						Value: "8",
					},
				},
			},
		},
	},
	{
		Name:        "rank",
		Description: "View your rank and experience points.",
	},
	{
		Name:        "leaderboard",
		Description: "View the server leaderboard.",
	},
}
