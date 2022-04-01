package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Notice = new(noticeService)

type noticeService struct {
}

// SelectCount 查询个数
func (s *noticeService) SelectCount(userId int64) (*result.NoticeCount, response.ResponseCode) {
	list, err := dao.SysNotice.
		Fields(dao.SysNotice.Columns.Type).
		Where(dao.SysNotice.Columns.Receiver, userId).
		Where(dao.SysNotice.Columns.Status, shared.StatusReview).All()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	var res result.NoticeCount
	for _, i := range list {
		if i.Type == shared.NoticeSystem {
			res.System += 1
		}
		//if i.Type == shared.NoticeFinance {
		//	res.Finance += 1
		//}
		if i.Type == shared.NoticeComment {
			res.Comment += 1
		}
		if i.Type == shared.NoticeAnswer {
			res.Answer += 1
		}
		if i.Type == shared.NoticeLike {
			res.Like += 1
		}
		if i.Type == shared.NoticeFollow {
			res.Follow += 1
		}
	}

	return &res, response.SUCCESS
}

// SelectList 查询列表
func (s *noticeService) SelectList(req *dto.QueryParam) (int, []*result.NoticeInfo, response.ResponseCode) {
	model := dao.SysNotice.SysNoticeDao.
		Where(dao.SysNotice.Columns.Receiver, req.UserId)
	if req.Type != 0 {
		model = model.Where(dao.SysNotice.Columns.Type, req.Type)
	}

	model = model.Order(dao.SysNotice.Columns.CreateTime, "desc")
	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)

	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.NoticeInfo
	for _, i := range list {
		var noticeInfo *result.NoticeInfo
		err = gconv.Struct(i, &noticeInfo)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}

		detailInfo, err := s.info(i)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		noticeInfo.DetailInfo = detailInfo

		if i.Status != 2 {
			_, err = dao.SysNotice.Update(g.Map{
				dao.SysNotice.Columns.Status: 2,
			}, dao.SysNotice.Columns.NoticeId, i.NoticeId)
		}
		res = append(res, noticeInfo)
	}
	return total, res, response.SUCCESS
}

// info 获取获取详细信息
func (s *noticeService) info(info *model.SysNotice) (interface{}, error) {
	switch info.Type {
	case shared.NoticeSystem:
		switch info.SystemType {
		case shared.NoticeSysTemDeleteContent:
			detailObj, err := s.moduleInfo(info.DetailModule, info.DetailId)
			if err != nil {
				return nil, err
			}
			return detailObj, nil
		case shared.NoticeUserTips:
			detail := make(map[string]interface{})
			order, err := dao.SysOrder.Fields(dao.SysOrder.Columns.UserId, dao.SysOrder.Columns.AuthorMoney).
				Where(dao.SysOrder.Columns.OrderId, info.DetailId).One()
			if err != nil {
				return nil, err
			}
			if order == nil {
				return nil, nil
			}
			//	获取打赏用户信息
			tipsUserInfo := make(map[string]interface{})
			userObj, err := dao.SysUser.
				Fields(dao.SysUser.Columns.NickName, dao.SysUser.Columns.Cover).
				Where(dao.SysUser.Columns.UserId, order.UserId).One()
			if err != nil {
				return nil, err
			}
			tipsUserInfo["nickName"] = userObj.NickName
			tipsUserInfo["cover"] = userObj.Cover
			tipsUserInfo["id"] = order.UserId
			detail["tipsUserInfo"] = tipsUserInfo
			detail["money"] = order.AuthorMoney
			return detail, nil
		case shared.NoticeUserBuyContent:
			detail := make(map[string]interface{})
			order, err := dao.SysOrder.Fields(dao.SysOrder.Columns.UserId,
				dao.SysOrder.Columns.AuthorMoney, dao.SysOrder.Columns.DetailId, dao.SysOrder.Columns.DetailModule).
				Where(dao.SysOrder.Columns.OrderId, info.DetailId).One()
			if err != nil {
				return nil, err
			}
			if order == nil {
				return nil, nil
			}
			detailObj, err := s.moduleInfo(order.DetailModule, order.DetailId)
			if err != nil {
				return nil, err
			}

			detail["detailInfo"] = detailObj
			//	获取购买用户信息
			userInfo := make(map[string]interface{})
			userObj, err := dao.SysUser.
				Fields(dao.SysUser.Columns.NickName, dao.SysUser.Columns.Cover).
				Where(dao.SysUser.Columns.UserId, order.UserId).One()
			if err != nil {
				return nil, err
			}
			userInfo["nickName"] = userObj.NickName
			userInfo["cover"] = userObj.Cover
			userInfo["id"] = order.UserId
			detail["userInfo"] = userInfo
			detail["money"] = order.AuthorMoney
			return detail, nil
		}
	case shared.NoticeComment:
		detail := make(map[string]interface{})
		//	获取评论用户信息
		commentInfo, err := dao.SysComment.
			Fields(dao.SysComment.Columns.Module, dao.SysComment.Columns.RelatedId, dao.SysComment.Columns.UserId).
			Where(dao.SysComment.Columns.CommentId, info.DetailId).
			One()
		if err != nil {
			return nil, err
		}
		//	获取点赞内容
		detailObj, err := s.moduleInfo(commentInfo.Module, commentInfo.RelatedId)
		if err != nil {
			return nil, err
		}
		detail["detail"] = detailObj

		userInfo, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Cover, dao.SysUser.Columns.NickName).
			Where(dao.SysUser.Columns.UserId, commentInfo.UserId).
			One()
		if err != nil {
			return nil, err
		}
		fromUser := make(map[string]interface{})
		fromUser["nickName"] = userInfo.NickName
		fromUser["cover"] = userInfo.Cover
		fromUser["id"] = commentInfo.UserId
		detail["fromUser"] = fromUser

		return detail, nil
	case shared.NoticeAnswer:
		detail := make(map[string]interface{})
		//	获取评论用户信息
		answerInfo, err := dao.SysAnswer.
			Fields(dao.SysAnswer.Columns.TopicId, dao.SysAnswer.Columns.UserId).
			Where(dao.SysAnswer.Columns.AnswerId, info.DetailId).
			One()
		if err != nil {
			return nil, err
		}
		//	获取内容
		detailObj, err := s.moduleInfo(shared.Question, answerInfo.TopicId)
		if err != nil {
			return nil, err
		}
		detail["detail"] = detailObj

		userInfo, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Cover, dao.SysUser.Columns.NickName).
			Where(dao.SysUser.Columns.UserId, answerInfo.UserId).
			One()
		if err != nil {
			return nil, err
		}
		fromUser := make(map[string]interface{})
		fromUser["nickName"] = userInfo.NickName
		fromUser["cover"] = userInfo.Cover
		fromUser["id"] = answerInfo.UserId
		detail["fromUser"] = fromUser

		return detail, nil
	case shared.NoticeLike:
		detail := make(map[string]interface{})
		//	获取点赞用户信息
		userInfo, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Cover, dao.SysUser.Columns.NickName).
			Where(dao.SysUser.Columns.UserId, info.FromUserId).
			One()
		if err != nil {
			return nil, err
		}
		fromUser := make(map[string]interface{})
		fromUser["nickName"] = userInfo.NickName
		fromUser["cover"] = userInfo.Cover
		fromUser["id"] = info.FromUserId
		detail["fromUser"] = fromUser

		//	获取点赞内容
		detailObj, err := s.moduleInfo(info.DetailModule, info.DetailId)
		if err != nil {
			return nil, err
		}
		detail["detail"] = detailObj

		return detail, nil
	case shared.NoticeFollow:
		detail := make(map[string]interface{})
		//	获取点赞用户信息
		userInfo, err := dao.SysUser.
			Fields(dao.SysUser.Columns.Cover, dao.SysUser.Columns.NickName).
			Where(dao.SysUser.Columns.UserId, info.FromUserId).
			One()
		if err != nil {
			return nil, err
		}
		detail["nickName"] = userInfo.NickName
		detail["cover"] = userInfo.Cover
		detail["id"] = info.FromUserId
		return detail, nil
	}
	return nil, nil
}

// moduleInfo 获取模块内容
func (s *noticeService) moduleInfo(module string, id int64) (interface{}, error) {

	switch module {
	case shared.Article:
		detailObj, err := dao.SysArticle.
			Fields(dao.SysArticle.Columns.Title, dao.SysArticle.Columns.Cover).
			Where(dao.SysArticle.Columns.ArticleId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Audio:
		detailObj, err := dao.SysAudio.
			Fields(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.Cover).
			Where(dao.SysAudio.Columns.AudioId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Resource:
		detailObj, err := dao.SysResource.
			Fields(dao.SysResource.Columns.Title, dao.SysResource.Columns.Cover).
			Where(dao.SysResource.Columns.ResourceId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Video:
		detailObj, err := dao.SysVideo.
			Fields(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.Cover).
			Where(dao.SysVideo.Columns.VideoId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Edu:
		detailObj, err := dao.SysEdu.
			Fields(dao.SysEdu.Columns.Title, dao.SysEdu.Columns.Cover).
			Where(dao.SysEdu.Columns.EduId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Group:
		detailObj, err := dao.SysGroup.
			Fields(dao.SysGroup.Columns.Title, dao.SysGroup.Columns.Cover).
			Where(dao.SysGroup.Columns.GroupId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Question:
		detailObj, err := dao.SysQuestion.
			Fields(dao.SysQuestion.Columns.Title).
			Where(dao.SysQuestion.Columns.QuestionId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Topic:
		detailObj, err := dao.SysTopic.
			Fields(dao.SysTopic.Columns.Title).
			Where(dao.SysTopic.Columns.TopicId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Module = module
			detailInfo.Id = id
			return detailInfo, nil
		}
	case shared.Answer:
		detailObj, err := dao.SysAnswer.
			Fields(dao.SysAnswer.Columns.Content, dao.SysAnswer.Columns.TopicId).
			Where(dao.SysAnswer.Columns.AnswerId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Title = detailObj.Content
			detailInfo.Module = module
			detailInfo.Id = detailObj.TopicId
			return detailInfo, nil
		}
	case shared.Comment:
		detailObj, err := dao.SysComment.
			Fields(dao.SysComment.Columns.Content, dao.SysComment.Columns.Module, dao.SysComment.Columns.RelatedId).
			Where(dao.SysComment.Columns.CommentId, id).
			One()
		if err != nil {
			return nil, err
		}
		if detailObj != nil {
			var detailInfo *result.NoticeDetailInfo
			err = gconv.Struct(detailObj, &detailInfo)
			if err != nil {
				return nil, err
			}
			detailInfo.Title = detailObj.Content
			detailInfo.Module = detailObj.Module
			detailInfo.Id = detailObj.RelatedId
			return detailInfo, nil
		}
	}
	return nil, gerror.New("类型不匹配")
}

// CheckHasNoticeLike 是否通知过点赞
func (s *noticeService) CheckHasNoticeLike(module string, id int64) bool {
	count, err := dao.SysNotice.
		Where(dao.SysNotice.Columns.DetailId, id).
		Where(dao.SysNotice.Columns.DetailModule, module).Count()
	if err == nil && count == 0 {
		return false
	}
	return true
}
