package rss

import (
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
	err := scene.UpdateScene(update.Message.From.ID, "newfeed", 1)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "好的，发送给我你想要新增的订阅链接吧。如果想要终止操作，可以发送 /cancel 来取消。")
	bot.Send(msg)
}

// ActionNewFeedName 新建订阅源的名字
func ActionNewFeedName(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	key := []byte(fmt.Sprintf("new_rss_url=%d", update.Message.From.ID))
	err := service.NutsDB.Set(scene.Bucket, key, []byte(update.Message.Text), 0)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	err = scene.UpdateScene(update.Message.From.ID, "newfeed", 2)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("新建的订阅源链接是： %s ，现在给它一个命名吧！如果想要终止操作，可以发送 /cancel 来取消。", update.Message.Text))
	bot.Send(msg)
}

// ActionConfirmFeed 确认新建订阅源
func ActionConfirmFeed(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	inlineMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("修改名字", "edit_newfeed_name"),
			tgbotapi.NewInlineKeyboardButtonData("修改链接", "edit_newfeed_uri"),
		), tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("确认", "confirm_newfeed"),
		))

	key := []byte(fmt.Sprintf("new_rss_url=%d", update.Message.From.ID))
	bytes, err := service.NutsDB.Get(scene.Bucket, key)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	err = scene.UpdateScene(update.Message.From.ID, "newfeed", 2)
	if err != nil {
		message.SendErrorFeedback(bot, update.Message.Chat.ID, newfeedError)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("新建的订阅源链接是：%s ，名称为 %s，如果想要终止操作，可以发送 /cancel 来取消。", string(bytes), update.Message.Text))
	msg.BaseChat.ReplyMarkup = inlineMarkup
	bot.Send(msg)
}

// ActionEditFeedName 编辑订阅源名称
func ActionEditFeedName(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "好哦")
	bot.Send(msg)
}

// ActionEditFeedURI 编辑订阅源地址
func ActionEditFeedURI(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "好哦")
	bot.Send(msg)
}

// ActionCreateFeed 创建订阅源
func ActionCreateFeed(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "创建成功")
	bot.Send(msg)
}
