package main

import (
	"fmt"
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/mrshankly/go-twitch/twitch"
	"net/http"
	"strings"
	"time"
)

func uptime(command *bot.Cmd) (msg string, err error) {

	client := twitch.NewClient(&http.Client{})

	// Replace with Channel/User id for v5 API
	stream, err := client.Streams.Channel("1903602")

	if err != nil {
		msg = "Error retrieving stream information from Twitch API"
		return
	}

	if stream.Stream.Id == 0 {
		msg = "Stream is currently offline"
		return
	}

	streamStartTime, _ := time.Parse(time.RFC3339, stream.Stream.CreatedAt)

	streamUptime := time.Since(streamStartTime)

	// Remove millseconds from Duration, we don't need that level of precision
	// This doesn't round, it just flat out removes everything after the dot
	uptimeString := strings.Split(streamUptime.String(), ".")[0]

	msg = fmt.Sprintf("Stream Uptime: %vs\n", uptimeString)

	return
}

func init() {
	bot.RegisterCommand(
		"uptime",
		"Returns the Channel Uptime",
		"",
		uptime)
}
