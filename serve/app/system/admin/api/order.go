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

var Order = new(orderApi)

type orderApi struct{}

func (c *orderApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/order", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/info", c.getInfo)
		group.Middleware(middleware.CheckTest)
		//group.POST("/review", c.postReview)
		group.POST("/remove", c.postRemove)
	})
}

func (c *orderApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.OrderQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Order.SelectList(req); code != response.SUCCESS {
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
			SetData(data).Send()
	}
}

func (c *orderApi) getInfo(r *ghttp.Request) {
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

	if info, code := service.Order.SelectInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		if info == nil {
			response.Success(r).
				SetCode(response.NOT_FOUND).
				SetMessage(response.CodeMsg(response.NOT_FOUND)).SetData(data).Send()
		}
		response.Success(r).
			SetCode(response.SUCCESS).SetData(data).Send()
	}
}

//func (c *orderApi) postReview(r *ghttp.Request) {
//	var req *dto.Review
//	if err := r.Parse(&req); err != nil {
//		response.Error(r).
//			SetCode(response.PARAM_INVALID).
//			SetMessage(err.Error()).
//			Send()
//	}
//	if req.Remark == "" {
//		response.Error(r).
//			SetCode(response.PARAM_INVALID).
//			SetMessage("请填写处理结果").
//			Send()
//	}
//
//	if err := service.Report.Review(req); err != nil {
//		response.Error(r).
//			SetCode(response.UPDATE_FAILED).
//			SetMessage(response.CodeMsg(response.UPDATE_FAILED)).Send()
//	} else {
//		response.Success(r).SetCode(response.SUCCESS).Send()
//	}
//}

func (c *orderApi) postRemove(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置删除的id",
	}
	ids := r.GetFormInts("idList")
	if err := gvalid.Check(ids, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if code := service.Order.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
