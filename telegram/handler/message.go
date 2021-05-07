package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/commands"
	"github.com/nekomeowww/pero/telegram/scene"
)

// HandleMessage 处理消息
func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Infof("来自用户 [%s]：%s", update.Message.From.UserName, update.Message.Text)

	// 命令处理
	isCommand, command, _ := commands.Parse(*update.Message)
	if isCommand {
		commands.Exec(bot, update, command)
		return
	}

	// 场景处理
	s := scene.NewScene(update.Message)
	sceneName := s.IsScene()
	if sceneName != "" {
		logger.Info("用户 %d 在场景：%s", update.Message.From.ID, sceneName)
		return
	}

	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
	// bot.Send(msg)
}
