package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/system/admin/shared"
	"fiber/app/system/admin/valid"
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
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/create/meta", c.getCreateMeta)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/edit", c.postEdit)
		group.POST("/review", c.postReview)
		group.POST("/remove", c.postRemove)
	})
}

func (c *userApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.UserQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if total, result, code := service.User.SelectList(req); code != response.SUCCESS {
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

func (c *userApi) getCreateMeta(r *ghttp.Request) {
	data := gmap.New(true)
	// 获取加入分类
	gradeList, code := service.Grade.SelectList()
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	}
	data.Set("gradeList", gradeList)

	// 获取加入分类
	vipList, code := service.Vip.SelectList()
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	}
	data.Set("vipList", vipList)

	response.Success(r).
		SetCode(response.SUCCESS).SetData(data).Send()
}

func (c *userApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.UserCreate
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
		if valid.User.CheckPhoneUniqueAll(req.Phone) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("手机号已存在").Send()
		}
	}

	if req.Phone != "" {
		if valid.User.CheckEmailUniqueAll(req.Email) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("邮箱已存在").Send()
		}
	}

	if code := service.User.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {

		response.Success(r).SetCode(response.SUCCESS).SetMessage("创建成功").Send()
	}
}

func (c *userApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.User.EditInfo(id); code != response.SUCCESS {
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

func (c *userApi) postReview(r *ghttp.Request) {
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

	if code := service.User.Review(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *userApi) postEdit(r *ghttp.Request) {
	var req *dto.UserEdit
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
		if valid.User.CheckPhoneUniqueAll(req.Phone, req.UserId) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("手机号已存在").Send()
		}
	}

	if req.Email != "" {
		if valid.User.CheckEmailUniqueAll(req.Email, req.UserId) {
			response.Error(r).
				SetCode(response.PARAM_INVALID).
				SetMessage("邮箱已存在").Send()
		}
	}

	if code := service.User.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *userApi) postRemove(r *ghttp.Request) {
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
	if code := service.User.Remove(gconv.Int64s(ids)); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
