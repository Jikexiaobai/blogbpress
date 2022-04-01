package dto

type MangerQuery struct {
	Page     int    `p:"page" v:"required#请设置页数"`
	Limit    int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status   int    `p:"status"`
	NickName string `p:"nickName"`
	Email    string `p:"email"`
	Phone    string `p:"phone"`
}

type MangerCreate struct {
	NickName string  `p:"nickName"  v:"required#请设置用户昵称"`
	Password string  `p:"password"  v:"required#请设置用户密码"`
	Phone    string  `p:"phone"`
	Email    string  `p:"email"`
	Avatar   string  `p:"avatar"`
	Sex      int     `p:"sex"`
	RoleId   []int64 `p:"roleId"  v:"required#请设置管理角色"`
}

type MangerEdit struct {
	UserId   int64   `p:"id"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	NickName string  `p:"nickName"  v:"required#请设置用户昵称"`
	Password string  `p:"password"  v:"required#请设置用户密码"`
	Phone    string  `p:"phone"`
	Email    string  `p:"email"`
	Avatar   string  `p:"avatar"`
	Sex      int     `p:"sex"`
	RoleId   []int64 `p:"roleId"  v:"required#请设置管理角色"`
}
