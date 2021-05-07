package main

import (
	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/service"
	"github.com/nekomeowww/pero/telegram"
)

func main() {
	conf := config.ParseConfig()
	// rss.Feed(conf)
	service.Init(conf)
	telegram.NewTelegramBot(conf)
}
