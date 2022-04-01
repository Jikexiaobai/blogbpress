package dto

type NoticeQuery struct {
	Page   int   `p:"page" v:"required#请设置页数"`
	Limit  int   `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int   `p:"status"`
	Type   int   `p:"type"`
	CateId int64 `p:"cateId"`
}

type NoticeCreate struct {
	Type         int
	DetailId     int64
	DetailModule string
	Content      string
	UserId       int64
}
