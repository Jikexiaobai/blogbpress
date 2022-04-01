package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Auth = new(authApi)

type authApi struct {
}

// LoadRouter 加载 authController 路由
func (c *authApi) LoadRouter(group *ghttp.RouterGroup) {
	group.Group("/auth", func(group *ghttp.RouterGroup) {
		group.GET("/image/captcha", c.getImageCaptcha)
		group.GET("/option", c.getOption)
		group.POST("/send/captcha", c.sendCaptcha)
		group.POST("/register", c.register)
	})
}

// 获取 图片验证码
func (c *authApi) getImageCaptcha(r *ghttp.Request) {
	if rp, code := service.Captcha.CreateImageCaptcha(); code != response.SUCCESS {
		response.Error(r).SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("imageCaptcha", rp)
		response.Success(r).SetData(data).Send()
	}
}

// 获取注册配置
func (c *authApi) getOption(r *ghttp.Request) {
	// 业务处理
	if result, code := service.Auth.Option(); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("info", result)
		response.Success(r).SetData(data).Send()
	}
}

// 发送邮箱验证码
func (c *authApi) sendCaptcha(r *ghttp.Request) {
	var req *dto.SendCaptcha
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	//判断验证码是否正确
	//if !service.Captcha.CheckImageCaptcha(req.Key, req.Captcha) {
	//	response.Error(r).
	//		SetCode(response.PARAM_INVALID).
	//		SetMessage("验证码错误").
	//		Send()
	//}
	ip := r.GetClientIp()
	//判断邮箱是否锁了
	if lock_utils.CheckLock(shared.UserRegisterAccountLock + req.Account) {
		response.Error(r).
			SetCode(response.FAILD).
			SetMessage("该账户已经锁住，请30分钟后再来测试").Send()
	}

	//判断IP是否被锁
	if lock_utils.CheckLock(shared.UserRegisterIpLock + ip) {
		response.Error(r).
			SetCode(response.FAILD).
			SetMessage("您的ip已被上锁，请30分钟后再来测试").Send()
	}

	// 业务逻辑
	if code := service.Captcha.SendCaptcha(req.Account, ip); code != response.SUCCESS {
		response.Error(r).
			SetCode(code).
			SetMessage(response.CodeMsg(code)).
			Send()
	} else {
		response.Success(r).Send()
	}
}

func (c *authApi) sendPhoneCaptcha(r *ghttp.Request) {
	userName := r.GetForm("number")
	if err := gvalid.Check(userName, "required", "请设置用户名"); err != nil {
		response.Error(r).SetMessage(err.Error()).Send()
	}
	// 获取验证码方式验证用户名格式
	g.Dump(userName)
}

func (c *authApi) register(r *ghttp.Request) {
	var req *dto.RegisterDto
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	val, err := service.Config.FindValue(shared.AuthSetting)
	if err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}
	j := gjson.New(val)
	valid := gconv.String(j.Get("registerMode"))

	if valid == "email" {
		if err := gvalid.Check(req.Account, "required|email", "请设置邮箱|邮箱格式不正确"); err != nil {
			response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
		}
		//判断邮箱是否注册
		if service.User.CheckEmailUniqueAll(req.Account) {
			response.Error(r).SetCode(response.INVALID).SetMessage(req.Account + "已存在").Send()
		}
	} else if valid == "phone" {
		if err := gvalid.Check(req.Account, "required|phone", "请设置手机号|手机号格式不正确"); err != nil {
			response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
		}
		//判断邮箱是否注册
		if service.User.CheckPhoneUniqueAll(req.Account) {
			response.Error(r).SetCode(response.INVALID).SetMessage(req.Account + "已存在").Send()
		}
	}

	//验证 验证码是否正确
	if err := service.Captcha.CheckCaptcha(req.Account, req.Captcha); err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	req.IP = r.GetClientIp()
	// 业务处理
	if code := service.Auth.Register(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		response.Success(r).Send()
	}
}
