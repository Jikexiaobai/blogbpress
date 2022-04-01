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

var Order = new(orderApi)

type orderApi struct{}

// LoadRouter 加载 authController 路由
func (c *orderApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/order", func(group *ghttp.RouterGroup) {
		group.GET("/list", c.getList)
		group.POST("/create", c.postCreate)
		group.POST("/pay", c.postPay)
		group.POST("/status", c.postCheckOrderStatus)
	})
}

func (c *orderApi) postPay(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置订单号",
	}
	orderNum := r.GetFormString("orderNum")
	if err := gvalid.Check(orderNum, rule, msg); err != nil {
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

	if service.Order.CheckPay(tokenUserId, orderNum) {
		response.Error(r).SetCode(response.INVALID).
			SetMessage("该订单已经支付过了").Send()
	}

	if info, code := service.Order.Pay(tokenUserId, orderNum); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).
			SetCode(response.SUCCESS).SetData(data).Send()
	}
}

func (c *orderApi) postCreate(r *ghttp.Request) {
	var req *dto.OrderCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	//if req.OrderType == shared.OrderCZ && req.PayMethod == shared.OrderBL {
	//	response.Error(r).
	//		SetCode(response.FAILD).SetMessage("充值订单不允许余额支付").Send()
	//}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.OrderCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁创建订单").Send()
	}
	req.UserId = tokenUserId
	if result, code := service.Order.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("orderNum", result)
		response.Success(r).SetData(data).Send()
	}

}

func (c *orderApi) postCheckOrderStatus(r *ghttp.Request) {

	rule := "required"
	msg := map[string]string{
		"required": "请设置订单号",
	}
	orderNum := r.GetFormString("orderNum")
	if err := gvalid.Check(orderNum, rule, msg); err != nil {
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

	data := gmap.New(true)
	data.Set("status", service.Order.CheckPay(tokenUserId, orderNum))
	response.Success(r).SetMessage("获取成功").SetData(data).Send()
}

func (c *orderApi) getList(r *ghttp.Request) {
	var req *dto.OrderQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Order.SelectList(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("total", total)
		data.Set("list", result)
		response.Success(r).SetData(data).Send()
	}
}
