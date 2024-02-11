package commands

import (
	"strings"

	db "R2/internal/db/json"

	"github.com/bwmarrin/discordgo"
)

type CommandHandlerFn func(botSession *discordgo.Session, i *discordgo.InteractionCreate)

type Command struct {
	Information *discordgo.ApplicationCommand
	Handler CommandHandlerFn
}

var /* const */ COMMANDS = [...]Command{
	{
		Information: &discordgo.ApplicationCommand{
			Name: "ping",
			Description: "<test> returns back pong",
		},
		Handler: func(botSession *discordgo.Session, i *discordgo.InteractionCreate) {
			botSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
	},
	{
		Information: &discordgo.ApplicationCommand{
			Name: "react",
			Description: "Add a new reaction role",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message_link",
					Description: "Message link",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "Role",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "emoji",
					Description: "Emoji",
					Required:    true,
				},
			},
		},
		Handler: func(botSession *discordgo.Session, i *discordgo.InteractionCreate) {
			_args := i.ApplicationCommandData().Options

			args := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
			for _, arg := range _args {
				args[arg.Name] = arg
			}

			messageLink := args["message_link"].Value.(string)
			// TODO: update go to ver 1.20, use CutPrefix
			_, messageLink, _ = strings.Cut(messageLink, "https://discord.com/channels/")
			_, messageLink, _ = strings.Cut(messageLink, "/")
			channelID, messageLink, _ := strings.Cut(messageLink, "/")

			messageID := messageLink
			role := args["role"].Value.(string)
			emoji := args["emoji"].Value.(string)

			db.SaveMessage(channelID, messageID, role, emoji)
			botSession.MessageReactionAdd(channelID, messageID, emoji)

			botSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "OK",
				},
			})

			botSession.InteractionResponseDelete(i.Interaction)
		},
	},
}
