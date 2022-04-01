package result

import "github.com/gogf/gf/os/gtime"

type CardList struct {
	CardId     int64       `json:"id"`         //
	NickName   string      `json:"nickName"`   // 使用者id
	SecretKey  string      `json:"secretKey"`  //
	Money      float64     `json:"money"`      //
	Status     int         `json:"status"`     // 状态: 1未使用，2已使用
	CreateTime *gtime.Time `json:"createTime"` //
}
