package rss

import (
	"github.com/nekomeowww/pero/telegram/handler"
)

// CommandHandlers rss 挂载点
func CommandHandlers() handler.Handler {
	handlers := handler.Handler{
		"newfeed": ActionNewFeed,
	}
	return handlers
}

// SceneHandlers rss 挂载点
func SceneHandlers() handler.SceneHandler {
	handlers := handler.SceneHandler{
		"newfeed": handler.StageHandler{
			1: ActionNewFeedName,
			2: ActionConfirmFeed,
		},
	}
	return handlers
}

// CallbackQueryHandlers rss 挂载点
func CallbackQueryHandlers() handler.Handler {
	handler := handler.Handler{
		"edit_newfeed_name": ActionEditFeedName,
		"edit_newfeed_uri":  ActionEditFeedURI,
		"confirm_newfeed":   ActionCreateFeed,
	}
	return handler
}

// Handlers rss 处理器
func Handlers() handler.Handlers {
	handlers := handler.Handlers{
		Scenes:          SceneHandlers(),
		Commands:        CommandHandlers(),
		CallbackQueries: CallbackQueryHandlers(),
	}

	return handlers
}
