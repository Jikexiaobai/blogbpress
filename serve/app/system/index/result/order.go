package result

import "github.com/gogf/gf/os/gtime"

// OrderInfo 订单信息
type OrderInfo struct {
	OrderNum     string      `json:"orderNum"`  // 订单编号
	Title        string      `json:"title"`     // 订单编号
	Money        float64     `json:"money"`     // 支付金额
	OrderType    int         `json:"orderType"` // 订单类型: 0 充值，
	Status       int         `json:"status"`    // 订单状态 0未支付，1已支付
	IsIncome     bool        `json:"isIncome"`
	DetailId     int64       `json:"detailId"`
	DetailModule int64       `json:"detailModule"`
	CreateTime   *gtime.Time `json:"createTime"` // 创建时间
}

// OrderPayInfo 订单二维码信息
type OrderPayInfo struct {
	IsPay     bool   `json:"isPay"`
	PayMethod int    `json:"payMethod"`
	QrCode    string `json:"qrCode"`
	OrderNum  string `json:"orderNum"`
}
