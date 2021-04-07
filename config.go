package main

import (
	"github.com/go-chat-bot/bot"
)

func config(command *bot.Cmd) (msg string, err error) {
	msg = "https://gist.github.com/Bjorn248/7548a1a362aaa17272e2 and https://gist.github.com/Bjorn248/c00359af0dafbcf0572c"
	return
}

func init() {
	bot.RegisterCommand(
		"config",
		"Sends config links to the channel",
		"",
		config)
}
