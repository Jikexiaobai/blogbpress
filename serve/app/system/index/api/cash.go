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
)

var Cash = new(cashApi)

type cashApi struct{}

// LoadRouter 加载 authController 路由
func (c *cashApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/cash", func(group *ghttp.RouterGroup) {
		group.POST("/create", c.postCreate)
		group.GET("/list", c.getList)
	})
}

func (c *cashApi) getList(r *ghttp.Request) {

	//获取QueryParam
	var req *dto.CashQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Cash.SelectList(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()

	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).
			SetCode(response.SUCCESS).
			SetData(data).Send()
	}
}

func (c *cashApi) postCreate(r *ghttp.Request) {
	var req *dto.CashCreate
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
	if lock_utils.CheckLock(shared.CashCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.FAILD).SetMessage("请不要频繁发起提现操作").Send()
	}

	// 判断余额
	if !service.Cash.CheckBalanceHasCash(tokenUserId, req.Money) {
		response.Error(r).
			SetCode(response.FAILD).
			SetMessage("余额不足").Send()
	}

	// 判断提现是否超过最少提现额度
	if !service.Cash.CheckCashMin(req.Money) {
		response.Error(r).
			SetCode(response.FAILD).
			SetMessage("提现金额少于最少额度").Send()
	}

	// 判断是否认证
	if !service.Verify.CheckIsVerify(tokenUserId) {
		response.Error(r).
			SetCode(response.FAILD).
			SetMessage("请先实名认证").Send()
	}
	req.UserId = tokenUserId
	if code := service.Cash.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
