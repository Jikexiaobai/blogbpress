package dto

type OrderQuery struct {
	Page      int    `p:"page" v:"required#请设置页数"`
	Limit     int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	OrderNum  string `p:"orderNum"`
	OrderType int    `p:"orderType"`
	Status    int    `p:"status"`
}
