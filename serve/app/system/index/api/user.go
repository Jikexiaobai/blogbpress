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

var User = new(userApi)

type userApi struct {
}

// LoadRouter 加载 authController 路由
func (c *userApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/user", func(group *ghttp.RouterGroup) {
		group.GET("/info", c.getUserInfo)
		group.GET("/posts", c.getPosts)

		group.GET("/sign", c.getSign)
		group.GET("/fansOrFollows", c.getFansOrFollows)
		group.GET("/reward", c.getReward)
		group.POST("/follow", c.postFollow)
	})
}

func (c *userApi) getUserInfo(r *ghttp.Request) {

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
	token := r.Header.Get("Authorization")
	var tokenUserId int64
	if token != "" {
		tokenUserId, _ = service.Auth.GetTokenId(r)
	}

	if info, code := service.User.Info(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *userApi) getPosts(r *ghttp.Request) {

	var req *dto.QueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if total, result, code := service.System.SelectPostsList(req); code != response.SUCCESS {
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

func (c *userApi) getSign(r *ghttp.Request) {

	var req *dto.UserSignQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if total, result, code := service.User.SelectSignList(req); code != response.SUCCESS {
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

func (c *userApi) getFansOrFollows(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.UserFansOrFollowsQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	tokenUserId, _ := service.Auth.GetTokenId(r)

	if total, list, code := service.User.SelectFansOrFollowsList(tokenUserId, req); code != response.SUCCESS {
		response.Error(r).SetCode(response.DB_READ_ERROR).SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", list)
		response.Success(r).SetData(data).Send()
	}
}

func (c *userApi) getReward(r *ghttp.Request) {

	//获取QueryParam
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

	if total, list, code := service.User.SelectRewardList(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", list)
		response.Success(r).SetData(data).Send()
	}
}

func (c *userApi) postFollow(r *ghttp.Request) {
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
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.FollowLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}

	if id == tokenUserId {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage("请不要关注自己").Send()
	}

	if code := service.User.Follow(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
