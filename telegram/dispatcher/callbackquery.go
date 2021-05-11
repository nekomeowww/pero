package dispatcher

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/callbackquery"
	"github.com/nekomeowww/pero/telegram/handler"
)

// HandleCallbackQuery 处理回调请求
func HandleCallbackQuery(handlers handler.Groups, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Infof("来自用户 [%s] callback query：%s", update.CallbackQuery.From.UserName, update.CallbackQuery.Data)

	callbackquery.Exec(bot, update, update.CallbackQuery.Data, handlers)
}
