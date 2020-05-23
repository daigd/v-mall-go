package repository

import (
	"sync"

	"github.com/daigd/v-mall-go/datamodel"
)

// Query represents the visitor and action queries.
// Query 用来查询数据,充当数据的访问者和操作者
type Query func(model datamodel.User) bool

const (
	// ReadOnlyMode 数据访问使用只读模式
	ReadOnlyMode = iota
	// ReadWriteMode 数据访问使用读写模式
	ReadWriteMode
)

// UserRepository 提供基础函数定义
type UserRepository interface {
	// 按给定的条件来执行数据操作
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	// 按给定的查询条件选择数据
	Select(query Query) (user datamodel.User, found bool)
}

type userRepository struct {
	source map[int64]datamodel.User
	mu     sync.RWMutex
}

// NewUserRepository 创建一个用户Repository
func NewUserRepository(s map[int64]datamodel.User) UserRepository {
	return &userRepository{source: s}
}

// 实现 BaseRepository 接口的 Exec 函数
func (r *userRepository) Exec(query Query, action Query, limit int, mode int) (ok bool) {
	loop := 0
	if ReadOnlyMode == mode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, user := range r.source {
		// 执行查询函数
		ok = query(user)
		if ok {
			// 执行操作函数
			if action(user) {
				// 通过loop变量控制返回的数据数量
				loop++
				if limit >= loop {
					break
				}
			}
		}
	}
	return
}

// 实现 BaseRepository 接口的 Select 函数
// Select receives a query function
// which is fired for every single user model inside
// our imaginary data source.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known boolean value
// and the last known user model
// to help callers to reduce the LOC.
//
// It's actually a simple but very clever prototype function
// I'm using everywhere since I firstly think of it,
// hope you'll find it very useful as well.
func (r *userRepository) Select(query Query) (user datamodel.User, found bool) {
	found = r.Exec(query, func(m datamodel.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)

	// 如果没有找到，返回空的结构体
	if !found {
		user = datamodel.User{}
	}
	return
}
