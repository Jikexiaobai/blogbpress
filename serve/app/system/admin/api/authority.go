package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/system/admin/valid"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

var Authority = new(authorityApi)

type authorityApi struct{}

// LoadRouter 加载 authController 路由
func (c *authorityApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/authority", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/edit", c.postEdit)
		group.POST("/create", c.postCreate)
		group.POST("/remove", c.postRemove)
	})
}

func (c *authorityApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.AuthorityQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if result, code := service.Authority.SelectAll(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()

	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).
			SetData(data).Send()
	}
}

func (c *authorityApi) postCreate(r *ghttp.Request) {
	var req *dto.AuthorityCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Type != 3 && req.Path == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置路由地址").
			Send()
	}

	if req.Type != 3 && req.Component == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置组件").
			Send()
	}

	if req.Type != 3 && req.Hidden == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置是否隐藏").
			Send()
	}

	if req.Type != 3 && req.Target == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置跳转方式").
			Send()
	}

	if req.Type != 3 && req.Target == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置跳转方式").
			Send()
	}

	if valid.Authority.CheckAuthorityUnique(req.Perms) {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("权限已存在").
			Send()
	}

	if code := service.Authority.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {

		response.Success(r).Send()
	}
}

func (c *authorityApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.Authority.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *authorityApi) postEdit(r *ghttp.Request) {
	var req *dto.AuthorityEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Type != 3 && req.Path == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置路由地址").
			Send()
	}

	if req.Type != 3 && req.Component == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置组件").
			Send()
	}

	if req.Type != 3 && req.Hidden == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置是否隐藏").
			Send()
	}

	if req.Type != 3 && req.Target == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置跳转方式").
			Send()
	}

	if req.Type != 3 && req.Target == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置跳转方式").
			Send()
	}

	if valid.Authority.CheckAuthorityUniqueExceptYourself(req.AuthorityId, req.Perms) {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("权限已存在").
			Send()
	}

	if code := service.Authority.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *authorityApi) postRemove(r *ghttp.Request) {
	rule := "integer|min:1"
	msg := map[string]string{
		"integer": "类型不正确，请设置整型",
		"min":     "id长度:min位",
	}
	id := r.GetFormInt64("id")
	if err := gvalid.Check(id, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Authority.Remove(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
