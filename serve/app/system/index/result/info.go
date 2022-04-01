package result

import (
	"github.com/gogf/gf/os/gtime"
)

// 文件
type FileInfo struct {
	Link   string `json:"link"`   // 文件链接
	OrName string `json:"orName"` // 原始文件名称
	Size   string `json:"size"`   // 文件大小
}

//圈子内容
type ContentInfo struct {
	Id          int64  `json:"id"`          //
	Title       string `json:"title"`       // 小组名称
	Description string `json:"description"` // 小组名称
	Hots        int64  `json:"hots"`        // 封面
	Answers     int64  `json:"answers"`     // 回答数
	Views       int64  `json:"views"`       // 封面

	Cover      string       `json:"cover"`  // 封面
	Files      string       `json:"files"`  // 封面
	Module     string       `json:"module"` // 封面
	Status     int          `json:"status"` //
	ViewMode   int          `json:"viewMode"`
	Price      int          `json:"price"`
	IsView     bool         `json:"isView"`
	Type       int          `json:"type"`
	UserInfo   *AccountInfo `json:"userInfo"`   // 发布用户
	CreateTime *gtime.Time  `json:"createTime"` //
}
