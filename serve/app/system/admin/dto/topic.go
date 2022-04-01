package dto

type TopicQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int    `p:"status"`
	Title  string `p:"title"`
}

// 话题
type TopicCreate struct {
	Title    string  `p:"title"  v:"required|min-length:6#请输入话题|话题最小长度为:min位"`
	Type     int     `p:"type" v:"between:1,3#权限参数只允许1到3"`
	ViewMode int     `p:"viewMode"  v:"between:0,3#权限参数只允许0到3"`
	GroupId  []int64 `p:"groupId"`
	Price    float64 `p:"price"`  // 付费价格
	Files    string  `p:"files"`  // 属性
	UserId   int64   `p:"userId"` // 属性
}

type TopicTop struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	IsTop  int     `p:"isTop"  v:"required|between:1,2#请设置需要置顶的值|值为1或2"` // 描述
}
