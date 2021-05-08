package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/commands"
	"github.com/nekomeowww/pero/telegram/controllers/base"
	"github.com/nekomeowww/pero/telegram/controllers/rss"
	"github.com/nekomeowww/pero/telegram/scene"
)

var mStageHandlers = map[string]scene.StageHandlers{
	"newfeed": rss.StageHandlers(),
}

var mCommandHandlers = []commands.CommandHandlers{
	rss.CommandHandlers(),
	base.CommandHandlers(),
}

// HandleMessage 处理消息
func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Infof("来自用户 [%s]：%s", update.Message.From.UserName, update.Message.Text)

	// 命令处理
	isCommand, command, _ := commands.Parse(*update.Message)
	if isCommand {
		commands.Exec(bot, update, command, mCommandHandlers)
		return
	}

	// 场景处理
	s := scene.NewScene(update.Message)
	isScene, name, stage := s.GetScene()
	if isScene {
		logger.Infof("用户 %d 在场景：%s 的第 %d 阶段", update.Message.From.ID, name, stage)

		if shc, ok := mStageHandlers[name]; ok {
			if sfc, ok := shc[name]; ok {
				if fc, ok := sfc[stage]; ok {
					fc(bot, update)
				}
			}
		}

		return
	}

	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
	// bot.Send(msg)
}
