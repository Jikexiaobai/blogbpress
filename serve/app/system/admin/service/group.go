package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Group = new(groupService)

type groupService struct{}

// SelectList 查询圈子列表
func (s *groupService) SelectList(req *dto.GroupQuery) (int, []*result.GroupList, response.ResponseCode) {
	model := dao.SysGroup.SysGroupDao.Order(dao.SysGroup.Columns.UpdateTime + " desc")

	if req.Status == 4 {
		model = model.Where(dao.SysGroup.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysGroup.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysGroup.Columns.Status, req.Status)
	}
	if req.Title != "" {
		model = model.Where(dao.SysGroup.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.CateId != 0 {
		model = model.Where(dao.SysGroup.Columns.CateId, req.CateId)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysGroup.Columns.GroupId,
		dao.SysGroup.Columns.UserId,
		dao.SysGroup.Columns.CateId,
		dao.SysGroup.Columns.Title,
		dao.SysGroup.Columns.Cover,
		dao.SysGroup.Columns.Joins,
		dao.SysGroup.Columns.Contents,
		dao.SysGroup.Columns.Status,
		dao.SysGroup.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.GroupList
	for _, i := range list {
		var info *result.GroupList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		category, err := dao.SysCategory.Value(dao.SysCategory.Columns.Title, dao.SysCategory.Columns.CateId, i.CateId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.Category = gconv.String(category)

		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)

		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Create 创建小组
func (s *groupService) Create(req *dto.GroupCreate) (code response.ResponseCode) {
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

	var entity model.SysGroup
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.ADD_FAILED
	}

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReviewed

	group, err := tx.Insert(dao.SysGroup.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := group.LastInsertId()
	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}
	// 管理媒体库
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Group)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *groupService) EditInfo(id int64) (*result.GroupEditInfo, response.ResponseCode) {
	var editInfo *result.GroupEditInfo

	err := dao.SysGroup.
		Where(dao.SysGroup.Columns.GroupId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}

	return editInfo, response.SUCCESS

}

// Edit 编辑
func (s *groupService) Edit(req *dto.GroupEdit) (code response.ResponseCode) {

	entity := g.Map{
		dao.SysGroup.Columns.Cover:       req.Cover,
		dao.SysGroup.Columns.Title:       req.Title,
		dao.SysGroup.Columns.CateId:      req.CateId,
		dao.SysGroup.Columns.Description: req.Description,
		dao.SysGroup.Columns.Price:       req.Price,
		dao.SysGroup.Columns.JoinMode:    req.JoinMode,
		dao.SysGroup.Columns.SecretKey:   req.SecretKey,
		dao.SysGroup.Columns.Icon:        req.Icon,
		dao.SysGroup.Columns.UpdateTime:  gtime.Now(),
	}
	entity[dao.SysGroup.Columns.Hots] = req.Hots
	entity[dao.SysGroup.Columns.Views] = req.Views
	entity[dao.SysGroup.Columns.Contents] = req.Contents
	entity[dao.SysGroup.Columns.Joins] = req.Joins
	entity[dao.SysGroup.Columns.Status] = shared.StatusReviewed
	entity[dao.SysGroup.Columns.UserId] = req.UserId

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

	_, err = tx.Update(dao.SysGroup.Table, entity, dao.SysGroup.Columns.GroupId, req.GroupId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.GroupId}, shared.Group)
	if err != nil {
		return response.UPDATE_FAILED
	}
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.GroupId, shared.Group)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}
	return response.SUCCESS
}

// Review 更新状态
func (s *groupService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysGroup.Fields(
		dao.SysGroup.Columns.UserId,
		dao.SysGroup.Columns.GroupId,
		dao.SysGroup.Columns.Title,
		dao.SysGroup.Columns.Status).
		Where(dao.SysGroup.Columns.GroupId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}
	var entity = g.Map{
		dao.SysGroup.Columns.Status: req.Status,
		dao.SysGroup.Columns.Remark: req.Remark,
	}
	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_READ_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var tmpIds []int64
	var tmpInfoList []*model.SysGroup
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.GroupId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysGroup.Table,
		entity, dao.SysGroup.Columns.GroupId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.GroupId
		notice.DetailModule = shared.Group
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的圈子《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的圈子《" + i.Title + "》未通过审核，原因：" + req.Remark
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
func (s *groupService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysGroup.Fields(
		dao.SysGroup.Columns.UserId,
		dao.SysGroup.Columns.GroupId,
		dao.SysGroup.Columns.Title).
		Where(dao.SysGroup.Columns.GroupId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysGroup.Table, g.Map{
		dao.SysGroup.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysGroup.Columns.GroupId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_READ_ERROR
	}

	if req.Remark != "" {
		var noticeList []model.SysNotice
		for _, info := range list {
			var notice model.SysNotice
			notice.Type = shared.NoticeSystem
			notice.SystemType = shared.NoticeSysTemDeleteContent
			notice.DetailId = info.GroupId
			notice.DetailModule = shared.Group
			notice.Status = shared.NoticeStatusReview
			notice.Content = "您发布的圈子《" + info.Title + "》已被删除，原因：" + req.Remark
			notice.Receiver = info.UserId
			notice.CreateTime = gtime.Now()
			noticeList = append(noticeList, notice)
		}

		// 写入通知
		_, err = tx.Insert(dao.SysNotice.Table, noticeList)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// Reduction 还原
func (s *groupService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysGroup.Update(g.Map{
		dao.SysGroup.Columns.DeleteTime: nil,
	}, dao.SysGroup.Columns.GroupId+" IN(?)", idList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	return response.SUCCESS
}

// Remove 删除
func (s *groupService) Remove(idList []int64) response.ResponseCode {

	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Delete(dao.SysGroup.Table,
		dao.SysGroup.Columns.GroupId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Group)
	if err != nil {

		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *groupService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysGroup.Fields(
		dao.SysGroup.Columns.UserId,
		dao.SysGroup.Columns.GroupId,
		dao.SysGroup.Columns.Title).Where(dao.SysGroup.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysGroup.Table, dao.SysGroup.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.GroupId)
	}
	err = Media.RemoveRelated(tx, idList, shared.Group)
	if err != nil {
		return err
	}
	return nil
}
