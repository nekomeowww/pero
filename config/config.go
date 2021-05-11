package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/nekomeowww/pero/logger"
	"gopkg.in/yaml.v2"
)

// Config 配置文件结构体
type Config struct {
	Token string `yaml:"token"`
	Debug bool   `yaml:"debug"`

	NutsDB NutsDB `yaml:"nutsdb"`
}

// NutsDB 配置
type NutsDB struct {
	Dir string `yaml:"path"`
}

// ParseConfig 解析配置文件
func ParseConfig() Config {
	fileName, _ := filepath.Abs("./conf.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Fatal(err)
	}

	conf := new(Config)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return Config{}
	}
	return *conf
}
