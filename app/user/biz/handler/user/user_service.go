// Code generated by hertz generator.

package user

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/csrf"
	"zqzqsb.com/gomall/app/user/biz/service"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 调用 RPC 服务进行注册
	registerService := service.NewRegisterService(ctx)
	resp, err := registerService.Run(&req)
	if err != nil {
		c.String(consts.StatusInternalServerError, fmt.Sprintf("Registration failed: %v", err))
		return
	}

	// 返回注册响应
	c.JSON(consts.StatusOK, resp)

}

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	LoginService := service.NewLoginService(ctx)
	resp, err := LoginService.Run(&req)
	if err != nil {
		c.String(consts.StatusInternalServerError, fmt.Sprintf("Registration failed: %v", err))
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// Hello .
// @router /hello [GET]
func Hello(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.HelloResp)
	csrfToken := csrf.GetToken(c)
	resp.RespBody = csrfToken
	c.JSON(consts.StatusOK, resp)
}
