package main

import (
	"github.com/go-chat-bot/bot"
)

func donate(command *bot.Cmd) (msg string, err error) {
	msg = "https://www.extra-life.org/index.cfm?fuseaction=donorDrive.participant&participantID=387590"
	return
}

func init() {
	bot.RegisterCommand(
		"donate",
		"Provides donation information to the channel",
		"",
		donate)
}
