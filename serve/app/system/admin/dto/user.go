package dto

type UserQuery struct {
	Page     int    `p:"page" v:"required#请设置页数"`
	Limit    int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Status   int    `p:"status"`
	NickName string `p:"nickName"`
	Email    string `p:"email"`
	Phone    string `p:"phone"`
}

type UserCreate struct {
	NickName    string  `p:"nickName"  v:"required#请设置用户昵称"`
	Password    string  `p:"password"  v:"required#请设置用户密码"`
	Phone       string  `p:"phone"`
	Email       string  `p:"email"`
	Avatar      string  `p:"avatar"`
	Cover       string  `p:"cover"`
	Balance     float64 `p:"balance"`
	Integral    int64   `p:"integral"`
	Sex         int     `p:"sex"`
	Follows     int64   `p:"follows"`
	Fans        int64   `p:"fans"`
	Likes       int64   `p:"likes"`
	Grade       int64   `p:"grade"`
	Vip         int64   `p:"vip"`
	Description string  `p:"description"`
}

type UserEdit struct {
	UserId      int64   `p:"userId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	NickName    string  `p:"nickName"  v:"required#请设置用户昵称"`
	Password    string  `p:"password"  v:"required#请设置用户密码"`
	Phone       string  `p:"phone"`
	Email       string  `p:"email"`
	Avatar      string  `p:"avatar"`
	Cover       string  `p:"cover"`
	Balance     float64 `p:"balance"`
	Integral    int64   `p:"integral"`
	Sex         int     `p:"sex"`
	Follows     int64   `p:"follows"`
	Fans        int64   `p:"fans"`
	Likes       int64   `p:"likes"`
	Admin       int64   `p:"admin"`
	Grade       int64   `p:"grade"`
	Vip         int64   `p:"vip"`
	Description string  `p:"description"`
}
