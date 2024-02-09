package main

import (
	"fmt"
	"os"
	"os/signal"

	R2 "R2/internal/bot"
)

func main() {
	R2Bot, err := R2.New(os.Getenv("R2TOK"))

	if err != nil {
		fmt.Println(err)
		return
	}

	R2Bot.OpenSession()
	R2Bot.RegisterInteractionCommands()
	R2Bot.AddInteractionCommandHandler()

	defer R2Bot.DeregisterInteractionCommands()
	defer R2Bot.CloseSession()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("R2-BOT: Press CTRL + C to stop")
	<-stop
}
