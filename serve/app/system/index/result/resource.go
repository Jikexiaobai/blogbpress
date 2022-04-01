package result

import "github.com/gogf/gf/os/gtime"

type ResourceList struct {
	ResourceId  int64       `json:"id"`
	Title       string      `json:"title"`
	Cover       string      `json:"cover"`
	Description string      `json:"description"`
	Status      int         `json:"status"`
	CreateTime  *gtime.Time `json:"createTime"`
}

// 资源
type ResourceInfo struct {
	ResourceId  int64         `json:"id"`          //
	UserInfo    *UserInfo     `json:"userInfo"`    // 发布用户
	CateInfo    *CategoryInfo `json:"cateInfo"`    //
	Title       string        `json:"title"`       // 标题
	Content     string        `json:"content"`     // 内容
	Cover       string        `json:"cover"`       // 封面
	Description string        `json:"description"` // 描述
	HasDown     int           `json:"hasDown"`
	DownMode    int           `json:"downMode"`
	Price       int           `json:"price"`
	IsDown      bool          `json:"isDown"`
	DownUrl     string        `json:"downUrl"`   // 下载地址
	Attribute   string        `json:"attribute"` // 属性
	Purpose     string        `json:"purpose"`   // 属性
	Likes       int64         `json:"likes"`
	Comments    int64         `json:"comments"`
	Hots        int64         `json:"hots"`
	Views       int64         `json:"views"`
	Favorites   int64         `json:"favorites"`
	IsLike      bool          `json:"isLike"`     // 是否点赞
	IsFavorite  bool          `json:"isFavorite"` // 是否收藏
	Status      int           `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	Remark      string        `json:"remark"`     // 备注
	CreateTime  *gtime.Time   `json:"createTime"` // 创建时间
	TagList     []*TagList    `json:"tagList"`    // 备注
}

// 资源
type ResourceEditInfo struct {
	CateId      int64       `json:"cateId"`      // 分类id
	UserId      int64       `json:"userId"`      // 分类id
	Hots        int64       `json:"hots"`        //
	Favorites   int64       `json:"favorites"`   //
	Likes       int64       `json:"likes"`       //
	Views       int64       `json:"views"`       // 阅读量
	Title       string      `json:"title"`       // 标题
	Content     string      `json:"content"`     // 内容
	Cover       string      `json:"cover"`       // 封面
	HasDown     int         `json:"hasDown"`     // 是否有下载1没有，2有
	DownMode    int         `json:"downMode"`    // 下载权限 0公开下载，1付费下载，2评论下载，3登录下载
	Price       float64     `json:"price"`       //
	DownUrl     string      `json:"downUrl"`     //
	Purpose     string      `json:"purpose"`     // 用途
	Attribute   string      `json:"attribute"`   // 属性
	KeyWords    string      `json:"keyWords"`    // 关键字
	SeoTitle    string      `json:"seoTitle"`    // seo 标题
	Description string      `json:"description"` // 描述
	Status      int         `json:"status"`      // 状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"`  // 更新时间
	TagList     []*TagList  `json:"tagList"`     // 备注
	//GroupList   []*GroupJoinList `json:"groupList"`   // 圈子
	Remark string `json:"remark"` // 备注
}

// 资源
type ResourceListInfo struct {
	Id          int64         `c:"id" json:"id"`                   //
	Module      string        `c:"module" json:"module"`           //
	UserInfo    *UserInfo     `c:"userInfo" json:"userInfo"`       // 发布用户
	Title       string        `c:"title" json:"title"`             // 标题
	Cover       string        `c:"cover" json:"cover"`             // 封面
	Description string        `c:"description" json:"description"` // 描述
	Likes       int64         `c:"likes" json:"likes"`
	Hots        int64         `c:"hots" json:"hots"`
	Views       int64         `c:"views" json:"views"`
	Favorites   int64         `c:"favorites" json:"favorites"`
	Status      int           `json:"status"`                    // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime  *gtime.Time   `c:"createTime" json:"createTime"` // 创建时间
	CateInfo    *CategoryInfo `json:"cateInfo" json:"cateInfo"`  //
}
