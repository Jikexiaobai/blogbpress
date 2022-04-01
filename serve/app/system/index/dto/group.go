package dto

type GroupQuery struct {
	Page   int `p:"page" v:"required#请设置页数"`
	Limit  int `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int `p:"status"`
	UserId int64
}

type GroupPostsQuery struct {
	Page    int    `p:"page" v:"required#请设置页数"`
	Limit   int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Module  string `p:"module" v:"in:topic,question#参数只允许topic或question"` // 模块
	GroupId int64  `p:"groupId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"`
}

// 圈子
type GroupCreate struct {
	CateId      int64   `p:"cateId"  v:"required|integer|min:1#请设置分类id|分类id必须为整型|分类id最小为1"`
	Title       string  `p:"title"  v:"required#请设置名称"`          // 小组名称
	Icon        string  `p:"icon" v:"required#请设置图标"`            // 小组图标
	Cover       string  `p:"cover"  v:"required#请设置封面"`          // 封面
	JoinMode    int     `p:"joinMode"  v:"between:1,3#参数只允许1到3"` // 小组类型1 公共小组，2付费小组，3私密小组
	Price       float64 `p:"price"`                              // 费用
	SecretKey   string  `p:"secretKey"`                          // 专属密钥
	Description string  `p:"description"`                        // 小组描述
	UserId      int64   `p:"userId"`
}

// 圈子
type GroupEdit struct {
	GroupId     int64   `p:"groupId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` //
	CateId      int64   `p:"cateId"  v:"required|integer|min:1#请设置分类id|分类id必须为整型|分类id最小为1"`
	Title       string  `p:"title"  v:"required#请设置名称"`          // 小组名称
	Icon        string  `p:"icon" v:"required#请设置图标"`            // 小组图标
	Cover       string  `p:"cover"  v:"required#请设置封面"`          // 封面
	JoinMode    int     `p:"joinMode"  v:"between:1,3#参数只允许1到3"` // 小组类型1 公共小组，2付费小组，3私密小组
	Price       float64 `p:"price"`                              // 费用
	SecretKey   string  `p:"secretKey"`                          // 专属密钥
	Description string  `p:"description"`                        // 小组描述
	UserId      int64   `p:"userId"`
}
