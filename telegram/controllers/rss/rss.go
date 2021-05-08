package rss

import (
	"github.com/nekomeowww/pero/telegram/commands"
	"github.com/nekomeowww/pero/telegram/scene"
)

// StageHandlers rss 挂载点
func StageHandlers() scene.StageHandlers {
	handler := scene.StageHandlers{
		"newfeed": {
			1: ActionNewFeedName,
		},
	}

	return handler
}

// CommandHandlers rss 挂载点
func CommandHandlers() commands.CommandHandlers {
	handler := commands.CommandHandlers{
		"newfeed": ActionNewFeed,
	}
	return handler
}
