package dto

// 答案
type VerifyCreate struct {
	UserId int64  `p:"userId"`                       // 联系方式 1 qq, 2微信
	Name   string `p:"name"  v:"required#请输入名称"`     // 真实姓名
	Code   string `p:"code"  v:"required#请输入证件号"`    // 身份证号码
	Mode   int    `p:"mode"  v:"required#请设置联系方式"`   // 联系方式 1 qq, 2微信
	Number string `p:"number"  v:"required#请设置联系方式"` // 联系号码
}
