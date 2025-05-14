package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
    bot, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

    if err != nil {
        fmt.Println("error initializing bot:", err)
        os.Exit(1)
    }

    bot.Open()
    defer bot.Close()

	reader := bufio.NewReader(os.Stdin)

	channel := os.Getenv("CHANNEL")
	if channel == "" {
		channel, err = reader.ReadString('\n')

		if err != nil {
			fmt.Println("error reading channel:", err)
			os.Exit(1)
		}

		channel = channel[:len(channel)-1]
	}

	for {
		fmt.Print("> ")

		msg, err := reader.ReadString('\n')

		if err != nil {
			return
		}

		msg = msg[:len(msg)-1]

		_, err = bot.ChannelMessageSend(channel, msg)

		if err != nil {
			fmt.Println("error sending message:", err)
		}
	}
}
