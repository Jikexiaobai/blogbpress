package upload_lib

import "github.com/gogf/gf/os/gtime"

type MediaObject struct {
	UserId     int64       `orm:"user_id"          json:"user_id"`     // 上传的用户
	Link       string      `orm:"link"             json:"link"`        // 文件链接
	Path       string      `orm:"path"             json:"path"`        // 存放路径
	Name       string      `orm:"name"             json:"name"`        // 文件名称
	OrName     string      `orm:"or_name"          json:"or_name"`     // 原始文件名称
	Size       string      `orm:"size"             json:"size"`        // 文件大小
	UploadKey  int         `orm:"upload_key"       json:"upload_key"`  // 上传方式 1 为本地上传， 2为oss上传
	Status     int         `orm:"status"       json:"status"`          // 上
	Ext        string      `orm:"ext"              json:"ext"`         // 文件后缀
	MediaType  string      `orm:"media_type"       json:"media_type"`  // 文件类型
	CreateTime *gtime.Time `orm:"create_time"      json:"create_time"` // 创建时间
	DeleteTime *gtime.Time `orm:"delete_time"      json:"delete_time"` //
}
