package commands

import (
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
}
