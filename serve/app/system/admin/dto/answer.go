package dto

// 回答的列表查询参数
type AnswerQueryParam struct {
	RelatedId int64 `p:"relatedId"`
	UserId    int64 `p:"userId"`
	Status    int   `p:"status"` // 状态
	IsPay     int   `p:"isPay"`  // 状态
	Page      int   `p:"page" v:"required#请设置页数"`
	Limit     int   `p:"limit" v:"between:1,100#参数只允许1到100"`
}

// 答案
type AnswerCreate struct {
	QuestionId int64   `p:"questionId"`
	Content    string  `p:"content" v:"required|min-length:6#请输入标题|标题最小长度为:min位"`
	Price      float64 `p:"price"`
	Files      string  `p:"files"` //
	Likes      int64   `p:"likes"`
	Hots       int64   `p:"hots"`
	Views      int64   `p:"views"`
	Favorites  int64   `p:"favorites"`
	UserId     int64   `p:"userId"`
}
