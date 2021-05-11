package base

import (
	"github.com/nekomeowww/pero/telegram/handler"
)

// CommandHandlers 命令函数
func CommandHandlers() handler.Handler {
	handler := handler.Handler{
		"info":   ActionInfo,
		"cancel": ActionCancel,
	}

	return handler
}

// Handlers base 处理器
func Handlers() handler.Handlers {
	handlers := handler.Handlers{
		Commands: CommandHandlers(),
	}

	return handlers
}
