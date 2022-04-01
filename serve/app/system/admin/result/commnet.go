package result

import "github.com/gogf/gf/os/gtime"

//  评论
type CommentList struct {
	CommentId int64  `json:"id"`        //
	RelatedId int64  `json:"relatedId"` //
	Module    string `json:"module"`    // 模块
	//Title      string      `json:"title"`      //
	NickName   string      `json:"nickName"`   // 内容
	Content    string      `json:"content"`    // 内容
	Type       int         `json:"type"`       // 类型 1图片 2视频 3文字
	Status     int         `json:"status"`     //
	CreateTime *gtime.Time `json:"createTime"` //
}
