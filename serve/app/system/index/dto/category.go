package dto

type CategoryCreate struct {
	Module      string `p:"module"  v:"required#请设置所属模块"` // 封面
	Title       string `p:"title"  v:"required#请输入分类标题"`  // 标题
	Slug        string `p:"slug"  v:"required#请输入分类别名"`
	ParentId    int64  `p:"parentId"`
	Description string `p:"description"`
}
