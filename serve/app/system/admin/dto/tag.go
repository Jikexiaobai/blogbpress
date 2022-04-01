package dto

type TagQuery struct {
	Page  int    `p:"page" v:"required#请设置页数"`
	Limit int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Title string `p:"title"`
	IsTop int    `p:"isTop"`
}
type TagTop struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	IsTop  int     `p:"isTop"  v:"required|between:1,2#请设置是否推荐|请设置是否推荐"` // 描述
}
