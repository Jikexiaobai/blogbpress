package imageCaptcha

import captcha_lib "fiber/library/captcha"

// Check 检查图片验证码
func Check(Key, code string) bool {
	//校验验证码
	verifyCaptcha := captcha_lib.VerifyCaptchaHandle(captcha_lib.Config{
		ID:          Key,
		VerifyValue: code,
	})
	return verifyCaptcha
}
