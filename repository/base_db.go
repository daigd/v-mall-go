package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type baseRepositoryInDB struct {
	db *gorm.DB
}

func (r *baseRepositoryInDB) FindAll() (out interface{}, found bool) {
	db := r.db.Find(out)
	fmt.Println("查询结果", db)
	if db.Error == nil {
		found = true
	}
	return
}

func (r *baseRepositoryInDB) First(out interface{}) (found bool) {
	db := r.db.First(out)
	if db.Error == nil {
		found = true
	}
	return
}
