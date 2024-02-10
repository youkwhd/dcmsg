package bot

import (
	"R2/internal/bot/commands"
	"R2/internal/bot/guild"
	"R2/internal/db"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type R2Bot struct {
	session *discordgo.Session
	devmode bool
}

func New(token string) (bot R2Bot, err error) {
	session, err := discordgo.New("Bot " + token)

	return R2Bot{
		session: session,
	}, err
}

func (bot *R2Bot) OpenSession() {
	bot.session.Open()
}

func (bot *R2Bot) CloseSession() {
	bot.session.Close()
}

func (bot *R2Bot) SetDevelopmentMode(mode bool) {
	bot.devmode = mode
}

func (bot *R2Bot) RegisterInteractionCommands() {
	for _, cmd := range commands.COMMANDS {
		var guildTarget string = guild.GUILD_GLOBAL
		if bot.devmode {
			guildTarget = guild.GUILD_DEV
		}

		_, err := bot.session.ApplicationCommandCreate(bot.session.State.User.ID, guildTarget, cmd.Information)

		if err != nil {
			fmt.Errorf("ERR: Registering \"%v\" command", cmd.Information.Name)
		}
	}

}

func (bot *R2Bot) DeregisterInteractionCommands() {
	var guildTarget string = guild.GUILD_GLOBAL
	if bot.devmode {
		guildTarget = guild.GUILD_DEV
	}

	rcommands, err := bot.session.ApplicationCommands(bot.session.State.User.ID, guildTarget)

	if err != nil {
		fmt.Errorf("ERR: Retrieving commands")
		return
	}

	for _, cmd := range rcommands {
		bot.session.ApplicationCommandDelete(bot.session.State.User.ID, cmd.ID, guildTarget)
	}
}

func (bot *R2Bot) AddInteractionCommandHandler() {
	bot.session.AddHandler(func(botSession *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		interactionName := data.Name

		// TODO: maybe consider using hashmap
		for _, cmd := range commands.COMMANDS {
			if interactionName == cmd.Information.Name {
				cmd.Handler(botSession, i)
			}
		}
	})
}

func (bot *R2Bot) AddMessageReactionHandler() {
	bot.session.AddHandler(func(botSession *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if r.UserID == bot.session.State.User.ID {
			return
		}

		roleReactionMessage, found := db.Messages[db.MessageID(r.MessageID)]
		if !found {
			return
		}

		role, found := roleReactionMessage.Reactions[db.Emoji(r.Emoji.Name)]
		if !found {
			return
		}

		bot.session.GuildMemberRoleAdd(r.GuildID, r.UserID, string(role))
	})

	bot.session.AddHandler(func(botSession *discordgo.Session, r *discordgo.MessageReactionRemove) {
		if r.UserID == bot.session.State.User.ID {
			return
		}

		roleReactionMessage, found := db.Messages[db.MessageID(r.MessageID)]
		if !found {
			return
		}

		role, found := roleReactionMessage.Reactions[db.Emoji(r.Emoji.Name)]
		if !found {
			return
		}

		bot.session.GuildMemberRoleRemove(r.GuildID, r.UserID, string(role))
	})
}
