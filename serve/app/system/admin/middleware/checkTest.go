package middleware

import (
	"fiber/app/system/admin/service"
	"fiber/app/tools/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func CheckTest(r *ghttp.Request) {

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	isDemo := g.Cfg().GetBool("system.IsDemo")
	if isDemo && tokenUserId != 1 {
		response.Error(r).
			SetCode(response.AUTH_ERROR).
			SetMessage("演示请不要操作").Send()
	} else {
		r.Middleware.Next()
	}
}
