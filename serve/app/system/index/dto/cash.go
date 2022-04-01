package dto

type CashQuery struct {
	Page   int   `p:"page" v:"required#请设置页数"`
	Limit  int   `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int   `p:"status"` // 状态
	UserId int64 // 状态
}

type CashCreate struct {
	PayMethod int     `p:"payMethod" v:"required|between:1,2#请设置支付方式|类型参数只允许1和2"`
	Money     float64 `p:"money" v:"required#请设置金额"`
	Number    string  `p:"number" v:"required#请输入提现账户"`
	UserId    int64
}
