package result

// 后台路由
type MenuInfo struct {
	MenuId    int64  `json:"menuId"`    // 菜单ID
	Title     string `json:"title"`     // 菜单名称
	ParentId  int64  `json:"parentId"`  // 父菜单ID
	Path      string `json:"path"`      // 路由地址
	Redirect  string `json:"redirect"`  // 重定向路径
	Component string `json:"component"` // 组件路径
	Icon      string `json:"icon"`      // 菜单图标
	IsFrame   int    `json:"isFrame"`   // 是否为外链（0是 1否）
	Mode      string `json:"mode"`
	Visible   int    `json:"visible"` // 菜单状态（0显示 1隐藏）
}
