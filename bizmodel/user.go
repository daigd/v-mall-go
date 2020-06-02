package bizmodel

// User 用户业务模型
type User struct {
	UserID   int64  `josn:"userId" form:"userId2"`     // 用户ID
	UserName string `json:"userName" form:"userName2"` // 用户登录名
	NickName string `json:"nickName" form:"nickName2"` // 用户昵称
}
