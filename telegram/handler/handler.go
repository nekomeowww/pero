package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Handler 处理器
type Handler map[string]BaseHandler

// BaseHandler 基本处理器
type BaseHandler func(bot *tgbotapi.BotAPI, update tgbotapi.Update)

// SceneHandler 场景处理器
type SceneHandler map[string]StageHandler

// StageHandler 阶段处理器
type StageHandler map[int]BaseHandler

// Handlers 处理对象
type Handlers struct {
	Commands        Handler
	Scenes          SceneHandler
	Messages        Handler
	InlineQueries   Handler
	CallbackQueries Handler
}

// Groups 处理组
type Groups []Handlers

// Command 命令处理
func (h *Handlers) Command() Handler {
	return h.Commands
}

// Scene 场景处理
func (h *Handlers) Scene() SceneHandler {
	return h.Scenes
}

// Message 消息处理
func (h *Handlers) Message() Handler {
	return h.Messages
}

// InlineQuery 行内请求处理
func (h *Handlers) InlineQuery() Handler {
	return h.InlineQueries
}

// CallbackQuery 回调请求处理
func (h *Handlers) CallbackQuery() Handler {
	return h.CallbackQueries
}

// Mount 挂载
func (g *Groups) Mount() {

}
