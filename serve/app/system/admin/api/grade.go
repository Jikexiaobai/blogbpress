package api

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/middleware"
	"fiber/app/system/admin/service"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

var Grade = new(gradeApi)

type gradeApi struct {
}

// LoadRouter 加载 authController 路由
func (c *gradeApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/grade", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CheckAuth)
		group.GET("/list", c.getList)
		group.GET("/edit/info", c.getEditInfo)
		group.Middleware(middleware.CheckTest)
		group.POST("/create", c.postCreate)
		group.POST("/edit", c.postEdit)
	})
}

func (c *gradeApi) getList(r *ghttp.Request) {
	if result, code := service.Grade.SelectList(); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).SetCode(response.SUCCESS).SetMessage("获取成功").SetData(data).Send()
	}
}

func (c *gradeApi) postCreate(r *ghttp.Request) {
	// 判断用户是否有权限发布

	var req *dto.GradeCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if len(req.PostsModule) < 1 || len(req.CommonAuth) < 1 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置等级权限").
			Send()
	}
	if code := service.Grade.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {

		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}

func (c *gradeApi) getEditInfo(r *ghttp.Request) {
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

	if info, code := service.Grade.EditInfo(id); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).
			SetCode(response.SUCCESS).SetData(data).Send()
	}
}

func (c *gradeApi) postEdit(r *ghttp.Request) {
	var req *dto.GradeEdit
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	if len(req.PostsModule) < 1 || len(req.CommonAuth) < 1 {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置等级权限").
			Send()
	}
	if code := service.Grade.Edit(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetCode(response.SUCCESS).Send()
	}
}
