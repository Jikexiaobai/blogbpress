package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"fiber/app/tools/utils"
	"fiber/library/redis"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gvalid"
	"strings"
)

var Auth = new(authService)

type authService struct {
}

var Token *gtoken.GfToken

func (s *authService) LoadToken() {
	Token = &gtoken.GfToken{
		CacheMode:       g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:        g.Cfg().GetString("gToken.CacheKey"),
		Timeout:         g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:      g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:  g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:      g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:     g.Cfg().GetString("gToken.AuthFailMsg"),
		MultiLogin:      true,
		LoginPath:       "/api/v1/web/auth/login",
		LoginBeforeFunc: s.LoginBefore,
		LoginAfterFunc:  s.LoginAfter,
		LogoutPath:      "/api/v1/web/auth/logout",
		LogoutAfterFunc: s.LogoutAfter,
		AuthPaths: g.SliceStr{
			"/api/v1/web/*",
		},
		AuthExcludePaths: g.SliceStr{
			"/api/v1/web/auth/*",
			"/api/v1/web/system/*",
			"/api/v1/web/test/*",
			"/api/v1/web/pay/*",
			"/api/v1/web/user/reward",
			"/api/v1/web/user/info",
			"/api/v1/web/user/posts",
			"/api/v1/web/user/sign",
			"/api/v1/web/user/fansOrFollows",
			"/api/v1/web/user/answerList",
			"/api/v1/web/edu/info",
			"/api/v1/web/screen/list",
			"/api/v1/web/group/info",
			"/api/v1/web/group/posts",
			"/api/v1/web/resource/info",
			"/api/v1/web/audio/info",
			"/api/v1/web/video/info",
			"/api/v1/web/article/info",
			"/api/v1/web/question/info",
			"/api/v1/web/question/list",
			"/api/v1/web/answer/info",
			"/api/v1/web/answer/list",
			"/api/v1/web/topic/info",
			"/api/v1/web/topic/list",
			"/api/v1/web/topic/top",
			"/api/v1/web/comment/list",
		},
		AuthAfterFunc: s.AuthAfter,
		//LogoutBeforeFunc: service.LoginOut,
	}
	Token.Start()
}

func (s *authService) LoginBefore(r *ghttp.Request) (string, interface{}) {
	var req *dto.LoginDto
	if err := r.Parse(&req); err != nil {
		response.Error(r).SetCode(response.PARAM_INVALID).SetMessage(err.Error()).Send()
	}

	//检查是否用户被锁
	if lock_utils.CheckLock(shared.UserLoginPassWordLock + req.Account) {
		response.Error(r).SetCode(response.INVALID).SetMessage("账号已锁定，请30分钟后再试").Send()
	}
	req.IP = r.GetClientIp()

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

	return gconv.String(tokenId), tokenId
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

func (s *authService) LogoutAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		response.Error(r).SetCode(response.INVALID).SetMessage(respData.Msg).SetData(respData).Send()
	}
	response.Success(r).Send()
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

// Register 注册
func (s *authService) Register(req *dto.RegisterDto) (code response.ResponseCode) {
	//获取一个随机数

	var entity model.SysUser
	err := gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}
	userSetting, err := Config.FindValue(shared.UserSetting)
	if err != nil {
		return response.DB_READ_ERROR
	}
	userJson := gjson.New(userSetting)

	defaultAvatar := gconv.String(userJson.Get("defaultAvatar"))
	defaultCover := gconv.String(userJson.Get("defaultCover"))
	defaultRole := gconv.String(userJson.Get("defaultGrade"))
	entity.Avatar = defaultAvatar
	entity.Cover = defaultCover
	entity.Grade = gconv.Int64(defaultRole)

	integralSetting, err := Config.FindValue(shared.IntegralSetting)
	if err != nil {
		return response.DB_READ_ERROR
	}
	integralJson := gjson.New(integralSetting)
	registerIntegral := gconv.Int64(integralJson.Get("registerIntegral"))
	entity.Integral = registerIntegral

	authSetting, err := Config.FindValue(shared.AuthSetting)
	if err != nil {
		return response.DB_READ_ERROR
	}
	authJson := gjson.New(authSetting)
	registerMode := authJson.Get("registerMode")
	switch registerMode {
	case "email":
		entity.Email = req.Account
	case "phone":
		entity.Phone = req.Account
	}

	randNumber := grand.Digits(6)
	entity.NickName = "新用户" + randNumber

	newSalt := utils.RandS(6)
	newToken := req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	entity.Salt = newSalt
	entity.Password = newToken
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.LoginTime = gtime.Now()
	entity.LoginIp = req.IP
	entity.Sex = 3
	entity.Status = 2

	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	res, err := tx.Save(dao.SysUser.Table, entity)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	rid, _ := res.LastInsertId()
	if rid <= 0 {
		return response.DB_SAVE_ERROR
	}
	pathList := make([]string, 0)
	pathList = append(pathList, defaultAvatar)
	pathList = append(pathList, defaultCover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.User)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	// 设置通知
	if entity.Status == shared.StatusReviewed {
		noticeSetting, err := Config.FindValue(shared.NoticeSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		noticeJson := gjson.New(noticeSetting)
		register := gconv.String(noticeJson.Get("register"))

		baseSetting, err := Config.FindValue("BaseSetting")
		if err != nil {
			return response.DB_READ_ERROR
		}
		base := gjson.New(baseSetting)
		title := gconv.String(base.Get("title"))

		// 设置通知

		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.SystemType = shared.NoticeSysTemRegister
		notice.Receiver = rid
		notice.Content = gstr.Replace(register, "{title}", title, -1)
		notice.CreateTime = gtime.Now()
		notice.Status = shared.NoticeStatusReview

		_, err = tx.Insert(dao.SysNotice.Table, notice)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

func (s *authService) Option() (*result.Auth, response.ResponseCode) {

	var redisCom redis.Com
	redisCom.Key = "AuthOption"
	authOption, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if authOption != nil {
		var authInfo *result.Auth
		err := gconv.Struct(authOption, &authInfo)
		if err != nil {
			return nil, response.INVALID
		}
		return authInfo, response.SUCCESS
	}

	var auth *result.Auth
	info, err := dao.SysConfig.Fields(dao.SysConfig.Columns.ConfigValue).
		Where(dao.SysConfig.Columns.ConfigKey, shared.AuthSetting).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	err = gconv.Struct(info.ConfigValue, &auth)
	// 写入缓存
	redisCom.Time = "10"
	redisCom.Data = auth
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}
	return auth, response.SUCCESS
}

// CheckLoginInfoByEmail 检查登录用户信息
func (s *authService) CheckLoginInfoByEmail(account, password, ip string) (int64, error) {
	userObj, err := dao.SysUser.Where(dao.SysUser.Columns.Email, account).
		Fields(
			dao.SysUser.Columns.Salt,
			dao.SysUser.Columns.Password,
			dao.SysUser.Columns.UserId,
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
		errTimes, _ := lock_utils.SetCount(shared.UserLoginPassWordCount+account,
			shared.UserLoginPassWordLock+account, 1800, 5)
		having := 5 - errTimes
		msg := gconv.String("密码不正确" + ",还有" + gconv.String(having) + "次之后账号将锁定")
		if having == 0 {
			msg = "账号已锁定，请30分钟后再试"
		}

		return 0, gerror.New(msg)
	}

	_, err = dao.SysUser.Data(g.Map{
		dao.SysUser.Columns.LoginIp:   ip,
		dao.SysUser.Columns.LoginTime: gtime.Datetime(),
	}).Where(dao.SysUser.Columns.UserId, userObj.UserId).Update()

	return userObj.UserId, nil
}

func (s *authService) CheckLoginInfoByPhone(account, password, ip string) (int64, error) {
	userObj, err := dao.SysUser.Where(dao.SysUser.Columns.Phone, account).
		Fields(
			dao.SysUser.Columns.Salt,
			dao.SysUser.Columns.Password,
			dao.SysUser.Columns.UserId,
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
		errTimes, _ := lock_utils.SetCount(shared.UserLoginPassWordCount+account,
			shared.UserLoginPassWordLock+account, 1800, 5)
		having := 5 - errTimes
		msg := gconv.String("密码不正确" + ",还有" + gconv.String(having) + "次之后账号将锁定")
		if having == 0 {
			msg = "账号已锁定，请30分钟后再试"
		}

		return 0, gerror.New(msg)
	}

	_, err = dao.SysUser.Data(g.Map{
		dao.SysUser.Columns.LoginIp:   ip,
		dao.SysUser.Columns.LoginTime: gtime.Datetime(),
	}).Where(dao.SysUser.Columns.UserId, userObj.UserId).Update()

	return userObj.UserId, nil
}
