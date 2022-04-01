package dto

type AuthorityCreate struct {
	Title     string `p:"title" v:"required#请设置权限标题"`
	Type      int    `p:"type" v:"required#请设置类型"`
	Perms     string `p:"perms" v:"required#请设置权限标识"`
	ParentId  int64  `p:"parentId"`
	Path      string `p:"path"`
	Redirect  string `p:"redirect"`
	OrderBy   string `p:"orderBy"`
	Component string `p:"component"`
	Icon      string `p:"icon"`
	Hidden    int    `p:"hidden"`
	Target    int    `p:"target"`
}

type AuthorityEdit struct {
	AuthorityId int64  `p:"id"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"`
	Title       string `p:"title" v:"required#请设置权限标题"`
	Type        int    `p:"type" v:"required#请设置类型"`
	Perms       string `p:"perms" v:"required#请设置权限标识"`
	ParentId    int64  `p:"parentId"`
	Path        string `p:"path"`
	Redirect    string `p:"redirect"`
	OrderBy     string `p:"orderBy"`
	Component   string `p:"component"`
	Icon        string `p:"icon"`
	Hidden      int    `p:"hidden"`
	Target      int    `p:"target"`
}

type AuthorityQuery struct {
	Title string `p:"title"`
}
