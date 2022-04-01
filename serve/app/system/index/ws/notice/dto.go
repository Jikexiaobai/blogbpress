package notice

import "github.com/gogf/gf/net/ghttp"

type WSHandle struct {
	Handle *ghttp.WebSocket
	key    string
}
