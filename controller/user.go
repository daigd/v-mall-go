package controller

import (
	"fmt"

	"github.com/daigd/v-mall-go/viewmodel"

	"github.com/daigd/v-mall-go/bizmodel"
	"github.com/daigd/v-mall-go/service"
	"github.com/kataras/iris"
)

// UserController 用户控制器
type UserController struct {
	Ctx         iris.Context
	UserService service.UserService
}

// UserRegisterReq 用户注册请求参数
type userRegisterReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

// GetBy 根据用户ID查询用户信息 /user/1
func (c *UserController) GetBy(id int64) (user bizmodel.User, found bool) {
	user, found = c.UserService.QueryByID(id)
	if !found {
		c.Ctx.Values().Set("message", fmt.Sprintf("根据ID[%d]找不到用户信息", id))
	}
	return
}

// PostRegister 用户注册
func (c *UserController) PostRegister() {
	req := userRegisterReq{}
	var rv *viewmodel.ResultVO
	err := c.Ctx.ReadJSON(&req)

	if err != nil {
		rv = viewmodel.ResultErrorCode(viewmodel.Fail, "请求参数异常")
	}
	userID, err := c.UserService.Create(req.UserName, req.Password, req.NickName, "sys")
	if err != nil {
		rv = viewmodel.ResultErrorCode(viewmodel.Fail, err.Error())
	} else {
		rv = viewmodel.ResultSuccess()
		rv.Data = userID
	}
	c.Ctx.JSON(rv)
	// content, err := json.Marshal(rv)
	// return mvc.Response{
	// 	Err:         err,
	// 	ContentType: "application/json;charset=utf-8",
	// 	Content:     content,
	// }
}
