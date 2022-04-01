package result

import "github.com/gogf/gf/os/gtime"

type GroupList struct {
	GroupId    int64       `json:"id"`         //
	Category   string      `json:"category"`   // 发布用户
	NickName   string      `json:"nickName"`   // 发布用户
	Title      string      `json:"title"`      // 小组名称
	Cover      string      `json:"cover"`      // 封面
	Joins      int64       `json:"joins"`      //
	Contents   int64       `json:"contents"`   //
	Status     int         `json:"status"`     //
	CreateTime *gtime.Time `json:"createTime"` //
}

// 圈子
type GroupEditInfo struct {
	GroupId     int64   `json:"groupId"`     //
	UserId      int64   `json:"userId"`      //
	CateId      int64   `json:"cateId"`      // 小组分类
	Title       string  `json:"title"`       // 小组名称
	Cover       string  `json:"cover"`       // 封面
	Price       float64 `json:"price"`       // 费用
	SecretKey   string  `json:"secretKey"`   // 私人密钥
	JoinMode    int     `json:"joinMode"`    // 小组类型1 公共小组，2付费小组，3私密小组
	ViewMode    int     `json:"viewMode"`    // 小组权限，1内容公开显示 2加入后显示
	Icon        string  `json:"icon"`        // 小组图标
	Description string  `json:"description"` // 小组描述
	Status      int     `json:"status"`      // 状态 0 全部  1待审   2拒绝 3拒绝
	Remark      string  `json:"remark"`      // 备注
	Views       int64   `json:"views"`       //
	Hots        int64   `json:"hots"`        //
	Joins       int64   `json:"joins"`       //
	Contents    int64   `json:"contents"`    //
}
