package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/system/admin/shared"
	"fiber/app/system/admin/valid"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Cash = new(cashApi)

type cashApi struct{}

func (c *cashApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/cash", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.Middleware(middleware.CheckTest)
		group.POST("/review", c.postReview)
		group.POST("/remove", c.postRemove)
	})
}

func (c *cashApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.CashQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Cash.SelectList(req); code != response.SUCCESS {
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

func (c *cashApi) postReview(r *ghttp.Request) {
	var req *dto.CashReview
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Status == shared.StatusReviewed && req.ReceiptNum == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置转账收据号").
			Send()
	}

	if valid.Cash.CheckIsCash(req.Code) {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("订单已经提现").
			Send()
	}

	if code := service.Cash.Review(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *cashApi) postRemove(r *ghttp.Request) {
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
	if code := service.Cash.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
