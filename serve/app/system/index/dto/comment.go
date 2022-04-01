package dto

// 评论的列表查询参数
type CommentQuery struct {
	RelatedId int64  `p:"relatedId"`
	Module    string `p:"module"`
	Content   string `p:"content"`
	Status    int    `p:"status"` // 状态
	Page      int    `p:"page" v:"required#请设置页数"`
	Limit     int    `p:"limit" v:"between:1,100#参数只允许1到100"`
}

// 评论
type CommentCreate struct {
	Content   string `p:"content"  v:"required|min-length:6#请输入评论内容|最小长度为:min位"`
	RelatedId int64  `p:"relatedId" v:"required|not-in:0#请设置评论所属|请设置评论所属"`
	Module    string `p:"module"  v:"required#请设置评论模块"`
	//Type      int    `p:"type"  v:"in:1,2,3#请设置评论类型"`
	ReplyId  int64  `p:"replyId"`
	TopId    int64  `p:"topId"`
	ParentId int64  `p:"parentId"`
	Files    string `p:"files"` // 属性
}
