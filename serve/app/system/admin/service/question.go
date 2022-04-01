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

var Question = new(questionService)

type questionService struct {
}

// SelectList 获取列表
func (s *questionService) SelectList(req *dto.QuestionQuery) (int, []*result.QuestionList, response.ResponseCode) {

	model := dao.SysQuestion.SysQuestionDao.Order(dao.SysQuestion.Columns.UpdateTime + " desc")
	if req.Status == 4 {
		model = model.Where(dao.SysQuestion.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysQuestion.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysQuestion.Columns.Status, req.Status)
	}
	if req.Title != "" {
		model = model.Where(dao.SysQuestion.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysQuestion.Columns.UserId,
		dao.SysQuestion.Columns.QuestionId,
		dao.SysQuestion.Columns.Title,
		dao.SysQuestion.Columns.Status,
		dao.SysQuestion.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.QuestionList
	for _, i := range list {
		var info *result.QuestionList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.FAILD
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

// Review 更新状态
func (s *questionService) Review(req *dto.Review) (code response.ResponseCode) {
	list, err := dao.SysQuestion.Fields(
		dao.SysQuestion.Columns.UserId,
		dao.SysQuestion.Columns.Title,
		dao.SysQuestion.Columns.QuestionId,
		dao.SysQuestion.Columns.Status).
		Where(dao.SysQuestion.Columns.QuestionId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
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
	var tmpInfoList []*model.SysQuestion
	for _, i := range list {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.QuestionId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysQuestion.Table,
		g.Map{
			dao.SysQuestion.Columns.Status: req.Status,
			dao.SysQuestion.Columns.Remark: req.Remark,
		}, dao.SysQuestion.Columns.QuestionId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.QuestionId
		notice.DetailModule = shared.Question
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的问题《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的问题《" + i.Title + "》未通过审核，原因：" + req.Remark
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
func (s *questionService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysQuestion.Fields(
		dao.SysQuestion.Columns.UserId,
		dao.SysQuestion.Columns.QuestionId,
		dao.SysQuestion.Columns.Title).
		Where(dao.SysQuestion.Columns.QuestionId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysQuestion.Table, g.Map{
		dao.SysQuestion.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysQuestion.Columns.QuestionId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	if req.Remark != "" {
		var noticeList []model.SysNotice
		for _, info := range list {
			var notice model.SysNotice
			notice.Type = shared.NoticeSystem
			notice.SystemType = shared.NoticeSysTemDeleteContent
			notice.DetailId = info.QuestionId
			notice.DetailModule = shared.Question
			notice.Status = shared.NoticeStatusReview
			notice.Content = "您发布的问题《" + info.Title + "》已被删除，原因：" + req.Remark
			notice.Receiver = info.UserId
			notice.CreateTime = gtime.Now()
			noticeList = append(noticeList, notice)
		}

		// 写入通知
		_, err = dao.SysNotice.Insert(noticeList)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// Reduction 还原
func (s *questionService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysQuestion.Update(g.Map{
		dao.SysQuestion.Columns.DeleteTime: nil,
	}, dao.SysQuestion.Columns.QuestionId+" IN(?)", idList)
	if err != nil {
		return response.DB_READ_ERROR
	}
	return response.SUCCESS
}

// Remove 删除
func (s *questionService) Remove(idList []int64) response.ResponseCode {
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

	_, err = tx.Delete(dao.SysQuestion.Table,
		dao.SysQuestion.Columns.QuestionId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联圈子

	err = Media.RemoveRelated(tx, idList, shared.Question)
	if err != nil {

		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *questionService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysQuestion.Fields(
		dao.SysQuestion.Columns.UserId,
		dao.SysQuestion.Columns.QuestionId,
		dao.SysQuestion.Columns.Title).Where(dao.SysQuestion.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysQuestion.Table, dao.SysQuestion.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.QuestionId)
	}

	err = Media.RemoveRelated(tx, idList, shared.Question)
	if err != nil {

		return err
	}

	return nil
}
