package dto

type RechargeQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Code   string `p:"code"`
	Mode   int    `p:"mode"`
	Status int    `p:"status"`
}

type RechargeReview struct {
	Status int    `p:"status" v:"required|between:2,3#请设置需要更新的状态|参数仅为2或3"`
	Code   string `p:"code" v:"required#请设置充值单号"`
	Remark string `p:"remark" v:"required#请填写充值通知信息"`
}
