package dto

type RechargeQuery struct {
	Page   int   `p:"page" v:"required#请设置页数"`
	Limit  int   `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int   `p:"status"`
	UserId int64 `p:"userId"` // 收益
}

type RechargeCreate struct {
	UserId  int64
	Mode    int     `p:"mode" v:"required|between:1,4#请设置充值模式|参数只允许1到4"`
	Money   float64 `p:"money"`
	CardKey string  `p:"cardKey"`
	Name    string  `p:"name"`   //
	Type    int     `p:"type"`   //
	Number  string  `p:"number"` //
}
