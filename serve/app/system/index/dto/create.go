package dto

// 配置
type ConfigCreate struct {
	ConfigName  string `p:"config_name"`                       // 参数名称
	ConfigKey   string `p:"config_key" v:"required#请输入权限字符"`   // 参数键名
	ConfigValue string `p:"config_value" v:"required#请输入权限字符"` // 参数键值
	Remark      string `p:"remark"`                            // 备注
}

// 账单
type BillCreate struct {
	OrderId    int64   // 关联订单编号
	UserId     int64   // 用户id
	Title      string  // 流水号标题
	BillNum    string  // 流水号（渠道返回）
	Money      float64 // 账单金额
	ServiceFee float64 // 服务费
	PayMethod  int     // 支付方式（1支付宝，2微信）
	BillType   int     // 账单类型 1支出，2收入
}
