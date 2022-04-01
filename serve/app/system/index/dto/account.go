package dto

// AccountBase 基础设置的dto
type AccountBase struct {
	NickName    string `p:"nickName"  v:"required#请输入用户昵称"`
	Sex         int    `p:"sex" v:"required#请设置用户性别"`
	Description string `p:"description"`
	Cover       string `p:"cover"`
	Avatar      string `p:"avatar"`
	UserId      int64
}
