package result

import "github.com/gogf/gf/os/gtime"

type TagListInfo struct {
	TagId      int64       `json:"tagId"`      //
	Title      string      `json:"title"`      // 标题
	Top        int         `json:"top"`        // 是否推荐 1不推荐 2推荐
	Slug       string      `json:"slug"`       // 别名
	CreateTime *gtime.Time `json:"createTime"` // 创建日期
}
