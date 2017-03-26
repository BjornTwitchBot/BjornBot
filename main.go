package main

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
	"log"
	"os"
	"strings"
)

func main() {

	if os.Getenv("TWITCH_OAUTH_TOKEN") == "" {
		log.Fatal("TWITCH_OAUTH_TOKEN not set")
	}
	if os.Getenv("TWITCH_CHANNEL_NAME") == "" {
		log.Fatal("TWITCH_CHANNEL_NAME not set")
	}

	bot.Run(&bot.Config{
		Server:   "irc.chat.twitch.tv:6697",
		Channels: strings.Split(os.Getenv("TWITCH_CHANNEL_NAME"), ","),
		User:     "BjornTwitchBot",
		Nick:     "bjorntwitchbot",
		Password: os.Getenv("TWITCH_OAUTH_TOKEN"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != ""})
}
