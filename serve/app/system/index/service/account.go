package service

import (
	"fiber/app/dao"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	"fiber/app/tools/response"
	"fiber/library/redis"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

var Account = new(accountService)

type accountService struct {
}

// Sign 用户签到
func (s *accountService) Sign(userId int64) (int, response.ResponseCode) {
	val, err := Config.FindValue(shared.IntegralSetting)
	if err != nil {
		return 0, response.DB_READ_ERROR
	}
	j := gjson.New(val)
	signInIntegral := gconv.String(j.Get("signInIntegral"))
	arr := gstr.Split(signInIntegral, "-")
	integralRand := grand.N(gconv.Int(arr[0]), gconv.Int(arr[1]))

	var redisCom redis.Com
	redisCom.Filed = gconv.String(userId)
	redisCom.Data = gconv.String(integralRand) + "_" + gconv.String(gtime.Now())
	redisCom.Key = shared.UserSignIn
	err = redisCom.SetHashField()
	if err != nil {
		return 0, response.CACHE_SAVE_ERROR
	}

	redisCom.Key = shared.UserSignToday
	err = redisCom.SetHashField()
	if err != nil {
		return 0, response.CACHE_SAVE_ERROR
	}

	redisCom.Data = gconv.String(userId)
	redisCom.Key = shared.UserSignInList
	err = redisCom.LPush()
	if err != nil {
		return 0, response.CACHE_SAVE_ERROR
	}
	return integralRand, response.SUCCESS
}

// SelectInfo 获取账户信息
func (s *accountService) SelectInfo(userId int64) (*result.AccountInfo, response.ResponseCode) {

	var redisCom redis.Com
	redisCom.Key = shared.AccountInfo + gconv.String(userId)
	//accountObj, err := redisCom.GetString()
	//if err != nil {
	//	return nil, response.CACHE_READ_ERROR
	//}
	//if accountObj != nil {
	//	var accountInfo *result.AccountInfo
	//	err := gconv.Struct(accountObj, &accountInfo)
	//	if err != nil {
	//		return nil, response.CACHE_READ_ERROR
	//	}
	//	return accountInfo, response.SUCCESS
	//}

	// 获取登录用户的信息
	info, err := dao.SysUser.Fields(
		dao.SysUser.Columns.Fans,
		dao.SysUser.Columns.Follows,
		dao.SysUser.Columns.Likes,
		dao.SysUser.Columns.Vip,
		dao.SysUser.Columns.Grade,
		dao.SysUser.Columns.Sex,
		dao.SysUser.Columns.Avatar,
		dao.SysUser.Columns.NickName,
		dao.SysUser.Columns.UserId,
		dao.SysUser.Columns.Cover,
		dao.SysUser.Columns.Description,
		dao.SysUser.Columns.Integral,
	).Where(dao.SysUser.Columns.UserId, userId).One()
	if err != nil || info == nil {
		return nil, response.NOT_FOUND
	}

	var accountInfo result.AccountInfo
	accountInfo.UserId = info.UserId
	accountInfo.NickName = info.NickName
	accountInfo.Follows = info.Follows
	accountInfo.Fans = info.Fans
	accountInfo.Likes = info.Likes
	accountInfo.Cover = info.Cover
	accountInfo.Integral = info.Integral
	accountInfo.Avatar = info.Avatar
	accountInfo.Description = info.Description
	accountInfo.Sex = info.Sex

	// 判断用户是否认证
	accountInfo.IsVerify = Verify.CheckIsVerify(userId)

	// 获取会员信息
	if info.Vip != 0 {
		vip, err := Vip.SelectAccountVip(userId, info.Vip)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		accountInfo.Vip = vip
	}

	// 获取等级信息
	grade, err := Grade.SelectUserGrade(info.Grade)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	accountInfo.Grade = grade

	// 获取用户投稿数
	var count int
	// 文章数量
	articleCount, err := dao.SysArticle.Where(dao.SysArticle.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	count += articleCount
	//资源数量
	resourceCount, err := dao.SysResource.Where(dao.SysResource.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	count += resourceCount
	//音频数量
	audioCount, err := dao.SysAudio.Where(dao.SysAudio.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	count += audioCount
	//视频数量
	videoCount, err := dao.SysVideo.Where(dao.SysVideo.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	count += videoCount
	//课程数量
	eduCount, err := dao.SysEdu.Where(dao.SysEdu.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	count += eduCount
	accountInfo.Posts = count

	// 是否去签到
	redisCom.Filed = gconv.String(userId)
	redisCom.Key = shared.UserSignToday
	res := redisCom.CheckHashField()
	if res {
		integral := redisCom.GetHashFieldString()
		arr := gstr.Split(integral, "_")
		sign := result.AccountSign{
			IsSign:   true,
			Integral: gconv.Int64(arr[0]),
		}
		accountInfo.Sign = &sign
	}

	// 检查会员是否过期
	if accountInfo.Vip != nil {
		if gtime.Now().After(gconv.GTime(accountInfo.Vip.FinishTime)) {
			_, err := dao.SysUser.Update(g.Map{
				dao.SysUser.Columns.Vip: nil,
			}, dao.SysUser.Columns.UserId, userId)
			if err != nil {
				return nil, response.UPDATE_FAILED
			}

			err = Vip.RemoveUserVip(info.Vip, userId)
			if err != nil {
				return nil, response.DELETE_FAILED
			}
		}
	}

	//存入缓存
	redisCom.Key = shared.AccountInfo + gconv.String(userId)
	redisCom.Time = "600"
	redisCom.Data = accountInfo
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	return &accountInfo, response.SUCCESS
}

// Edit 修改账户信息
func (s *accountService) Edit(req *dto.AccountBase) response.ResponseCode {

	var userDto = g.Map{
		dao.SysUser.Columns.Avatar:      req.Avatar,
		dao.SysUser.Columns.Cover:       req.Cover,
		dao.SysUser.Columns.Description: req.Description,
		dao.SysUser.Columns.NickName:    req.NickName,
		dao.SysUser.Columns.Sex:         req.Sex,
	}
	tx, err := g.DB().Begin()
	if err != nil {
		return response.EXCEPTION
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = dao.SysUser.Update(userDto, dao.SysUser.Columns.UserId, req.UserId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.UserId, shared.User)
	if err != nil {
		return response.DELETE_FAILED
	}
	//pathList, err := utils.GetSrcLink(req.Content)
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Avatar)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.UserId, "user")
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	// 设置缓存
	var redisCom redis.Com
	redisCom.Key = shared.AccountInfo + gconv.String(req.UserId)
	err = redisCom.DELString()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}
	return response.SUCCESS
}

// SelectSecurity 获取用户安全信息
func (s *accountService) SelectSecurity(userId int64) (*result.AccountSecurity, response.ResponseCode) {

	var security result.AccountSecurity
	user, err := dao.SysUser.Fields(
		dao.SysUser.Columns.Email,
		dao.SysUser.Columns.Phone).Where(dao.SysUser.Columns.UserId, userId).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	security.Email = user.Email
	security.Phone = user.Phone

	return &security, response.SUCCESS
}

// GetBalance 获取用户余额
func (s *accountService) GetBalance(userId int64) (float64, response.ResponseCode) {
	res, err := s.SelectBalance(userId)
	if err != nil {
		return 0, response.DB_READ_ERROR
	}
	return res, response.SUCCESS
}

// SelectBalance 获取用户余额
func (s *accountService) SelectBalance(userId int64) (float64, error) {
	res, err := dao.SysUser.
		Value(dao.SysUser.Columns.Balance,
			dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return 0, err
	}
	return gconv.Float64(res), nil
}

// EditBalance 修改用户余额
func (s *accountService) EditBalance(tx *gdb.TX, userId int64, money float64) error {
	_, err := tx.Update(dao.SysUser.Table, g.Map{
		dao.SysUser.Columns.Balance: money,
	}, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return err
	}
	return nil
}

// EditIntegral 修改用户积分
func (s *accountService) EditIntegral(tx *gdb.TX, userId, Integral int64) error {
	_, err := tx.Update(dao.SysUser.Table, g.Map{
		dao.SysUser.Columns.Integral: Integral,
	}, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return err
	}
	return nil
}

// EditGrade 修改用户等级
func (s *accountService) EditGrade(tx *gdb.TX, userId, gradeId int64) error {
	_, err := tx.Update(dao.SysUser.Table, g.Map{
		dao.SysUser.Columns.Grade: gradeId,
	}, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return err
	}
	return nil
}
