package datamodel

import "time"

// User 用户
type User struct {
	UserID         int64     `gorm:"primary_key"` // 用户ID
	UserName       string    // 用户登录名
	NickName       string    // 用户昵称
	HashedPassword string    // 用户密码，经过Hash解析
	CreatedAt      time.Time // 创建时间
	CreatedBy      string    // 创建人
	UpdatedAt      time.Time // 更新时间
	UpdatedBy      string    // 更新人
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}
