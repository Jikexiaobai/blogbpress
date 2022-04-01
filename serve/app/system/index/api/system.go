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

var System = new(systemApi)

type systemApi struct{}

func (c *systemApi) LoadRouter(group *ghttp.RouterGroup) {
	group.GET("/system/kr", c.getFiber)
	group.GET("/system/home", c.getHome)
	group.GET("/system/h5/home", c.getHFiveHome)
	group.GET("/system/info", c.getInfo)
	group.GET("/system/vipAndGrade", c.getVipAndGrade)
	group.GET("/system/hotTag", c.getHotTag)
	group.GET("/system/cate", c.getCate)
	group.GET("/system/filter", c.getFilter)
	group.GET("/system/search", c.getSearch)
	group.GET("/system/hotSearch", c.getHotSearch)
	group.GET("/system/hotUser", c.getHotUser)
}

func (c *systemApi) getHFiveHome(r *ghttp.Request) {

	if result, code := service.System.HFiveHome(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetData(result).Send()
	}
}

func (c *systemApi) getHome(r *ghttp.Request) {

	if result, code := service.System.Home(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).SetData(result).Send()
	}
}

func (c *systemApi) getFiber(r *ghttp.Request) {
	data := gmap.New(true)
	data.Set("welcome", "欢迎来到氪讯")
	response.Success(r).SetData(data).Send()
}

func (c *systemApi) getHotTag(r *ghttp.Request) {
	// 获取热门标签
	result, err := service.Tag.SelectHotTagList()
	if err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).
			Send()
	}
	data := gmap.New(true)
	data.Set("list", result)
	response.Success(r).SetData(data).Send()
}

func (c *systemApi) getCate(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置模块",
	}
	module := r.GetQueryString("module")
	if err := gvalid.Check(module, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	data := gmap.New(true)
	// 获取分类
	result, code := service.Category.SelectListByModule(module)
	if code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	}
	data.Set("list", result)
	response.Success(r).SetData(data).Send()
}

func (c *systemApi) getFilter(r *ghttp.Request) {
	//获取QueryParam
	var req *dto.QueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Module == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("模块不能为空").
			Send()
	}
	if req.Mode == shared.ModeFollow && req.Module == shared.Topic {

		tokenUserId, err := service.Auth.GetTokenId(r)
		if err != nil {
			response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
		}
		req.UserId = tokenUserId
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
		response.Success(r).SetData(data).Send()
	}
}

func (c *systemApi) getSearch(r *ghttp.Request) {
	//获取QueryParam
	var req *dto.QueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if req.Title == "" {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("你到底要搜啥啊？").
			Send()
	}
	req.IsSearch = true

	// 这里是临时使用ip锁
	req.Related = r.GetClientIp()

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.SearchLock + gconv.String(r.GetClientIp()) + gconv.String(r.GetClientIp())) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
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

		response.Success(r).SetData(data).Send()
	}
}

func (c *systemApi) getHotSearch(r *ghttp.Request) {
	//检查是否用户被锁
	if lock_utils.CheckLock(shared.SearchLock + gconv.String(r.GetClientIp()) + gconv.String(r.GetClientIp())) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁操作").Send()
	}

	if result, code := service.System.HotSearch(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *systemApi) getHotUser(r *ghttp.Request) {

	if result, code := service.System.HotUser(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *systemApi) getInfo(r *ghttp.Request) {
	if result, code := service.System.Info(); code != response.SUCCESS {

		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *systemApi) getVipAndGrade(r *ghttp.Request) {
	if result, code := service.System.VipAndGradeList(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()
	}
}
