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

var Question = new(questionApi)

type questionApi struct {
}

func (c *questionApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/question", func(group *ghttp.RouterGroup) {
		group.GET("/info", c.getInfo)
		group.GET("/list", c.getList)
		group.POST("/like", c.postLike)
		group.POST("/favorite", c.postFavorite)
		group.GET("/create/meta", c.getCreatMeta)
		group.POST("/create", c.postCreate)
		group.POST("/remove", c.postRemove)
	})
}

func (c *questionApi) getInfo(r *ghttp.Request) {
	rule := "integer|min:1"
	msg := map[string]string{
		"integer": "类型不正确，请设置整型",
		"min":     "id最小:min位",
	}
	id := r.GetQueryInt64("id")
	if err := gvalid.Check(id, rule, msg); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}
	tokenUserId, _ := service.Auth.GetTokenId(r)

	if result, code := service.Question.SelectInfo(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *questionApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.QuestionQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Question.SelectList(req); code != response.SUCCESS {
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

func (c *questionApi) postLike(r *ghttp.Request) {
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
	if lock_utils.CheckLock(shared.QuestionLikeLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Question.Like(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *questionApi) postFavorite(r *ghttp.Request) {
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
	if lock_utils.CheckLock(shared.QuestionFavoriteLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Question.Favorite(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *questionApi) getCreatMeta(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)

	if err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置id").
			Send()
	}
	data := gmap.New(true)
	// 获取加入的小组
	groupList, err := service.Group.MyJoinGroupList(tokenUserId)
	if err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).
			Send()
	}
	data.Set("groupList", groupList)
	response.Success(r).SetData(data).Send()
}

func (c *questionApi) postCreate(r *ghttp.Request) {
	var req *dto.QuestionCreate
	if err := r.Parse(&req); err != nil {
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
	if lock_utils.CheckLock(shared.QuestionCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}

	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, shared.Question) {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	// 检查用户是否加入圈子
	//if !service.Group.CheckIsJoin(tokenUserId, req.GroupId) {
	//	response.Error(r).SetCode(response.INVALID).SetMessage("你还没有加入圈子").Send()
	//}

	req.UserId = tokenUserId
	if code := service.Question.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetMessage("创建成功").Send()
	}
}

func (c *questionApi) postRemove(r *ghttp.Request) {
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

	if code := service.Question.Remove(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
