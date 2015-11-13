package main

import (
	"github.com/fabioxgn/go-bot"
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
