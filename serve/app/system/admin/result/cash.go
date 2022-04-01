package result

import "github.com/gogf/gf/os/gtime"

type CashList struct {
	CashId       int64       `json:"id"`           // 提现单号
	Code         string      `json:"code"`         // 提现单号
	Mode         int         `json:"mode"`         // 提现单号
	NickName     string      `json:"nickName"`     //
	ReceiptNum   string      `json:"receiptNum"`   //
	CashMoney    float64     `json:"cashMoney"`    //
	ServiceMoney float64     `json:"serviceMoney"` // 服务费
	Money        float64     `json:"money"`        // 实际金额
	PayMethod    int         `json:"payMethod"`    // 支付方式（1支付宝，2微信）
	Status       int         `json:"status"`       // 状态 1待审核，2审核通过，3审核不通过
	CreateTime   *gtime.Time `json:"createTime"`   //
}
