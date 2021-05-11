package dispatcher

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/handler"
)

// HandleInlineQuery 处理行内请求
func HandleInlineQuery(handlers handler.Groups, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Infof("来自用户 [%s] inline query：%s", update.InlineQuery.From, update.InlineQuery.Query)
}
