package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config 配置文件结构体
type Config struct {
	Token string `yaml:"token"`
	Debug bool   `yaml:"debug"`
	RSS   []RSS  `yaml:"rss"`
}

// RSS 配置结构
type RSS struct {
	Name string `yaml:"name"`
	URI  string `yaml:"uri"`
}

// ParseConfig 解析配置文件
func ParseConfig() Config {
	fileName, _ := filepath.Abs("./conf.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}

	conf := new(Config)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return Config{}
	}
	return *conf
}
