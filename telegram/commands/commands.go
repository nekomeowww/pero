package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/telegram/handler"
)

// Handlers 命令函数
type Handlers map[string]func(*tgbotapi.BotAPI, tgbotapi.Update)

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
func Exec(bot *tgbotapi.BotAPI, update tgbotapi.Update, commandFuncs handler.Groups) {
	command := update.Message.Command()
	for i := range commandFuncs {
		if fc, ok := commandFuncs[i].Commands[command]; ok {
			fc(bot, update)
		}
	}
}
