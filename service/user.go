package service

import (
	"errors"
	"fmt"
	"github.com/daigd/v-mall-go/bizmodel"
	"github.com/daigd/v-mall-go/datamodel"
	"github.com/daigd/v-mall-go/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService 定义用户服务相关函数
type UserService interface {
	QueryByID(id int64) (bizmodel.User, bool)
	QueryByUserName(username string) (user bizmodel.User, found bool)
	QueryByUserNameAndPwd(username string, password string) (user bizmodel.User, found bool)
	Create(username string, password string, nickname string, createdBy string) (id int64, err error)
}

type userService struct {
	repo repository.BaseRepository
}

// NewUserService 创建用户 Service
func NewUserService(repo repository.BaseRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) QueryByUserName(username string) (user bizmodel.User, found bool) {
	dataUser := datamodel.User{}
	u.repo.FirstByCondition(&dataUser, "user_name=?", username)
	if dataUser.UserID < 1 {
		user = bizmodel.User{}
		return
	}
	found = true
	user = bizmodel.User{UserID: dataUser.UserID, UserName: dataUser.UserName, NickName: dataUser.NickName}
	return
}

func (u *userService) QueryByUserNameAndPwd(username string, password string) (user bizmodel.User, found bool) {
	dataUser := datamodel.User{}
	u.repo.FirstByCondition(&dataUser, "user_name=?", username)
	if dataUser.UserID < 1 {
		user = bizmodel.User{}
		return
	}
	err := compareHashedPassword(dataUser.HashedPassword, password)
	if err != nil {
		fmt.Println("密码不正确", err)
		user = bizmodel.User{}
		return
	}
	found = true
	user = bizmodel.User{UserID: dataUser.UserID, UserName: dataUser.UserName, NickName: dataUser.NickName}
	return
}

func (u *userService) QueryByID(id int64) (user bizmodel.User, found bool) {
	fmt.Printf("查询用户信息，userID:%d\n", id)
	dataUser := datamodel.User{}
	u.repo.FirstByCondition(&dataUser, "user_id=?", id)

	if dataUser.UserID < 1 {
		user = bizmodel.User{}
		return
	}
	found = true
	user = bizmodel.User{UserID: dataUser.UserID, UserName: dataUser.UserName, NickName: dataUser.NickName}
	return
}

func (u *userService) Create(username string, password string, nickname string, operatedBy string) (id int64, err error) {
	err = checkUserInfo(username, password, nickname)
	if err != nil {
		return 0, err
	}
	hashedPwd, err := generateHashedPassword(password)
	if err != nil {
		return 0, err
	}
	du := datamodel.User{
		UserName:       username,
		NickName:       nickname,
		HashedPassword: hashedPwd,
		CreatedBy:      operatedBy,
		UpdatedBy:      operatedBy,
	}
	err = u.repo.Create(&du)
	return du.UserID, err
}

func checkUserInfo(username string, password string, nickname string) (err error) {
	if len(username) == 0 || len(password) == 0 || len(nickname) == 0 {
		err = errors.New("用户名或密码或昵称长度不能为0")
		return
	}
	if len(username) >= 16 || len(password) >= 32 || len(nickname) >= 32 {
		err = errors.New("用户名或密码或昵称长度超出限制范围")
		return
	}
	return
}

func generateHashedPassword(password string) (hashedPwd string, err error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPwd = string(pwd)
	return
}

func compareHashedPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
