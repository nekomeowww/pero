package main

import (
	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/service"
	"github.com/nekomeowww/pero/telegram"
	"github.com/nekomeowww/pero/telegram/controllers/base"
	"github.com/nekomeowww/pero/telegram/controllers/rss"
	"github.com/nekomeowww/pero/telegram/dispatcher"
	"github.com/nekomeowww/pero/telegram/handler"
)

func main() {
	conf := config.ParseConfig()
	// rss.Feed(conf)
	service.Init(conf)

	handlers := handler.Groups{
		base.Handlers(),
		rss.Handlers(),
	}

	dispatcher.Dispatch(handlers, telegram.NewTelegramBot(conf))
}
