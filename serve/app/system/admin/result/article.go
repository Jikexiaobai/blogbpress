package result

import "github.com/gogf/gf/os/gtime"

// 文章
type ArticleList struct {
	ArticleId  int64       `json:"id"`
	NickName   string      `json:"nickName"`
	Title      string      `json:"title"`
	Cover      string      `json:"cover"`
	Status     int         `json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `json:"createTime"`
	Category   string      `json:"category"`
}

// 文章
type ArticleEditInfo struct {
	UserId      int64       `json:"userId"`      // 分类id
	CateId      int64       `json:"cateId"`      // 分类id
	Hots        int64       `json:"hots"`        //
	Favorites   int64       `json:"favorites"`   //
	Likes       int64       `json:"likes"`       //
	Views       int64       `json:"views"`       // 阅读量
	Title       string      `json:"title"`       // 标题
	Content     string      `json:"content"`     // 内容
	Cover       string      `json:"cover"`       // 封面
	Description string      `json:"description"` // 描述
	Status      int         `json:"status"`      // 状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"`  // 更新时间
	TagList     []*TagList  `json:"tagList"`     // 备注
	Remark      string      `json:"remark"`      // 备注
}
