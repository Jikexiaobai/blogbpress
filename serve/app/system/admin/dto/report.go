package dto

type ReportQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status int    `p:"status"`
	Type   int    `p:"type"`
	Title  string `p:"title"`
	CateId int64  `p:"cateId"`
}

// 举报
type ReportCreate struct {
	Description string `p:"description"`
	RelatedId   int64  `p:"relatedId" v:"required|not-in:0#请设置举报所属|请设置举报所属"`
	Module      string `p:"module"  v:"required#请设置举报模块"`
	Type        int    `p:"type"  v:"in:1,2,3,4#请设置举报类型"`
}
