package dto

type ArticleQuery struct {
	Page   int `p:"page" v:"required#请设置页数"`
	Limit  int `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int `p:"status"`
	UserId int64
}

// 文章
type ArticleEdit struct {
	ArticleId int64  `p:"articleId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	Cover     string `p:"cover"  v:"required#请设置封面图"`                                 // 封面
	Title     string `p:"title"  v:"required#请输入标题"`                                  // 标题
	Content   string `p:"content"  v:"required#请输入内容"`
	CateId    int64  `p:"cateId"  v:"required|integer|min:1#请设置分类id|分类id必须为整型|分类id最小为1"`
	//GroupId     []int64  `p:"groupId"  v:"required#请设置圈子"` // 圈子id
	Tags        []string `p:"tags"`
	Description string   `p:"description"`
	Views       int64    `p:"views"`
	Hots        int64    `p:"hots"`
	Favorites   int64    `p:"favorites"`
	Likes       int64    `p:"likes"`
	UserId      int64    `p:"userId"`
}

// 文章
type ArticleCreate struct {
	Cover   string `p:"cover"  v:"required#请设置封面图"` // 封面
	Title   string `p:"title"  v:"required#请输入标题"`  // 标题
	Content string `p:"content"  v:"required#请输入内容"`
	CateId  int64  `p:"cateId"  v:"required|integer|min:1#请设置分类id|分类id必须为整型|分类id最小为1"`
	//GroupId     []int64  `p:"groupId"  v:"required#请设置圈子"` // 圈子id
	Tags        []string `p:"tags"`
	Description string   `p:"description"`
	Views       int64    `p:"views"`
	Hots        int64    `p:"hots"`
	Favorites   int64    `p:"favorites"`
	Likes       int64    `p:"likes"`
	UserId      int64    `p:"userId"`
}
