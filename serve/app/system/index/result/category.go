package result

import "github.com/gogf/gf/os/gtime"

// 分类
type CategoryInfo struct {
	CateId   int64  ` json:"cateId"`  //
	ParentId int64  `json:"parentId"` // 顶级分类
	Title    string `json:"title"`    // 分类名称
}

type CategoryListInfo struct {
	CateId      int64       `json:"cateId"`      //
	ParentId    int64       `json:"parentId"`    // 顶级分类
	Module      string      `json:"module"`      // 所属模块
	Title       string      `json:"title"`       // 分类名称
	Slug        string      `json:"slug"`        // 分类别名
	Cover       string      `json:"cover"`       // 分类背景图
	Sort        int         `json:"sort"`        // 分类排序
	Keywords    string      `json:"keywords"`    // 分类关键字
	Description string      `json:"description"` // 分类描述
	Status      int         `json:"status"`      // 分类状态 1 禁用 2启用
	Remark      string      `json:"remark"`      // 备注
	CreateTime  *gtime.Time `json:"createTime"`  // 创建日期
}
