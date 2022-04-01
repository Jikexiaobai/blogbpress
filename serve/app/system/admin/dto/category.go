package dto

type CategoryQuery struct {
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Title  string `p:"title"`
	Module string `p:"module"`
}

type CategoryCreate struct {
	Module      string `p:"module"  v:"required#请设置所属模块"` // 封面
	Title       string `p:"title"  v:"required#请输入分类标题"`  // 标题
	Slug        string `p:"slug"  v:"required#请输入分类别名"`
	IsTop       int    `p:"isTop"`
	Cover       string `p:"cover"`
	ParentId    int64  `p:"parentId"`
	Description string `p:"description"`
}

type CategoryEdit struct {
	CateId      int64  `p:"id"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"`
	Module      string `p:"module"  v:"required#请设置所属模块"` // 封面
	Title       string `p:"title"  v:"required#请输入分类标题"`  // 标题
	Slug        string `p:"slug"  v:"required#请输入分类别名"`
	IsTop       int    `p:"isTop"`
	Cover       string `p:"cover"`
	ParentId    int64  `p:"parentId"`
	Description string `p:"description"`
}
