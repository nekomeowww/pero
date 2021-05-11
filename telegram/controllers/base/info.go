package base

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// ActionInfo 信息
func ActionInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "你好哦")
	bot.Send(msg)
}
