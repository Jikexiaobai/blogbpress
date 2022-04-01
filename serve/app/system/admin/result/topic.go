package result

import "github.com/gogf/gf/os/gtime"

// 话题
type TopicList struct {
	TopicId    int64       `json:"id"` //
	NickName   string      `json:"nickName"`
	Title      string      `json:"title"`
	IsTop      int         `json:"isTop"`
	Status     int         `json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `c:"createTime" json:"createTime"`
}
