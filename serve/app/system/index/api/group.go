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

var Group = new(groupApi)

type groupApi struct {
}

// LoadRouter 加载 authController 路由
func (c *groupApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/group", func(group *ghttp.RouterGroup) {
		group.GET("/info", c.getInfo)
		group.GET("/posts", c.getPosts)
		group.GET("/list", c.getList)
		group.POST("/join", c.postJoin)
		group.GET("/create/meta", c.getCreatMeta)
		group.POST("/create", c.postCreate)
		group.GET("/edit/info", c.getEditInfo)
		group.POST("/edit", c.postEdit)
		group.POST("/remove", c.postRemove)

	})
}

func (c *groupApi) getInfo(r *ghttp.Request) {
	id := r.GetQueryInt64("id")
	if id == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置id").
			Send()
	}

	tokenUserId, _ := service.Auth.GetTokenId(r)
	if result, code := service.Group.SelectInfo(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *groupApi) getPosts(r *ghttp.Request) {

	var req *dto.QueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	//tokenUserId, _ := service.Auth.GetTokenId(r)
	//req.UserId = tokenUserId

	if req.GroupId == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(response.CodeMsg(response.PARAM_INVALID)).
			Send()
	}

	if req.Module == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(response.CodeMsg(response.PARAM_INVALID)).
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

func (c *groupApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.GroupQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Group.SelectList(req); code != response.SUCCESS {
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

func (c *groupApi) postJoin(r *ghttp.Request) {
	rule := "integer|min:1"
	msg := map[string]string{
		"integer": "类型不正确，请设置整型",
		"min":     "id长度:min",
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
	secretKey := r.GetFormString("secretKey")
	// 检查是否加入
	if !service.Group.CheckIsJoin(tokenUserId, id) {
		// 检查能否加入

		if !service.Group.CheckCanJoin(tokenUserId, id, secretKey) {
			response.Error(r).
				SetCode(response.FAILD).
				SetMessage(response.CodeMsg(response.FAILD)).Send()
		}
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.JoinGroupLock + gconv.String(tokenUserId) + gconv.String(id)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	if code := service.Group.Join(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *groupApi) getCreatMeta(r *ghttp.Request) {
	data := gmap.New(true)
	// 获取加入分类
	categoryList, code := service.Category.SelectListByModule(shared.Group)
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	}
	data.Set("cateList", categoryList)

	response.Success(r).SetData(data).Send()
}

func (c *groupApi) postCreate(r *ghttp.Request) {
	var req *dto.GroupCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.JoinMode == 2 && req.Price == 0 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置付费加入价格").Send()
	}

	if req.JoinMode == 3 && req.SecretKey == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置加入密钥").Send()
	}
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否有权限操作
	if !service.Grade.CheckHasPosts(tokenUserId, shared.Group) {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	if !service.Grade.CheckHasCreateGroup(tokenUserId) {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("超出等级的创建数量").Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.GroupCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	// 业务处理
	req.UserId = tokenUserId
	if code := service.Group.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *groupApi) getEditInfo(r *ghttp.Request) {
	rule := "integer|min:1"
	msg := map[string]string{
		"integer": "类型不正确，请设置整型",
		"min":     "id长度:min位",
	}
	Id := r.GetQueryInt64("id")
	if err := gvalid.Check(Id, rule, msg); err != nil {
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
	if info, code := service.Group.EditInfo(tokenUserId, Id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *groupApi) postEdit(r *ghttp.Request) {
	var req *dto.GroupEdit
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
	if lock_utils.CheckLock(shared.GroupEditLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}
	req.UserId = tokenUserId
	if code := service.Group.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}

}

func (c *groupApi) postRemove(r *ghttp.Request) {
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

	if code := service.Group.Remove(tokenUserId, id); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
