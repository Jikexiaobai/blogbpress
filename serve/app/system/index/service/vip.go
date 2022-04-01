package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/result"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Vip = new(vipService)

type vipService struct {
}

// SelectUserVip 获取用户的vip
func (s *vipService) SelectUserVip(vipId int64) (*result.UserVip, error) {
	info, err := dao.SysVip.Fields(
		dao.SysVip.Columns.Title,
		dao.SysVip.Columns.Color,
		dao.SysVip.Columns.Icon,
	).Where(dao.SysVip.Columns.VipId, vipId).One()
	if err != nil {
		return nil, err
	}
	var vip result.UserVip
	err = gconv.Struct(info, &vip)
	if err != nil {
		return nil, err
	}
	return &vip, nil
}

// SelectAccountVip 获取账户和vip
func (s *vipService) SelectAccountVip(userId, vipId int64) (*result.AccountVip, error) {
	info, err := dao.SysVip.Fields(
		dao.SysVip.Columns.Title,
		dao.SysVip.Columns.Color,
		dao.SysVip.Columns.Icon,
		dao.SysVip.Columns.Discount,
	).Where(dao.SysVip.Columns.VipId, vipId).One()

	if err != nil {
		return nil, err
	}
	var vip result.AccountVip
	err = gconv.Struct(info, &vip)
	if err != nil {
		return nil, err
	}
	userVip, err := dao.SysUserVip.Fields(dao.SysUserVip.Columns.FinishTime).
		Where(dao.SysUserVip.Columns.UserId, userId).
		Where(dao.SysUserVip.Columns.VipId, vipId).One()
	if err != nil {
		return nil, err
	}
	vip.FinishTime = userVip.FinishTime
	return &vip, nil
}

// SelectList vip列表
func (s *vipService) SelectList() ([]*result.Vip, error) {
	var list []*result.Vip
	err := dao.SysVip.Structs(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// SelectByHomeList 查询首页列表
func (s *vipService) SelectByHomeList(ids string) ([]*result.Vip, error) {
	idList := gstr.Split(ids, ",")
	var res []*result.Vip
	err := dao.SysVip.Where(dao.SysVip.Columns.VipId+" IN(?)", idList).Structs(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SelectVipDiscount 获取会员折扣
func (s *vipService) SelectVipDiscount(userId int64) (float64, error) {
	vip, err := dao.SysUser.Value(dao.SysUser.Columns.Vip, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return 0, err
	}

	discount, err := dao.SysVip.
		Value(dao.SysVip.Columns.Discount, dao.SysVip.Columns.VipId, gconv.Int64(vip))
	if err != nil {
		return 0, err
	}
	return gconv.Float64(discount), nil
}

// RemoveUserVip 删除用户开通记录
func (s *vipService) RemoveUserVip(id, userId int64) error {
	_, err := dao.SysUserVip.
		Where(dao.SysUserVip.Columns.UserId, userId).
		And(dao.SysUserVip.Columns.VipId, id).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

// OpenVip 设置用户会员开通记录
func (s *vipService) OpenVip(tx *gdb.TX, id, userId int64) error {
	// 获取会员日期
	role, err := dao.SysVip.Fields(
		dao.SysVip.Columns.Day,
		dao.SysVip.Columns.VipId,
	).Where(dao.SysVip.Columns.VipId, id).One()
	if err != nil {
		return err
	}

	// 检擦 用户是否为 Vip
	if s.CheckIsVip(userId) {
		// 获取用户结束日期
		userVip, err := dao.SysUserVip.
			Fields(dao.SysUserVip.Columns.FinishTime).
			Where(dao.SysUserVip.Columns.UserId, userId).One()
		if err != nil {
			return err
		}
		var newFinishTime = userVip.FinishTime.AddDate(0, 0, role.Day)
		if role.Day == 0 {
			newFinishTime = nil
		}
		_, err = tx.Update(dao.SysUserVip.Table, g.Map{
			dao.SysUserVip.Columns.FinishTime: newFinishTime,
			dao.SysUserVip.Columns.StartTime:  gtime.Now(),
			dao.SysUserVip.Columns.VipId:      id,
		}, dao.SysUserVip.Columns.UserId, userId)
		if err != nil {
			return err
		}

		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Vip: id,
		}, dao.SysUserVip.Columns.UserId, userId)
		if err != nil {
			return err
		}
		return nil
	}

	var entity model.SysUserVip
	entity.VipId = id
	entity.UserId = userId
	entity.StartTime = gtime.Now()
	if role.Day != 0 {
		entity.FinishTime = gtime.Now().AddDate(0, 0, role.Day)
	}

	_, err = tx.Save(dao.SysUserVip.Table, entity)
	if err != nil {
		return err
	}
	_, err = tx.Update(dao.SysUser.Table, g.Map{
		dao.SysUser.Columns.Vip: id,
	}, dao.SysUserVip.Columns.UserId, userId)
	if err != nil {
		return err
	}
	return nil
}

// CheckIsVip 检查是否为会员
func (s *vipService) CheckIsVip(userId int64) bool {
	user, err := dao.SysUser.
		Value(dao.SysUser.Columns.Vip, dao.SysUser.Columns.UserId, userId)
	if err != nil {
		return true
	}
	if gconv.Int64(user) != 0 {
		return true
	}
	return false
}
