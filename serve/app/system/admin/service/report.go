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

var Report = new(reportService)

type reportService struct {
}

// SelectList 查询列表
func (s *reportService) SelectList(req *dto.ReportQuery) (int, []*result.ReportList, response.ResponseCode) {

	model := dao.SysReport.SysReportDao.Order(dao.SysReport.Columns.CreateTime, "desc")

	if req.Status != 0 {
		model = model.Where(dao.SysReport.Columns.Status, req.Status)
	}

	if req.Type != 0 {
		model = model.Where(dao.SysReport.Columns.Type, req.Type)
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
	var res []*result.ReportList
	for _, i := range list {
		var contentInfo *result.ReportList
		err = gconv.Struct(i, &contentInfo)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}

		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		contentInfo.NickName = gconv.String(nickName)

		switch i.Module {
		case shared.Group:
			title, err := dao.SysGroup.Value(dao.SysGroup.Columns.Title,
				dao.SysGroup.Columns.GroupId, i.RelatedId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Title = gconv.String(title)
		case shared.Question:
			title, err := dao.SysQuestion.Value(dao.SysQuestion.Columns.Title,
				dao.SysQuestion.Columns.QuestionId, i.RelatedId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Title = gconv.String(title)
		case shared.Topic:
			title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title,
				dao.SysTopic.Columns.TopicId, i.RelatedId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Title = gconv.String(title)
		case shared.Answer:
			title, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.Content,
				dao.SysAnswer.Columns.AnswerId, i.RelatedId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Title = gconv.String(title)
		case shared.Comment:
			title, err := dao.SysComment.Value(dao.SysComment.Columns.Content,
				dao.SysComment.Columns.CommentId, i.RelatedId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Title = gconv.String(title)
		}

		res = append(res, contentInfo)
	}
	return total, res, response.SUCCESS
}

// Review 更新状态
func (s *reportService) Review(req *dto.Review) (code response.ResponseCode) {
	list, err := dao.SysReport.Fields(
		dao.SysReport.Columns.UserId,
		dao.SysReport.Columns.ReportId,
		dao.SysReport.Columns.RelatedId,
		dao.SysReport.Columns.Module,
		dao.SysReport.Columns.Status).
		Where(dao.SysReport.Columns.ReportId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}
	var dto = g.Map{
		dao.SysReport.Columns.Status: 2,
		dao.SysReport.Columns.Remark: req.Remark,
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
	var tmpInfoList []*model.SysReport
	for _, i := range list {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.ReportId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.DB_READ_ERROR
	}
	_, err = tx.Update(dao.SysReport.Table,
		dto, dao.SysReport.Columns.ReportId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.ReportId
		notice.DetailModule = shared.Report
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		switch i.Module {
		case shared.Topic:
			title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title, dao.SysTopic.Columns.TopicId, i.RelatedId)
			if err != nil {
				return response.DB_READ_ERROR
			}
			notice.Content = "您举报的帖子《" + gconv.String(title) + "》"
		case shared.Question:
			title, err := dao.SysQuestion.Value(dao.SysQuestion.Columns.Title, dao.SysQuestion.Columns.QuestionId, i.RelatedId)
			if err != nil {
				return response.DB_READ_ERROR
			}
			notice.Content = "您举报的问题《" + gconv.String(title) + "》"
		case shared.Comment:
			title, err := dao.SysComment.Value(dao.SysComment.Columns.Content, dao.SysComment.Columns.CommentId, i.RelatedId)
			if err != nil {
				return response.DB_READ_ERROR
			}
			notice.Content = "您举报的评论《" + gconv.String(title) + "》"
		case shared.Answer:
			title, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.Content, dao.SysAnswer.Columns.AnswerId, i.RelatedId)
			if err != nil {
				return response.DB_READ_ERROR
			}
			notice.Content = "您举报的回答《" + gconv.String(title) + "》"
		}
		if req.Status == shared.StatusReviewed {
			notice.Content = notice.Content + "已处理：" + req.Remark
			err = Integral.SetUserReportIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
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

// Remove 删除
func (s *reportService) Remove(idList []int64) response.ResponseCode {
	_, err := dao.SysReport.Delete(dao.SysReport.Columns.ReportId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *reportService) RemoveByUser(tx *gdb.TX, Ids []int64) error {
	_, err := tx.Delete(dao.SysReport.Table, dao.SysReport.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}
	return nil
}
