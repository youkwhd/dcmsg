package main

import (
    "fmt"
    "os"
    "os/signal"

    bot "rrolls/internal/bot"
)

func main() {
    r2, err := bot.New(os.Getenv("R2TOK"))

    if err != nil {
        fmt.Println("ERR:", err)
        os.Exit(1)
    }

    r2.SetDevelopmentMode(os.Getenv("R2DEV") == "1")

    r2.OpenSession()
    r2.RegisterCommands(os.Getenv("R2AID"), "")
    r2.AddMessageReactionHandler()
    r2.AddCommandHandler()

    defer r2.CloseSession()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    fmt.Println("Press CTRL + C to stop")
    <-stop
}
