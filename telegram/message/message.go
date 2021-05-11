package message

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/logger"
)

// SendErrorFeedback 发送错误信息给用户
func SendErrorFeedback(bot *tgbotapi.BotAPI, chatID int64, v ...interface{}) {
	logger.Error(v...)
	str := fmt.Sprintln(v...)

	msg := tgbotapi.NewMessage(chatID, str)
	bot.Send(msg)
}

// SendErrorfFeedback 发送错误信息给用户
func SendErrorfFeedback(bot *tgbotapi.BotAPI, chatID int64, format string, v ...interface{}) {
	logger.Errorf(format, v...)
	str := fmt.Sprintf(format, v...)

	msg := tgbotapi.NewMessage(chatID, str)
	bot.Send(msg)
}
