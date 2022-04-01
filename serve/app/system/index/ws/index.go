package ws

import (
	"fiber/app/system/index/ws/notice"
	"github.com/gogf/gf/net/ghttp"
)

func SetUp(s *ghttp.Server) {
	notice.Notice.LoadRouter(s)
}
