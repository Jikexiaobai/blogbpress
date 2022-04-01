package result

import "github.com/gogf/gf/os/gtime"

type Role struct {
	RoleId int64  `json:"id"`    // 角色名称
	Title  string `json:"title"` // 角色名称
}

// RoleList 角色
type RoleList struct {
	RoleId     int64       `json:"id"`
	Title      string      `json:"title"`
	Status     int         `json:"status"`
	CreateTime *gtime.Time `json:"createTime"`
	//Authority  []string    `json:"authority"` // 积分
}

type RoleEditInfo struct {
	Title     string  `json:"title"`     // 角色名称
	Status    int     `json:"status"`    // 积分
	Authority []int64 `json:"authority"` // 积分
}
