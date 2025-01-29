package mw

import (
	"context"
	"log"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"zqzqsb.com/gomall/app/user/biz/service"
)

var permissionSvc *service.PermissionService

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

	// 初始化权限服务
	permissionSvc = service.NewPermissionService(enforcer)

	return enforcer, nil
}

func NewCasbinMiddleware(enforcer *casbin.Enforcer) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取用户信息
		userID := c.GetInt64("identity")
		user := strconv.FormatInt(userID, 10)

		// 获取请求方法和路径
		method := string(c.Method())
		path := string(c.Path())

		log.Printf("Casbin checking - user: %s, method: %s, path: %s", user, method, path)

		// 首先检查是否在黑名单中
		if res, err := permissionSvc.IsBlacklisted(user); res || err != nil {
			log.Printf("User %s is blacklisted", user)
			c.JSON(403, utils.H{
				"code":    403,
				"user":    user,
				"message": "用户已被封禁",
			})
			c.Abort()
			return
		}

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
