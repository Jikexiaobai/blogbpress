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

var Edu = new(eduApi)

type eduApi struct{}

func (c *eduApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/edu", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/joinList", c.getJoin)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/review", c.postReview)
		group.POST("/edit", c.postEdit)
		group.POST("/recover", c.postRecover)
		group.POST("/reduction", c.postReduction)
		group.POST("/remove", c.postRemove)
	})
}

func (c *eduApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.EduQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Edu.SelectList(req); code != response.SUCCESS {
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

func (c *eduApi) getJoin(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.EduJoinQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Edu.SelectUserJoinList(req); code != response.SUCCESS {
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
			SetData(data).Send()
	}
}

func (c *eduApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.EduCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if len(req.Tags) > 5 {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage("标签太多，最多只能设置5个").Send()
	}

	if req.JoinMode == 2 && req.Price == 0 {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage("请设置报名价格").Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	if req.UserId == 0 {
		req.UserId = tokenUserId
	}
	if code := service.Edu.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {

		response.Success(r).SetCode(response.SUCCESS).SetMessage("创建成功").Send()
	}
}

func (c *eduApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.Edu.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		if info == nil {
			response.Success(r).
				SetCode(response.NOT_FOUND).
				SetMessage(response.CodeMsg(response.NOT_FOUND)).SetData(data).Send()
		}
		response.Success(r).
			SetCode(response.SUCCESS).
			SetMessage("获取成功").SetData(data).Send()
	}
}

func (c *eduApi) postReview(r *ghttp.Request) {
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

	if code := service.Edu.Review(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *eduApi) postEdit(r *ghttp.Request) {
	var req *dto.EduEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if len(req.Tags) > 5 {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage("标签太多，最多只能设置5个").Send()
	}

	if req.JoinMode == 2 && req.Price == 0 {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage("请设置报名价格").Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	if req.UserId == 0 {
		req.UserId = tokenUserId
	}
	if code := service.Edu.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *eduApi) postRecover(r *ghttp.Request) {
	var req *dto.Remove
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if code := service.Edu.Recover(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *eduApi) postReduction(r *ghttp.Request) {
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
	if code := service.Edu.Reduction(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *eduApi) postRemove(r *ghttp.Request) {
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
	if code := service.Edu.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
