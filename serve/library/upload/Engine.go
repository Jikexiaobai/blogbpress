package upload_lib

import (
	"github.com/gogf/gf/net/ghttp"
)

// 存储引擎
type Engine interface {
	Upload(Id int64, files ghttp.UploadFiles) ([]MediaObject, error)
	Remove() error
}
