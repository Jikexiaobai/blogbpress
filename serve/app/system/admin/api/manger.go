package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/system/admin/valid"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Manger = new(mangerApi)

type mangerApi struct {
}

// LoadRouter 加载 authController 路由
func (c *mangerApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/manger", func(group *ghttp.RouterGroup) {
		group.GET("/info", c.getAccount) //获取当前登录用户的信息
		group.GET("/authority", c.getAuthority)
		group.Middleware(middleware.CheckAuth)
		//group.GET("/create/meta", c.getCreateMeta)
		group.GET("/edit/info", c.getEditInfo)
		group.GET("/list", c.getList)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/edit", c.postEdit)
		//group.POST("/review", c.postReview)
		group.POST("/remove", c.postRemove)
	})
}

func (c *mangerApi) getAccount(r *ghttp.Request) {

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	//userAgent := r.Header.Get("User-Agent")
	//ipAddr := r.GetClientIp()

	if res, code := service.Manger.SelectInfo(tokenUserId); err != nil {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetData(res).Send()
	}
}

func (c *mangerApi) getAuthority(r *ghttp.Request) {

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	//userAgent := r.Header.Get("User-Agent")
	//ipAddr := r.GetClientIp()

	if res, code := service.Authority.SelectMangerAuthority(tokenUserId); err != nil {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetData(res).Send()
	}
}

func (c *mangerApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.MangerQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.Manger.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.
				CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).
			SetData(data).Send()
	}
}

func (c *mangerApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.MangerCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Email == "" && req.Phone == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置邮箱或手机号").Send()
	}

	if req.Phone != "" {
		if valid.Manger.CheckPhoneUniqueAll(req.Phone) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("手机号已存在").Send()
		}
	}

	if req.Email != "" {
		if valid.Manger.CheckEmailUniqueAll(req.Email) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("邮箱已存在").Send()
		}
	}

	if code := service.Manger.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *mangerApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.Manger.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

//func (c *mangerApi) postReview(r *ghttp.Request) {
//	var req *dto.Review
//	if err := r.Parse(&req); err != nil {
//		response.Error(r).
//			SetCode(response.PARAM_INVALID).
//			SetMessage(err.Error()).
//			Send()
//	}
//	if req.Status == 1 && req.Remark == "" {
//		response.Error(r).
//			SetCode(response.PARAM_INVALID).
//			SetMessage("请填写禁用原因").
//			Send()
//	}
//
//	if err := service.User.Review(req); err != nil {
//		response.Error(r).
//			SetCode(response.UPDATE_FAILED).
//			SetMessage(response.CodeMsg(response.UPDATE_FAILED)).Send()
//	} else {
//		response.Success(r).SetCode(response.SUCCESS).Send()
//	}
//}

func (c *mangerApi) postEdit(r *ghttp.Request) {
	var req *dto.MangerEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Email == "" && req.Phone == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置邮箱或手机号").Send()
	}

	if req.Phone != "" {
		if valid.Manger.CheckPhoneUniqueAll(req.Phone, req.UserId) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("手机号已存在").Send()
		}
	}

	if req.Email != "" {
		if valid.Manger.CheckEmailUniqueAll(req.Email, req.UserId) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("邮箱已存在").Send()
		}
	}

	if code := service.Manger.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *mangerApi) postRemove(r *ghttp.Request) {
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
	if code := service.Manger.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
