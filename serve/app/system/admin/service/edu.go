package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/regex"
	"fiber/app/tools/response"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Edu = new(eduService)

type eduService struct {
}

// SelectUserJoinList 查询报名加入的用户
func (s *eduService) SelectUserJoinList(req *dto.EduJoinQuery) (int, []*result.EduUserJoinList, response.ResponseCode) {

	model := dao.SysUserJoinEdu.SysUserJoinEduDao.
		Where(dao.SysUserJoinEdu.Columns.EduId, req.EduId)
	if req.Name != "" {
		model = model.Where(dao.SysUserJoinEdu.Columns.Name, req.Name)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.EduUserJoinList
	for _, i := range list {
		var info *result.EduUserJoinList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)
		res = append(res, info)
	}

	return total, res, response.SUCCESS
}

// SelectList 查询
func (s *eduService) SelectList(req *dto.EduQuery) (int, []*result.EduList, response.ResponseCode) {
	model := dao.SysEdu.SysEduDao.Order(dao.SysEdu.Columns.UpdateTime + " desc")
	if req.Status == 4 {
		model = model.Where(dao.SysEdu.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysEdu.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysEdu.Columns.Status, req.Status)
	}

	if req.Title != "" {
		model = model.Where(dao.SysEdu.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.CateId != 0 {
		model = model.Where(dao.SysEdu.Columns.CateId, req.CateId)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysEdu.Columns.EduId,
		dao.SysEdu.Columns.UserId,
		dao.SysEdu.Columns.CateId,
		dao.SysEdu.Columns.Title,
		dao.SysEdu.Columns.Cover,
		dao.SysEdu.Columns.Joins,
		dao.SysEdu.Columns.Status,
		dao.SysEdu.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	var res []*result.EduList
	for _, i := range list {
		var info *result.EduList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		category, err := dao.SysCategory.Value(dao.SysCategory.Columns.Title, dao.SysCategory.Columns.CateId, i.CateId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.Category = gconv.String(category)

		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
			dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)

		res = append(res, info)
	}

	return total, res, response.SUCCESS
}

// Create 创建
func (s *eduService) Create(req *dto.EduCreate) (code response.ResponseCode) {
	var entity model.SysEdu
	err := gconv.Struct(req, &entity)
	if err != nil {
		return response.FAILD
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReviewed

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

	rs, err := tx.Insert(dao.SysEdu.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, rid, shared.Edu)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 管理媒体库
	pathList, err := regex.GetStringLink(req.Section)
	contentPath, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, contentPath...)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Edu)
		if err != nil {
			return response.ADD_FAILED
		}
	}
	// 写入到es

	return response.SUCCESS
}

// EditInfo 获取信息
func (s *eduService) EditInfo(id int64) (*result.EduEditInfo, response.ResponseCode) {

	var editInfo *result.EduEditInfo

	err := dao.SysEdu.
		Where(dao.SysEdu.Columns.EduId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	// 获取标签
	tagList, err := Tag.SelectRelatedList(id, shared.Edu)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	editInfo.TagList = tagList

	return editInfo, response.SUCCESS

}

// Edit 编辑
func (s *eduService) Edit(req *dto.EduEdit) (code response.ResponseCode) {

	entity := g.Map{
		dao.SysEdu.Columns.Cover:       req.Cover,
		dao.SysEdu.Columns.Title:       req.Title,
		dao.SysEdu.Columns.Max:         req.Max,
		dao.SysEdu.Columns.Content:     req.Content,
		dao.SysEdu.Columns.Section:     req.Section,
		dao.SysEdu.Columns.CateId:      req.CateId,
		dao.SysEdu.Columns.Description: req.Description,
		dao.SysEdu.Columns.Price:       req.Price,
		dao.SysEdu.Columns.UpdateTime:  gtime.Now(),
	}
	entity[dao.SysEdu.Columns.UserId] = req.UserId
	entity[dao.SysEdu.Columns.Hots] = req.Hots
	entity[dao.SysEdu.Columns.Views] = req.Views
	entity[dao.SysEdu.Columns.Favorites] = req.Favorites
	entity[dao.SysEdu.Columns.Likes] = req.Likes
	entity[dao.SysEdu.Columns.Joins] = req.Joins
	entity[dao.SysEdu.Columns.JoinMode] = req.JoinMode
	entity[dao.SysEdu.Columns.Status] = shared.StatusReviewed

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
	_, err = tx.Update(dao.SysEdu.Table, entity, dao.SysEdu.Columns.EduId, req.EduId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, []int64{req.EduId}, shared.Edu)
	if err != nil {

		return response.DB_SAVE_ERROR
	}
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, req.EduId, shared.Edu)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.EduId}, shared.Edu)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	pathList, err := regex.GetStringLink(req.Section)
	contentPath, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, contentPath...)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.EduId, shared.Edu)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}
	return response.SUCCESS
}

// Review 更新状态
func (s *eduService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysEdu.Fields(
		dao.SysEdu.Columns.UserId,
		dao.SysEdu.Columns.EduId,
		dao.SysEdu.Columns.Title,
		dao.SysEdu.Columns.Status).
		Where(dao.SysEdu.Columns.EduId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}
	var entity = g.Map{
		dao.SysEdu.Columns.Status: req.Status,
		dao.SysEdu.Columns.Remark: req.Remark,
	}
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

	var tmpIds []int64
	var tmpInfoList []*model.SysEdu
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.EduId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysEdu.Table,
		entity, dao.SysEdu.Columns.EduId+" IN(?)", tmpIds)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.EduId
		notice.DetailModule = shared.Edu
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的课程《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的课程《" + i.Title + "》未通过审核，原因：" + req.Remark
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

// Recover 软删除
func (s *eduService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysEdu.Fields(
		dao.SysEdu.Columns.UserId,
		dao.SysEdu.Columns.EduId,
		dao.SysEdu.Columns.Title).
		Where(dao.SysEdu.Columns.EduId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysEdu.Table, g.Map{
		dao.SysEdu.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysEdu.Columns.EduId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_READ_ERROR
	}

	var noticeList []model.SysNotice
	for _, info := range list {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.SystemType = shared.NoticeSysTemDeleteContent
		notice.DetailId = info.EduId
		notice.DetailModule = shared.Edu
		notice.Status = shared.NoticeStatusReview
		notice.Content = "您发布的课程《" + info.Title + "》已被删除，原因：" + req.Remark
		notice.Receiver = info.UserId
		notice.CreateTime = gtime.Now()
		noticeList = append(noticeList, notice)
	}
	// 写入通知
	_, err = tx.Insert(dao.SysNotice.Table, noticeList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	return response.SUCCESS
}

// Reduction 还原
func (s *eduService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysEdu.Update(g.Map{
		dao.SysEdu.Columns.DeleteTime: nil,
	}, dao.SysEdu.Columns.EduId+" IN(?)", idList)
	if err != nil {
		return response.UPDATE_FAILED
	}
	return response.SUCCESS
}

// Remove 删除
func (s *eduService) Remove(idList []int64) (code response.ResponseCode) {

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
	_, err = tx.Delete(dao.SysEdu.Table, dao.SysEdu.Columns.EduId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// RemoveByUser 删除
func (s *eduService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysEdu.Fields(
		dao.SysEdu.Columns.UserId,
		dao.SysEdu.Columns.EduId,
		dao.SysEdu.Columns.Title).
		Where(dao.SysEdu.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysEdu.Table, dao.SysEdu.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.EduId)
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Edu)
	if err != nil {
		return err
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Edu)
	if err != nil {
		return err
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Edu)
	if err != nil {
		return err
	}

	err = Media.RemoveRelated(tx, idList, shared.Edu)
	if err != nil {

		return err
	}
	return nil
}
