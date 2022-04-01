package dto

type TopicQuery struct {
	Page   int `p:"page" v:"required#请设置页数"`
	Limit  int `p:"limit" v:"between:1,100#参数只允许1到100"`
	Mode   int `p:"mode"` // 筛选方式
	Type   int `p:"type"` // 筛选方式
	UserId int64
}

type TopicCreate struct {
	Title     string `p:"title"  v:"required|min-length:6|max-length:256#请输入话题|话题最小长度为:min位|话题最大长度为:max位"`
	Type      int    `p:"type" v:"between:1,4#权限参数只允许1到4"`
	GroupId   int64  `p:"groupId"  v:"required|integer|min:1#请设置圈子|id必须为整型|id最小为1"`
	Module    string `p:"module"`
	RelatedId int64  `p:"relatedId"`
	Files     string `p:"files"`
	UserId    int64
}
