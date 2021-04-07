package main

import (
	"github.com/go-chat-bot/bot"
)

func sens(command *bot.Cmd) (msg string, err error) {
	msg = "0.95 in game @ 1000 DPI"
	return
}

func init() {
	bot.RegisterCommand(
		"sens",
		"Sends information about mouse sensitivity to the channel.",
		"",
		sens)
}
