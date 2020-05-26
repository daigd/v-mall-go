package repository

import (
	datasoure "github.com/daigd/v-mall-go/datasource"
	"github.com/jinzhu/gorm"
)

// BaseRepository 基础查询接口定义
type BaseRepository interface {
	FindAll() (out interface{}, found bool)
	First(out interface{}) (found bool)
}

// NewBaseRepository 创建一个BaseRepository
func NewBaseRepository(de datasoure.DataEngine) BaseRepository {
	if datasoure.Memory == de {
		return &baseRepositoryInMemory{}
	}
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return &baseRepositoryInDB{db: db}
}
