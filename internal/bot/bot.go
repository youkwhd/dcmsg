package bot

import (
    "fmt"
    "strings"

    "R2/internal/bot/commands"
    "R2/internal/bot/guild"
    db "R2/internal/db/json"
    "R2/internal/message"

    "github.com/bwmarrin/discordgo"
)

type R2Bot struct {
    session *discordgo.Session
    devmode bool
}

func New(token string) (bot R2Bot, err error) {
    token = strings.Trim(token, " ")
    if token == "" {
        return R2Bot{}, fmt.Errorf("bot token cannot be empty")
    }

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
        _, err := bot.session.ApplicationCommandCreate(bot.session.State.User.ID, guild.GetGuild(bot.devmode), cmd.Information)

        if err != nil {
            fmt.Printf("ERR: Registering \"%v\" command\n", cmd.Information.Name)
        }
    }
}

func (bot *R2Bot) DeregisterInteractionCommands() {
    rcommands, err := bot.session.ApplicationCommands(bot.session.State.User.ID, guild.GetGuild(bot.devmode))

    if err != nil {
        fmt.Printf("ERR: Retrieving commands\n")
        return
    }

    for _, cmd := range rcommands {
        bot.session.ApplicationCommandDelete(bot.session.State.User.ID, cmd.ID, guild.GetGuild(bot.devmode))
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
                break
            }
        }
    })
}

func (bot *R2Bot) AddMessageReactionHandler() {
    bot.session.AddHandler(func(botSession *discordgo.Session, r *discordgo.MessageReactionAdd) {
        if r.UserID == bot.session.State.User.ID {
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

        bot.session.GuildMemberRoleAdd(r.GuildID, r.UserID, string(role))
    })

    bot.session.AddHandler(func(botSession *discordgo.Session, r *discordgo.MessageReactionRemove) {
        if r.UserID == bot.session.State.User.ID {
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

        bot.session.GuildMemberRoleRemove(r.GuildID, r.UserID, string(role))
    })

    bot.session.AddHandler(func(botSession *discordgo.Session, m *discordgo.MessageDelete) {
        db.RemoveMessage(m.ID)
    })
}
