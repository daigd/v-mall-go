package repository

import (
	"fmt"
	"github.com/daigd/v-mall-go/config"
	datasoure "github.com/daigd/v-mall-go/datasource"
	"github.com/jinzhu/gorm"
)

// BaseRepository 基础查询接口定义
type BaseRepository interface {
	FindAll() (out interface{}, found bool)
	First(out interface{}) (found bool)
	FirstByCondition(out interface{}, query interface{}, values ...interface{})
	Create(value interface{}) error
}

// NewBaseRepository 创建一个BaseRepository
func NewBaseRepository(de datasoure.DataEngine) BaseRepository {
	if datasoure.Memory == de {
		return &baseRepositoryInMemory{}
	}
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.DBConfig.UserName, config.Config.DBConfig.Password, config.Config.DBConfig.Host, config.Config.DBConfig.Port, config.Config.DBConfig.DataBase)
	fmt.Println("数据库连接地址", dbURL)
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	return &baseRepositoryInDB{db: db}
}
