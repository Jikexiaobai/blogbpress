package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Tag = new(tagService)

type tagService struct {
}

// 获取热门标题列表
func (s *tagService) SelectHotTagList() ([]*result.TagList, error) {
	// 获取热门标签
	var tagList []*result.TagList
	err := dao.SysTag.
		Where(dao.SysTag.Columns.Top, 2).
		Fields(dao.SysTag.Columns.TagId, dao.SysTag.Columns.Title).Structs(&tagList)
	if err != nil {
		return nil, err
	}
	return tagList, nil
}

func (s *tagService) SelectRelatedList(rid int64, module string) ([]*result.TagList, error) {
	var listId []int64
	relatedIds, err := dao.SysTagRelated.
		Where(dao.SysTagRelated.Columns.Module, module).
		And(dao.SysTagRelated.Columns.RelatedId, rid).
		All()
	if err != nil {
		return nil, err
	}
	for _, i := range relatedIds {
		listId = append(listId, i.TagId)
	}

	var list []*result.TagList
	err = dao.SysTag.Where(dao.SysTag.Columns.TagId+" IN(?)", listId).Structs(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// AddTags 添加标签
func (s *tagService) AddTags(tx *gdb.TX, tags []string, rid int64, module string) error {

	var tagIdList []int64
	for _, i := range tags {
		res, err := dao.SysTag.Where(dao.SysTag.Columns.Title, i).Count()
		if err != nil {
			return err
		}

		if res < 1 {
			var entity model.SysTag
			entity.Title = i
			entity.CreateTime = gtime.Now()
			rs, err := dao.SysTag.Insert(entity)
			if err != nil {
				return err
			}
			rid, _ := rs.LastInsertId()
			tagIdList = append(tagIdList, rid)
		} else {
			rid, err := dao.SysTag.
				Value(dao.SysTag.Columns.TagId, dao.SysTag.Columns.Title, i)
			if err != nil {
				return err
			}
			tagIdList = append(tagIdList, gconv.Int64(rid))
		}
	}

	tagRelated := make([]model.SysTagRelated, 0)
	for _, i := range tagIdList {
		var tmp model.SysTagRelated
		tmp.Module = module
		tmp.TagId = i
		tmp.RelatedId = rid
		tagRelated = append(tagRelated, tmp)
	}
	if len(tagRelated) > 0 {
		_, err := tx.InsertIgnore(dao.SysTagRelated.Table, tagRelated)
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveRelated 删除标签关联
func (s *tagService) RemoveRelated(tx *gdb.TX, relatedId int64, module string) error {
	_, err := tx.Model(dao.SysTagRelated.Table).
		Where(dao.SysTagRelated.Columns.RelatedId, relatedId).
		And(dao.SysTagRelated.Columns.Module, module).Delete()
	if err != nil {
		return err
	}
	return nil
}

// 查询列表
func (s *tagService) SelectList(req *dto.QueryParam) (int, []*result.TagListInfo, error) {
	model := dao.SysTag.
		SysTagDao.
		Order(dao.SysTag.Columns.CreateTime + " desc")

	if req.Title != "" {
		model = model.Where(dao.SysTag.Columns.Title+" like ?", "%"+req.Title+"%")
	}
	if req.IsTop != 0 {
		model = model.Where(dao.SysTag.Columns.Top+" IN(?)", req.IsTop)
	}
	if req.Page != 0 && req.Limit != 0 {
		model = model.Page(req.Page, req.Limit)
	}
	var res []*result.TagListInfo
	err := model.Structs(&res)
	if err != nil {
		return 0, nil, err
	}

	return 0, res, nil
}

//
func (s *tagService) SetTop(req *dto.TagTop) error {
	_, err := dao.SysTag.Update(g.Map{
		dao.SysTag.Columns.Top: req.IsTop,
	}, dao.SysTag.Columns.TagId+" IN(?)", req.IdList)
	if err != nil {
		return err
	}
	return nil
}

// 删除
func (s *tagService) Remove(ids []int64) error {

	tx, err := g.DB().Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Delete(dao.SysTag.Table, dao.SysTag.Columns.TagId+" IN(?)", ids)
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysTagRelated.Table, dao.SysTagRelated.Columns.TagId+" IN(?)", ids)
	if err != nil {
		return err
	}
	return nil
}
