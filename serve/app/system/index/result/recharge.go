package result

import "github.com/gogf/gf/os/gtime"

// RechargerPayInfo 订单二维码信息
type RechargerPayInfo struct {
	IsPay  bool   `json:"isPay"`
	Mode   int    `json:"mode"`
	QrCode string `json:"qrCode"`
	Code   string `json:"code"`
}

type RechargerInfo struct {
	Code       string      `json:"code"`
	Money      float64     `json:"money"`
	Mode       int         `json:"mode"`
	Status     int         `json:"status"`
	Remark     string      `json:"remark"`
	CreateTime *gtime.Time `json:"createTime"`
}
