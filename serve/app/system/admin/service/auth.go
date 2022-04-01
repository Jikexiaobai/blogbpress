package service

import (
	"fiber/app/dao"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"strings"
)

var Auth = new(authService)

type authService struct {
}

var Token *gtoken.GfToken

func (s *authService) LoadToken() {
	Token = &gtoken.GfToken{
		CacheMode:      g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:       g.Cfg().GetString("gToken.CacheKey"),
		Timeout:        g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:     g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter: g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:     g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:    g.Cfg().GetString("gToken.AuthFailMsg"),
		//MultiLogin:      true,
		LoginPath:       "/api/v1/admin/auth/login",
		LoginBeforeFunc: s.LoginBefore,
		LoginAfterFunc:  s.LoginAfter,
		LogoutPath:      "/api/v1/admin/auth/logout",
		AuthPaths: g.SliceStr{
			"/api/v1/admin/*",
		},
		AuthExcludePaths: g.SliceStr{
			"/api/v1/admin/auth/*",
			"/api/v1/admin/system/info",
			"/api/v1/admin/system/imageCaptcha",
		},
		AuthAfterFunc: s.AuthAfter,
		//LogoutBeforeFunc: service.LoginOut,
	}
	Token.Start()
}

func (s *authService) LoginBefore(r *ghttp.Request) (string, interface{}) {
	var req *dto.AdminLogin
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.MangerLoginPassWordLock + req.Account) {
		response.Error(r).SetCode(response.INVALID).SetMessage("账号已锁定，请30分钟后再试").Send()
	}
	req.IP = r.GetClientIp()

	//// 检查图片验证码
	//if !imageCaptcha.Check(req.Key, req.Captcha) {
	//	response.Error(r).SetCode(response.EXCEPTION).SetMessage("验证码错误").Send()
	//}
	var tokenId int64
	err := gvalid.Check(req.Account, "email", nil)
	if err == nil {
		tokenUser, err := s.CheckLoginInfoByEmail(req.Account, req.Password, req.IP)
		// 检查用户名是否正确
		if err != nil {
			response.Error(r).SetCode(response.EXCEPTION).SetMessage(err.Error()).Send()
		}
		tokenId = tokenUser
	} else {
		tokenUser, err := s.CheckLoginInfoByPhone(req.Account, req.Password, req.IP)
		// 检查用户名是否正确
		if err != nil {
			response.Error(r).SetCode(response.EXCEPTION).SetMessage(err.Error()).Send()
		}
		tokenId = tokenUser
	}

	return gconv.String(tokenId) + "admin", tokenId
}

func (s *authService) LoginAfter(r *ghttp.Request, respData gtoken.Resp) {

	if !respData.Success() {
		response.Error(r).SetCode(response.INVALID).SetMessage(respData.Msg).SetData(respData).Send()
	}

	token := respData.GetString("token")
	data := gmap.New(true)
	data.Set("token", token)
	response.Success(r).SetCode(response.SUCCESS).SetData(data).Send()
}

func (s *authService) AuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	if r.Method == "OPTIONS" || respData.Success() {
		r.Middleware.Next()
	} else {
		respData.Msg = "请求错误或登录超时"
		res := r.Response
		options := res.DefaultCORSOptions()
		res.CORS(options)
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).SetMessage(respData.Msg).Send()
	}
}

func (s *authService) GetTokenId(r *ghttp.Request) (int64, error) {
	respData := Token.GetTokenData(r)
	if !respData.Success() {
		return 0, gerror.New(respData.Msg)
	}
	return gconv.Int64(respData.Get("data")), nil
}

// CheckLoginInfoByEmail 检查登录用户信息
func (s *authService) CheckLoginInfoByEmail(account, password, ip string) (int64, error) {
	userObj, err := dao.SysManger.Where(dao.SysManger.Columns.Email, account).
		Fields(
			dao.SysManger.Columns.Salt,
			dao.SysManger.Columns.Password,
			dao.SysManger.Columns.UserId,
		).
		One()
	if err != nil {
		return 0, err
	}
	if userObj == nil {
		return 0, gerror.New("找不到用户")
	}
	//校验密码
	randPwd := password + userObj.Salt
	randPwd = gmd5.MustEncryptString(randPwd)

	if !strings.EqualFold(userObj.Password, randPwd) {
		// 设置密码错误次数
		errTimes, _ := lock_utils.SetCount(shared.MangerLoginPassWordCount+account,
			shared.MangerLoginPassWordLock+account, 1800, 5)
		having := 5 - errTimes
		msg := gconv.String("密码不正确" + ",还有" + gconv.String(having) + "次之后账号将锁定")
		if having == 0 {
			msg = "账号已锁定，请30分钟后再试"
		}

		return 0, gerror.New(msg)
	}

	_, err = dao.SysManger.Data(g.Map{
		dao.SysManger.Columns.LoginIp:   ip,
		dao.SysManger.Columns.LoginTime: gtime.Datetime(),
	}).Where(dao.SysManger.Columns.UserId, userObj.UserId).Update()

	return userObj.UserId, nil
}

func (s *authService) CheckLoginInfoByPhone(account, password, ip string) (int64, error) {
	userObj, err := dao.SysManger.Where(dao.SysManger.Columns.Phone, account).
		Fields(
			dao.SysManger.Columns.Salt,
			dao.SysManger.Columns.Password,
			dao.SysManger.Columns.UserId,
		).
		One()
	if err != nil {
		return 0, err
	}
	if userObj == nil {
		return 0, gerror.New("找不到用户")
	}
	//校验密码
	randPwd := password + userObj.Salt
	randPwd = gmd5.MustEncryptString(randPwd)

	if !strings.EqualFold(userObj.Password, randPwd) {
		// 设置密码错误次数
		errTimes, _ := lock_utils.SetCount(shared.MangerLoginPassWordCount+account,
			shared.MangerLoginPassWordLock+account, 1800, 5)
		having := 5 - errTimes
		msg := gconv.String("密码不正确" + ",还有" + gconv.String(having) + "次之后账号将锁定")
		if having == 0 {
			msg = "账号已锁定，请30分钟后再试"
		}

		return 0, gerror.New(msg)
	}

	_, err = dao.SysManger.Data(g.Map{
		dao.SysManger.Columns.LoginIp:   ip,
		dao.SysManger.Columns.LoginTime: gtime.Datetime(),
	}).Where(dao.SysManger.Columns.UserId, userObj.UserId).Update()

	return userObj.UserId, nil
}
