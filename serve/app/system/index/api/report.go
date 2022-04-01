package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Report = new(reportApi)

type reportApi struct {
}

func (c *reportApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/report", func(group *ghttp.RouterGroup) {
		group.POST("/create", c.postCreate)
	})
}

func (c *reportApi) postCreate(r *ghttp.Request) {
	var req *dto.ReportCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.ReportCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}

	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, shared.Report) {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	if err := service.Report.Create(tokenUserId, req); err != nil {
		response.Error(r).
			SetCode(response.ADD_FAILED).
			SetMessage(response.CodeMsg(response.ADD_FAILED)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).SetMessage("操作成功").Send()
	}
}
