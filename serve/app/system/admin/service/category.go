package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/tools/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Category = new(categoryService)

type categoryService struct {
}

// SelectModuleList 获取模块类型下的分类类型
func (s *categoryService) SelectModuleList(module string) ([]*result.CategoryInfo, response.ResponseCode) {
	// 获取分类
	var cateList []*result.CategoryInfo
	err := dao.SysCategory.
		Where(dao.SysCategory.Columns.Module, module).
		Fields(dao.SysCategory.Columns.CateId,
			dao.SysCategory.Columns.Title,
			dao.SysCategory.Columns.ParentId).Structs(&cateList)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	return cateList, response.SUCCESS
}

// SelectList 查询列表
func (s *categoryService) SelectList(req *dto.CategoryQuery) (int, []*result.CategoryList, response.ResponseCode) {
	model := dao.SysCategory.
		SysCategoryDao.
		Order(dao.SysCategory.Columns.UpdateTime + " desc")

	if req.Title != "" {
		model = model.Where(dao.SysCategory.Columns.Title+" like ?", "%"+req.Title+"%")
	}
	if req.Module != "" {
		model = model.Where(dao.SysCategory.Columns.Module+" IN(?)", req.Module)
	}
	if req.Page != 0 && req.Limit != 0 {
		model = model.Page(req.Page, req.Limit)
	}
	var res []*result.CategoryList
	err := model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return 0, res, response.SUCCESS
}

// Create 创建
func (s *categoryService) Create(req *dto.CategoryCreate) response.ResponseCode {
	// 加入锁限制

	var entity model.SysCategory
	err := gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}

	entity.Keywords = req.Title
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	_, err = dao.SysCategory.Save(entity)
	if err != nil {
		return response.ADD_FAILED
	}
	return response.SUCCESS
}

//EditInfo 编辑信息
func (s *categoryService) EditInfo(id int64) (*result.CategoryEditInfo, response.ResponseCode) {
	var editInfo *result.CategoryEditInfo
	err := dao.SysCategory.
		Where(dao.SysCategory.Columns.CateId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	return editInfo, response.SUCCESS
}

//Edit 编辑
func (s *categoryService) Edit(req *dto.CategoryEdit) response.ResponseCode {

	entity := make(map[string]interface{})
	entity[dao.SysCategory.Columns.Module] = req.Module
	entity[dao.SysCategory.Columns.Title] = req.Title
	entity[dao.SysCategory.Columns.Slug] = req.Slug
	entity[dao.SysCategory.Columns.IsTop] = req.IsTop
	entity[dao.SysCategory.Columns.Cover] = req.Cover
	entity[dao.SysCategory.Columns.ParentId] = req.ParentId
	entity[dao.SysCategory.Columns.Description] = req.Description

	_, err := dao.SysCategory.Update(entity, dao.SysCategory.Columns.CateId, req.CateId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	return response.SUCCESS
}

// Remove 删除
func (s *categoryService) Remove(ids []int64) (code response.ResponseCode) {

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
	_, err = tx.Delete(dao.SysCategory.Table, dao.SysCategory.Columns.CateId+" IN(?)", ids)
	if err != nil {
		return response.DELETE_FAILED
	}

	_, err = tx.Update(dao.SysCategory.Table, g.Map{
		dao.SysCategory.Columns.ParentId: 0,
	}, dao.SysCategory.Columns.ParentId+" IN(?)", ids)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}
