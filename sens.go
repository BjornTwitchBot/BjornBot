package main

import (
	"fmt"
	"github.com/fabioxgn/go-bot"
)

func sens(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("1.0 in game @ 1000 DPI")
	return
}

func init() {
	bot.RegisterCommand(
		"sens",
		"Sends information about mouse sensitivity to the channel.",
		"",
		sens)
}
