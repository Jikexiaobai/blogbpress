package dto

type GradeCreate struct {
	Title       string   `p:"title"  v:"required#请输入标题"`
	Integral    int      `p:"integral"`
	CreateGroup int      `p:"createGroup"`
	Icon        string   `p:"icon"  v:"required#请上传图标"`
	PostsModule []string `p:"postsModule"`
	CommonAuth  []string `p:"commonAuth"`
}

type GradeEdit struct {
	GradeId     int64    `p:"gradeId"  v:"required|integer|min:1#请设置id|id必须为整型|id最小为1"` // 封面
	Title       string   `p:"title"  v:"required#请输入标题"`
	Integral    int      `p:"integral"`
	CreateGroup int      `p:"createGroup"`
	Icon        string   `p:"icon"  v:"required#请上传图标"`
	PostsModule []string `p:"postsModule"`
	CommonAuth  []string `p:"commonAuth"`
}
