package result

import "github.com/gogf/gf/os/gtime"

type NoticeCount struct {
	System  int `json:"system"`
	Comment int `json:"comment"`
	Answer  int `json:"answer"`
	Finance int `json:"finance"`
	Like    int `json:"like"`
	Follow  int `json:"follow"`
}

type NoticeInfo struct {
	Status     int         `json:"status"`
	Type       int         `json:"type"`
	SystemType int         `json:"systemType"`
	UserInfo   *UserInfo   `json:"userInfo"` // 发布用户
	Content    string      `json:"content"`
	DetailInfo interface{} `json:"detailInfo"`
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
}

type NoticeDetailInfo struct {
	Cover  string `json:"cover"`
	Title  string `json:"title"`
	Module string `json:"module"`
	Id     int64  `json:"id"`
}
