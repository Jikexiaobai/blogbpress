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

var Resource = new(resourceApi)

type resourceApi struct{}

func (c *resourceApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/resource", func(group *ghttp.RouterGroup) {
		group.GET("/info", c.getInfo)
		group.GET("/list", c.getList)
		group.POST("/like", c.postLike)
		group.POST("/favorite", c.postFavorite)
		group.GET("/create/meta", c.getCreatMeta)
		group.POST("/create", c.postCreate)
		group.GET("/edit/info", c.getEditInfo)
		group.POST("/edit", c.postEdit)
		group.POST("/remove", c.postRemove)
	})
}

func (c *resourceApi) getInfo(r *ghttp.Request) {
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

	if result, code := service.Resource.SelectInfo(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		if result == nil {
			response.Success(r).SetCode(response.NOT_FOUND).SetMessage(response.CodeMsg(response.NOT_FOUND)).SetData(data).Send()
		}
		response.Success(r).SetCode(response.SUCCESS).SetMessage("获取成功").SetData(data).Send()
	}
}

func (c *resourceApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.ResourceQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Resource.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).
			SetData(data).Send()
	}
}

func (c *resourceApi) postLike(r *ghttp.Request) {
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
	if lock_utils.CheckLock(shared.ResourceLikeLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Resource.Like(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *resourceApi) postFavorite(r *ghttp.Request) {
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
	if lock_utils.CheckLock(shared.ResourceFavoriteLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Resource.Favorite(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *resourceApi) getCreatMeta(r *ghttp.Request) {
	data := gmap.New(true)
	// 获取加入分类
	categoryList, code := service.Category.SelectListByModule(shared.Resource)
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	}
	data.Set("cateList", categoryList)

	// 获取热门标签
	tagList, err := service.Tag.SelectHotTagList()
	if err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	}
	data.Set("tagList", tagList)

	response.Success(r).SetData(data).Send()
}

func (c *resourceApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.ResourceCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if len(req.Tags) > 5 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("标签太多，最多只能设置5个").Send()
	}

	//if len(req.GroupId) > 5 {
	//	response.Error(r).
	//		SetCode(response.PARAM_INVALID).
	//		SetMessage("圈子太多，最多只能设置5个").Send()
	//}

	// 有下载模块
	if req.HasDown == 2 {
		if req.DownMode == shared.DownModePay && req.Price == 0 {
			response.Error(r).SetCode(response.PARAM_INVALID).
				SetMessage("请设置下载价格").Send()
		}
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	//检查是否有权限操作
	if !service.Grade.CheckHasPosts(tokenUserId, shared.Resource) {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.ResourceCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	req.UserId = tokenUserId
	if code := service.Resource.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *resourceApi) getEditInfo(r *ghttp.Request) {
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
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	if info, code := service.Resource.EditInfo(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *resourceApi) postEdit(r *ghttp.Request) {
	var req *dto.ResourceEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if len(req.Tags) > 5 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("标签太多，最多只能设置5个").Send()
	}

	//if len(req.GroupId) > 5 {
	//	response.Error(r).
	//		SetCode(response.PARAM_INVALID).
	//		SetMessage("圈子太多，最多只能设置5个").Send()
	//}

	// 有下载模块

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	//检查是否用户被锁
	if lock_utils.CheckLock(shared.ResourceEditLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	req.UserId = tokenUserId
	if code := service.Resource.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *resourceApi) postRemove(r *ghttp.Request) {
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

	if code := service.Resource.Remove(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
