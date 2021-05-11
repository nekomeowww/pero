package dispatcher

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/telegram/commands"
	"github.com/nekomeowww/pero/telegram/handler"
	"github.com/nekomeowww/pero/telegram/scene"
)

// HandleMessage 处理消息
func HandleMessage(handlers handler.Groups, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Infof("来自用户 [%s]：%s", update.Message.From.UserName, update.Message.Text)

	// 命令处理
	if isCommand, _, _ := commands.Parse(*update.Message); isCommand {
		commands.Exec(bot, update, handlers)
		return
	}

	// 场景处理
	if isScene, _, _ := scene.NewScene(update.Message).GetScene(); isScene {
		scene.Exec(bot, update, handlers)
		return
	}
}
