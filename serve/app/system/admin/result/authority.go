package result

import "github.com/gogf/gf/os/gtime"

type AuthorityList struct {
	AuthorityId int64       `json:"id"`         // 菜单ID
	Title       string      `json:"title"`      // 菜单名称
	ParentId    int64       `json:"parentId"`   // 父菜单ID
	Component   string      `json:"component"`  // 组件地址
	Path        string      `json:"path"`       // 请求地址
	Redirect    string      `json:"redirect"`   // 请求地址
	Perms       string      `json:"perms"`      // 权限标识
	Target      int         `json:"target"`     // 打开方式（1页签 2新窗口）
	Type        int         `json:"type"`       // 菜单类型（1目录 2菜单 3按钮）
	Hidden      int         `json:"hidden"`     // 菜单状态（2显示 1隐藏）
	Icon        string      ` json:"icon"`      // 菜单图标
	CreateTime  *gtime.Time `json:"createTime"` // 创建时间
}

type AuthorityEditInfo struct {
	Title     string `json:"title"`     // 菜单名称
	ParentId  int64  `json:"parentId"`  // 父菜单ID
	OrderBy   int    `json:"orderBy"`   // 排序
	Component string `json:"component"` // 组件地址
	Path      string `json:"path"`      // 请求地址
	Redirect  string `json:"redirect"`  // 请求地址
	Perms     string `json:"perms"`     // 权限标识
	Target    int    `json:"target"`    // 打开方式（1页签 2新窗口）
	Type      int    `json:"type"`      // 菜单类型（1目录 2菜单 3按钮）
	Hidden    int    `json:"hidden"`    // 菜单状态（2显示 1隐藏）
	Icon      string ` json:"icon"`     // 菜单图标
}
