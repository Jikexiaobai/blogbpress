package dto

type CarouselQuery struct {
	Page     int    `p:"page" v:"required#请设置页数"`
	Limit    int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status   int    `p:"status"`
	Position string `p:"position"`
	Type     int    `p:"Type"`
	Mode     int    `p:"mode"`
}

//
type CarouselCreate struct {
	Cover     string `p:"cover"  v:"required#请设置封面图"` // 封面
	Link      string `p:"link"  v:"required#请输入地址"`   // 标题
	Type      int    `p:"type"  v:"required#请设置轮播平台"`
	Mode      int    `p:"mode"  v:"required#请设置内容所属"`
	Module    string `p:"module"`
	RelatedId int64  `p:"relatedId"`
}
