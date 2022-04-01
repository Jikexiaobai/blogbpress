package service

import (
	"fiber/app/system/index/shared"
	"fiber/app/tools/response"
	"fiber/library/redis"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/mojocn/base64Captcha"

	lock_utils "fiber/app/tools/lock"
	captcha_lib "fiber/library/captcha"
)

var Captcha = new(captchaService)

type captchaService struct {
}

// 创建图片验证码
func (s *captchaService) CreateImageCaptcha() (*captcha_lib.Data, response.ResponseCode) {
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
		return nil, response.INVALID
	}
	return &rp, response.SUCCESS
}

// 检查图片验证码
func (s *captchaService) CheckImageCaptcha(Key, code string) bool {
	//校验验证码
	verifyCaptcha := captcha_lib.VerifyCaptchaHandle(captcha_lib.Config{
		ID:          Key,
		VerifyValue: code,
	})
	return verifyCaptcha
}

// SendCaptcha 创建发送验证码
func (s *captchaService) SendCaptcha(account string, Ip string) response.ResponseCode {
	//设置账户发送次数
	_, err := lock_utils.SetCount(shared.UserRegisterAccountCount+account,
		shared.UserRegisterAccountLock+account, 1800, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}
	// 设置IP发送次数
	_, err = lock_utils.SetCount(shared.UserRegisterIpCount+Ip,
		shared.UserRegisterIpLock+Ip, 1800, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	authSetting, err := Config.FindValue(shared.AuthSetting)
	if err != nil {
		return response.FAILD
	}
	authJson := gjson.New(authSetting)
	registerMode := gconv.String(authJson.Get("registerMode"))

	switch registerMode {
	case "email":
		err = s.sendEmailCaptcha(account)
		if err != nil {
			return response.FAILD
		}
	case "phone":
	}

	return response.SUCCESS
}

//发送邮箱验证码
func (s *captchaService) sendEmailCaptcha(account string) error {
	// 生成验证码
	captcha := grand.Digits(6)

	//获取网站名称
	result, err := Config.FindValue("BaseSetting")
	if err != nil {
		return err
	}
	j := gjson.New(result)
	title := gconv.String(j.Get("title"))

	template := `<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Document</title>
   <style type="text/css">::-webkit-scrollbar{ display: none; }</style>
   <style id="cloudAttachStyle" type="text/css">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>
   <style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style>
   <style type="text/css">
       body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}
       td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}
       pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}
       th,td{font-family:arial,verdana,sans-serif;line-height:1.666}
       img{ border:0}
       header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}
       blockquote{margin-right:0px}
   </style>
   <style>
       .container{
           height: 100vh;
           display: flex;
           justify-content: center;
           align-items: center;
           background-color: #e2e2e2;
       }
       .email-captcha-box{
           width: 600px;
           height: 400px;
           background-color: white;
           border-radius: 5px;
       }
       .email-captcha-title{
           height: 60px;
           border-top-right-radius: 5px;
           border-top-left-radius: 5px;
           background: #2980B9;  /* fallback for old browsers */
           background: -webkit-linear-gradient(to left, #FFFFFF, #6DD5FA, #2980B9);  /* Chrome 10-25, Safari 5.1-6 */
           background: linear-gradient(to left, #FFFFFF, #6DD5FA, #2980B9); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
           display: flex;
           align-items: center;
           font-size: 18px;
           color: white;
           padding-left: 40px;
       }
       .email-captcha-content{
           padding: 40px;
           font-size: 20px;
       }
       .email-captcha-1{
           font-size: 16px;
           margin-bottom: 30px;
       }
       .email-captcha > span{
           border: 1px solid orange;
           border-radius: 10px;
           color: orange;
           padding: 10px;
       }
       .email-captcha-2{
           font-size: 16px;
           margin-top: 30px;
       }
       .email-captcha-site{
           font-size: 14px;
           color: rgb(165, 165, 165);
       }
       .email-captcha-hr{
           margin: 10px 0;
           width: 100%;
           height: 1px;
           background-color: rgb(226, 226, 226);
       }
   </style>
</head>

<body>
   <div class="container">
       <div class="email-captcha-box">
           <div class="email-captcha-title">
               欢迎加入我们，请验证您的邮箱
           </div>
           <div class="email-captcha-content">
              <div>你好：</div>
              <div class="email-captcha-1">您的邮箱为：` + account + `，验证码为：</div>
              <div class="email-captcha"><span>` + captcha + `</span></div>
              <div class="email-captcha-2">验证码的有效期为5分钟，请在有效期内输入！</div>
              <div class="email-captcha-site">
               — ` + title + `
              </div>
              <div class="email-captcha-hr"></div>
              <div>
               本邮件为系统邮件不能回复，请勿回复。
              </div>
           </div>
       </div>
   </div>
</body>
</html>`
	Email.Tos = account
	Email.Subject = "请查收你的验证码"
	Email.Body = template
	Email.ContentType = "html"
	//发送验证码
	err = Email.Send()
	if err != nil {
		return err
	}

	//验证码存入缓存
	var redisCom redis.Com
	redisCom.Key = "user_register_captcha_" + account
	redisCom.Data = captcha
	redisCom.Time = "900"
	_ = redisCom.SetStringEX()

	return nil
}

//发送邮箱验证码
func (s *captchaService) SendPhoneCaptcha(account string) error {
	// 生成验证码
	captcha := grand.Digits(6)

	//获取网站名称
	result, err := Config.FindValue("BaseSetting")
	if err != nil {
		return err
	}
	j := gjson.New(result)
	title := gconv.String(j.Get("title"))
	g.Dump(title)
	//验证码存入缓存
	var redisCom redis.Com
	redisCom.Key = "user_register_captcha_" + account
	redisCom.Data = captcha
	redisCom.Time = "900"
	_ = redisCom.SetStringEX()

	return nil
}

// 检查邮箱验证码
func (s *captchaService) CheckCaptcha(key, code string) error {
	var redisCom redis.Com
	redisCom.Key = "user_register_captcha_" + key

	captcha, err := redisCom.GetString()

	if err != nil {
		return err
	}
	if captcha == nil {
		return gerror.New("验证码错误")
	}
	captchaStr := gconv.String(captcha)

	if captchaStr != code {
		return gerror.New("验证码不正确")
	}
	err = redisCom.DELString()
	if err != nil {
		return err
	}
	return nil
}
