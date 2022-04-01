package dto

type UserFansOrFollowsQuery struct {
	Page    int    `p:"page" v:"required#请设置页数"`
	Limit   int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	UserId  int64  `p:"userId" v:"required|not-in:0|min:1#请设置用户id|用户id不能为0|用户id不能为0"`  // 所属用户
	Related string `p:"related" v:"required|in:fans,follow#请设置查询关系|参数只允许为fans,follow"` // 所属用户
}

type UserSignQuery struct {
	Page  int `p:"page" v:"required#请设置页数"`
	Limit int `p:"limit" v:"between:1,100#参数只允许1到100"`
	Type  int `p:"type" v:"required|between:1,2#请设置查询关系|参数只允许为1,2"` // 所属用户
}

// 修改密码的数据
type PassWordDto struct {
	OldPass string `p:"oldPass"  v:"required|length:6,30#请输入旧密码|密码长度为:min到:max位"`
	NewPass string `p:"newPass" v:"required|length:6,30#请输入新密码|密码长度为:min到:max位"`
}

// 设置邮箱数据
type EmailDto struct {
	Email   string `p:"email"  v:"required|email#请输入邮箱|邮箱格式不正确"`
	Captcha string `p:"captcha" v:"required#请输入验证码"`
}
