package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
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
		"info": commandInfo,
	}

	if fc, ok := mCommand[command]; ok {
		fc(bot, update)
	}
}

func commandInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger.Info(fmt.Sprint("执行命令 info: ", update.Message.Text))
}
