package r2

import (
	"fmt"
	"strings"

	"rrolls/internal/bot/commands"
	db "rrolls/internal/db/json"
	"rrolls/internal/message"

	"github.com/bwmarrin/discordgo"
)

type rrolls struct {
    session *discordgo.Session
    devmode bool
}

func New(token string) (r2 rrolls, err error) {
    token = strings.Trim(token, " ")
    if token == "" {
        return rrolls{}, fmt.Errorf("bot token cannot be empty")
    }

    session, err := discordgo.New("Bot " + token)

    return rrolls{
        session: session,
    }, err
}

func (r2 *rrolls) OpenSession() {
    r2.session.Open()
}

func (r2 *rrolls) CloseSession() {
    r2.session.Close()
}

func (r2 *rrolls) SetDevelopmentMode(mode bool) {
    r2.devmode = mode
}

func (r2 *rrolls) RegisterCommands(appId, guildId string) {
    for _, cmd := range commands.COMMANDS {
		_, err := r2.session.ApplicationCommandCreate(appId, guildId, cmd.Information)
		if err != nil {
			fmt.Printf("ERR: Registering command '%s'\n", cmd.Information.Name)
			// TODO: exit here
		}
    }
}

func (r2 *rrolls) AddCommandHandler() {
    r2.session.AddHandler(func(botSession *discordgo.Session, i *discordgo.InteractionCreate) {
        data := i.ApplicationCommandData()
        interactionName := data.Name

        // TODO: maybe consider using hashmap
        for _, cmd := range commands.COMMANDS {
            if interactionName == cmd.Information.Name {
                cmd.Handler(botSession, i)
                break
            }
        }
    })
}

func (r2 *rrolls) AddMessageReactionHandler() {
    r2.session.AddHandler(func(botSession *discordgo.Session, r *discordgo.MessageReactionAdd) {
        if r.UserID == r2.session.State.User.ID {
            return
        }

        roleReactionMessage, found := db.GetMessage(r.MessageID)
        if !found {
            return
        }

        role, found := roleReactionMessage.Reactions[message.Emoji(r.Emoji.Name)]
        if !found {
            return
        }

        r2.session.GuildMemberRoleAdd(r.GuildID, r.UserID, string(role))
    })

    r2.session.AddHandler(func(botSession *discordgo.Session, r *discordgo.MessageReactionRemove) {
        if r.UserID == r2.session.State.User.ID {
            return
        }

        roleReactionMessage, found := db.GetMessage(r.MessageID)
        if !found {
            return
        }

        role, found := roleReactionMessage.Reactions[message.Emoji(r.Emoji.Name)]
        if !found {
            return
        }

        r2.session.GuildMemberRoleRemove(r.GuildID, r.UserID, string(role))
    })

    r2.session.AddHandler(func(botSession *discordgo.Session, m *discordgo.MessageDelete) {
        db.RemoveMessage(m.ID)
    })
}
