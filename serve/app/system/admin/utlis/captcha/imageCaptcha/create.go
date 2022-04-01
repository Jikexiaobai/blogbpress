package imageCaptcha

import (
	captcha_lib "fiber/library/captcha"
	"github.com/mojocn/base64Captcha"
)

// Create 创建图片验证码
func Create() (*captcha_lib.Data, error) {
	driver := base64Captcha.DriverString{
		Height:          47,
		Width:           160,
		NoiseCount:      15,
		ShowLineOptions: base64Captcha.OptionShowSineLine,
		Length:          4,
		Source:          "qwertyuipkjhgfdsazxcvbnm",
		BgColor:         nil,
		Fonts:           []string{"RitaSmith.ttf", "chromohv.ttf"},
	}
	param := captcha_lib.Config{
		ID:            "",
		CaptchaType:   "string",
		VerifyValue:   "",
		DriverAudio:   nil,
		DriverString:  &driver,
		DriverChinese: nil,
		DriverMath:    nil,
		DriverDigit:   nil,
	}
	rp, err := captcha_lib.GenerateCaptchaHandler(param)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}
