package main

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

func rank(command *bot.Cmd) (msg string, err error) {
	// We are registering this command only to show up in !help
	// The actual text is handled elsewhere (in the bot.rank function)
	msg = ""
	return
}

func init() {
	bot.RegisterCommand(
		"rank",
		"Sends information about rankings to the channel.",
		"",
		rank)
}
