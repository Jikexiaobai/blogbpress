package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

var Vip = new(vipApi)

type vipApi struct {
}

// LoadRouter 加载 authController 路由
func (c *vipApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/vip", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/edit", c.postEdit)
	})
}

func (c *vipApi) getList(r *ghttp.Request) {
	if result, code := service.Vip.SelectList(); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *vipApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.VipCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Vip.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {

		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *vipApi) getEditInfo(r *ghttp.Request) {
	rule := "integer|min:1"
	msg := map[string]string{
		"integer": "类型不正确，请设置整型",
		"min":     "id长度:min位",
	}
	id := r.GetQueryInt64("id")
	if err := gvalid.Check(id, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if info, code := service.Vip.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).
			SetCode(response.SUCCESS).SetData(data).Send()
	}
}

func (c *vipApi) postEdit(r *ghttp.Request) {
	var req *dto.VipEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Vip.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
