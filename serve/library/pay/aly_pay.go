package pay_lib

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// 支付包官方支付
type AliPayBody struct {
	AppId            string
	NotifyUrl        string
	PrivateKey       string
	AlyRootCert      string
	AlyCertPublicKey string
	AppCertPublicKey string
	Subject          string  // 标题
	OutTradeNo       string  // 订单或者流水号
	TotalAmount      float64 // 金额
}

// WebPay  二维码支付
func (e *AliPayBody) WebPay() (string, error) {

	client := alipay.NewClient(e.AppId, e.PrivateKey, false)
	e.AppCertPublicKey, _ = alipay.GetCertSN(e.AppCertPublicKey)
	e.AlyCertPublicKey, _ = alipay.GetCertSN(e.AlyCertPublicKey)
	e.AlyRootCert, _ = alipay.GetRootCertSN(e.AlyRootCert)

	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetAppCertSN(e.AppCertPublicKey).
		SetAliPayPublicCertSN(e.AlyCertPublicKey).
		SetAliPayRootCertSN(e.AlyRootCert).
		SetPrivateKeyType(alipay.PKCS8).
		SetNotifyUrl(e.NotifyUrl)

	//请求参数
	body := make(gopay.BodyMap)
	body.Set("subject", e.Subject)
	body.Set("out_trade_no", e.OutTradeNo)
	body.Set("total_amount", e.TotalAmount)
	g.Dump(client)
	aliRsp, err := client.TradePrecreate(body)
	if err != nil {
		glog.Info(err.Error())
		return "", err
	}

	return aliRsp.Response.QrCode, nil
}
