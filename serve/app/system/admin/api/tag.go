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

var Tag = new(tagApi)

type tagApi struct{}

func (c *tagApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/tag", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/hots", c.getHots)
		group.GET("/list", c.getList)
		group.Middleware(middleware.CheckTest)
		group.POST("/top", c.postTop)
		group.POST("/remove", c.postRemove)
	})
}

func (c *tagApi) getHots(r *ghttp.Request) {

	// 获取热门标签
	tagList, code := service.Tag.SelectHotList()
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	}

	data := gmap.New(true)
	// 获取加入分类
	data.Set("list", tagList)
	response.Success(r).
		SetCode(response.SUCCESS).
		SetMessage("获取成功").SetData(data).Send()
}

func (c *tagApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.TagQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Tag.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.
				CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).
			SetData(data).Send()
	}
}

func (c *tagApi) postTop(r *ghttp.Request) {
	var req *dto.TagTop
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Tag.SetTop(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *tagApi) postRemove(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置删除id列表",
	}
	ids := r.GetFormInts("idList")

	if err := gvalid.Check(ids, rule, msg); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	idList := gconv.Int64s(ids)
	if code := service.Tag.Remove(idList); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
