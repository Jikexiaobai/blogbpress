package dto

// 举报
type ReportCreate struct {
	Description string `p:"description"`
	RelatedId   int64  `p:"relatedId" v:"required|not-in:0#请设置举报所属|请设置举报所属"`
	Module      string `p:"module"  v:"required#请设置举报模块"`
	Type        int    `p:"type"  v:"in:1,2,3,4#请设置举报类型"`
}
