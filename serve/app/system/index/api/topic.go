package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Topic = new(topicApi)

type topicApi struct {
}

func (c *topicApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/topic", func(group *ghttp.RouterGroup) {
		group.GET("/info", c.getInfo)
		group.GET("/list", c.getList)
		group.GET("/top", c.getTop)
		group.POST("/like", c.postLike)
		group.POST("/create", c.postCreate)
		group.POST("/remove", c.postRemove)
	})
}

func (c *topicApi) getInfo(r *ghttp.Request) {
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

	tokenUserId, _ := service.Auth.GetTokenId(r)
	if result, code := service.Topic.SelectInfo(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *topicApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.TopicQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	tokenUserId, _ := service.Auth.GetTokenId(r)
	if req.Mode == 5 && tokenUserId == 0 {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage("你还没有登录").Send()
	}
	req.UserId = tokenUserId

	if total, result, code := service.Topic.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).
			SetData(data).Send()
	}
}

func (c *topicApi) getTop(r *ghttp.Request) {
	if result, code := service.Topic.SelectTop(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).
			SetData(data).Send()
	}
}

func (c *topicApi) postLike(r *ghttp.Request) {
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

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.TopicLikeLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Topic.Like(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *topicApi) postCreate(r *ghttp.Request) {
	var req *dto.TopicCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	//if req.Type == 2 && (req.Module == "" || req.RelatedId == 0) && valid.Topic.CheckModule(req.Module) {
	//	response.Error(r).SetCode(response.INVALID).SetMessage("请设置转发内容信息").Send()
	//}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.TopicCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, shared.Topic) {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	// 检查用户是否加入圈子
	//if !service.Group.CheckIsJoin(tokenUserId, req.GroupId) {
	//	response.Error(r).SetCode(response.INVALID).SetMessage("你还没有加入圈子").Send()
	//}

	req.UserId = tokenUserId
	if code := service.Topic.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *topicApi) postRemove(r *ghttp.Request) {
	rule := "integer|min:1"
	msg := map[string]string{
		"integer": "类型不正确，请设置整型",
		"min":     "id最小:min位",
	}
	id := r.GetFormInt64("id")
	if err := gvalid.Check(id, rule, msg); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	if code := service.Topic.Remove(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
