package main

import (
	"fmt"
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

func ranklol(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("League of Legends - Gold V")
	return
}

func init() {
	bot.RegisterCommand(
		"rank",
		"Sends information about rankings to the channel.",
		"",
		ranklol)
}
