package controller

import (
	"fmt"

	"github.com/daigd/v-mall-go/bizmodel"
	"github.com/daigd/v-mall-go/service"
	"github.com/kataras/iris"
)

// UserController 用户控制器
type UserController struct {
	Ctx         iris.Context
	Userservice service.UserService
}

// GetBy 根据用户ID查询用户信息 /user/1
func (c *UserController) GetBy(id int64) (user bizmodel.User, found bool) {
	user, found = c.Userservice.QueryByID(id)
	if !found {
		c.Ctx.Values().Set("message", fmt.Sprintf("根据ID[%d]找不到用户信息", id))
	}
	return
}
