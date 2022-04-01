package dto

type AdminAuthorityCreate struct {
	Description string `p:"description" v:"required#请设置描述"` // 描述
	Type        string `p:"type" v:"required#请设置类型"`        // 描述
	Mode        string `p:"mode" v:"required#请设置模式"`        // 描述
	Path        string `p:"path" v:"required#请设置标识"`        // 描述
}

type AdminAuthorityQuery struct {
	Page  int    `p:"page" v:"required#请设置页数"`
	Limit int    `p:"limit" v:"between:1,100#参数只允许1到100"`
	Mode  string `p:"mode"`
	Type  string `p:"type"`
}
