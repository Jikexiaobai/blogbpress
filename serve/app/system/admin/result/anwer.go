package result

import "github.com/gogf/gf/os/gtime"

// 回答
type AnswerList struct {
	AnswerId   int64       `json:"id"`       //
	Content    string      `json:"content"`  //
	NickName   string      `json:"nickName"` //
	Price      float64     `json:"price"`
	Status     int         `json:"status"`     //
	CreateTime *gtime.Time `json:"createTime"` //
}
