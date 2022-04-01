package dto

type RegisterDto struct {
	Account  string `p:"account"  v:"required#请输入账号"`
	Captcha  string `p:"captcha" v:"required#请输入验证码"`
	Password string `p:"password" v:"required#请输入密码"`
	IP       string `p:"ip"`
}

type LoginDto struct {
	Account  string `p:"account"  v:"required#请输入账号"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
	IP       string `p:"ip"`
}

type SendCaptcha struct {
	//Key     string `p:"key"  v:"required#请输入验证码"`
	//Captcha string `p:"captcha"  v:"required#请输入验证码"`
	Account string `p:"account"  v:"required#请输入账号"`
}
