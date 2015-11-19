package bot

import (
	"fmt"
)

const (
	rankCommand = "rank"
)

func rank(c *Cmd, channel, senderNick string, conn connection) {
	conn.Privmsg(channel, fmt.Sprintf("League of Legends - Gold V"))
	conn.Privmsg(channel, fmt.Sprintf("CS:GO - Supreme Master First Class"))
}
