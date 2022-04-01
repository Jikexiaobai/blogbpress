package dto

type CashQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Code   string `p:"code"`
	Status int    `p:"status"`
}

type CashReview struct {
	ReceiptNum string `p:"receiptNum"`
	Status     int    `p:"status" v:"required|between:2,3#请设置需要更新的状态|参数仅为2或3"`
	Code       string `p:"code" v:"required#请设置提现单号"`
	Remark     string `p:"remark" v:"required#请填写通知信息"`
}
