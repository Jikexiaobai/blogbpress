package api

import (
	"fiber/app/system/index/service"
	"fiber/app/tools/response"
	"github.com/go-pay/gopay/alipay"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Pay = new(payApi)

type payApi struct{}

func (c *payApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/pay", func(group *ghttp.RouterGroup) {

		group.POST("/aly/notify", c.postAlyNotify)
	})
}

func (c *payApi) postAlyNotify(r *ghttp.Request) {

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
