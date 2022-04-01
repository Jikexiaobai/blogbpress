package result

import "github.com/gogf/gf/os/gtime"

// 订单信息
type RechargeList struct {
	RechargeId int64       `json:"id"`
	Code       string      `json:"code"`
	NickName   string      `json:"nickName"`
	Mode       int         `json:"mode"`
	Money      float64     `json:"money"`
	CardKey    string      `json:"cardKey"`
	Name       string      `json:"name"`
	Type       int         `json:"type"`
	Number     string      `json:"number"`
	Remark     string      `json:"remark"`
	Status     int         `json:"status"`     //
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
}
