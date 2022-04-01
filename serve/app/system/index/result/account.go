package result

import "github.com/gogf/gf/os/gtime"

type AccountSecurity struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type AccountInfo struct {
	UserId      int64        `json:"userId"`
	NickName    string       `json:"nickName"`
	Follows     int64        `json:"follows"`
	Fans        int64        `json:"fans"`
	Posts       int          `json:"posts"`
	Likes       int64        `json:"likes"`
	Cover       string       `json:"cover"`
	Integral    int64        `json:"integral"`
	Avatar      string       `json:"avatar"`
	Description string       `json:"description"`
	Sex         int          `json:"sex"`
	IsVerify    bool         `json:"isVerify"`
	Sign        *AccountSign `json:"sign"`
	Grade       *UserGrade   `json:"grade"`
	Vip         *AccountVip  `json:"vip"`
}

type AccountSign struct {
	IsSign   bool  `json:"isSign"`
	Integral int64 `json:"integral"`
}

type AccountGrade struct {
	Title string `json:"title"` // 角色名称
	Icon  string `json:"icon"`  // 角色图标
}

type AccountVip struct {
	Title      string      `json:"title"`      // 角色名称
	Icon       string      `json:"icon"`       // 角色图标
	Color      string      `json:"color"`      // 角色图标
	Discount   float64     `json:"discount"`   // 角色图标
	FinishTime *gtime.Time `json:"finishTime"` // 结束时间
}

type AccountRewardInfo struct {
	UserId int64  `json:"userId"`
	Avatar string `json:"avatar"`
}

type AccountVerifyStatusIsPayPrice struct {
	IsPay  bool    `json:"isPay"`
	Status int     `json:"status"`
	Price  float64 `json:"price"`
}
