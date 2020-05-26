package repository

import (
	"fmt"
	"reflect"
)

type baseRepositoryInMemory struct {
}

func (r *baseRepositoryInMemory) FindAll() (out interface{}, found bool) {

	return nil, false
}

func (r *baseRepositoryInMemory) First(out interface{}) (found bool) {
	t := reflect.TypeOf(out)
	fmt.Println("接口类型", t)
	k := t.Kind()
	fmt.Println("接口种类", k)
	// 获取接口指向的元素
	v := reflect.ValueOf(out).Elem()
	// 获取字段数量
	for i := 0; i < v.NumField(); i++ {
		// 取出每个值
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			f.SetString("dgd")
		case reflect.Bool:
			f.SetBool(false)
		}
	}
	found = true
	return
}
