package dto

// 筛选列表参数
type QueryParam struct {
	Page       int    `p:"page" v:"required#请设置页数"`
	Limit      int    `p:"limit"`
	Module     string `p:"module"`     // 模块
	Mode       int    `p:"mode"`       // 筛选方式
	UserId     int64  `p:"userId"`     // 所属用户
	CateId     int64  `p:"cateId"`     // 分类id
	TagId      int64  `p:"TagId"`      // 标签id
	GroupId    int64  `p:"groupId"`    // 圈子id
	IsDown     int    `p:"isDown"`     // 是否涵盖下载
	IsTop      int    `p:"isTop"`      // 是否置顶
	IsFavorite bool   `p:"isFavorite"` // 我收藏的
	IsBuy      bool   `p:"isBuy"`      // 我购买的
	IsJoin     bool   `p:"isJoin"`     // 我加入的
	IsSearch   bool
	Type       int    `p:"type"`    // 模块对应的类型
	Title      string `p:"title"`   // 标题
	Related    string `p:"related"` // 关联方式
}

// 用户报名查询参数
type EduJoinQueryParam struct {
	EduId  int64  `p:"eduId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	UserId int64  `p:"userId"`
	Name   string `p:"name"`
	Number string `p:"number"`
	Mode   int    `p:"mode"`
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
}
