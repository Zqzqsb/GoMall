package mw

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
)

func InitCSRF(h *server.Hertz) {
	h.Use(csrf.New(
		csrf.WithSecret("scrf-secret"),
		csrf.WithKeyLookUp("header:csrf"),
		// csrf.WithNext(func(c context.Context, ctx *app.RequestContext) bool {
		// 	// 如果当前请求路径是 /login，则跳过 CSRF 校验
		// 	// 你也可以按需检查请求方法，或其他自定义逻辑
		// 	if string(ctx.Request.URI().Path()) == "/login" {
		// 		return true
		// 	}
		// 	return false
		// }),
		csrf.WithErrorFunc(func(c context.Context, ctx *app.RequestContext) {
			ctx.String(400, ctx.Errors.Last().Error())
			ctx.Abort()
		}),
	))
	log.Println("init csrf success")
}
