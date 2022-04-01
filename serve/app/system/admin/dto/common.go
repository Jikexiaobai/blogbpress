package dto

type Review struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	Status int     `p:"status"  v:"required#请设置需要更新的状态"` // 描述
	Remark string  `p:"remark"`
}

type Remove struct {
	IdList []int64 `p:"idList"  v:"required#请设置Id列表"`
	Remark string  `p:"remark" v:"required#请输入处理通知信息"`
}
