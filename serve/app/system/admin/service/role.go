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

var Role = new(roleService)

type roleService struct {
}

// SelectMangerRoles 获取管理员所有角色
func (s *roleService) SelectMangerRoles(userId int64) ([]*result.Role, error) {
	var tmp []*result.Role
	err := dao.SysRole.As("a").
		LeftJoin("sys_manger_role b", "b.role_id = a.role_id").
		Fields("a.title", "a.role_id").
		Where("b.user_id", userId).Structs(&tmp)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// Create 创建
func (s *roleService) Create(req *dto.RoleCreate) (code response.ResponseCode) {
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
	var entity model.SysRole
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	res, err := tx.Save(dao.SysRole.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := res.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	//增加 关联权限
	if len(req.Authority) > 0 {
		err = Authority.AddRelated(tx, rid, req.Authority)
		if err != nil {
			return response.ADD_FAILED
		}
	}
	return response.SUCCESS
}

//EditInfo 编辑信息
func (s *roleService) EditInfo(id int64) (*result.RoleEditInfo, response.ResponseCode) {
	var editInfo *result.RoleEditInfo
	err := dao.SysRole.
		Where(dao.SysRole.Columns.RoleId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	// 获取权限
	list, err := dao.SysRoleAuthority.
		Fields(dao.SysRoleAuthority.Columns.AuthorityId).
		Where(dao.SysRoleAuthority.Columns.RoleId, id).All()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	var authorityIds []int64
	for _, i := range list {
		authorityIds = append(authorityIds, i.AuthorityId)
	}
	editInfo.Authority = authorityIds
	return editInfo, response.SUCCESS
}

//Edit 编辑
func (s *roleService) Edit(req *dto.RoleEdit) (code response.ResponseCode) {

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
	entity[dao.SysRole.Columns.RoleId] = req.RoleId
	entity[dao.SysRole.Columns.Title] = req.Title
	entity[dao.SysRole.Columns.Status] = req.Status
	entity[dao.SysRole.Columns.UpdateTime] = gtime.Now()
	_, err = tx.Update(dao.SysRole.Table, entity, dao.SysRole.Columns.RoleId, req.RoleId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联权限
	err = Authority.RemoveRelated(tx, []int64{req.RoleId})
	if err != nil {
		return response.DELETE_FAILED
	}
	if len(req.Authority) > 0 {
		err = Authority.AddRelated(tx, req.RoleId, req.Authority)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	return response.SUCCESS
}

// SelectList 等级列表
func (s *roleService) SelectList() ([]*result.RoleList, response.ResponseCode) {
	var list []*result.RoleList
	err := dao.SysRole.Fields(
		dao.SysRole.Columns.RoleId,
		dao.SysRole.Columns.Title,
		dao.SysRole.Columns.Status,
		dao.SysRole.Columns.CreateTime,
	).Structs(&list)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	return list, response.SUCCESS
}

// AddRelated 添加关联
func (s *roleService) AddRelated(tx *gdb.TX, userId int64, ids []int64) error {
	var entityList []model.SysMangerRole
	for _, i := range ids {
		var entity model.SysMangerRole
		entity.UserId = userId
		entity.RoleId = i
		entityList = append(entityList, entity)
	}

	_, err := tx.Save(dao.SysMangerRole.Table, entityList)
	if err != nil {
		return err
	}
	return nil
}

// RemoveRelated 删除关联
func (s *roleService) RemoveRelated(tx *gdb.TX, ids []int64) error {
	_, err := tx.Delete(dao.SysMangerRole.Table, dao.SysMangerRole.Columns.UserId+" IN(?)", ids)
	if err != nil {
		return err
	}
	return nil
}
