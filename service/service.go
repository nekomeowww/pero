package service

import (
	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/service/database"
)

// service
var (
	NutsDB *database.NutsDB
)

// Init 初始化服务
func Init(conf config.Config) {
	NutsDB = database.InitNutsDB(conf)
}
