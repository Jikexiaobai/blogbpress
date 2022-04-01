package archive

import (
	"fiber/app/system/archive/api"
	"github.com/gogf/gf/net/ghttp"
)

func Init(s *ghttp.Server) {
	s.Group("public/", func(group *ghttp.RouterGroup) {
		api.Thumbnail.LoadRouter(group)
	})
}
