package result

import "github.com/gogf/gf/os/gtime"

// TopicTopList 置顶话题列表
type TopicTopList struct {
	TopicId int64  `json:"id"` //
	Title   string `json:"title"`
}

type TopicInfo struct {
	TopicId     int64       `json:"id"`
	UserInfo    *UserInfo   `json:"userInfo"`    // 发布用户
	GroupInfo   *GroupJoin  `json:"groupInfo"`   // 圈子
	RelatedInfo RelatedInfo `json:"relatedInfo"` // 圈子
	Title       string      `json:"title"`
	Type        int         `json:"type"`
	Files       string      `json:"files"`
	Hots        int64       `json:"hots"`     // 热度
	Likes       int64       `json:"likes"`    // 点赞数
	Views       int64       `json:"views"`    // 播放量
	Comments    int64       `json:"comments"` // 播放量
	IsLike      bool        `json:"isLike"`
	IsTop       int         `json:"isTop"`
	Status      int         `json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime  *gtime.Time `json:"createTime"`
}

// 话题
type TopicListInfo struct {
	Id          int64       `c:"id" json:"id"`                   //
	Module      string      `c:"module" json:"module"`           //
	UserInfo    *UserInfo   `c:"userInfo" json:"userInfo"`       // 发布用户
	GroupInfo   GroupJoin   `c:"groupInfo" json:"groupInfo"`     // 圈子
	RelatedInfo RelatedInfo `c:"relatedInfo" json:"relatedInfo"` // 圈子
	Title       string      `c:"title" json:"title"`
	Type        int         `c:"type" json:"type"`
	Files       string      `c:"files" json:"files"`
	Hots        int64       `c:"hots" json:"hots"`         // 热度
	Comments    int64       `c:"comments" json:"comments"` // 热度
	Likes       int64       `c:"likes" json:"likes"`       // 点赞数
	Views       int64       `c:"views" json:"views"`       // 播放量
	IsLike      bool        `c:"isLike" json:"isLike"`
	IsTop       int         `c:"isTop" json:"isTop"`
	Status      int         `c:"status"  json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime  *gtime.Time `c:"createTime" json:"createTime"`
}

// TopicList 话题列表
type TopicList struct {
	TopicId    int64       `json:"id"`        //
	UserInfo   *UserInfo   `json:"userInfo"`  // 发布用户
	GroupInfo  GroupJoin   `json:"groupInfo"` // 圈子
	Title      string      `json:"title"`
	Type       int         `json:"type"`
	Files      string      `json:"files"`
	Hots       int64       `json:"hots"`  // 热度
	Likes      int64       `json:"likes"` // 点赞数
	Views      int64       `json:"views"` // 播放量
	IsLike     bool        `json:"isLike"`
	IsTop      int         `json:"isTop"`
	Status     int         `json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	CreateTime *gtime.Time `json:"createTime"`
}

type RelatedInfo struct {
	Id     int64  `json:"id"`     //
	Module string `json:"module"` //
	Title  string `json:"title"`
	Cover  string `json:"cover"`
}
