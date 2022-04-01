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

var Comment = new(commentApi)

type commentApi struct {
}

func (c *commentApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/comment", func(group *ghttp.RouterGroup) {
		group.GET("/list", c.getList)
		group.POST("/like", c.postLike)
		group.POST("/create", c.postCreate)
		group.POST("/remove", c.postRemove)
	})
}

func (c *commentApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.CommentQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, _ := service.Auth.GetTokenId(r)
	if total, result, err := service.Comment.SelectList(tokenUserId, req); err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).
			Send()
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

func (c *commentApi) postCreate(r *ghttp.Request) {
	var req *dto.CommentCreate
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
	if lock_utils.CheckLock(shared.CommentCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}

	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, shared.Comment) {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	if res, code := service.Comment.Create(tokenUserId, req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", res)
		response.Success(r).SetData(data).Send()
	}
}

func (c *commentApi) postLike(r *ghttp.Request) {
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
	if lock_utils.CheckLock(shared.CommentLikeLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Comment.Like(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *commentApi) postRemove(r *ghttp.Request) {
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

	if code := service.Comment.Remove(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
