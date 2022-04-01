package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

var Notice = new(noticeApi)

type noticeApi struct{}

// LoadRouter 加载 authController 路由
func (c *noticeApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/notice", func(group *ghttp.RouterGroup) {
		group.GET("/list", c.getList)
		group.GET("/count", c.getCount)
	})
}

func (c *noticeApi) getCount(r *ghttp.Request) {
	tokenUserId, _ := service.Auth.GetTokenId(r)
	if result, code := service.Notice.SelectCount(tokenUserId); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()

	}
}

func (c *noticeApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.QueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Type == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置对应消息类型").
			Send()
	}

	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Notice.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()

	}
}
