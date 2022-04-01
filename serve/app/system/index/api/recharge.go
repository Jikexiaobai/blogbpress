package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/system/index/shared"
	"fiber/app/system/index/valid"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"github.com/go-pay/gopay/alipay"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Recharge = new(rechargeApi)

type rechargeApi struct{}

// LoadRouter 加载 authController 路由
func (c *rechargeApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/recharge", func(group *ghttp.RouterGroup) {
		group.POST("/alipay/notice", c.postNotice)
		group.GET("/list", c.getList)
		group.POST("/create", c.postCreate)
		group.POST("/pay", c.postPay)
		group.POST("/status", c.postCheckOrderStatus)
	})
}

func (c *rechargeApi) postNotice(r *ghttp.Request) {

	notifyReq, err := alipay.ParseNotifyToBodyMap(r.Request)

	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	// 获取文件保存的存储引擎
	alyPayOptions, err := service.Config.FindValue("AlyPayOptions")
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}
	alyPayOptionsJson := gjson.New(alyPayOptions)
	alyCertPublicKey := service.Media.SelectMediaPath(gconv.String(alyPayOptionsJson.Get("alyCertPublicKey")))
	ok, err := alipay.VerifySignWithCert("."+alyCertPublicKey, notifyReq)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	if ok && notifyReq.Get("trade_status") == "TRADE_SUCCESS" {
		orderNum := gconv.String(notifyReq.Get("out_trade_no"))
		code := service.Order.UpdateStatus(orderNum)
		if code != response.SUCCESS {
			response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
		}
		response.Success(r).Send()
	}
	response.Error(r).SetCode(response.INVALID).SetMessage("验签出现了问题").Send()

}

func (c *rechargeApi) postCreate(r *ghttp.Request) {
	var req *dto.RechargeCreate
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.RechargeCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁充值").Send()
	}
	req.UserId = tokenUserId

	//判断充值方式如果为支付宝或微信必须设置金额
	if (req.Mode == shared.RechargeModeOne || req.Mode == shared.RechargeModeTwo) && req.Money == 0 {
		response.Error(r).
			SetCode(response.FAILD).SetMessage("请输入充值金额").Send()
	}

	//判断充值方式如果为卡密必须输入卡密信息
	if req.Mode == shared.RechargeModeThree {

		if req.CardKey == "" {
			response.Error(r).
				SetCode(response.FAILD).SetMessage("请输入卡密").Send()
		}

		// 判断卡密是否有效
		if valid.Card.CheckIsUse(req.CardKey) {
			response.Error(r).
				SetCode(response.FAILD).SetMessage("卡密已被使用").Send()
		}

	}

	//判断充值方式如果为人工转账必须输入转账信息
	if req.Mode == shared.RechargeModeFour && (req.Money == 0 || req.Name == "" || req.Number == "" || req.Type == 0) {
		response.Error(r).
			SetCode(response.FAILD).SetMessage("请设置转账信息").Send()
	}

	if result, code := service.Recharge.Create(req); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("code", result)
		response.Success(r).SetData(data).Send()
	}

}

func (c *rechargeApi) postPay(r *ghttp.Request) {
	rule := "required"
	msg := map[string]string{
		"required": "请设置订单号",
	}
	codeNum := r.GetFormString("code")
	if err := gvalid.Check(codeNum, rule, msg); err != nil {
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

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.RechargeCreateLock + gconv.String(tokenUserId)) {
		response.Error(r).SetCode(response.INVALID).SetMessage("请不要频繁充值").Send()
	}

	if valid.Recharge.CheckPay(tokenUserId, codeNum) {
		response.Error(r).SetCode(response.INVALID).
			SetMessage("该订单已经支付过了").Send()
	}

	if valid.Recharge.CheckRechargeModeOneOrTwo(tokenUserId, codeNum) {
		response.Error(r).SetCode(response.INVALID).
			SetMessage("支付方式不正确").Send()
	}

	if info, code := service.Recharge.Pay(tokenUserId, codeNum); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", info)
		response.Success(r).
			SetCode(response.SUCCESS).SetData(data).Send()
	}
}

func (c *rechargeApi) postCheckOrderStatus(r *ghttp.Request) {

	rule := "required"
	msg := map[string]string{
		"required": "请设置订单号",
	}
	codeNum := r.GetFormString("code")
	if err := gvalid.Check(codeNum, rule, msg); err != nil {
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
	data.Set("status", valid.Recharge.CheckPay(tokenUserId, codeNum))
	response.Success(r).SetMessage("获取成功").SetData(data).Send()
}

func (c *rechargeApi) getList(r *ghttp.Request) {
	var req *dto.RechargeQuery
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	tokenUserId, _ := service.Auth.GetTokenId(r)
	req.UserId = tokenUserId
	if total, result, code := service.Recharge.SelectList(req); code != response.SUCCESS {
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
