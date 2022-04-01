package result

import "github.com/gogf/gf/os/gtime"

type MangerInfo struct {
	UserId    int64    `json:"userId"`
	NickName  string   `json:"nickName"`
	Avatar    string   `json:"avatar"`
	Authority []string `json:"authority"`
	Roles     []string `json:"roles"`
}

type MangerList struct {
	UserId     int64       `json:"id"`
	NickName   string      `json:"nickName"`
	Avatar     string      `json:"avatar"`
	LoginIp    string      `json:"loginIp"`
	LoginTime  *gtime.Time `json:"loginTime"`
	Status     int         `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
}

type MangerEditInfo struct {
	NickName string  `json:"nickName"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Email    string  `json:"email"`
	Avatar   string  `json:"avatar"`
	Sex      int     `json:"sex"`
	RoleId   []int64 `json:"roleId"`
}
