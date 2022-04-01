package result

import "github.com/gogf/gf/os/gtime"

// 音频
type MediaListInfo struct {
	Id         int64       `json:"id"`       //
	OrName     string      `json:"orName"`   //
	Size       string      `json:"size"`     // 发布用户
	Ext        string      `json:"ext"`      // 标题
	NickName   string      `json:"nickName"` // 封面
	Path       string      `json:"path"`
	Link       string      `json:"link"`
	UploadKey  int         `json:"uploadKey"`
	Status     int         `json:"status"`
	CreateTime *gtime.Time `c:"createTime"  json:"createTime"` // 创建时间
}
