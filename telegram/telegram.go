package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/handler"
)

// NewTelegramBot 创建新的 Telegram Bot 实例
func NewTelegramBot(conf config.Config) {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		logger.Fatal(err)
		return
	}
	bot.Debug = conf.Debug
	logger.Infof("已连接到 @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.Message != nil {
			handler.HandleMessage(bot, update)
		}

		if update.InlineQuery != nil {
			handler.HandleInlineQuery(bot, update)
		}

		if update.CallbackQuery != nil {
			handler.HandleCallbackQuery(bot, update)
		}
	}
}
