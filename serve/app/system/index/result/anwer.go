package result

import "github.com/gogf/gf/os/gtime"

type AnswerInfo struct {
	AnswerId   int64       `json:"id"`       //
	UserInfo   *UserInfo   `json:"userInfo"` // 发布用户
	Content    string      `json:"content"`  //
	Files      string      `json:"files"`    //
	Likes      int64       `json:"likes"`
	Price      float64     `json:"price"`
	IsLike     bool        `json:"isLike"` // 是否点赞
	IsAdoption bool        `json:"isAdoption"`
	Status     int         `json:"status"`     //
	CreateTime *gtime.Time `json:"createTime"` //
}
