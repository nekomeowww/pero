package scene

import (
	"encoding/json"
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

// StageSubHandlers 阶段控制子函数
type StageSubHandlers map[int]func(*tgbotapi.BotAPI, tgbotapi.Update)

// StageHandlers 阶段控制
type StageHandlers map[string]StageSubHandlers

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

// GetScene 是否在场景中
func (s *Scene) GetScene() (bool, string, int) {
	switch s.sceneType.(type) {
	case *tgbotapi.Message:
		messageScene := s.sceneType.(*tgbotapi.Message)
		us := ParseSceneInfo(messageScene.From.ID)
		if us == nil {
			return false, "", 0
		}
		return true, us.Name, us.Stage
	case *tgbotapi.InlineQuery:
		inlineScene := s.sceneType.(*tgbotapi.InlineQuery)
		us := ParseSceneInfo(inlineScene.From.ID)
		if us == nil {
			return false, "", 0
		}
		return true, us.Name, us.Stage
	case *tgbotapi.CallbackQuery:
		callbackScene := s.sceneType.(*tgbotapi.CallbackQuery)
		us := ParseSceneInfo(callbackScene.From.ID)
		if us == nil {
			return false, "", 0
		}
		return true, us.Name, us.Stage
	default:
		return false, "", 0
	}
}

// ParseSceneInfo 解析场景信息
func ParseSceneInfo(id int) *UserScene {
	b, err := service.NutsDB.Get(Bucket, UserSceneKey(id))
	if err != nil {
		logger.Error(err)
		return nil
	}

	us := new(UserScene)
	err = json.Unmarshal(b, us)
	if err != nil {
		logger.Error(err)
		return nil
	}

	return us
}
