// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "zqzqsb.com/gomall/app/user/biz/handler/user"
	mw "zqzqsb.com/gomall/app/user/biz/router/middleware"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {
    publicGroup := r.Group("/", rootMw()...)
    {
        publicGroup.POST("/login", _loginMw()...)
        publicGroup.POST("/register", append(_registerMw(), user.Register)...)
    }

    // 2. 私有路由组（需要 JWT）
    privateGroup := r.Group("/", rootMw()...)
    {

		// 先注册 jwt 中间件
		privateGroup.Use(mw.JwtMiddleware.MiddlewareFunc())
		// 再注册 casbin 中间件
		enforce, err := mw.InitCasbin()
		if err != nil {
			panic(err)
		}
		privateGroup.Use(mw.NewCasbinMiddleware(enforce))
        privateGroup.GET("/hello", append(_helloMw(), user.Hello)...)
        // ... 其他需要鉴权的路由
    }
}
