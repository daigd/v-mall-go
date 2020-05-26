package service

import (
	"fmt"

	"github.com/daigd/v-mall-go/bizmodel"
	"github.com/daigd/v-mall-go/datamodel"
	"github.com/daigd/v-mall-go/repository"
)

// UserService 定义用户服务相关函数
type UserService interface {
	QueryByID(id int64) (bizmodel.User, bool)
	QueryByUserNameAndPwd(username string, password string) bizmodel.User
}

type userService struct {
	repo repository.BaseRepository
}

// NewUserService 创建用户 Service
func NewUserService(repo repository.BaseRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) QueryByUserNameAndPwd(username string, password string) bizmodel.User {
	return bizmodel.User{}
}

func (u *userService) QueryByID(id int64) (user bizmodel.User, found bool) {
	fmt.Printf("查询用户信息，userID:%d\n", id)
	dataUser := datamodel.User{UserID: id}
	found = u.repo.First(&dataUser)
	if !found {
		user = bizmodel.User{}
		return
	}
	user = bizmodel.User{UserID: dataUser.UserID, UserName: dataUser.UserName, NickName: dataUser.NickName}
	return
}
