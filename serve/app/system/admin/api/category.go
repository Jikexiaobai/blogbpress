package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Category = new(categoryApi)

type categoryApi struct{}

func (c *categoryApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/category", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/module", c.getModuleList)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/edit", c.postEdit)
		group.POST("/remove", c.postRemove)
	})
}

func (c *categoryApi) getModuleList(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置分类模块",
	}
	module := r.GetString("module")

	if err := gvalid.Check(module, rule, msg); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	categoryList, code := service.Category.SelectModuleList(module)
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	}
	data := gmap.New(true)
	// 获取加入分类
	data.Set("list", categoryList)
	response.Success(r).
		SetCode(response.SUCCESS).
		SetMessage("获取成功").SetData(data).Send()
}

func (c *categoryApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.CategoryQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Category.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.
				CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).
			SetCode(response.SUCCESS).
			SetMessage("获取成功").
			SetData(data).Send()
	}
}

func (c *categoryApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.CategoryCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Category.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *categoryApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.Category.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *categoryApi) postEdit(r *ghttp.Request) {
	var req *dto.CategoryEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Category.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *categoryApi) postRemove(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置删除id列表",
	}
	ids := r.GetFormInts("idList")

	if err := gvalid.Check(ids, rule, msg); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	idList := gconv.Int64s(ids)
	if code := service.Category.Remove(idList); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
