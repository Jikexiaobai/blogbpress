package dto

type EduQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int    `p:"status"`
	Title  string `p:"title"`
	CateId int64  `p:"cateId"`
}

// 用户报名查询参数
type EduJoinQuery struct {
	EduId  int64  `p:"eduId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	UserId int64  `p:"userId"`
	Name   string `p:"name"`
	Number string `p:"number"`
	Mode   int    `p:"mode"`
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
}

// 互动
type EduCreate struct {
	Title       string   `p:"title"  v:"required#请输入标题"`  // 标题
	Cover       string   `p:"cover"  v:"required#请设置封面图"` // 封面
	Description string   `p:"description"`                // 描述
	Content     string   `p:"content"`                    // 描述
	CateId      int64    `p:"cateId"  v:"required|integer|min:1#请设置分类id|分类id必须为整型|分类id最小为1"`
	Tags        []string `p:"tags"`                                   // 标签
	Section     string   `p:"section"  v:"json|required#请设置章节|请设置章节"` // 封面
	JoinMode    int      `p:"joinMode"  v:"between:0,1#权限参数只允许0到1"`
	Type        int      `p:"type"  v:"between:1,2#权限参数只允许1和2"`
	Max         int      `p:"max"  v:"required#请设置最大报名数"`
	Price       float64  `p:"price"`
	Joins       int64    `p:"joins"`
	Views       int64    `p:"views"`
	Hots        int64    `p:"hots"`
	Favorites   int64    `p:"favorites"`
	Likes       int64    `p:"likes"`
	UserId      int64    `p:"userId"`
}

// 互动
type EduEdit struct {
	EduId       int64    `p:"eduId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	Title       string   `p:"title"  v:"required#请输入标题"`                              // 标题
	Cover       string   `p:"cover"  v:"required#请设置封面图"`                             // 封面
	Description string   `p:"description"`                                            // 描述
	Content     string   `p:"content"`                                                // 描述
	CateId      int64    `p:"cateId"  v:"required|integer|min:1#请设置分类id|分类id必须为整型|分类id最小为1"`
	Tags        []string `p:"tags"`                                       // 标签
	Section     string   `p:"section"  v:"json|required#请设置下载地址|请设置下载地址"` // 封面
	JoinMode    int      `p:"joinMode"  v:"between:1,2#权限参数只允许1到2"`
	Type        int      `p:"type"  v:"between:1,2#权限参数只允许1和2"`
	Max         int      `p:"max"  v:"required#请设置最大报名数"`
	Price       float64  `p:"price"`
	Joins       int64    `p:"joins"`
	Views       int64    `p:"views"`
	Hots        int64    `p:"hots"`
	Favorites   int64    `p:"favorites"`
	Likes       int64    `p:"likes"`
	UserId      int64    `p:"userId"`
}
