package callbackquery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nekomeowww/pero/telegram/handler"
)

// Exec 执行命令
func Exec(bot *tgbotapi.BotAPI, update tgbotapi.Update, data string, cbqFuncs handler.Groups) {
	for i := range cbqFuncs {
		if fc, ok := cbqFuncs[i].CallbackQueries[data]; ok {
			fc(bot, update)
		}
	}
}
