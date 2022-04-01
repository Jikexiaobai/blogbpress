package result

import "github.com/gogf/gf/os/gtime"

type GroupList struct {
	GroupId     int64       `json:"id"`
	Title       string      `json:"title"`
	Cover       string      `json:"cover"`
	Description string      `json:"description"`
	Status      int         `json:"status"`
	CreateTime  *gtime.Time `json:"createTime"`
}

type GroupInfo struct {
	GroupId     int64         `json:"id"`          //
	UserInfo    *UserInfo     `json:"userInfo"`    // 发布用户
	CateInfo    *CategoryInfo `json:"cateInfo"`    // 发布用户
	Title       string        `json:"title"`       // 小组名称
	Icon        string        `json:"icon"`        // 小组图标
	Cover       string        `json:"cover"`       // 封面
	JoinMode    int           `json:"joinMode"`    // 小组类型1 公共小组，2付费小组，3私密小组
	Price       float64       `json:"price"`       // 费用
	Joins       int64         `json:"joins"`       //
	Contents    int64         `json:"contents"`    //
	Hots        int64         `json:"hots"`        //
	Views       int64         `json:"views"`       //
	IsJoin      bool          `json:"isJoin"`      // 是否加入
	Description string        `json:"description"` // 小组描述
	Remark      string        `json:"remark"`      // 备注
	Status      int           `json:"status"`      // 状态 0 全部  1待审   2通过  3拒绝
	CreateTime  *gtime.Time   `json:"createTime"`  //
}

type GroupListInfo struct {
	Id          int64         `c:"id" json:"id"`                   //
	Module      string        `c:"module" json:"module"`           //
	CateInfo    *CategoryInfo `json:"cateInfo"`                    // 发布用户
	UserInfo    *UserInfo     `c:"userInfo" json:"userInfo"`       // 发布用户
	Title       string        `c:"title" json:"title"`             // 小组名称
	Icon        string        `c:"icon" json:"icon"`               // 小组图标
	Cover       string        `c:"cover" json:"cover"`             // 封面
	Views       int64         `c:"views" json:"views"`             //
	Hots        int64         `c:"hots" json:"hots"`               //
	Joins       int64         `c:"joins" json:"joins"`             //
	Contents    int64         `c:"contents" json:"contents"`       //
	Status      int           `c:"status" json:"status"`           //
	Description string        `c:"description" json:"description"` // 小组描述
	CreateTime  *gtime.Time   `c:"createTime" json:"createTime"`   //
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
