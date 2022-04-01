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

var Authority = new(authorityService)

type authorityService struct {
}

// SelectRolesAuthority 获取角色下的权限列表
func (s *authorityService) SelectRolesAuthority(id []int64) ([]string, error) {
	res, err := dao.SysAuthority.As("a").
		LeftJoin("sys_role_authority b", "b.authority_id = a.authority_id").
		Where("b.role_id IN(?)", id).
		Fields("a.perms").All()
	if err != nil {
		return nil, err
	}
	var perms []string
	for _, i := range res {
		perms = append(perms, i.Perms)
	}
	return perms, nil
}

// SelectMangerAuthority 获取管理下的权限列表
func (s *authorityService) SelectMangerAuthority(id int64) ([]*result.AuthorityList, response.ResponseCode) {
	roleRes, err := dao.SysMangerRole.Where(dao.SysMangerRole.Columns.UserId, id).All()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	var roleIds []int64
	for _, i := range roleRes {
		roleIds = append(roleIds, i.RoleId)
	}

	var list []*result.AuthorityList

	err = dao.SysAuthority.As("a").
		LeftJoin("sys_role_authority b", "b.authority_id = a.authority_id").
		Where("b.role_id IN(?)", roleIds).
		Fields("a.authority_id", "a.redirect",
			"a.title", "a.parent_id",
			"a.component", "a.path",
			"a.perms", "a.target",
			"a.type", "a.hidden",
			"a.icon", "a.create_time",
		).Order("a.order_by desc").Structs(&list)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	return list, response.SUCCESS
}

// SelectAll 获取所有权限列表
func (s *authorityService) SelectAll(req *dto.AuthorityQuery) ([]*result.AuthorityList, response.ResponseCode) {
	model := dao.SysAuthority.SysAuthorityDao.Order(dao.SysAuthority.Columns.AuthorityId, "desc")
	if req.Title != "" {
		model = model.Where(dao.SysAuthority.Columns.Title, req.Title)
	}
	var list []*result.AuthorityList
	err := model.Structs(&list)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	return list, response.SUCCESS
}

// Create 创建权限
func (s *authorityService) Create(req *dto.AuthorityCreate) response.ResponseCode {
	var entity model.SysAuthority
	err := gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysAuthority.Save(entity)
	if err != nil {
		return response.ADD_FAILED
	}
	return response.SUCCESS
}

// EditInfo 编辑信息
func (s *authorityService) EditInfo(id int64) (*result.AuthorityEditInfo, response.ResponseCode) {
	var editInfo *result.AuthorityEditInfo

	err := dao.SysAuthority.
		Where(dao.SysAuthority.Columns.AuthorityId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}

	return editInfo, response.SUCCESS
}

// Edit 编辑
func (s *authorityService) Edit(req *dto.AuthorityEdit) response.ResponseCode {
	entity := make(map[string]interface{})
	entity[dao.SysAuthority.Columns.AuthorityId] = req.AuthorityId
	entity[dao.SysAuthority.Columns.Title] = req.Title
	entity[dao.SysAuthority.Columns.Type] = req.Type
	entity[dao.SysAuthority.Columns.Perms] = req.Perms
	entity[dao.SysAuthority.Columns.ParentId] = req.ParentId
	entity[dao.SysAuthority.Columns.Path] = req.Path
	entity[dao.SysAuthority.Columns.Redirect] = req.Redirect
	entity[dao.SysAuthority.Columns.OrderBy] = req.OrderBy
	entity[dao.SysAuthority.Columns.Component] = req.Component
	entity[dao.SysAuthority.Columns.Icon] = req.Icon
	entity[dao.SysAuthority.Columns.Hidden] = req.Hidden
	entity[dao.SysAuthority.Columns.Target] = req.Target
	entity[dao.SysAuthority.Columns.UpdateTime] = gtime.Now()
	_, err := dao.SysAuthority.
		Where(dao.SysAuthority.Columns.AuthorityId, req.AuthorityId).Update(entity)
	if err != nil {
		return response.UPDATE_FAILED
	}
	return response.SUCCESS
}

// Remove 删除
func (s *authorityService) Remove(id int64) (code response.ResponseCode) {

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
	list, err := dao.SysAuthority.
		Where(dao.SysAuthority.Columns.AuthorityId, id).
		Or(dao.SysAuthority.Columns.ParentId, id).
		Fields(dao.SysAuthority.Columns.AuthorityId).All()
	if err != nil {
		return response.DB_READ_ERROR
	}
	var ids []int64
	for _, i := range list {
		ids = append(ids, i.AuthorityId)
	}

	_, err = tx.Model(dao.SysAuthority.Table).
		Delete(dao.SysAuthority.Columns.AuthorityId+" IN(?)", ids)
	if err != nil {
		return response.DELETE_FAILED
	}
	_, err = tx.Model(dao.SysRoleAuthority.Table).
		Delete(dao.SysRoleAuthority.Columns.AuthorityId+" IN(?)", ids)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// AddRelated 添加关联
func (s *authorityService) AddRelated(tx *gdb.TX, roleId int64, authorityId []int64) error {
	var entityList []model.SysRoleAuthority
	for _, i := range authorityId {
		var entity model.SysRoleAuthority
		entity.RoleId = roleId
		entity.AuthorityId = i
		entityList = append(entityList, entity)
	}

	_, err := tx.Save(dao.SysRoleAuthority.Table, entityList)
	if err != nil {
		return err
	}
	return nil
}

// RemoveRelated 删除关联
func (s *authorityService) RemoveRelated(tx *gdb.TX, roleId []int64) error {
	_, err := tx.Delete(dao.SysRoleAuthority.Table, dao.SysRoleAuthority.Columns.RoleId+" IN(?)", roleId)
	if err != nil {
		return err
	}
	return nil
}
