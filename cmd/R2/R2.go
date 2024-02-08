package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

const GUILD_GLOBAL string = ""

func main() {
	dg, err := discordgo.New("Bot " + "")

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		fmt.Println(i.ApplicationCommandData().Name)
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == dg.State.User.ID {
			return
		}

		dg.ChannelMessageSendReply(m.ChannelID, "Your message: " + m.Content, m.Reference())
	})

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Press Ctrl+C to exit")
	<-stop
}
