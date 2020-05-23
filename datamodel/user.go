package datamodel

import "time"

// User 用户数据库模型
type User struct {
	UserID         int64     // 用户ID
	UserName       string    // 用户登录名
	NickName       string    // 用户昵称
	HashedPassword []byte    // 用户密码，经过Hash解析
	CreatedAt      time.Time // 创建时间
	CreatedBy      string    // 创建人
	UpdatedAt      time.Time // 更新时间
	UpdatedBy      string    // 更新人
}
