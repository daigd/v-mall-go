package repository

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daigd/v-mall-go/config"
	datasoure "github.com/daigd/v-mall-go/datasource"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/yaml.v3"
)

// BaseRepository 基础查询接口定义
type BaseRepository interface {
	FindAll() (out interface{}, found bool)
	First(out interface{}) (found bool)
	FirstByConditon(out interface{}, query interface{}, values ...interface{})
	Create(value interface{}) error
}

// NewBaseRepository 创建一个BaseRepository
func NewBaseRepository(de datasoure.DataEngine) BaseRepository {
	if datasoure.Memory == de {
		return &baseRepositoryInMemory{}
	}
	file, err := os.Open("app.yaml")
	if err != nil {
		iris.Default().Logger().Fatal("打开数据库配置文件失败", err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		iris.Default().Logger().Fatal("读取数据库配置文件失败", err)
	}
	var appConfig config.AppConfig
	err = yaml.Unmarshal(bytes, &appConfig)
	if err != nil {
		iris.Default().Logger().Fatal("数据库配置文件解析失败", err)
	}
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", appConfig.DBConfig.UserName, appConfig.DBConfig.Password, appConfig.DBConfig.Host, appConfig.DBConfig.Port, appConfig.DBConfig.DataBase)

	fmt.Println("数据库连接地址", dbURL)
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	return &baseRepositoryInDB{db: db}
}
