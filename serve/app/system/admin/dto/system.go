package dto

type SystemCreate struct {
	ConfigName  string `p:"configName"`                       // 参数名称
	ConfigKey   string `p:"configKey" v:"required#请输入权限字符"`   // 参数键名
	ConfigValue string `p:"configValue" v:"required#请输入权限字符"` // 参数键值
	Remark      string `p:"remark"`                           // 备注
}
