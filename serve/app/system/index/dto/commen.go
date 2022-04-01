package dto

type Review struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	Status int     `p:"status"  v:"required#请设置需要更新的状态"` // 描述
	Remark string  `p:"remark"`
	Node   string
	UserId int64
}

type Remove struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	Remark string  `p:"remark"`
	Node   string
	UserId int64
}
