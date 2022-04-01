package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/tools/response"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Vip = new(vipService)

type vipService struct {
}

// SelectList vip列表
func (s *vipService) SelectList() ([]*result.Vip, response.ResponseCode) {
	var list []*result.Vip
	err := dao.SysVip.Structs(&list)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	return list, response.SUCCESS
}

// Create 创建
func (s *vipService) Create(req *dto.VipCreate) (code response.ResponseCode) {
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
	var entity model.SysVip
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}

	res, err := tx.Insert(dao.SysVip.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := res.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 媒体库
	var pathList []string
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, "vip")
		if err != nil {
			return response.ADD_FAILED
		}
	}
	return response.SUCCESS
}

// EditInfo 编辑信息
func (s *vipService) EditInfo(id int64) (*result.Vip, response.ResponseCode) {

	info, err := dao.SysVip.
		Where(dao.SysVip.Columns.VipId, id).One()
	if info == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	var res *result.Vip
	err = gconv.Struct(info, &res)
	if err != nil {
		return nil, response.FAILD
	}

	// 获取权限
	return res, response.SUCCESS
}

// Edit 编辑
func (s *vipService) Edit(req *dto.VipEdit) (code response.ResponseCode) {

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
	var entity model.SysVip
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}

	_, err = tx.Update(dao.SysVip.Table, entity, dao.SysVip.Columns.VipId, req.VipId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.VipId}, "vip")
	if err != nil {

		return response.UPDATE_FAILED
	}
	var pathList []string
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.VipId, "vip")
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	return response.SUCCESS
}

// AddRelated 设置关联
func (s *vipService) AddRelated(tx *gdb.TX, id int64, userId int64) error {
	day, err := dao.SysVip.Value(dao.SysVip.Columns.Day, dao.SysVip.Columns.VipId, id)
	if err != nil {
		return err
	}
	var tmp model.SysUserVip
	tmp.VipId = id
	tmp.UserId = userId
	tmp.StartTime = gtime.Now()
	tmp.FinishTime = gtime.Now().AddDate(0, 0, gconv.Int(day))

	_, err = tx.Save(dao.SysUserVip.Table, tmp)
	if err != nil {
		return err
	}

	return nil
}

func (s *vipService) RemoveUserVip(tx *gdb.TX, Ids []int64) error {
	_, err := tx.Delete(dao.SysUserVip.Table, dao.SysUserVip.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}
	return nil
}
