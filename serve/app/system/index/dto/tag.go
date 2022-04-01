package dto

type TagTop struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	IsTop  int     `p:"isTop"  v:"required|between:1,2#请设置是否推荐|请设置是否推荐"` // 描述
}
