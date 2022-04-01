package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Report = new(reportService)

type reportService struct {
}

// 查询列表
func (s *reportService) SelectList(req *dto.QueryParam) (int, []*result.ReportListInfo, error) {

	model := dao.SysReport.SysReportDao.Order(dao.SysAudio.Columns.CreateTime, "desc")

	if req.Type != 0 {
		model = model.Where(dao.SysReport.Columns.Type, req.Type)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, err
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, nil
	}
	var res []*result.ReportListInfo
	for _, i := range list {
		var contentInfo *result.ReportListInfo
		err = gconv.Struct(i, &contentInfo)
		if err != nil {
			return 0, nil, err
		}

		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.NickName = gconv.String(nickName)

		switch i.Module {
		case shared.Article:
			title, err := dao.SysArticle.Value(dao.SysArticle.Columns.Title,
				dao.SysArticle.Columns.ArticleId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Audio:
			title, err := dao.SysAudio.Value(dao.SysAudio.Columns.Title,
				dao.SysAudio.Columns.AudioId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Video:
			title, err := dao.SysVideo.Value(dao.SysVideo.Columns.Title,
				dao.SysVideo.Columns.VideoId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Resource:
			title, err := dao.SysResource.Value(dao.SysResource.Columns.Title,
				dao.SysResource.Columns.ResourceId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Edu:
			title, err := dao.SysEdu.Value(dao.SysEdu.Columns.Title,
				dao.SysEdu.Columns.EduId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Group:
			title, err := dao.SysGroup.Value(dao.SysGroup.Columns.Title,
				dao.SysGroup.Columns.GroupId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Question:
			title, err := dao.SysQuestion.Value(dao.SysQuestion.Columns.Title,
				dao.SysQuestion.Columns.QuestionId, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Topic:
			title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title,
				dao.SysTopic.Columns.Title, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Answer:
			title, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.Content,
				dao.SysAnswer.Columns.Content, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		case shared.Comment:
			title, err := dao.SysComment.Value(dao.SysComment.Columns.Content,
				dao.SysComment.Columns.Content, i.RelatedId)
			if err != nil {
				return 0, nil, err
			}
			contentInfo.Title = gconv.String(title)
		}

		res = append(res, contentInfo)
	}
	return total, res, nil
}

// 创建
func (s *reportService) Create(userId int64, req *dto.ReportCreate) error {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ReportCreateCount+gconv.String(userId),
		shared.ReportCreateLock+gconv.String(userId), 60, 5)
	if err != nil {
		return nil
	}

	res, err := dao.SysReport.
		Where(dao.SysReport.Columns.RelatedId, req.RelatedId).
		Where(dao.SysReport.Columns.Module, req.Module).
		Where(dao.SysReport.Columns.UserId, userId).
		One()
	if err != nil {
		return err
	}
	if res != nil {
		return gerror.New("你已经对该内容举报过了")
	}
	var entity model.SysReport
	err = gconv.Struct(req, &entity)
	if err != nil {
		return err
	}
	entity.UserId = userId
	entity.Status = 2
	entity.UpdateTime = gtime.Now()
	entity.CreateTime = gtime.Now()
	_, err = dao.SysReport.Insert(entity)
	if err != nil {
		return err
	}

	return nil
}
