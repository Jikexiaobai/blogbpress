package result

import "github.com/gogf/gf/os/gtime"

// 问题
type QuestionInfo struct {
	QuestionId int64       `json:"id"`        //
	UserInfo   *UserInfo   `json:"userInfo"`  // 发布用户
	GroupInfo  *GroupJoin  `json:"groupInfo"` // 圈子
	Title      string      `json:"title"`     //
	Content    string      `json:"content"`   //
	Likes      int64       `json:"likes"`
	Hots       int64       `json:"hots"`
	Views      int64       `json:"views"`
	Favorites  int64       `json:"favorites"`
	Answers    int64       `json:"answers"`
	IsLike     bool        `json:"isLike"`     // 是否点赞
	IsFavorite bool        `json:"isFavorite"` // 是否收藏
	Anonymous  int         `json:"anonymous"`  // 是否匿名
	Status     int         `json:"status"`     //
	CreateTime *gtime.Time `json:"createTime"` //
}

// 问题
type QuestionList struct {
	QuestionId int64       `json:"id"`                    //
	UserInfo   *UserInfo   `c:"userInfo" json:"userInfo"` // 发布用户
	Title      string      `c:"title" json:"title"`       //
	Likes      int64       `c:"likes" json:"likes"`
	Hots       int64       `c:"hots" json:"hots"`
	Views      int64       `c:"views" json:"views"`
	Favorites  int64       `c:"favorites" json:"favorites"`
	Status     int         `c:"status"  json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	Answers    int64       `c:"answers" json:"answers"`
	CreateTime *gtime.Time `c:"createTime" json:"createTime"` //
}

// 问题
type QuestionListInfo struct {
	Id        int64     `c:"id" json:"id"`             //
	Module    string    `c:"module" json:"module"`     //
	UserInfo  *UserInfo `c:"userInfo" json:"userInfo"` // 发布用户
	Title     string    `c:"title" json:"title"`       //
	Likes     int64     `c:"likes" json:"likes"`
	Hots      int64     `c:"hots" json:"hots"`
	Views     int64     `c:"views" json:"views"`
	Favorites int64     `c:"favorites" json:"favorites"`
	Status    int       `c:"status"  json:"status"` // 状态：0 待审核 ，1已发布 ，3 草稿
	//Answers    int64        `c:"answers" json:"answers"`
	CreateTime *gtime.Time `c:"createTime" json:"createTime"` //
}

type QuestionEditInfo struct {
	UserId    int64      `json:"userId"`    //
	Title     string     `json:"title"`     //
	Content   string     `json:"content"`   //
	Hots      int64      `json:"hots"`      //
	Favorites int64      `json:"favorites"` //
	Likes     int64      `json:"likes"`     //
	Views     int64      `json:"views"`     //
	Anonymous int        `json:"anonymous"` // 是否匿名
	GroupInfo *GroupJoin `json:"groupInfo"` // 圈子
}
