package captcha_lib

import "github.com/mojocn/base64Captcha"

// Data json request body.
type Data struct {
	Id  string
	B64 string
}

// Config json request body.
type Config struct {
	ID            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptchaHandler json request body.
func GenerateCaptchaHandler(param Config) (Data, error) {
	//parse request parameters
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	rp := Data{
		Id:  id,
		B64: b64s,
	}
	return rp, err
}

// VerifyCaptchaHandle base64Captcha verify http handler
func VerifyCaptchaHandle(param Config) bool {
	return store.Verify(param.ID, param.VerifyValue, true)
}
