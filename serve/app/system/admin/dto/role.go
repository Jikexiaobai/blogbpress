package dto

type RoleCreate struct {
	Title     string  `p:"title"  v:"required#请输入标题"`
	Status    int     `p:"status"  v:"required#请设置状态"`
	Authority []int64 `p:"authority"`
}

type RoleEdit struct {
	RoleId    int64   `p:"id"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	Title     string  `p:"title"  v:"required#请输入标题"`
	Status    int     `p:"status"  v:"required#请设置状态"`
	Authority []int64 `p:"authority"`
}
