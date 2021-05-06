package main

import (
	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/telegram"
)

func main() {
	conf := config.ParseConfig()
	// rss.Feed(conf)
	telegram.NewTelegramBot(conf)
}
