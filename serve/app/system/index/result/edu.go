package result

import "github.com/gogf/gf/os/gtime"

type EduList struct {
	EduId       int64       `json:"id"`
	Title       string      `json:"title"`
	Cover       string      `json:"cover"`
	Description string      `json:"description"`
	Status      int         `json:"status"`
	CreateTime  *gtime.Time `json:"createTime"`
}

type EduInfo struct {
	EduId       int64         `json:"id"`       //
	UserInfo    *UserInfo     `json:"userInfo"` // 发布用户
	CateInfo    *CategoryInfo `json:"cateInfo"` // 发布用户
	Title       string        `json:"title"`    // 标题
	Section     string        `json:"section"`  // 标题
	Cover       string        `json:"cover"`    // 封面
	JoinMode    float64       `json:"joinMode"`
	Price       float64       `json:"price"`
	Description string        `json:"description"` // 描述
	Content     string        `json:"content"`     //
	Joins       int64         `json:"joins"`
	Hots        int64         `json:"hots"`
	Comments    int64         `json:"comments"`
	Views       int64         `json:"views"`
	Likes       int64         `json:"likes"`
	Favorites   int64         `json:"favorites"`
	IsJoin      bool          `json:"isJoin"`
	IsPay       bool          `json:"isPay"`
	IsLike      bool          `json:"isLike"`     // 是否点赞
	IsFavorite  bool          `json:"isFavorite"` // 是否收藏
	Type        int           `json:"type"`
	Max         int           `json:"max"`
	Status      int           `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	Remark      string        `json:"remark"`     // 备注
	CreateTime  *gtime.Time   `json:"createTime"` // 创建时间
	TagList     []*TagList    `json:"tagList"`    // 备注
}

type EduListInfo struct {
	Id          int64         `c:"id"          json:"id"`          //
	Module      string        `c:"module"      json:"module"`      //
	UserInfo    *UserInfo     `c:"userInfo"    json:"userInfo"`    // 发布用户
	Title       string        `c:"title"       json:"title"`       // 标题
	Cover       string        `c:"cover"       json:"cover"`       // 封面
	Description string        `c:"description" json:"description"` // 描述
	JoinMode    int           `c:"joinMode"        json:"joinMode"`
	Price       float64       `c:"price"        json:"price"`
	Hots        int64         `c:"hots"        json:"hots"`
	Joins       int64         `c:"joins"        json:"joins"`
	Views       int64         `c:"views" json:"views"`
	Likes       int64         `c:"likes" json:"likes"`
	Favorites   int64         `c:"favorites" json:"favorites"`
	Status      int           `json:"status"`                    // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime  *gtime.Time   `c:"createTime" json:"createTime"` // 创建时间
	CateInfo    *CategoryInfo `c:"cateInfo" json:"cateInfo"`     // 发布用户
}

type EduUserJoinListInfo struct {
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
