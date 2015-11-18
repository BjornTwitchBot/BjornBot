package main

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

func rank(command *bot.Cmd) (msg string, err error) {
	rankcs(command)
	ranklol(command)
	return "", nil
}

func init() {
	bot.RegisterCommand(
		"rank",
		"Sends information about rankings to the channel.",
		"",
		rank)
}
