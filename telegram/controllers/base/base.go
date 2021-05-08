package base

import "github.com/nekomeowww/pero/telegram/commands"

// CommandHandlers 命令函数
func CommandHandlers() commands.CommandHandlers {
	handler := commands.CommandHandlers{
		"info":   ActionInfo,
		"cancel": ActionCancel,
	}

	return handler
}
