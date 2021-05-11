package base

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/service"
	"github.com/nekomeowww/pero/telegram/scene"
)

// ActionCancel 取消
func ActionCancel(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	err := service.NutsDB.Del(scene.Bucket, scene.UserSceneKey(update.Message.From.ID))
	if err != nil {
		logger.Error(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "当前在进行的操作都已经结束了哦")
	bot.Send(msg)
}
