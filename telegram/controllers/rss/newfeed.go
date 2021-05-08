package rss

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/service"
	"github.com/nekomeowww/pero/telegram/message"
	"github.com/nekomeowww/pero/telegram/scene"
)

const (
	newfeedError = "新建订阅源时发生错误，如果时有发生此错误请联系开发者或管理员"
)

// ActionNewFeed 新建订阅源
func ActionNewFeed(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var us scene.UserScene
	us.Name = "newfeed"
	us.Stage = 1
	val, err := json.Marshal(us)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	err = service.NutsDB.Set(scene.Bucket, scene.UserSceneKey(update.Message.From.ID), val, 0)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "好的，发送给我你想要新增的订阅链接吧。如果想要终止操作，可以发送 /cancel 来取消。")
	bot.Send(msg)
}

// ActionNewFeedName 新建订阅源的名字
func ActionNewFeedName(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var us scene.UserScene
	us.Name = "newfeed"
	us.Stage = 2
	val, err := json.Marshal(us)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	key := []byte(fmt.Sprintf("new_rss_url=%s", update.Message.Text))
	err = service.NutsDB.Set(scene.Bucket, key, []byte(update.Message.Text), 0)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	err = service.NutsDB.Set(scene.Bucket, scene.UserSceneKey(update.Message.From.ID), val, 0)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("新建的订阅源链接是：%s ，现在给它一个命名吧！如果想要终止操作，可以发送 /cancel 来取消。", update.Message.Text))
	bot.Send(msg)
}
