package dispatcher

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/handler"
)

// Dispatch 分发
func Dispatch(handlers handler.Groups, bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.Fatal(err)
	}
	for update := range updates {
		if update.Message != nil {
			HandleMessage(handlers, bot, update)
		}

		if update.InlineQuery != nil {
			HandleInlineQuery(handlers, bot, update)
		}

		if update.CallbackQuery != nil {
			HandleCallbackQuery(handlers, bot, update)
		}
	}
}
