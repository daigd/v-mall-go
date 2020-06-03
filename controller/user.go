package controller

import (
	"encoding/json"
	"fmt"

	"github.com/daigd/v-mall-go/bizmodel"
	"github.com/daigd/v-mall-go/middleware"
	"github.com/daigd/v-mall-go/viewmodel"
	"github.com/kataras/iris/sessions"

	"github.com/daigd/v-mall-go/service"
	"github.com/kataras/iris"
)

const (
	userInfo = "userInfo"
)

// UserController 用户控制器
type UserController struct {
	Ctx         iris.Context
	UserService service.UserService
	// Session, binded using dependency injection from the iris.go.
	Session *sessions.Session
}

// userLoginReq 用户登录请求参数
type userLoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// userRegisterReq 用户注册请求参数
type userRegisterReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

// GetBy 根据用户ID查询用户信息 /user/1
func (c *UserController) GetBy(id int64) (rv *viewmodel.ResultVO) {
	user, found := c.UserService.QueryByID(id)
	if !found {
		rv = viewmodel.ResultErrorCode(viewmodel.FAIL, fmt.Sprintf("根据ID[%d]找不到用户信息", id))
	} else {
		rv = viewmodel.ResultSuccessData(&user)
	}
	return
}

// GetMe 获取用户信息 /user/me
func (c *UserController) GetMe() (rv *viewmodel.ResultVO) {
	user, found := c.isLoggedIn()
	if !found {
		rv = viewmodel.ResultErrorMsg("请重新登录")
		return
	}
	rv = viewmodel.ResultSuccessData(user)
	return
}

// PostLogin 用户登录 POST /user/login
func (c *UserController) PostLogin() (rv *viewmodel.ResultVO) {
	req := userLoginReq{}
	err := c.Ctx.ReadJSON(&req)
	if err != nil {
		rv = viewmodel.ResultErrorMsg("请求参数异常")
		return
	}
	user, found := c.UserService.QueryByUserNameAndPwd(req.UserName, req.Password)
	if !found {
		rv = viewmodel.ResultErrorMsg("用户名或密码不正确")
		return
	}
	rv = viewmodel.ResultSuccessData(user)
	// 将用户信息缓存到Session中
	bytes, _ := json.Marshal(&user)
	fmt.Printf("将用户信息缓存起来:%q\n", string(bytes))
	err = middleware.Set(userInfo, string(bytes))
	if err != nil {
		fmt.Println("缓存用户信息出现错误", err)
	}
	return
}

// PostRegister 用户注册 POST /user/register
func (c *UserController) PostRegister() (rv *viewmodel.ResultVO) {
	req := userRegisterReq{}
	err := c.Ctx.ReadJSON(&req)

	if err != nil {
		rv = viewmodel.ResultErrorMsg("请求参数异常")
		return
	}
	_, found := c.UserService.QueryByUserName(req.UserName)
	if found {
		rv = viewmodel.ResultErrorMsg("用户名已存在")
		return
	}
	userID, err := c.UserService.Create(req.UserName, req.Password, req.NickName, "sys")

	if err != nil {
		rv = viewmodel.ResultErrorMsg(err.Error())
		return
	}
	rv = viewmodel.ResultSuccessData(userID)
	return
}

// 判断用户是否登录
func (c *UserController) isLoggedIn() (user bizmodel.User, found bool) {
	userInfoStr := middleware.Get(userInfo)
	fmt.Println("缓存读取到的用户信息", userInfoStr)
	if len(userInfoStr) == 0 {
		return
	}
	err := json.Unmarshal([]byte(userInfoStr), &user)
	if err != nil {
		fmt.Println("用户登录信息解析异常", err)
		return
	}
	found = true
	return
}
