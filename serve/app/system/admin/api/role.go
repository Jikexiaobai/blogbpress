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

var Role = new(roleApi)

type roleApi struct {
}

// LoadRouter 加载 authController 路由
func (c *roleApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/role", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/edit", c.postEdit)
	})
}

func (c *roleApi) getList(r *ghttp.Request) {
	if result, code := service.Role.SelectList(); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *roleApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.RoleCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if len(req.Authority) < 1 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置角色权限").
			Send()
	}
	if code := service.Role.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *roleApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.Role.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *roleApi) postEdit(r *ghttp.Request) {
	var req *dto.RoleEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if len(req.Authority) < 1 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置角色权限").
			Send()
	}

	if code := service.Role.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
