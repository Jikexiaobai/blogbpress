package dto

// 评论的列表查询参数
type CommentQueryParam struct {
	RelatedId int64  `p:"relatedId"`
	Module    string `p:"module"`
	Content   string `p:"content"`
	Status    int    `p:"status"` // 状态
	Page      int    `p:"page" v:"required#请设置页数"`
	Limit     int    `p:"limit" v:"between:1,100#参数只允许1到100"`
}
