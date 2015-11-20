package main

import (
	"fmt"
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

func sens(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("0.95 in game @ 1000 DPI")
	return
}

func init() {
	bot.RegisterCommand(
		"sens",
		"Sends information about mouse sensitivity to the channel.",
		"",
		sens)
}
