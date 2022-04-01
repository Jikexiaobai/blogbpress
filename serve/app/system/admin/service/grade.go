package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/tools/response"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Grade = new(gradeService)

type gradeService struct {
}

// SelectList 等级列表
func (s *gradeService) SelectList() ([]*result.Grade, response.ResponseCode) {
	list, err := dao.SysGrade.All()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	var res []*result.Grade
	for _, i := range list {
		var info *result.Grade
		err = gconv.Struct(i, &info)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		postsModule := gjson.New(i.PostsModule)
		postsModuleList := gconv.Strings(postsModule.Value())
		if len(postsModuleList) > 0 {
			info.PostsModule = postsModuleList
		}

		commonAuth := gjson.New(i.CommonAuth)
		commonAuthList := gconv.Strings(commonAuth.Value())
		if len(commonAuthList) > 0 {
			info.CommonAuth = commonAuthList
		}
		res = append(res, info)
	}
	return res, response.SUCCESS
}

// Create 创建
func (s *gradeService) Create(req *dto.GradeCreate) (code response.ResponseCode) {
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
	var entity model.SysGrade
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.ADD_FAILED
	}

	res, err := tx.Insert(dao.SysGrade.Table, entity)
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
		err = Media.AddRelated(tx, pathList, rid, "grade")
		if err != nil {
			return response.ADD_FAILED
		}
	}
	return response.SUCCESS
}

// EditInfo 编辑信息
func (s *gradeService) EditInfo(id int64) (*result.Grade, response.ResponseCode) {

	info, err := dao.SysGrade.
		Where(dao.SysGrade.Columns.GradeId, id).One()
	if info == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	var res *result.Grade
	err = gconv.Struct(info, &res)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	postsModule := gjson.New(info.PostsModule)
	postsModuleList := gconv.Strings(postsModule.Value())
	if len(postsModuleList) > 0 {
		res.PostsModule = postsModuleList
	}

	commonAuth := gjson.New(info.CommonAuth)
	commonAuthList := gconv.Strings(commonAuth.Value())
	if len(commonAuthList) > 0 {
		res.CommonAuth = commonAuthList
	}

	// 获取权限
	return res, response.SUCCESS
}

// Edit 编辑
func (s *gradeService) Edit(req *dto.GradeEdit) (code response.ResponseCode) {

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
	var entity model.SysGrade
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.UPDATE_FAILED
	}

	_, err = tx.Update(dao.SysGrade.Table, entity, dao.SysGrade.Columns.GradeId, req.GradeId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.GradeId}, "grade")
	if err != nil {

		return response.UPDATE_FAILED
	}
	var pathList []string
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.GradeId, "grade")
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	return response.SUCCESS
}

func (s *gradeService) Remove(ids []int64) error {

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
	_, err = tx.Delete(dao.SysRole.Table, dao.SysRole.Columns.RoleId+" IN(?)", ids)
	if err != nil {
		return err
	}
	// 增加 关联权限
	//err = Authority.RemoveRelated(ids)
	//if err != nil {
	//
	//	return err
	//}
	// 删除关联的权限
	return nil
}
