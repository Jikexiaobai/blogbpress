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

var Answer = new(answerService)

type answerService struct {
}

// SelectList 获取列表
func (s *answerService) SelectList(req *dto.AnswerQueryParam) (int, []*result.AnswerList, response.ResponseCode) {

	model := dao.SysAnswer.SysAnswerDao.Order(dao.SysAnswer.Columns.UpdateTime + " desc")

	if req.RelatedId != 0 {
		model = model.Where(dao.SysAnswer.Columns.TopicId, req.RelatedId)
	}

	if req.IsPay == 2 {
		model = model.Where(dao.SysAnswer.Columns.Price+" >", 0)
	}

	if req.UserId != 0 {
		model = model.Where(dao.SysAnswer.Columns.UserId, req.UserId)
	}

	if req.Status == 4 {
		model = model.Where(dao.SysAnswer.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysAnswer.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysAnswer.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Order(dao.SysComment.Columns.UpdateTime + " desc")
	model = model.Page(req.Page, req.Limit)
	list, err := model.
		Fields(
			dao.SysAnswer.Columns.AnswerId,
			dao.SysAnswer.Columns.UserId,
			dao.SysAnswer.Columns.Content,
			dao.SysAnswer.Columns.Price,
			dao.SysAnswer.Columns.Status,
			dao.SysAnswer.Columns.CreateTime,
		).
		All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.AnswerList
	for _, i := range list {
		var info *result.AnswerList
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

// Recover 软删除
func (s *answerService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysAnswer.Fields(
		dao.SysAnswer.Columns.UserId,
		dao.SysAnswer.Columns.AnswerId,
		dao.SysAnswer.Columns.Content).
		Where(dao.SysAnswer.Columns.AnswerId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysAnswer.Table, g.Map{
		dao.SysAnswer.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysAnswer.Columns.AnswerId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_READ_ERROR
	}

	var noticeList []model.SysNotice
	for _, info := range list {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.SystemType = shared.NoticeSysTemDeleteContent
		notice.DetailId = info.AnswerId
		notice.DetailModule = shared.Answer
		notice.Status = shared.NoticeStatusReview
		notice.Content = "你发布的答案《" + info.Content + "》已被删除，原因：" + req.Remark
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
func (s *answerService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysAnswer.Update(g.Map{
		dao.SysAnswer.Columns.DeleteTime: nil,
	}, dao.SysAnswer.Columns.AnswerId+" IN(?)", idList)
	if err != nil {
		return response.DB_READ_ERROR
	}
	return response.SUCCESS
}

// Remove 删除
func (s *answerService) Remove(idList []int64) (code response.ResponseCode) {
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

	_, err = tx.Delete(dao.SysAnswer.Table,
		dao.SysAnswer.Columns.AnswerId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Answer)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Answer)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// Review 更新状态
func (s *answerService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysAnswer.Fields(
		dao.SysAnswer.Columns.UserId,
		dao.SysAnswer.Columns.AnswerId,
		dao.SysAnswer.Columns.TopicId,
		dao.SysAnswer.Columns.Content,
		dao.SysAnswer.Columns.Status).
		Where(dao.SysAnswer.Columns.AnswerId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
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
	var tmpInfoList []*model.SysAnswer
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.AnswerId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.DB_READ_ERROR
	}
	_, err = tx.Update(dao.SysAnswer.Table,
		g.Map{
			dao.SysAnswer.Columns.Status: req.Status,
			dao.SysAnswer.Columns.Remark: req.Remark,
		}, dao.SysAnswer.Columns.AnswerId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_READ_ERROR
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		// 通知问题作者
		notice.Type = shared.NoticeAnswer
		question, err := dao.SysQuestion.
			Fields(dao.SysQuestion.Columns.UserId, dao.SysQuestion.Columns.Title).
			Where(dao.SysQuestion.Columns.QuestionId, i.TopicId).One()
		if err != nil {
			return response.DB_READ_ERROR
		}

		//获取作者名称
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, question.UserId)
		if err != nil {
			return response.DB_READ_ERROR
		}
		notice.Receiver = question.UserId
		notice.DetailId = i.TopicId
		notice.DetailModule = shared.Question
		notice.Status = shared.NoticeStatusReview
		notice.CreateTime = gtime.Now()

		if req.Status == shared.StatusReviewed {
			notice.Content = gconv.String(nickName) + "回答了你发布的《" + question.Title + "》问题已通过审核" + req.Remark
			notice.SystemType = shared.NoticeSysTemReview
			err = Integral.SetUserAnswerIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = gconv.String(nickName) + "回答了你发布的《" + question.Title + "》问题未通过审核" + req.Remark
			notice.SystemType = shared.NoticeSysTemReview
		}
		noticeList = append(noticeList, notice)
	}

	if len(noticeList) > 0 {
		_, err = tx.Insert(dao.SysNotice.Table, noticeList)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *answerService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysAnswer.Fields(
		dao.SysAnswer.Columns.UserId,
		dao.SysAnswer.Columns.AnswerId,
		dao.SysAnswer.Columns.Content).
		Where(dao.SysAnswer.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysAnswer.Table, dao.SysAnswer.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.AnswerId)
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Answer)
	if err != nil {
		return err
	}

	err = Media.RemoveRelated(tx, idList, shared.Answer)
	if err != nil {

		return err
	}
	return nil
}
