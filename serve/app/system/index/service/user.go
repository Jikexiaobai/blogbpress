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
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var User = new(userService)

type userService struct {
}

func (s *userService) RemoveUserLike(tx *gdb.TX, related int64, module string) error {
	_, err := tx.Model(dao.SysUserLike.Table).
		Where(dao.SysUserLike.Columns.Module, module).
		Where(dao.SysUserLike.Columns.RelatedId, related).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) RemoveUserFavorite(tx *gdb.TX, related int64, module string) error {
	_, err := tx.Model(dao.SysUserFavorite.Table).
		Where(dao.SysUserFavorite.Columns.Module, module).
		Where(dao.SysUserFavorite.Columns.FavoriteId, related).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Follow(userId, id int64) (code response.ResponseCode) {

	// 加入锁限制
	_, err := lock_utils.SetCount(shared.FollowCount+gconv.String(userId)+gconv.String(id),
		shared.FollowLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否关注
	count, err := dao.SysUserFollow.
		Where(dao.SysUserFollow.Columns.UserId, userId).
		Where(dao.SysUserFollow.Columns.FollowId, id).
		Count()
	if err != nil {
		return response.DB_READ_ERROR
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var redisCom redis.Com
	redisCom.Key = shared.UserInfo + gconv.String(id)
	accountObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	if count == 0 {

		//	修改用户信息
		var entity model.SysUserFollow
		entity.UserId = userId
		entity.FollowId = id
		_, err := tx.Insert(dao.SysUserFollow.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}
		// 更新关注数量
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Follows: gdb.Raw("follows+1"),
		}, dao.SysUser.Columns.UserId, userId)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 更新关注用户粉丝数量
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Fans: gdb.Raw("fans+1"),
		}, dao.SysUser.Columns.UserId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if accountObj != nil {
			var accountInfo *result.UserInfo
			err := gconv.Struct(accountObj, &accountInfo)
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
			accountInfo.IsFollow = true
			accountInfo.Fans = accountInfo.Fans + 1
			redisCom.Time = "180"
			redisCom.Data = accountInfo
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}
		// 通知被点赞用户
		var notice model.SysNotice
		notice.Type = shared.NoticeFollow
		notice.FromUserId = userId
		notice.Status = shared.NoticeStatusReview
		notice.Content = "关注了你"
		notice.Receiver = id
		notice.CreateTime = gtime.Now()
		_, err = tx.Insert(dao.SysNotice.Table, notice)
		if err != nil {
			return response.DB_SAVE_ERROR
		}

		//	设置点赞积分
		err = Integral.SetUserFollowIntegral(redisCom, tx, userId)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	} else {
		_, err := tx.Model(dao.SysUserFollow.Table).
			Where(dao.SysUserFollow.Columns.UserId, userId).
			Where(dao.SysUserFollow.Columns.FollowId, id).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新关注数量
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Follows: gdb.Raw("follows-1"),
		}, dao.SysUser.Columns.UserId, userId)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 更新关注用户粉丝数量
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Fans: gdb.Raw("fans-1"),
		}, dao.SysUser.Columns.UserId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}
		if accountObj != nil {
			var accountInfo *result.UserInfo
			err := gconv.Struct(accountObj, &accountInfo)
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
			accountInfo.IsFollow = false
			if accountInfo.Fans > 0 {
				accountInfo.Fans = accountInfo.Fans - 1
			} else {
				accountInfo.Fans = 0
			}

			redisCom.Time = "180"
			redisCom.Data = accountInfo
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}
	}

	//	设置用户活跃度
	redisCom.Key = shared.UserHot
	redisCom.Data = userId
	err = redisCom.ADDSet()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	return response.SUCCESS
}

// SelectFilterList 获取过滤列表
func (s *userService) SelectFilterList(req *dto.QueryParam) (int, []*result.UserListInfo, error) {

	model := dao.SysUser.SysUserDao.Where(dao.SysUser.Columns.DeleteTime, nil)
	model = model.Where(dao.SysUser.Columns.Status, shared.StatusReviewed)

	if req.Title != "" {
		model = model.Where(dao.SysUser.Columns.NickName+" like ?", "%"+req.Title+"%")
	}

	// 查询签到的用户列表
	//switch req.Type {
	//case shared.UserToday:
	//	var redisCom redis.Com
	//	redisCom.Key = shared.UserClockInList
	//	total, err := redisCom.LLength()
	//	if err != nil {
	//		return 0, nil, err
	//	}
	//	if req.Page <= 0 {
	//		req.Page = 1
	//	}
	//	start := (req.Page - 1) * req.Limit
	//	resUserId, err := redisCom.LRange(start, req.Limit)
	//	if err != nil {
	//		return 0, nil, err
	//	}
	//	var res []*result.UserListInfo
	//	for _, i := range resUserId {
	//		var contentInfo result.UserClockUserListInfoInInfo
	//		redisCom.Filed = i
	//		redisCom.Key = shared.UserClockIn
	//		clockInInfo := redisCom.GetHashFieldString()
	//		arr := gstr.Split(clockInInfo, "_")
	//		contentInfo.Integral = gconv.Int64(arr[0])
	//		contentInfo.CreateTime = arr[1]
	//		info, err := model.
	//			Fields(dao.SysUser.Columns.Avatar,
	//				dao.SysUser.Columns.NickName).
	//			Where(dao.SysUser.Columns.UserId, gconv.Int64(i)).One()
	//		if err != nil {
	//			return 0, nil, err
	//		}
	//		if info != nil {
	//			contentInfo.NickName = info.NickName
	//			contentInfo.Avatar = info.Avatar
	//			contentInfo.UserId = gconv.Int64(i)
	//			res = append(res, &contentInfo)
	//		}
	//
	//	}
	//
	//	return total, res, nil
	//	//case shared.UserContinuous:
	//	//	clockInModel := dao.SysUserClockIn.SysUserClockInDao.Order(dao.SysUserClockIn.Columns.Count + " desc")
	//	//	total, err := clockInModel.Count()
	//	//	if err != nil {
	//	//		return 0, nil, err
	//	//	}
	//	//	clockInModel = clockInModel.Page(req.Page, req.Limit)
	//	//	list, err := clockInModel.All()
	//	//	if err != nil {
	//	//		return 0, nil, err
	//	//	}
	//	//	var res []*result.AccountClockInInfo
	//	//	for _, i := range list {
	//	//		var contentInfo result.AccountClockInInfo
	//	//		contentInfo.Count = i.Count
	//	//		contentInfo.CreateTime = gconv.String(i.CreateTime)
	//	//		info, err := model.
	//	//			Fields(dao.SysUser.Columns.Avatar,
	//	//				dao.SysUser.Columns.NickName).
	//	//			Where(dao.SysUser.Columns.UserId, i.UserId).One()
	//	//		if err != nil {
	//	//			return 0, nil, err
	//	//		}
	//	//		if info != nil {
	//	//			contentInfo.NickName = info.NickName
	//	//			contentInfo.Avatar = info.Avatar
	//	//			contentInfo.UserId = gconv.Int64(i)
	//	//			res = append(res, &contentInfo)
	//	//		}
	//	//	}
	//	//	return total, res, nil
	//}

	model = model.Order(dao.SysUser.Columns.UpdateTime + " desc")
	total, err := model.Count()
	if err != nil {
		return 0, nil, err
	}
	model = model.Fields(
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
	)
	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, nil
	}

	var res []*result.UserListInfo
	for _, i := range list {
		var contentInfo *result.UserListInfo
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.Module = shared.User
		res = append(res, contentInfo)
	}
	return total, res, nil
}

// SelectInfo 获取用户信息
func (s *userService) SelectInfo(userId, id int64) (*result.UserInfo, error) {

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
	).Where(dao.SysUser.Columns.UserId, id).One()
	if err != nil || info == nil {
		return nil, err
	}
	res, err := s.info(userId, info)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Info 获取用户信息
func (s *userService) Info(userId, id int64) (*result.UserInfo, response.ResponseCode) {
	info, err := s.SelectInfo(userId, id)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	return info, response.SUCCESS
}

// resetInfo 获取用户信息
func (s *userService) info(userId int64, info *model.SysUser) (*result.UserInfo, error) {
	var redisCom redis.Com
	redisCom.Key = shared.UserInfo + gconv.String(info.UserId)
	//userObj, err := redisCom.GetString()
	//if err != nil {
	//	return nil, err
	//}
	//if userObj != nil {
	//	var userInfo *result.UserInfo
	//	err := gconv.Struct(userObj, &userInfo)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return userInfo, nil
	//}

	var userInfo result.UserInfo
	userInfo.UserId = info.UserId
	userInfo.NickName = info.NickName
	userInfo.Follows = info.Follows
	userInfo.Fans = info.Fans
	userInfo.Likes = info.Likes
	userInfo.Cover = info.Cover
	userInfo.Integral = info.Integral
	userInfo.Avatar = info.Avatar
	userInfo.Description = info.Description
	userInfo.Sex = info.Sex

	// 判断用户是否认证
	userInfo.IsVerify = Verify.CheckIsVerify(info.UserId)

	// 获取会员信息
	if info.Vip != 0 {
		vip, err := Vip.SelectUserVip(info.Vip)
		if err != nil {
			return nil, err
		}
		userInfo.Vip = vip
	}

	// 获取等级信息
	grade, err := Grade.SelectUserGrade(info.Grade)
	if err != nil {
		return nil, err
	}
	userInfo.Grade = grade

	// 获取用户投稿数
	var count int
	// 文章数量
	articleCount, err := dao.SysArticle.Where(dao.SysArticle.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, err
	}
	count += articleCount
	//资源数量
	resourceCount, err := dao.SysResource.Where(dao.SysResource.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, err
	}
	count += resourceCount
	//音频数量
	audioCount, err := dao.SysAudio.Where(dao.SysAudio.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, err
	}
	count += audioCount
	//视频数量
	videoCount, err := dao.SysVideo.Where(dao.SysVideo.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, err
	}
	count += videoCount
	//课程数量
	eduCount, err := dao.SysEdu.Where(dao.SysEdu.Columns.UserId, info.UserId).Count()
	if err != nil {
		return nil, err
	}
	count += eduCount
	userInfo.Posts = count

	// 判断是否关注
	userInfo.IsFollow = s.CheckUserFollow(userId, info.UserId)
	//存入缓存
	redisCom.Key = shared.UserInfo + gconv.String(info.UserId)
	redisCom.Time = "180"
	redisCom.Data = userInfo
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// SelectFansOrFollowsList 获取用户关注或粉丝
func (s *userService) SelectFansOrFollowsList(userId int64, req *dto.UserFansOrFollowsQuery) (int, []*result.UserInfo, response.ResponseCode) {

	model := dao.SysUser.SysUserDao.
		Where(dao.SysUser.Columns.DeleteTime, nil)
	model = model.Where(dao.SysUser.Columns.Status, shared.StatusReviewed)
	switch req.Related {
	case "follow":
		var ids []int64
		relateIds, err := dao.SysUserFollow.
			Where(dao.SysUserFollow.Columns.UserId, req.UserId).
			All()
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		for _, i := range relateIds {
			ids = append(ids, i.FollowId)
		}
		model = model.Where(dao.SysUser.Columns.UserId+" IN(?)", ids)
	case "fans":
		var ids []int64
		relateIds, err := dao.SysUserFollow.
			Where(dao.SysUserFollow.Columns.FollowId, req.UserId).
			All()
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		for _, i := range relateIds {
			ids = append(ids, i.UserId)
		}
		model = model.Where(dao.SysUser.Columns.UserId+" IN(?)", ids)
	}

	model = model.Order(dao.SysUser.Columns.UpdateTime + " desc")
	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Fields(
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
	)
	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	var res []*result.UserInfo
	var redisCom redis.Com
	for _, i := range list {
		redisCom.Key = shared.AccountInfo + gconv.String(i.UserId)
		accountObj, err := redisCom.GetString()
		if err != nil {
			return 0, nil, response.CACHE_READ_ERROR
		}

		if accountObj != nil {
			var contentInfo *result.UserInfo
			err := gconv.Struct(accountObj, &contentInfo)
			if err != nil {
				return 0, nil, response.INVALID
			}
			res = append(res, contentInfo)
		} else {
			contentInfo, err := s.info(userId, i)
			if err != nil {
				return 0, nil, response.INVALID
			}
			res = append(res, contentInfo)
		}
	}

	return total, res, response.SUCCESS
}

// SelectSignList 获取用户签到列表
func (s *userService) SelectSignList(req *dto.UserSignQuery) (int, []*result.UserSignList, response.ResponseCode) {

	//查询签到的用户列表
	switch req.Type {
	case shared.UserToday:
		var redisCom redis.Com
		redisCom.Key = shared.UserSignInList
		total, err := redisCom.LLength()
		if err != nil {
			return 0, nil, response.CACHE_READ_ERROR
		}
		if req.Page <= 0 {
			req.Page = 1
		}
		start := (req.Page - 1) * req.Limit
		resUserId, err := redisCom.LRange(start, req.Limit)
		if err != nil {
			return 0, nil, response.CACHE_READ_ERROR
		}
		var res []*result.UserSignList
		for _, i := range resUserId {
			var contentInfo result.UserSignList
			redisCom.Filed = i
			redisCom.Key = shared.UserSignToday
			clockInInfo := redisCom.GetHashFieldString()
			arr := gstr.Split(clockInInfo, "_")
			contentInfo.Integral = gconv.Int64(arr[0])
			contentInfo.CreateTime = arr[1]
			info, err := dao.SysUser.
				Fields(dao.SysUser.Columns.Avatar,
					dao.SysUser.Columns.NickName).
				Where(dao.SysUser.Columns.UserId, gconv.Int64(i)).One()
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			if info != nil {
				contentInfo.NickName = info.NickName
				contentInfo.Avatar = info.Avatar
				contentInfo.UserId = gconv.Int64(i)
				res = append(res, &contentInfo)
			}

		}
		return total, res, response.SUCCESS
	case shared.UserContinuous:
		signModel := dao.SysSign.SysSignDao.Order(dao.SysSign.Columns.Count + " desc")
		total, err := signModel.Count()
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		signModel = signModel.Page(req.Page, req.Limit)
		list, err := signModel.All()
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		var res []*result.UserSignList
		for _, i := range list {
			var contentInfo result.UserSignList
			contentInfo.Count = i.Count
			contentInfo.CreateTime = gconv.String(i.CreateTime)
			info, err := dao.SysUser.
				Fields(dao.SysUser.Columns.Avatar,
					dao.SysUser.Columns.NickName).
				Where(dao.SysUser.Columns.UserId, i.UserId).One()
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			if info != nil {
				contentInfo.NickName = info.NickName
				contentInfo.Avatar = info.Avatar
				contentInfo.UserId = gconv.Int64(i)
				res = append(res, &contentInfo)
			}
		}
		return total, res, response.SUCCESS
	}

	return 0, nil, response.DB_READ_ERROR
}

// SelectRewardList 获取用户的打赏列表
func (s *userService) SelectRewardList(userId int64) (int, []*result.AccountRewardInfo, response.ResponseCode) {
	total, err := dao.SysOrder.
		Where(dao.SysOrder.Columns.OrderType, shared.OrderTypeTwo).
		Where(dao.SysOrder.Columns.AuthorId, userId).
		Fields("DISTINCT "+dao.SysOrder.Columns.UserId).
		Where(dao.SysOrder.Columns.Status, shared.StatusReviewed).Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	list, err := dao.SysOrder.
		Where(dao.SysOrder.Columns.OrderType, shared.OrderTypeTwo).
		Where(dao.SysOrder.Columns.AuthorId, userId).
		Where(dao.SysOrder.Columns.Status, shared.StatusReviewed).
		Fields(dao.SysOrder.Columns.UserId).
		Fields("DISTINCT " + dao.SysOrder.Columns.UserId).
		All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.AccountRewardInfo
	for _, i := range list {
		var contentInfo result.AccountRewardInfo
		avatar, err := dao.SysUser.
			Value(dao.SysUser.Columns.Avatar, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		contentInfo.Avatar = gconv.String(avatar)
		contentInfo.UserId = i.UserId
		res = append(res, &contentInfo)
	}

	return total, res, response.SUCCESS
}

// EditAccountPassWord 修改用户账户密码
func (s *userService) EditAccountPassWord(userId int64, req *dto.PassWordDto) error {

	userObj, err := dao.SysUser.Where(dao.SysUser.Columns.UserId, userId).One()
	if err != nil {
		return err
	}
	// 判断旧密码是否正确
	randPwd := "sd" + req.OldPass + userObj.Salt
	randPwd = gmd5.MustEncryptString(randPwd)
	if !strings.EqualFold(userObj.Password, randPwd) {
		return gerror.New("密码错误")
	}
	// 更新密码
	newSalt := utils.RandS(6)
	newToken := "userObj.UserName" + req.NewPass + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	userObj.Salt = newSalt
	userObj.Password = newToken
	_, err = dao.SysUser.Where(dao.SysUser.Columns.UserId, userId).Update(userObj)
	if err != nil {
		return err
	}
	return nil
}

// EditAccountEmail 修改用户账户邮箱
func (s *userService) EditAccountEmail(userId int64, req *dto.EmailDto) error {

	userObj, err := dao.SysUser.Where(dao.SysUser.Columns.UserId, userId).One()
	if err != nil {
		return err
	}
	//验证 验证码是否正确
	//err = Captcha.CheckEmailCaptcha(userObj.UserName, req.Captcha)
	//if err != nil {
	//	return err
	//}

	userObj.Email = req.Email
	_, err = dao.SysUser.Update(userObj, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------------------------

// CheckEmailUniqueAll 检查用户邮箱是否唯一
func (s *userService) CheckEmailUniqueAll(email string) bool {
	rs, err := dao.SysUser.Where(dao.SysUser.Columns.Email, email).Count()
	if err != nil || rs > 0 {
		return true
	}
	return false
}

// CheckPhoneUniqueAll 检查用户手机号是否唯一
func (s *userService) CheckPhoneUniqueAll(phone string) bool {
	rs, err := dao.SysUser.Where(dao.SysUser.Columns.Phone, phone).Count()
	if err != nil || rs > 0 {
		return true
	}

	return false
}

// CheckUserFavorite 检查用户是否收藏
func (s *userService) CheckUserFavorite(userId, favoriteId int64, module string) bool {
	res, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.Module, module).
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, favoriteId).Count()

	if err != nil || res > 0 {
		return true
	}
	return false
}

// CheckUserLike 检查用户是否点赞
func (s *userService) CheckUserLike(userId, relatedId int64, module string) bool {
	res, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.Module, module).
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, relatedId).Count()

	if err != nil || res > 0 {
		return true
	}
	return false
}

// CheckUserEdu 检查用户是否报名了课程
func (s *userService) CheckUserEdu(userId, relatedId int64) bool {
	res, err := dao.SysUserJoinEdu.
		Where(dao.SysUserJoinEdu.Columns.UserId, userId).
		Where(dao.SysUserJoinEdu.Columns.EduId, relatedId).Count()

	if err != nil || res > 0 {
		return true
	}
	return false
}

// CheckUserFollow 检查用户是关注
func (s *userService) CheckUserFollow(userId, followId int64) bool {
	res, err := dao.SysUserFollow.
		Where(dao.SysUserFollow.Columns.UserId, userId).
		Where(dao.SysUserFollow.Columns.FollowId, followId).Count()
	if err != nil || res > 0 {
		return true
	}
	return false
}

// CheckUserGroup 检查用户是否加入圈子
func (s *userService) CheckUserGroup(userId, relatedId int64) bool {
	res, err := dao.SysUserJoinGroup.
		Where(dao.SysUserJoinGroup.Columns.UserId, userId).
		Where(dao.SysUserJoinGroup.Columns.GroupId, relatedId).Count()
	if err != nil || res > 0 {
		return true
	}
	return false
}
