package result

import "github.com/gogf/gf/os/gtime"

type AuthorityListInfo struct {
	AuthorityId int64       `json:"authorityId"` // 菜单ID
	Type        string      `json:"type"`        //
	Description string      `json:"description"` // 菜单名称
	Path        string      `json:"path"`        // 路由地址
	Mode        string      `json:"mode"`        // 菜单类型（1前台，2后台）
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
}
