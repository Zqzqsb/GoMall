package mw

import (
	"context"
	"log"
	"strconv"

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
		// 从请求中获取用户信息
		userID := c.GetInt64("identity")

		// 将 userID 转换为字符串，因为 policy 中的主体是字符串
		user := strconv.FormatInt(userID, 10)

		// 获取请求方法和路径
		method := string(c.Method())
		path := string(c.Path())

		log.Printf("Casbin checking - user: %s, method: %s, path: %s", user, method, path)

		// 检查权限
		ok, err := enforcer.Enforce(user, path, method)

		log.Printf("Casbin result - ok: %v, err: %v", ok, err)

		if err != nil {
			log.Printf("Casbin error: %v", err)
			c.JSON(403, utils.H{
				"code":    403,
				"message": "权限验证错误",
			})
			c.Abort()
			return
		}

		if !ok {
			log.Printf("Casbin denied access")
			c.JSON(403, utils.H{
				"code":    403,
				"user":    user,
				"message": "没有访问权限",
			})
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
