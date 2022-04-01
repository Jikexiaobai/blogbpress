package result

import "github.com/gogf/gf/os/gtime"

type TagList struct {
	TagId int64  ` json:"tagId"`
	Name  string `json:"title"`
}

type TagListInfo struct {
	TagId      int64       `json:"id"`         //
	Title      string      `json:"title"`      // 标题
	Top        int         `json:"top"`        // 是否推荐 1不推荐 2推荐
	CreateTime *gtime.Time `json:"createTime"` // 创建日期
}
