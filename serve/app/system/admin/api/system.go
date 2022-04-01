package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/system/admin/utlis/captcha/imageCaptcha"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

var System = new(systemApi)

type systemApi struct {
}

// LoadRouter 加载 authController 路由
func (c *systemApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.GET("/imageCaptcha", c.getImageCaptcha)
		group.GET("/info", c.getInfo)
		group.Middleware(middleware.CheckAuth)
		group.GET("/setting", c.getKey)
		group.Middleware(middleware.CheckTest)
		group.POST("/save", c.postSave)
	})
}
func (c *systemApi) getInfo(r *ghttp.Request) {
	// 业务处理
	if res, err := service.System.Info(); err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	} else {
		response.Success(r).SetData(res).Send()
	}
}

func (c *systemApi) getKey(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置配置key",
	}
	key := r.GetQueryString("key")
	if err := gvalid.Check(key, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	// 业务处理
	if res, code := service.System.Value(key); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", res)
		response.Success(r).SetData(data).Send()
	}
}

func (c *systemApi) postSave(r *ghttp.Request) {

	var req *dto.SystemCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	if code := service.System.Save(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}

}

// 获取 图片验证码
func (c *systemApi) getImageCaptcha(r *ghttp.Request) {
	if rp, err := imageCaptcha.Create(); err != nil {
		response.Error(r).SetCode(response.INVALID).
			SetMessage(response.CodeMsg(response.INVALID)).Send()
	} else {
		data := gmap.New(true)
		data.Set("imageCaptcha", rp)
		response.Success(r).SetData(data).Send()
	}
}
