package result

import "github.com/gogf/gf/os/gtime"

// 视频
type VideoList struct {
	VideoId    int64       `json:"id"`         //
	NickName   string      `json:"nickName"`   // 发布用户
	Title      string      `json:"title"`      // 标题
	Cover      string      `json:"cover"`      // 封面
	Status     int         `json:"status"`     // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
	Category   string      `json:"category"`   // 发布用户
}

// 视频
type VideoEditInfo struct {
	UserId      int64       `json:"userId"`      // 分类id
	CateId      int64       `json:"cateId"`      // 分类
	Title       string      `json:"title"`       // 标题
	Cover       string      `json:"cover"`       // 封面
	VideoMode   int         `json:"videoMode"`   // 查看权限 0公开下载，1付费下载，2评论下载，3登录下载
	VideoPrice  float64     `json:"videoPrice"`  //
	Link        string      `json:"link"`        // 视频地址
	Hots        int64       `json:"hots"`        //
	Likes       int64       `json:"likes"`       // 点赞数
	Favorites   uint64      `json:"favorites"`   // 收藏
	Views       int64       `json:"views"`       // 播放量
	HasDown     int         `json:"hasDown"`     // 是否有下载1没有，2有
	DownMode    int         `json:"downMode"`    // 下载权限 0公开下载，1付费下载，2评论下载，3登录下载
	Price       float64     `json:"price"`       //
	DownUrl     string      `json:"downUrl"`     //
	Purpose     string      `json:"purpose"`     //
	Attribute   string      `json:"attribute"`   //
	Description string      `json:"description"` // 描述
	Status      int         `json:"status"`      // 状态：0全部,1待审核 ，2已发布 ，3拒绝，4草稿
	CreateTime  *gtime.Time `json:"createTime"`  //
	TagList     []*TagList  `json:"tagList"`     // 备注
}
