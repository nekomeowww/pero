package commands

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/service"
	"github.com/nekomeowww/pero/telegram/message"
	"github.com/nekomeowww/pero/telegram/scene"
)

// Parse 解析 Telegram 消息命令
func Parse(message tgbotapi.Message) (bool, string, string) {
	isCommand := message.IsCommand()
	if isCommand {
		return isCommand, message.Command(), message.CommandArguments()
	}

	return false, "", ""
}

// ParseWithAt 解析 Telegram 消息命令 并且 不移除 @
func ParseWithAt(message tgbotapi.Message) (bool, string, string) {
	command := message.CommandWithAt()
	args := message.CommandArguments()

	if command != "" {
		return true, command, args
	}
	return false, "", ""
}

// Exec 执行命令
func Exec(bot *tgbotapi.BotAPI, update tgbotapi.Update, command string) {
	mCommand := map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
		"info":        commandInfo,
		"cancel":      commandCancel,
		"newfeed":     commandNewFeed,
		"currentfeed": commandCurrentFeed,
	}
	if fc, ok := mCommand[command]; ok {
		fc(bot, update)
	}
}

func commandInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Info(fmt.Sprint("执行命令 /info: ", update.Message.Text))
}

func commandCancel(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	err := service.NutsDB.Del(scene.Bucket, scene.UserSceneKey(update.Message.From.ID))
	if err != nil {
		logger.Error(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "当前在进行的操作都已经结束了哦")
	bot.Send(msg)
}

func commandNewFeed(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var us scene.UserScene
	us.Name = "newfeed"
	us.Stage = 1
	val, err := json.Marshal(us)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, "新建订阅源时发生错误，如果时有发生此错误请联系开发者或管理员")
		logger.Error(err)
		return
	}

	err = service.NutsDB.Set(scene.Bucket, scene.UserSceneKey(update.Message.From.ID), val, 0)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, "新建订阅源时发生错误，如果时有发生此错误请联系开发者或管理员")
		logger.Error(err)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "好的，发送给我你想要新增的订阅链接吧。如果想要终止操作，可以发送 /cancel 来取消。")
	bot.Send(msg)
}

func commandCurrentFeed(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Info(fmt.Sprint("执行命令 /currentfeed: ", update.Message.Text))
}
