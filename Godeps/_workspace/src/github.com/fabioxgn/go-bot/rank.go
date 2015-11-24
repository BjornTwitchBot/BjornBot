package bot

const (
	rankCommand = "rank"
)

func rank(c *Cmd, channel, senderNick string, conn connection) {
	conn.Privmsg(channel, "League of Legends - Gold V")
	conn.Privmsg(channel, "CS:GO - Legendary Eagle Master")
}
