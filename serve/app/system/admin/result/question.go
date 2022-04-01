package result

import "github.com/gogf/gf/os/gtime"

// 问题
type QuestionList struct {
	QuestionId int64       `json:"id"`                        //
	NickName   string      `json:"nickName"`                  // 发布用户
	Title      string      `c:"title" json:"title"`           //
	Status     int         `c:"status"  json:"status"`        // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `c:"createTime" json:"createTime"` //
}

type QuestionEditInfo struct {
	UserId    int64  `json:"userId"`    //
	Title     string `json:"title"`     //
	Content   string `json:"content"`   //
	Hots      int64  `json:"hots"`      //
	Favorites int64  `json:"favorites"` //
	Likes     int64  `json:"likes"`     //
	Views     int64  `json:"views"`     //
	Anonymous int    `json:"anonymous"` // 是否匿名
}
