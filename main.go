package main

import (
	_ "github.com/Bjorn248/BjornBot/commands/rank"
	_ "github.com/Bjorn248/BjornBot/commands/sens"
	"github.com/fabioxgn/go-bot"
	// Import all the commands you wish to use
	"os"
	"strings"
)

func main() {
	bot.Run(&bot.Config{
		Server:   "irc.twitch.tv:6667",
		Channels: strings.Split(os.Getenv("TWITCH_CHANNEL_NAME"), ","),
		User:     "BjornTwitchBot",
		Nick:     "bjorntwitchbot",
		Password: os.Getenv("TWITCH_OAUTH_TOKEN"),
		UseTLS:   false,
		Debug:    os.Getenv("DEBUG") != ""})
}
