package dto

type CardQuery struct {
	Page      int    `p:"page" v:"required#请设置页数"`
	Limit     int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	SecretKey string `p:"secretKey"`
	Status    int    `p:"status"`
}

type CardCreate struct {
	Count int     `p:"count" v:"required|between:1,100#请设置创建张数|卡密张数范围1~100"`
	Money float64 `p:"money" v:"required#请设置对应金额"`
}
