package result

import "github.com/gogf/gf/os/gtime"

// 文章
type VerifyList struct {
	VerifyId   int64       `json:"id"`
	NickName   string      `json:"nickName"`
	Name       string      `json:"name"`       // 真实姓名
	Code       string      `json:"code"`       // 身份证号码
	Mode       int         `json:"mode"`       // 联系方式 1 qq, 2微信
	Number     string      `json:"number"`     // 联系号码
	CreateTime *gtime.Time `json:"createTime"` // 认证时间
	Status     int         `json:"status"`     // 状态（1待审，2成功，3拒绝）
}
