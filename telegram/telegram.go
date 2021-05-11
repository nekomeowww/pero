package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/logger"
)

// NewTelegramBot 创建新的 Telegram Bot 实例
func NewTelegramBot(conf config.Config) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	bot.Debug = conf.Debug
	logger.Infof("已连接到 @%s", bot.Self.UserName)
	return bot
}
