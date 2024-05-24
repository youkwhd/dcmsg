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
        fmt.Println("R2-ERR:", err)
        os.Exit(1)
    }

    R2Bot.SetDevelopmentMode(os.Getenv("R2DEV") == "1")

    R2Bot.OpenSession()
    R2Bot.RegisterCommands(os.Getenv("R2AID"), "")
    R2Bot.AddMessageReactionHandler()
    R2Bot.AddCommandHandler()

    defer R2Bot.CloseSession()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    fmt.Println("R2-BOT: Press CTRL + C to stop")
    <-stop
}
