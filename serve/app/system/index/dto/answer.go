package dto

// 回答的列表查询参数
type AnswerQuery struct {
	TopicId int64 `p:"topicId" v:"required#请设置查询对应id参数"`
	Page    int   `p:"page" v:"required#请设置页数"`
	Limit   int   `p:"limit" v:"between:1,100#参数只允许1到100"`
}

// 答案
type AnswerCreate struct {
	TopicId int64   `p:"TopicId"`
	Content string  `p:"content" v:"required|min-length:6#请输入标题|标题最小长度为:min位"`
	Price   float64 `p:"price"`
	Files   string  `p:"files"` //
	UserId  int64   `p:"userId"`
}
