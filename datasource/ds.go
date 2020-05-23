package datasoure

import (
	"errors"

	"github.com/daigd/v-mall-go/datamodel"
)

// DataEngine 数据引擎
type DataEngine uint8

const (
	// Memory 数据从内存加载
	Memory DataEngine = iota
	// Mysql 数据从Mysql加载
	Mysql
)

// LoadMemoryData 从内存中加载数据
func LoadMemoryData(engine DataEngine) (map[int64]datamodel.User, error) {
	if Memory != engine {
		return nil, errors.New("当前方式只支持从内存中加载数据")
	}
	ds := make(map[int64]datamodel.User, 0)
	user := datamodel.User{UserID: 1, UserName: "dai", NickName: "努力的Coder"}
	ds[user.UserID] = user
	return ds, nil
}
