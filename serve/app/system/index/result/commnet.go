package result

import "github.com/gogf/gf/os/gtime"

type CommentInfo struct {
	CommentId  int64          `json:"id"`           //
	ParentId   int64          `json:"parentId"`     //
	Children   []*CommentInfo `json:"children"`     //
	UserInfo   *UserInfo      `json:"userInfo"`     // 发布用户
	RelatedId  int64          `json:"relatedId"`    //
	Title      string         `json:"relatedTitle"` //
	Module     string         `json:"module"`       //
	Content    string         `json:"content"`      // 内容
	Type       int            `json:"type"`         // 类型 1图片 2视频 3文字
	Status     int            `json:"status"`       //
	Files      string         `json:"files"`        // 文件链接
	IsLike     bool           `json:"isLike"`       // 是否点赞
	Likes      int64          `json:"likes"`        // 点赞数
	CreateTime *gtime.Time    `json:"createTime"`   //
}
