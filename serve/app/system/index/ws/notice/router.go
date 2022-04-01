package notice

import (
	"encoding/json"
	"fmt"
	"gitee.com/xushenghao/glib"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gutil"
	"time"
)

var Notice = new(noticeWs)
var maxTimeOut = 3600

type noticeWs struct {
	ConnMap        *gmap.TreeMap // 连接数
	SendQueue      *gqueue.Queue // 发送队列
	BroadcastQueue *gqueue.Queue // 广播队列
	Receive        func(data WSData)
}

func (c *noticeWs) LoadRouter(s *ghttp.Server) {
	c.ConnMap = gmap.NewTreeMap(gutil.ComparatorString, true)
	c.SendQueue = gqueue.New(1000)
	c.BroadcastQueue = gqueue.New(10)
	c.loop()
	s.BindHandler("/ws/notice", c.start)
}

func (c *noticeWs) start(r *ghttp.Request) {

	var wh WSHandle
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}

	wh.Handle = ws
	key := fmt.Sprintf("%s-%s", ws.RemoteAddr().String(), grand.Digits(10))
	wh.key = key
	c.ConnMap.Set(key, wh)

	// 后端心跳
	var cTimer *gtimer.Entry
	heartbeatTime := gtime.Now()
	cTimer = gtimer.Add(10*time.Second, func() {
		if glib.TimeSubSecond(time.Now(), heartbeatTime.Time) > maxTimeOut {
			_ = wh.Handle.Close()
			wh.Handle = nil
		}
	})

	defer func() {
		if wh.Handle != nil {
			_ = wh.Handle.Close()
		}
		c.ConnMap.Remove(key)
		cTimer.Close()
	}()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}

		if msgType == ghttp.WS_MSG_TEXT {
			var rp, response WSData
			g.Dump(rp)
			if err := gjson.DecodeTo(msg, &rp); err != nil {
				glog.Debug(err)
				continue
			}
			rp.WsKey = key
			switch rp.Code {
			case "ping":
				heartbeatTime = gtime.Now()
				response.WsKey = key
				response.Code = "pong"
				c.Send(response, false)
			default:
				g.Dump("asdasd")
			}
		} else if msgType == ghttp.WS_MSG_CLOSE {
			return
		}
	}
}

// WsCount 获取在线数量
func (c *noticeWs) WsCount() int {
	return c.ConnMap.Size()
}

// Send 发送消息
func (c *noticeWs) Send(data WSData, broadcast bool) {
	if broadcast {
		if c.BroadcastQueue.Size() < 10 {
			c.BroadcastQueue.Push(data)
		}
	} else {
		if c.SendQueue.Size() < 1000 {
			c.SendQueue.Push(data)
		}
	}
}

// loop 循环消费队列
func (c *noticeWs) loop() {
	go func() {
		for {
			select {
			case v := <-c.SendQueue.C:
				if v != nil {
					response := v.(WSData)
					bs, err := json.Marshal(response)
					if err != nil {
						glog.Info(err)
					} else {
						if c.ConnMap.Contains(response.WsKey) {
							t := c.ConnMap.Get(response.WsKey)
							if t != nil {
								ws := t.(WSHandle)
								if ws.Handle != nil {
									if err := ws.Handle.WriteMessage(ghttp.WS_MSG_TEXT, bs); err != nil {
										if ws.Handle != nil {
											ws.Handle.Close()
											c.ConnMap.Remove(ws.key)
										}
										glog.Info(err)
									}
								}
							}
						}
					}
				}
			case v := <-c.BroadcastQueue.C:
				if v != nil {
					response := v.(WSData)
					bs, err := json.Marshal(response)
					if err != nil {
						glog.Info(err)
					} else {
						if c.ConnMap.Size() != 0 {
							array := c.ConnMap.Values()
							for _, v := range array {
								ws := v.(WSHandle)
								if ws.Handle != nil {
									if err := ws.Handle.WriteMessage(ghttp.WS_MSG_TEXT, bs); err != nil {
										glog.Info(err)
										if ws.Handle != nil {
											_ = ws.Handle.Close()
											c.ConnMap.Remove(ws.key)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}()
}
