package mw // 定义包名为 mw，通常用于存放中间件相关的代码

import (
	"context"
	"fmt"
	"net/http"
	"time"

	// 导入项目内部的包
	// 数据库访问层，用于检查用户信息
	"github.com/cloudwego/hertz/pkg/app"          // Hertz 框架的 app 包，处理请求上下文
	"github.com/cloudwego/hertz/pkg/common/hlog"  // Hertz 的日志包
	"github.com/cloudwego/hertz/pkg/common/utils" // Hertz 的工具包，包含辅助函数
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt" // Hertz 的 JWT 中间件包
	"zqzqsb.com/gomall/app/user/biz/model"
	"zqzqsb.com/gomall/app/user/biz/service"
	"zqzqsb.com/gomall/app/user/kitex_gen/user"
)

// 全局变量，用于存储初始化后的 JWT 中间件实例
var (
	JwtMiddleware *jwt.HertzJWTMiddleware // JWT 中间件实例
	IdentityKey   = "identity"            // 用于在 JWT 载荷中存储用户身份的键
)

// InitJwt 初始化 JWT 中间件
func InitJwt() {
	var err error
	// 创建新的 JWT 中间件实例并配置相关参数
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",                                        // 认证领域，用于在 WWW-Authenticate 头中返回
		Key:           []byte("secret key"),                               // 用于签名 JWT 的密钥，请确保使用足够复杂且安全的密钥
		Timeout:       time.Hour,                                          // JWT 的有效期，此处设置为 1 小时
		MaxRefresh:    time.Hour,                                          // 允许刷新 JWT 的最大时间，此处设置为 1 小时
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // 定义从哪里查找 JWT
		TokenHeadName: "Bearer",                                           // JWT 在请求头中的前缀
		// 自定义登录成功后的响应格式
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,                        // 状态码
				"token":   token,                       // 生成的 JWT Token
				"expire":  expire.Format(time.RFC3339), // 过期时间
				"message": "success",                   // 消息
			})
		},
		// 认证函数，用于验证用户登录信息
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req user.LoginReq
			err = c.BindAndValidate(&req)
			if err != nil {
				c.String(consts.StatusBadRequest, err.Error())
				return nil, err
			}
			LoginService := service.NewLoginService(ctx)
			resp, err := LoginService.Run(&req)
			if err != nil {
				c.String(consts.StatusInternalServerError, fmt.Sprintf("Registration failed: %v", err))
				return nil, err
			}
			return resp.UserId, nil
		},
		IdentityKey: IdentityKey, // 设置用于标识用户身份的键
		// 从 JWT 载荷中提取用户身份信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c) // 提取 JWT 载荷中的声明
			return &model.User{
				ID: claims[IdentityKey].(uint), // 使用声明中的身份键获取用户ID
			}
		},
		// 将用户数据转换为 JWT 载荷中的声明
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID, // 将用户ID存储在 JWT 载荷中
				}
			}
			return jwt.MapClaims{} // 返回空的声明
		},
		// 自定义 HTTP 状态消息函数，用于记录错误日志并返回错误消息
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error()) // 记录错误日志
			return e.Error()                                    // 返回错误消息
		},
		// 自定义未授权响应
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,    // 状态码
				"message": message, // 错误消息
			})
		},
	})
	// 如果初始化过程中出现错误，程序将崩溃并输出错误信息
	if err != nil {
		panic(err)
	}
}
