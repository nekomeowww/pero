package handler

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HandleCallbackQuery 处理回调请求
func HandleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("来自用户 [%s]：%s", update.Message.From.UserName, update.Message.Text)
	log.Printf("来自用户 [%s] callback query：%s", update.CallbackQuery.Data, update.CallbackQuery.From)
}
