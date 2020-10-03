package main

import (
	"fmt"
	"os"

	bot "github.com/kinbiko/twitch-bot"
)

func main() {
	if err := bot.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
