package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Comment = new(commentApi)

type commentApi struct{}

func (c *commentApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/comment", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.Middleware(middleware.CheckTest)
		group.POST("/review", c.postReview)
		group.POST("/recover", c.postRecover)
		group.POST("/reduction", c.postReduction)
		group.POST("/remove", c.postRemove)
	})
}

func (c *commentApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.CommentQueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Comment.SelectList(req); code != response.SUCCESS {
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

func (c *commentApi) postReview(r *ghttp.Request) {
	var req *dto.Review
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if req.Remark == "" && req.Status == shared.StatusRefuse {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请填写，处理通知信息").
			Send()
	}

	if code := service.Comment.Review(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *commentApi) postRecover(r *ghttp.Request) {
	var req *dto.Remove
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Comment.Recover(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *commentApi) postReduction(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置还原的id",
	}
	ids := r.GetFormInts("idList")
	if err := gvalid.Check(ids, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if code := service.Comment.Reduction(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *commentApi) postRemove(r *ghttp.Request) {
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
	if code := service.Comment.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
