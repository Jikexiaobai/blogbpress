package dto

type VerifyQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int    `p:"status"`
	Name   string `p:"name"`
	Code   string `p:"code"`
}

// 答案
type VerifyCreate struct {
	UserId int64  `p:"userId"`                       // 联系方式 1 qq, 2微信
	Name   string `p:"name"  v:"required#请输入名称"`     // 真实姓名
	Code   string `p:"code"  v:"required#请输入证件号"`    // 身份证号码
	Mode   int    `p:"mode"  v:"required#请设置联系方式"`   // 联系方式 1 qq, 2微信
	Number string `p:"number"  v:"required#请设置联系方式"` // 联系号码
}
type VerifyEdit struct {
	VerifyId int64  `p:"verifyId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"`
	Name     string `p:"name"  v:"required#请输入名称"`     // 真实姓名
	Code     string `p:"code"  v:"required#请输入证件号"`    // 身份证号码
	Mode     int    `p:"mode"  v:"required#请设置联系方式"`   // 联系方式 1 qq, 2微信
	Number   string `p:"number"  v:"required#请设置联系方式"` // 联系号码
	UserId   int64  `p:"userId"`                       // 联系方式 1 qq, 2微信
}
