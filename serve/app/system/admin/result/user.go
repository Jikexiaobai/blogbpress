package result

import "github.com/gogf/gf/os/gtime"

type UserList struct {
	UserId     int64       `json:"id"`
	NickName   string      `json:"nickName"`
	Integral   int64       `json:"integral"`
	Balance    float64     `json:"balance"`
	Avatar     string      `json:"avatar"`
	LoginIp    string      `json:"loginIp"`
	LoginTime  *gtime.Time `json:"loginTime"`
	Status     int         `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
}

//
type UserEditInfo struct {
	NickName    string  `json:"nickName"`
	Password    string  `json:"password"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Avatar      string  `json:"avatar"`
	Cover       string  `json:"cover"`
	Balance     float64 `json:"balance"`
	Integral    int64   `json:"integral"`
	Sex         int     `json:"sex"`
	Follows     int64   `json:"follows"`
	Fans        int64   `json:"fans"`
	Likes       int64   `json:"likes"`
	Admin       int64   `json:"admin"`
	Grade       int64   `json:"grade"`
	Vip         int64   `json:"vip"`
	Description string  `json:"description"`
}
