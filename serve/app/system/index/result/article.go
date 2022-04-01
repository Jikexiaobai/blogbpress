package result

import "github.com/gogf/gf/os/gtime"

// 文章
type ArticleInfo struct {
	ArticleId   int64         `json:"id"`          //
	UserInfo    *UserInfo     `json:"userInfo"`    // 发布用户
	CateInfo    *CategoryInfo `json:"cateInfo"`    //
	Title       string        `json:"title"`       // 标题
	Content     string        `json:"content"`     // 内容
	Cover       string        `json:"cover"`       // 封面
	Description string        `json:"description"` // 描述
	Likes       int64         `json:"likes"`
	Hots        int64         `json:"hots"`
	Comments    int64         `json:"comments"`
	Views       int64         `json:"views"`
	Favorites   int64         `json:"favorites"`
	IsLike      bool          `json:"isLike"`     // 是否点赞
	IsFavorite  bool          `json:"isFavorite"` // 是否收藏
	Status      int           `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime  *gtime.Time   `json:"createTime"` // 创建时间
	TagList     []*TagList    `json:"tagList"`    // 备注
}

// 文章
type ArticleList struct {
	ArticleId   int64       `json:"id"`
	Title       string      `json:"title"`
	Cover       string      `json:"cover"`
	Description string      `json:"description"`
	Status      int         `json:"status"`
	CreateTime  *gtime.Time `json:"createTime"`
}

// 文章
type ArticleFilterList struct {
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
	CateInfo    *CategoryInfo `json:"cateInfo"`                  //
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
