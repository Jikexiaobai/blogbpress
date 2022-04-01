package dto

type QuestionQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int    `p:"status"`
	Title  string `p:"title"`
}

// 问题
type QuestionCreate struct {
	Title     string `p:"title"  v:"required|min-length:6#请输入标题|标题最小长度为:min位"`
	Content   string `p:"content"`
	Views     int64  `p:"views"`
	Hots      int64  `p:"hots"`
	Favorites int64  `p:"favorites"`
	Likes     int64  `p:"likes"`
	UserId    int64  `p:"userId"`
	Anonymous int    `p:"anonymous"`
}

// 问题
type QuestionEdit struct {
	QuestionId int64  `p:"questionId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	Title      string `p:"title"  v:"required|min-length:6#请输入标题|标题最小长度为:min位"`
	Content    string `p:"content"`
	Views      int64  `p:"views"`
	Hots       int64  `p:"hots"`
	Favorites  int64  `p:"favorites"`
	Likes      int64  `p:"likes"`
	UserId     int64  `p:"userId"`
	Anonymous  int    `p:"anonymous"`
}
