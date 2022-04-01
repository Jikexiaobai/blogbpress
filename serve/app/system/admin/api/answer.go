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

var Answer = new(answerApi)

type answerApi struct{}

func (c *answerApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/answer", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.Middleware(middleware.CheckTest)
		group.POST("/review", c.postReview)
		group.POST("/recover", c.postRecover)
		group.POST("/reduction", c.postReduction)
		group.POST("/remove", c.postRemove)
	})
}

func (c *answerApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.AnswerQueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Answer.SelectList(req); code != response.SUCCESS {
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

func (c *answerApi) postReview(r *ghttp.Request) {
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

	if code := service.Answer.Review(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *answerApi) postRecover(r *ghttp.Request) {
	var req *dto.Remove
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Answer.Recover(req); code != response.SUCCESS {
		response.Error(r).SetCode(response.DELETE_FAILED).SetMessage(response.CodeMsg(response.DELETE_FAILED)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *answerApi) postReduction(r *ghttp.Request) {
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
	if code := service.Answer.Reduction(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(response.FAILD).SetMessage(response.CodeMsg(response.FAILD)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *answerApi) postRemove(r *ghttp.Request) {
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
	if code := service.Answer.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(response.DELETE_FAILED).SetMessage(response.CodeMsg(response.DELETE_FAILED)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
