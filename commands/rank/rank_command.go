package rank

import (
	"fmt"
	"github.com/fabioxgn/go-bot"
)

func rank(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("CS:GO - Supreme Master First Class")
	return
}

func init() {
	bot.RegisterCommand(
		"rank",
		"Sends information about rankings to the channel.",
		"",
		rank)
}
