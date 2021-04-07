package main

import (
	"github.com/go-chat-bot/bot"
)

func rank(command *bot.Cmd) (msg string, err error) {
	// We are registering this command only to show up in !help
	// The actual text is handled elsewhere (in the bot.rank function)
	msg = "Silver something, I don't really play CSGO anymore"
	return
}

func init() {
	bot.RegisterCommand(
		"rank",
		"Sends information about rankings to the channel.",
		"",
		rank)
}
