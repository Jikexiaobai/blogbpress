package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/system/index/valid"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

var Account = new(accountApi)

type accountApi struct {
}

// LoadRouter 加载 authController 路由
func (c *accountApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/account", func(group *ghttp.RouterGroup) {
		group.GET("/group", c.getGroup)
		group.GET("/favorites", c.getAccountPosts)
		group.GET("/join/edu", c.getAccountPosts)
		group.GET("/buy/posts", c.getAccountPosts)
		group.GET("/info", c.getAccountInfo)
		group.POST("/edit", c.postAccountEdit)
		group.GET("/security", c.getAccountSecurity)
		//group.POST("/update/password", c.postAccountUpdatePassword)
		//group.POST("/update/email", c.postAccountUpdateEmail)
		group.GET("/verify/statusAndIsPayAndPrice", c.getAccountVerifyStatusIsPayPrice)
		group.POST("/verify", c.postAccountVerify)
		group.GET("/verify", c.getAccountVerify)
		group.GET("/balance", c.getAccountBalance)
		group.POST("/sign", c.postSign)
	})
}

func (c *accountApi) getGroup(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage("请设置id").
			Send()
	}
	data := gmap.New(true)
	// 获取加入的小组
	list, err := service.Group.MyJoinGroupList(tokenUserId)
	if err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).
			Send()
	}
	data.Set("list", list)
	response.Success(r).SetData(data).Send()
}

func (c *accountApi) postSign(r *ghttp.Request) {

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	if valid.Account.CheckSign(tokenUserId) {
		response.Error(r).SetCode(response.ADD_FAILED).
			SetMessage("今日已打卡签到").Send()
	}

	if res, code := service.Account.Sign(tokenUserId); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("integral", res)
		response.Success(r).
			SetCode(response.SUCCESS).
			SetData(data).
			SetMessage("签到打卡成功").Send()
	}
}

func (c *accountApi) getAccountPosts(r *ghttp.Request) {

	var req *dto.QueryParam
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	req.UserId = tokenUserId
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

func (c *accountApi) getAccountInfo(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	if info, code := service.Account.SelectInfo(tokenUserId); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).SetData(data).Send()
	}
}

func (c *accountApi) postAccountEdit(r *ghttp.Request) {
	var req *dto.AccountBase
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetBuniss(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	req.UserId = tokenUserId
	if code := service.Account.Edit(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *accountApi) getAccountSecurity(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	if result, code := service.Account.SelectSecurity(tokenUserId); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *accountApi) postAccountUpdatePassword(r *ghttp.Request) {
	var req *dto.PassWordDto
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//获取前台提交的数据

	if err := service.User.EditAccountPassWord(tokenUserId, req); err != nil {
		response.Error(r).SetCode(response.UPDATE_FAILED).SetMessage(response.CodeMsg(response.UPDATE_FAILED)).Send()
	} else {
		response.Success(r).SetMessage("提交成功").Send()
	}
}

func (c *accountApi) postAccountUpdateEmail(r *ghttp.Request) {
	var req *dto.EmailDto
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	//获取前台提交的数据

	if err := service.User.EditAccountEmail(tokenUserId, req); err != nil {
		response.Error(r).SetCode(response.UPDATE_FAILED).SetMessage(response.CodeMsg(response.UPDATE_FAILED)).Send()
	} else {
		response.Success(r).SetMessage("提交成功").Send()
	}
}

func (c *accountApi) getAccountBalance(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}
	if result, code := service.Account.GetBalance(tokenUserId); code != response.SUCCESS {
		response.Error(r).SetCode(response.DB_READ_ERROR).SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	} else {
		data := gmap.New(true)
		data.Set("balance", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *accountApi) getAccountVerifyStatusIsPayPrice(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if result, code := service.Verify.SelectStatus(tokenUserId); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *accountApi) postAccountVerify(r *ghttp.Request) {
	var req *dto.VerifyCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).
			Send()
	}

	if service.Verify.CheckHasInfo(req.Code, req.Name) {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("认证信息已存在").
			Send()
	}

	if !service.Verify.CheckHasPay(tokenUserId) {
		response.Error(r).SetCode(response.ADD_FAILED).
			SetMessage("请支付认证服务费").
			Send()
	}

	req.UserId = tokenUserId
	if status, code := service.Verify.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		data := gmap.New(true)
		data.Set("status", status)
		response.Success(r).SetData(data).Send()
	}
}

func (c *accountApi) getAccountVerify(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.INVALID).
			SetMessage(err.Error()).
			Send()
	}

	if result, code := service.Verify.SelectInfo(tokenUserId); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}
