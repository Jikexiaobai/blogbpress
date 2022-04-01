package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"fiber/app/tools/utils"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var User = new(userService)

type userService struct {
}

// SelectList 查询列表
func (s *userService) SelectList(req *dto.UserQuery) (int, []*result.UserList, response.ResponseCode) {

	model := dao.SysUser.SysUserDao.
		Order(dao.SysUser.Columns.CreateTime + " desc")
	//model.Where(dao.SysArticle.Columns.Status, shared.REVIEWED)
	if req.Status != 0 {
		model = model.Where(dao.SysUser.Columns.Status, req.Status)
	}
	if req.NickName != "" {
		model = model.Where(dao.SysUser.Columns.NickName+" like ?", "%"+req.NickName+"%")
	}

	if req.Email != "" {
		model = model.Where(dao.SysUser.Columns.Email, req.Email)
	}

	if req.Phone != "" {
		model = model.Where(dao.SysUser.Columns.Phone, req.Phone)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysUser.Columns.UserId,
		dao.SysUser.Columns.NickName,
		dao.SysUser.Columns.Integral,
		dao.SysUser.Columns.Balance,
		dao.SysUser.Columns.Avatar,
		dao.SysUser.Columns.LoginIp,
		dao.SysUser.Columns.LoginTime,
		dao.SysUser.Columns.Status,
		dao.SysUser.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.UserList

	for _, i := range list {
		var info *result.UserList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Create 创建
func (s *userService) Create(req *dto.UserCreate) (code response.ResponseCode) {
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

	var entity model.SysUser
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	newSalt := utils.RandS(6)
	newToken := req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	entity.Salt = newSalt
	entity.Password = newToken

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReviewed

	if req.Cover == "" || req.Avatar == "" {
		val, err := System.FindValue(shared.UserSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		j := gjson.New(val)
		if req.Cover == "" {

			defaultCover := gconv.String(j.Get("defaultCover"))
			entity.Cover = defaultCover
		}
		if req.Avatar == "" {

			defaultAvatar := gconv.String(j.Get("defaultAvatar"))
			entity.Avatar = defaultAvatar
		}
	}

	if req.Grade == 0 {
		val, err := System.FindValue(shared.UserSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		j := gjson.New(val)

		// 设置用户等级
		defaultRole := gconv.String(j.Get("defaultGrade"))
		entity.Grade = gconv.Int64(defaultRole)
	}

	res, err := tx.Insert(dao.SysUser.Table, entity)
	if err != nil {
		return response.DB_READ_ERROR
	}
	rid, err := res.LastInsertId()
	if err != nil || rid <= 0 {
		return response.DB_SAVE_ERROR
	}

	// 增加
	if req.Vip != 0 {
		err = Vip.AddRelated(tx, req.Vip, rid)
		if err != nil {
			return response.DB_READ_ERROR
		}
	}

	pathList := make([]string, 0)
	pathList = append(pathList, entity.Avatar)
	pathList = append(pathList, entity.Cover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, "user")
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// EditInfo 编辑信息
func (s *userService) EditInfo(id int64) (*result.UserEditInfo, response.ResponseCode) {
	var editInfo *result.UserEditInfo
	err := dao.SysUser.
		Where(dao.SysUser.Columns.UserId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	return editInfo, response.SUCCESS
}

// Edit 编辑
func (s *userService) Edit(req *dto.UserEdit) (code response.ResponseCode) {

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

	entity := make(map[string]interface{})
	entity[dao.SysUser.Columns.UserId] = req.UserId
	entity[dao.SysUser.Columns.NickName] = req.NickName
	entity[dao.SysUser.Columns.Phone] = req.Phone
	entity[dao.SysUser.Columns.Email] = req.Email
	entity[dao.SysUser.Columns.Avatar] = req.Avatar
	entity[dao.SysUser.Columns.Cover] = req.Cover
	entity[dao.SysUser.Columns.Balance] = req.Balance
	entity[dao.SysUser.Columns.Integral] = req.Integral
	entity[dao.SysUser.Columns.Sex] = req.Sex
	entity[dao.SysUser.Columns.Follows] = req.Follows
	entity[dao.SysUser.Columns.Fans] = req.Fans
	entity[dao.SysUser.Columns.Likes] = req.Likes
	//entity[dao.SysUser.Columns.Admin] = req.Admin
	entity[dao.SysUser.Columns.Grade] = req.Grade
	entity[dao.SysUser.Columns.Vip] = req.Vip
	entity[dao.SysUser.Columns.Description] = req.Description
	entity[dao.SysUser.Columns.UpdateTime] = gtime.Now()

	info, err := dao.SysUser.Fields(
		dao.SysUser.Columns.Password,
		dao.SysUser.Columns.Vip,
	).Where(dao.SysUser.Columns.UserId, req.UserId).One()
	if err != nil {
		return response.DB_READ_ERROR
	}
	if req.Password != gconv.String(info.Password) {
		newSalt := utils.RandS(6)
		newToken := req.Password + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		entity[dao.SysUser.Columns.Salt] = newSalt
		entity[dao.SysUser.Columns.Password] = newToken
	}

	if req.Cover == "" || req.Avatar == "" {
		val, err := System.FindValue(shared.UserSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		j := gjson.New(val)
		if req.Cover == "" {

			defaultCover := gconv.String(j.Get("defaultCover"))
			entity[dao.SysUser.Columns.Cover] = defaultCover
		}
		if req.Avatar == "" {

			defaultAvatar := gconv.String(j.Get("defaultAvatar"))
			entity[dao.SysUser.Columns.Avatar] = defaultAvatar
		}
	}

	if req.Grade == 0 {
		val, err := System.FindValue(shared.UserSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		j := gjson.New(val)

		// 设置用户等级
		defaultRole := gconv.String(j.Get("defaultGrade"))
		entity[dao.SysUser.Columns.Grade] = gconv.Int64(defaultRole)
	}

	_, err = tx.Update(dao.SysUser.Table, entity, dao.SysUser.Columns.UserId, req.UserId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	if info.Vip != req.Vip {
		if req.Vip != 0 {
			err = Vip.RemoveUserVip(tx, []int64{req.UserId})
			if err != nil {
				return response.DB_SAVE_ERROR
			}
			err = Vip.AddRelated(tx, req.Vip, req.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}

		if req.Vip == 0 {
			err = Vip.RemoveUserVip(tx, []int64{req.UserId})
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.UserId}, "user")
	if err != nil {
		return response.DELETE_FAILED
	}
	pathList := make([]string, 0)
	pathList = append(pathList, gconv.String(entity[dao.SysUser.Columns.Avatar]))
	pathList = append(pathList, gconv.String(entity[dao.SysUser.Columns.Cover]))
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.UserId, "user")
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// Review 更新状态
func (s *userService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysUser.Fields(
		dao.SysUser.Columns.UserId,
		dao.SysUser.Columns.NickName,
		dao.SysUser.Columns.Status).
		Where(dao.SysUser.Columns.UserId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	entity := g.Map{
		dao.SysUser.Columns.Status: req.Status,
		dao.SysUser.Columns.Remark: req.Remark,
	}
	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var tmpInfoList []*model.SysUser
	for _, i := range info {
		tmpInfoList = append(tmpInfoList, i)
	}

	_, err = tx.Update(dao.SysUser.Table,
		entity, dao.SysUser.Columns.UserId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.UserId
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemUserReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您的账户已启用"
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您的账户已禁用，原因：" + req.Remark
		}
		noticeList = append(noticeList, notice)
	}
	if len(noticeList) > 0 {
		_, err = tx.Insert(dao.SysNotice.Table, noticeList)
		if err != nil {
			return response.DB_TX_ERROR
		}
	}
	return response.SUCCESS
}

// Remove 删除
func (s *userService) Remove(idList []int64) (code response.ResponseCode) {

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

	_, err = tx.Delete(dao.SysUser.Table,
		dao.SysUser.Columns.UserId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Answer.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Article.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Audio.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Cash.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Comment.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Edu.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Group.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Order.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Question.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Report.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Resource.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Vip.RemoveUserVip(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Topic.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Video.RemoveByUser(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 实名

	err = Media.RemoveRelated(tx, idList, "user")
	if err != nil {

		return response.DELETE_FAILED
	}

	// 删除用户关注
	_, err = tx.Delete(dao.SysUserFollow.Table, dao.SysUserFollow.Columns.UserId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除签到记录
	_, err = tx.Delete(dao.SysSign.Table, dao.SysSign.Columns.UserId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户加入的课程
	_, err = tx.Delete(dao.SysUserJoinEdu.Table, dao.SysUserJoinEdu.Columns.UserId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户加入的圈子
	_, err = tx.Delete(dao.SysUserJoinGroup.Table, dao.SysUserJoinGroup.Columns.UserId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// EditAccountIntegral 修改用户积分
func (s *userService) EditAccountIntegral(tx *gdb.TX, userId, Integral int64) error {
	_, err := tx.Update(dao.SysUser.Table, g.Map{
		dao.SysUser.Columns.Integral: Integral,
	}, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------------------------

func (s *userService) RemoveUserLike(tx *gdb.TX, related []int64, module string) error {
	_, err := tx.Model(dao.SysUserLike.Table).
		Where(dao.SysUserLike.Columns.Module, module).
		Where(dao.SysUserLike.Columns.RelatedId+" IN(?)", related).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) RemoveUserFavorite(tx *gdb.TX, related []int64, module string) error {
	_, err := tx.Model(dao.SysUserFavorite.Table).
		Where(dao.SysUserFavorite.Columns.Module, module).
		Where(dao.SysUserFavorite.Columns.FavoriteId+" IN(?)", related).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

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
