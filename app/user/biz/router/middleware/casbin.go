package mw

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func InitCasbin() (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer("conf/rbac_model.conf", "conf/rbac_policy.csv")
	if err != nil {
		return nil, err
	}

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}

func NewCasbinMiddleware(enforcer *casbin.Enforcer) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取用户信息（这里需要根据你的认证系统进行修改）
		user := c.GetString("identity")
		if user == "" {
			user = "anonymous"
		}

		// log.Println("user:", user)

		// 获取请求方法和路径
		method := string(c.Method())
		path := string(c.Path())

		// 检查权限
		ok, err := enforcer.Enforce(user, path, method)
		if err != nil {
			c.JSON(403, utils.H{
				"code":    403,
				"message": "权限验证错误",
			})
			c.Abort()
			return
		}

		if !ok {
			// c.JSON(403, utils.H{
			// 	"code":    403,
			// 	"message": "没有访问权限",
			// })
			// c.Abort()
			// return
			// log.Println("没有访问权限")
		}

		c.Next(ctx)
	}
}
