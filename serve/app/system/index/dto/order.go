package dto

type OrderCreate struct {
	UserId          int64
	OrderMoney      float64 `p:"orderMoney" v:"required#请输入金额"`
	PayMethod       int     `p:"payMethod" v:"required#请选择支付方式"`
	OrderType       int     `p:"orderType" v:"required#请设置订单类型"`
	OrderMode       int
	DetailId        int64  `p:"detailId"`     //
	DetailModule    string `p:"detailModule"` //
	ShippingAddress string `p:"shippingAddress"`
	ShippingName    string `p:"shippingName"`
	ShippingPhone   string `p:"shippingPhone"`
}

type OrderQuery struct {
	Page      int   `p:"page" v:"required#请设置页数"`
	Limit     int   `p:"limit" v:"between:1,100#参数只允许1到100"`
	OrderType int   `p:"orderType"`
	Status    int   `p:"status"`
	UserId    int64 `p:"userId"`   // 收益
	AuthorId  int64 `p:"authorId"` // 收益
}
