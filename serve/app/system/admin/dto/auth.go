package dto

type AdminLogin struct {
	IP string `p:"ip"`
	//Key      string `p:"key"  v:"required#请输入验证码"`
	//Captcha  string `p:"captcha"  v:"required#请输入验证码"`
	Account  string `p:"account"  v:"required#请输入账号"`
	Password string `p:"passWord"  v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
}
