package repository

import (
	"fmt"
	"reflect"

	"github.com/daigd/v-mall-go/datamodel"
)

type baseRepositoryInMemory struct {
}

func (r *baseRepositoryInMemory) FindAll() (out interface{}, found bool) {

	return nil, false
}

func (r *baseRepositoryInMemory) First(out interface{}) (found bool) {
	// 获取接口指向的元素
	v := reflect.ValueOf(out)
	vk := v.Kind()
	fmt.Printf("接口类型:%v,Kind:%v\n", v.Type(), vk)
	// 获取指针指向的数据
	if reflect.Ptr == vk {
		ve := v.Elem()
		switch ve.Kind() {
		case reflect.Struct:
			du := reflect.TypeOf(datamodel.User{})
			// 利用反射修改User的值
			if du == ve.Type() {
				if ve.FieldByName("UserID").Int() < 1 {
					return
				}
				ve.FieldByName("UserName").SetString("test")
				ve.FieldByName("NickName").SetString("Dandy")
			}
		}
	}
	found = true
	return
}

func (r *baseRepositoryInMemory) FirstByCondition(out interface{}, query interface{}, values ...interface{}) {
}

func (*baseRepositoryInMemory) Create(value interface{}) (err error) {
	return nil
}
