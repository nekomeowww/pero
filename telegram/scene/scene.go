package scene

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nekomeowww/pero/logger"
	"github.com/nekomeowww/pero/service"
)

// Bucket 场景命名空间
const Bucket = "bucket_scene"

type sceneType interface{}

// Scene 结构
type Scene struct {
	sceneType
}

// UserScene 用户场景
type UserScene struct {
	Name  string `json:"name"`
	Stage int    `json:"stage"`
}

// NewScene 新建场景对象
func NewScene(v interface{}) *Scene {
	s := new(Scene)
	s.sceneType = v
	return s
}

// UserSceneKey 获取用户的 scene 键
func UserSceneKey(userID int) []byte {
	return []byte(fmt.Sprintf("scene_%s", strconv.Itoa(userID)))
}

// IsScene 是否在场景中
func (s *Scene) IsScene() string {
	switch s.sceneType.(type) {
	case *tgbotapi.Message:
		messageScene := s.sceneType.(*tgbotapi.Message)
		b, err := service.NutsDB.Get(Bucket, UserSceneKey(messageScene.From.ID))
		if err != nil {
			logger.Error(err)
			return ""
		}
		return string(b)
	case *tgbotapi.InlineQuery:
		inlineScene := s.sceneType.(*tgbotapi.InlineQuery)
		b, err := service.NutsDB.Get(Bucket, UserSceneKey(inlineScene.From.ID))
		if err != nil {
			logger.Error(err)
			return ""
		}
		return string(b)
	case *tgbotapi.CallbackQuery:
		callbackScene := s.sceneType.(*tgbotapi.CallbackQuery)
		b, err := service.NutsDB.Get(Bucket, UserSceneKey(callbackScene.From.ID))
		if err != nil {
			logger.Error(err)
			return ""
		}
		return string(b)
	default:
		return ""
	}
}
