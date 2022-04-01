package dto

// 文章
type CarouselCreate struct {
	Cover     string `p:"cover"  v:"required#请设置封面图"` // 封面
	Link      string `p:"link"  v:"required#请输入地址"`   // 标题
	Type      int    `p:"type"  v:"required#请设置轮播平台"`
	Position  string `p:"position"  v:"required#请设置轮播位置"`
	Mode      int    `p:"mode"  v:"required#请设置内容所属"`
	Module    string `p:"module"`
	RelatedId int64  `p:"relatedId"`
}
