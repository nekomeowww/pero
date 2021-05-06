package handler

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HandleInlineQuery 处理行内请求
func HandleInlineQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("来自用户 [%s] inline query：%s", update.InlineQuery.From, update.InlineQuery.Query)
}
