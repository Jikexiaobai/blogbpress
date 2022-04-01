package result

import "github.com/gogf/gf/os/gtime"

type EduList struct {
	EduId      int64       `json:"id"`       //
	NickName   string      `json:"nickName"` // 发布用户
	Title      string      `json:"title"`    // 标题
	Cover      string      `json:"cover"`    // 封面
	Joins      int64       `json:"joins"`
	Status     int         `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
	Category   string      `json:"category"`   // 发布用户
}

type EduUserJoinList struct {
	UserId     int64       `json:"userId"`
	NickName   string      `json:"nickName"`
	Name       string      `json:"name"`
	Mode       int         `json:"mode"`
	Number     string      `json:"number"`
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
}

// 课程
type EduEditInfo struct {
	UserId      int64       `json:"userId"`      // 分类
	CateId      int64       `json:"cateId"`      // 分类
	Title       string      `json:"title"`       // 标题
	Cover       string      `json:"cover"`       // 封面
	Content     string      `json:"content"`     //
	Description string      `json:"description"` // 描述
	Section     string      `json:"section"`     // 视频地址
	Joins       int64       `json:"joins"`       //
	Hots        int64       `json:"hots"`        //
	Views       int64       `json:"views"`       // 播放量
	Favorites   int64       `json:"favorites"`   //
	Likes       int64       `json:"likes"`       //
	Type        int         `json:"type"`        // 课程类型 1线下，2线上
	Max         int         `json:"max"`
	JoinMode    int         `json:"joinMode"`   // 查看权限 0公开下载，1付费下载，2评论下载，3登录下载
	Price       float64     `json:"price"`      //
	Status      int         `json:"status"`     // 状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿
	CreateTime  *gtime.Time `json:"createTime"` //
	TagList     []*TagList  `json:"tagList"`    // 备注
}

type SectionInfo struct {
	Title    string       `json:"title"`    //
	Children []*ClassInfo `json:"children"` //
}
type ClassInfo struct {
	IsWatch bool   `json:"isWatch"` //
	Title   string `json:"title"`   //
	Link    string `json:"link"`    //
}
