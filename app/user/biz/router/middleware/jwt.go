package mw // 定义包名为 mw，通常用于存放中间件相关的代码

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	// 导入项目内部的包
	// 数据库访问层，用于检查用户信息
	"github.com/cloudwego/hertz/pkg/app"          // Hertz 框架的 app 包，处理请求上下文
	"github.com/cloudwego/hertz/pkg/common/hlog"  // Hertz 的日志包
	"github.com/cloudwego/hertz/pkg/common/utils" // Hertz 的工具包，包含辅助函数
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt" // Hertz 的 JWT 中间件包
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
		Realm:      "test zone",          // 认证领域，用于在 WWW-Authenticate 头中返回
		Key:        []byte("secret key"), // 用于签名 JWT 的密钥，请确保使用足够复杂且安全的密钥
		Timeout:    time.Hour,            // JWT 的有效期，此处设置为 1 小时
		MaxRefresh: time.Hour,            // 允许刷新 JWT 的最大时间，此处设置为 1 小时
		// 表示在解析请求时，会尝试从以下几处获取 Token：
		// HTTP Header 中的 Authorization 字段
		// URL 查询参数 ?token=xxx
		// Cookie 名为 jwt
		// （按照这个顺序依次查找）
		TokenLookup:   "cookie: jwt", // 定义从哪里查找 JWT
		TokenHeadName: "Bearer",      // JWT 在请求头中的前缀
		// 自定义登录成功后的响应格式
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			// 直接set一个http only cookie 给客户端
			c.SetCookie(
				"jwt", // Cookie名称
				token, // Cookie值
				3600,  // 过期时间(秒)
				"/",   // 路径
				"",    // 域名(留空表示当前域)
				protocol.CookieSameSiteDefaultMode,
				true, // Secure
				true, // HttpOnly
			)
			c.JSON(http.StatusOK, utils.H{
				"code":    code,                        // 状态码
				"token":   token,                       // 生成的 JWT Token
				"expire":  expire.Format(time.RFC3339), // 过期时间
				"message": "success",
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
				c.String(consts.StatusInternalServerError, fmt.Sprintf("Login failed: %v", err))
				return nil, err
			}
			return resp.UserId, nil
		},
		IdentityKey: IdentityKey, // 设置用于标识用户身份的键
		// 从 JWT 载荷中提取用户身份信息 当挂载中间键时 会自动提取cookie中的jwt到ctx中
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			// 先判断有没有 "identity" 字段
			val, ok := claims[IdentityKey]
			if !ok {
				// 这里说明没有 identity，可能是未携带 token
				// 你可以选择直接返回 nil，或者返回一个默认 user
				log.Printf("no identity in token")
				return nil
			}
			// 再断言为 float64 (JWT 解析数字通常变成 float64)
			f64, ok := val.(float64)
			if !ok {
				// 万一断言失败，也可视为未登录或解析失败
				log.Printf("invalid identity in token")
				return "-1"
			}
			// 再转成 uint / int64
			log.Printf("identity in token: %v", f64)
			return int64(f64)
		},
		// 将用户数据转换为 JWT 载荷中的声明
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			switch v := data.(type) {
			case uint:
				return jwt.MapClaims{IdentityKey: v}
			case int64:
				return jwt.MapClaims{IdentityKey: v}
			case int32:
				// 这里把 int32 转成 int64 或直接存 float64
				return jwt.MapClaims{IdentityKey: int64(v)}
			default:
				return jwt.MapClaims{}
			}
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
