package result

import "github.com/gogf/gf/os/gtime"

type CategoryInfo struct {
	CateId   int64  ` json:"cateId"`  //
	ParentId int64  `json:"parentId"` // 顶级分类
	Title    string `json:"title"`    // 分类名称
}

type CategoryList struct {
	CateId      int64       `json:"id"`          //
	ParentId    int64       `json:"parentId"`    // 顶级分类
	Module      string      `json:"module"`      // 所属模块
	Title       string      `json:"title"`       // 分类名称
	Slug        string      `json:"slug"`        // 分类别名
	Cover       string      `json:"cover"`       // 分类背景图
	Sort        int         `json:"sort"`        // 分类排序
	Description string      `json:"description"` // 分类描述
	IsTop       int         `json:"isTop"`       // 分类状态 1 禁用 2启用
	Remark      string      `json:"remark"`      // 备注
	CreateTime  *gtime.Time `json:"createTime"`  // 创建日期
}

type CategoryEditInfo struct {
	ParentId    int64  `json:"parentId"`    // 顶级分类
	Module      string `json:"module"`      // 所属模块
	Title       string `json:"title"`       // 分类名称
	Slug        string `json:"slug"`        // 分类别名
	Cover       string `json:"cover"`       // 分类背景图
	Sort        int    `json:"sort"`        // 分类排序
	Description string `json:"description"` // 分类描述
	IsTop       int    `json:"isTop"`       // 分类状态 1 禁用 2启用
}
