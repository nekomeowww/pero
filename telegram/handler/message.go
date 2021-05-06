package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/commands"
)

// HandleMessage 处理消息
func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Infof("来自用户 [%s]：%s", update.Message.From.UserName, update.Message.Text)

	isCommand, command, _ := commands.Parse(*update.Message)
	if isCommand {
		commands.Exec(bot, update, command)
	}

	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
	// bot.Send(msg)
}
