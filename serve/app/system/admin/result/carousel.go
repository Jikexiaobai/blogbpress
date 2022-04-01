package result

import "github.com/gogf/gf/os/gtime"

// 音频
type CarouselListInfo struct {
	Id          int64       `json:"id"`          // 类型1(投稿内容)，2(其他内容)
	Mode        int         `json:"mode"`        // 类型1(投稿内容)，2(其他内容)
	ContentInfo interface{} `json:"contentInfo"` //
	Link        string      `json:"link"`        // 链接
	Cover       string      `json:"cover"`       // 封面地址
	Position    string      `json:"position"`    //
	Type        int         `json:"type"`        //
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
}
