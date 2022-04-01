package result

type UserInfo struct {
	UserId      int64      `json:"id"`
	NickName    string     `json:"nickName"`
	Follows     int64      `json:"follows"`
	Fans        int64      `json:"fans"`
	Posts       int        `json:"posts"`
	Likes       int64      `json:"likes"`
	Cover       string     `json:"cover"`
	Integral    int64      `json:"integral"`
	Avatar      string     `json:"avatar"`
	Description string     `json:"description"`
	Sex         int        `json:"sex"`
	IsFollow    bool       `json:"isFollow"`
	IsVerify    bool       `json:"isVerify"`
	Grade       *UserGrade `json:"grade"`
	Vip         *UserVip   `json:"vip"`
}

type UserGrade struct {
	Title string `json:"title"` // 角色名称
	Icon  string `json:"icon"`  // 角色图标
}

type UserVip struct {
	Title string `json:"title"` // 角色名称
	Icon  string `json:"icon"`  // 角色图标
	Color string `json:"color"` // 角色图标

}

type UserListInfo struct {
	UserId      int64      `c:"id" json:"id"`
	Module      string     `c:"module" json:"module"` //
	NickName    string     `c:"nickName" json:"nickName"`
	Follows     int64      `c:"follows" json:"follows"`
	Fans        int64      `c:"fans" json:"fans"`
	Comments    int64      `c:"comments" json:"comments"`
	Likes       int64      `c:"likes" json:"likes"`
	Cover       string     `c:"cover" json:"cover"`
	Integral    int64      `c:"integral" json:"integral"`
	Avatar      string     `c:"avatar" json:"avatar"`
	Description string     `c:"description" json:"description"`
	Sex         int        `c:"sex" json:"sex"`
	IsFollow    bool       `c:"isFollow" json:"isFollow"`
	IsVerify    bool       `c:"isVerify" son:"isVerify"`
	Grade       *UserGrade `c:"grade" json:"grade"`
	Vip         *UserVip   `c:"vip" json:"vip"`
	CreateTime  string     `json:"createTime"`
}

type UserHotList struct {
	UserId   int64  `c:"id" json:"id"`
	NickName string `c:"nickName" json:"nickName"`
	Avatar   string `c:"avatar" json:"avatar"`
}

type UserSignList struct {
	UserId     int64  `json:"id"`
	NickName   string `json:"nickName"`
	Count      int    `json:"count"`
	Integral   int64  `json:"integral"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"createTime"`
}
