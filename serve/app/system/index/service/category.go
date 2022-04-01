package service

import (
	"fiber/app/dao"
	"fiber/app/system/index/result"
	"fiber/app/tools/response"
	"github.com/gogf/gf/util/gconv"
)

var Category = new(categoryService)

type categoryService struct {
}

// SelectListByModule 获取模块类型下的分类类型
func (s *categoryService) SelectListByModule(module string) ([]*result.CategoryInfo, response.ResponseCode) {
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

// SelectInfo 查询信息
func (s *categoryService) SelectInfo(id int64, module string) (*result.CategoryInfo, error) {

	info, err := dao.SysCategory.
		Where(dao.SysCategory.Columns.CateId, id).
		Where(dao.SysCategory.Columns.Module, module).
		Where(dao.SysCategory.Columns.DeleteTime, nil).
		One()
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	var res *result.CategoryInfo
	err = gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
