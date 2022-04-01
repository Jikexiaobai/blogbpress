package dto

type QuestionQuery struct {
	Page   int `p:"page" v:"required#请设置页数"`
	Limit  int `p:"limit" v:"between:1,100#参数只允许1到100"`
	Mode   int `p:"mode"` // 筛选方式
	UserId int64
}

// 问题
type QuestionCreate struct {
	Title     string `p:"title"  v:"required|min-length:6#请输入标题|标题最小长度为:min位"`
	GroupId   int64  `p:"groupId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"`
	Content   string `p:"content"`
	UserId    int64  `p:"userId"`
	Anonymous int    `p:"anonymous"`
}
