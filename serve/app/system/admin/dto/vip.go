package dto

type VipCreate struct {
	Title    string `p:"title"  v:"required#请输入标题"`
	Day      int    `p:"day"`
	Price    int    `p:"price" v:"required|not-in:0#请设置会员开通价格|请设置会员开通价格"`
	Discount int    `p:"discount"`
	Color    string `p:"color"`
	Icon     string `p:"icon"  v:"required#请上传图标"`
}

type VipEdit struct {
	VipId    int64  `p:"vipId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	Title    string `p:"title"  v:"required#请输入标题"`
	Day      int    `p:"day"`
	Price    int    `p:"price" v:"required|not-in:0#请设置会员开通价格|请设置会员开通价格"`
	Discount int    `p:"discount"`
	Color    string `p:"color"`
	Icon     string `p:"icon"  v:"required#请上传图标"`
}
