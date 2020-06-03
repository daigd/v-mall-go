package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
	"gopkg.in/yaml.v2"
)

// Config 应用配置信息
var Config *AppConfig

// AppConfig 应用配置
type (
	AppConfig struct {
		DBConfig struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			UserName string `yaml:"username"`
			Password string `yaml:"password"`
			DataBase string `yaml:"database"`
		}
		RedisConfig struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
			DB   string `yaml:"db"`
		}
	}
)

func init() {
	file, err := os.Open("app.yaml")

	if err != nil {
		iris.Default().Logger().Fatal("打开应用配置文件失败", err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		iris.Default().Logger().Fatal("读取应用配置信息失败", err)
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		iris.Default().Logger().Fatal("应用配置信息解析失败", err)
	}
	fmt.Println("应用配置信息", Config)
}
