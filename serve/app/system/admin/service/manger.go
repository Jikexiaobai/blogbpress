package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"fiber/app/tools/utils"
	"fiber/library/redis"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Manger = new(mangerService)

type mangerService struct {
}

// SelectInfo 查询信息
func (s *mangerService) SelectInfo(userId int64) (*result.MangerInfo, response.ResponseCode) {
	var redisCom redis.Com
	redisCom.Key = shared.MangerInfo + gconv.String(userId)
	accountObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	if accountObj != nil {
		var accountInfo *result.MangerInfo
		err := gconv.Struct(accountObj, &accountInfo)
		if err != nil {
			return nil, response.FAILD
		}

		return accountInfo, response.SUCCESS
	}

	// 获取登录用户的信息
	var mangerInfo *result.MangerInfo
	user, err := dao.SysManger.
		Fields(dao.SysManger.Columns.UserId,
			dao.SysManger.Columns.NickName,
			dao.SysManger.Columns.Avatar).
		Where(dao.SysManger.Columns.UserId, userId).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	err = gconv.Struct(user, &mangerInfo)
	if err != nil {
		return nil, response.FAILD
	}

	roles, err := Role.SelectMangerRoles(userId)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	var userRole []string
	for _, i := range roles {
		userRole = append(userRole, i.Title)
	}
	mangerInfo.Roles = userRole

	var userRoleId []int64
	for _, i := range roles {
		userRoleId = append(userRoleId, i.RoleId)
	}
	authorityList, err := Authority.SelectRolesAuthority(userRoleId)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	mangerInfo.Authority = authorityList

	//存入缓存
	redisCom.Time = "10"
	redisCom.Data = mangerInfo
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	return mangerInfo, response.SUCCESS
}

// SelectList 查询列表
func (s *mangerService) SelectList(req *dto.MangerQuery) (int, []*result.MangerList, response.ResponseCode) {

	model := dao.SysManger.SysMangerDao.
		Order(dao.SysManger.Columns.CreateTime + " desc")
	//model.Where(dao.SysArticle.Columns.Status, shared.REVIEWED)
	if req.Status != 0 {
		model = model.Where(dao.SysManger.Columns.Status, req.Status)
	}
	if req.NickName != "" {
		model = model.Where(dao.SysManger.Columns.NickName+" like ?", "%"+req.NickName+"%")
	}

	if req.Email != "" {
		model = model.Where(dao.SysManger.Columns.Email, req.Email)
	}

	if req.Phone != "" {
		model = model.Where(dao.SysManger.Columns.Phone, req.Phone)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysManger.Columns.UserId,
		dao.SysManger.Columns.NickName,
		dao.SysManger.Columns.Avatar,
		dao.SysManger.Columns.LoginIp,
		dao.SysManger.Columns.LoginTime,
		dao.SysManger.Columns.Status,
		dao.SysManger.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.MangerList

	for _, i := range list {
		var info *result.MangerList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.FAILD
		}
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Create 创建
func (s *mangerService) Create(req *dto.MangerCreate) (code response.ResponseCode) {
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

	var entity model.SysManger
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}

	newSalt := utils.RandS(6)
	newToken := req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	entity.Salt = newSalt
	entity.Password = newToken

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	if req.Avatar == "" {
		val, err := System.FindValue(shared.UserSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		j := gjson.New(val)

		if req.Avatar == "" {
			defaultAvatar := gconv.String(j.Get("defaultAvatar"))
			entity.Avatar = defaultAvatar
		}
	}

	res, err := tx.Insert(dao.SysManger.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := res.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	//增加 关联权限
	if len(req.RoleId) > 0 {
		err = Role.AddRelated(tx, rid, req.RoleId)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	pathList := make([]string, 0)
	pathList = append(pathList, entity.Avatar)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, "manger")
		if err != nil {
			return response.FILE_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// EditInfo 编辑信息
func (s *mangerService) EditInfo(id int64) (*result.MangerEditInfo, response.ResponseCode) {
	var editInfo *result.MangerEditInfo
	err := dao.SysManger.
		Where(dao.SysManger.Columns.UserId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}

	roles, err := Role.SelectMangerRoles(id)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	var userRoleId []int64
	for _, i := range roles {
		userRoleId = append(userRoleId, i.RoleId)
	}
	editInfo.RoleId = userRoleId
	return editInfo, response.SUCCESS
}

// Edit 编辑
func (s *mangerService) Edit(req *dto.MangerEdit) (code response.ResponseCode) {

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
	entity[dao.SysManger.Columns.UserId] = req.UserId
	entity[dao.SysManger.Columns.NickName] = req.NickName
	entity[dao.SysManger.Columns.Password] = req.Password
	entity[dao.SysManger.Columns.Phone] = req.Phone
	entity[dao.SysManger.Columns.Email] = req.Email
	entity[dao.SysManger.Columns.Avatar] = req.Avatar
	entity[dao.SysManger.Columns.Sex] = req.Sex
	entity[dao.SysManger.Columns.UpdateTime] = gtime.Now()

	info, err := dao.SysManger.Fields(
		dao.SysManger.Columns.Password,
	).Where(dao.SysManger.Columns.UserId, req.UserId).One()
	if err != nil {
		return response.DB_READ_ERROR
	}
	if req.Password != gconv.String(info.Password) {
		newSalt := utils.RandS(6)
		newToken := req.Password + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		entity[dao.SysManger.Columns.Salt] = newSalt
		entity[dao.SysManger.Columns.Password] = newToken
	}

	if req.Avatar == "" {
		val, err := System.FindValue(shared.UserSetting)
		if err != nil {
			return response.DB_READ_ERROR
		}
		j := gjson.New(val)
		if req.Avatar == "" {

			defaultAvatar := gconv.String(j.Get("defaultAvatar"))
			entity[dao.SysUser.Columns.Avatar] = defaultAvatar
		}
	}

	_, err = tx.Update(dao.SysManger.Table, entity, dao.SysManger.Columns.UserId, req.UserId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联权限
	err = Role.RemoveRelated(tx, []int64{req.UserId})
	if err != nil {
		return response.DELETE_FAILED
	}
	if len(req.RoleId) > 0 {
		err = Role.AddRelated(tx, req.UserId, req.RoleId)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.UserId}, "manger")
	if err != nil {
		return response.DELETE_FAILED
	}
	pathList := make([]string, 0)
	pathList = append(pathList, gconv.String(entity[dao.SysManger.Columns.Avatar]))
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.UserId, "manger")
		if err != nil {
			return response.FILE_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// Remove 删除
func (s *mangerService) Remove(idList []int64) (code response.ResponseCode) {
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

	_, err = tx.Delete(dao.SysManger.Table,
		dao.SysManger.Columns.UserId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Role.RemoveRelated(tx, idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, "manger")
	if err != nil {

		return response.DELETE_FAILED
	}

	return response.SUCCESS
}
