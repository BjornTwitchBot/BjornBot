package main

import (
	"fmt"
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

func sens(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("2.1 in game @ 400 DPI")
	return
}

func init() {
	bot.RegisterCommand(
		"sens",
		"Sends information about mouse sensitivity to the channel.",
		"",
		sens)
}
