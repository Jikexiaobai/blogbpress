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

var Comment = new(commentService)

type commentService struct {
}

// SelectList 查询列表
func (s *commentService) SelectList(req *dto.CommentQueryParam) (int, []*result.CommentList, response.ResponseCode) {
	model := dao.SysComment.SysCommentDao.Order(dao.SysComment.Columns.UpdateTime + " desc")
	if req.Status == 4 {
		model = model.Where(dao.SysComment.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysComment.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysComment.Columns.Status, req.Status)
	}

	if req.RelatedId != 0 && req.Module != "" {
		model = model.Where(dao.SysComment.Columns.RelatedId, req.RelatedId)
		model = model.Where(dao.SysComment.Columns.Module, req.Module)
	}

	if req.Module != "" {
		model = model.Where(dao.SysComment.Columns.Module, req.Module)
	}

	if req.Content != "" {
		model = model.Where(dao.SysComment.Columns.Content+" like ?", "%"+req.Content+"%")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.
		Fields(
			dao.SysComment.Columns.CommentId,
			dao.SysComment.Columns.UserId,
			dao.SysComment.Columns.Module,
			dao.SysComment.Columns.RelatedId,
			dao.SysComment.Columns.Content,
			dao.SysComment.Columns.Type,
			dao.SysComment.Columns.Status,
			dao.SysComment.Columns.CreateTime,
		).
		All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.CommentList
	for _, i := range list {
		var info *result.CommentList
		err := gconv.Struct(i, &info)
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
func (s *commentService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysComment.Fields(
		dao.SysComment.Columns.UserId,
		dao.SysComment.Columns.CommentId,
		dao.SysComment.Columns.Content).
		Where(dao.SysComment.Columns.CommentId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysComment.Table, g.Map{
		dao.SysComment.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysComment.Columns.CommentId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	var noticeList []model.SysNotice
	for _, info := range list {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.SystemType = shared.NoticeSysTemDeleteContent
		notice.DetailId = info.CommentId
		notice.DetailModule = shared.Comment
		notice.Status = shared.NoticeStatusReview
		notice.Content = "你发布的评论《" + info.Content + "》已被删除，原因：" + req.Remark
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
func (s *commentService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysComment.Update(g.Map{
		dao.SysComment.Columns.DeleteTime: nil,
	}, dao.SysComment.Columns.CommentId+" IN(?)", idList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	return response.SUCCESS
}

// Review 更新状态
func (s *commentService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysComment.Fields(
		dao.SysComment.Columns.UserId,
		dao.SysComment.Columns.CommentId,
		dao.SysComment.Columns.Content,
		dao.SysComment.Columns.RelatedId,
		dao.SysComment.Columns.Module,
		//dao.SysComment.Columns.ParentId,
		dao.SysComment.Columns.Status).
		Where(dao.SysComment.Columns.CommentId+" IN(?)", req.IdList).All()
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
	var tmpInfoList []*model.SysComment
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.CommentId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysComment.Table,
		g.Map{
			dao.SysComment.Columns.Status: req.Status,
			dao.SysComment.Columns.Remark: req.Remark,
		}, dao.SysComment.Columns.CommentId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		// 通知模块所属作者
		var notice model.SysNotice
		notice.Type = shared.NoticeComment
		switch i.Module {
		case shared.Article:
			tmpInfo, err := dao.SysArticle.
				Fields(dao.SysArticle.Columns.UserId, dao.SysArticle.Columns.Title).
				Where(dao.SysArticle.Columns.ArticleId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}
			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Title + "》文章"
		case shared.Answer:
			tmpInfo, err := dao.SysAnswer.
				Fields(dao.SysAnswer.Columns.UserId, dao.SysAnswer.Columns.Content).
				Where(dao.SysAnswer.Columns.AnswerId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}

			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Content + "》答案"
		case shared.Audio:
			tmpInfo, err := dao.SysAudio.
				Fields(dao.SysAudio.Columns.UserId, dao.SysAudio.Columns.Title).
				Where(dao.SysAudio.Columns.AudioId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}

			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Title + "》音频"
		case shared.Video:
			tmpInfo, err := dao.SysVideo.
				Fields(dao.SysVideo.Columns.UserId, dao.SysVideo.Columns.Title).
				Where(dao.SysVideo.Columns.VideoId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}

			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Title + "》视频"
		case shared.Resource:
			tmpInfo, err := dao.SysResource.
				Fields(dao.SysResource.Columns.UserId, dao.SysResource.Columns.Title).
				Where(dao.SysResource.Columns.ResourceId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}

			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Title + "》资源"
		case shared.Topic:
			tmpInfo, err := dao.SysTopic.
				Fields(dao.SysTopic.Columns.UserId, dao.SysTopic.Columns.Title).
				Where(dao.SysTopic.Columns.TopicId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}

			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Title + "》帖子"
		case shared.Edu:
			tmpInfo, err := dao.SysEdu.
				Fields(dao.SysEdu.Columns.UserId, dao.SysEdu.Columns.Title).
				Where(dao.SysEdu.Columns.EduId, i.RelatedId).One()
			if err != nil {
				return response.DB_READ_ERROR
			}

			notice.Receiver = gconv.Int64(tmpInfo.UserId)
			notice.Content = "评论了你发布的《" + tmpInfo.Title + "》课程"
		}
		notice.DetailId = i.CommentId
		notice.DetailModule = shared.Comment
		notice.Status = shared.NoticeStatusReview
		notice.CreateTime = gtime.Now()
		noticeList = append(noticeList, notice)

		// 通知评论用户
		notice.Type = shared.NoticeSystem
		notice.Receiver = i.UserId
		notice.DetailId = i.CommentId
		notice.DetailModule = shared.Comment
		notice.Status = shared.NoticeStatusReview
		notice.CreateTime = gtime.Now()
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的评论《" + i.Content + "》已通过审核" + req.Remark
			err = Integral.SetUserCommentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的评论《" + i.Content + "》未通过审核，原因：" + req.Remark
		}
		noticeList = append(noticeList, notice)

		// 通知回复
		//if i.ParentId != 0 && req.Status == shared.StatusReviewed {
		//	parent, err := dao.SysComment.
		//		Fields(dao.SysComment.Columns.UserId, dao.SysComment.Columns.Content).
		//		Where(dao.SysComment.Columns.CommentId, i.ParentId).One()
		//	if err != nil {
		//		return response.DB_READ_ERROR
		//	}
		//	if gconv.Int64(parent.ParentId) != notice.Receiver {
		//
		//		notice.Receiver = gconv.Int64(parent.UserId)
		//		notice.Content = "回复了你的评论《" + i.Content + "》"
		//		notice.DetailModule = shared.Comment
		//		notice.DetailId = parent.ParentId
		//		notice.Type = shared.NoticeComment
		//		noticeList = append(noticeList, notice)
		//	}
		//}
	}
	if len(noticeList) > 0 {
		_, err = tx.Insert(dao.SysNotice.Table, noticeList)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}
	return response.SUCCESS
}

// Remove 删除
func (s *commentService) Remove(idList []int64) (code response.ResponseCode) {

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

	_, err = tx.Delete(dao.SysComment.Table,
		dao.SysComment.Columns.CommentId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Comment)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Comment)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Comment)
	if err != nil {
		return response.DELETE_FAILED
	}

	//_, err = tx.Update(dao.SysComment.Table,
	//	g.Map{
	//		dao.SysComment.Columns.ParentId: 0,
	//	}, dao.SysComment.Columns.ParentId+" IN(?)", idList)
	//if err != nil {
	//	return response.DB_SAVE_ERROR
	//}

	return response.SUCCESS
}

// RemoveByUser 删除
func (s *commentService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysComment.Fields(
		dao.SysComment.Columns.UserId,
		dao.SysComment.Columns.CommentId,
		dao.SysComment.Columns.Content).
		Where(dao.SysComment.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysComment.Table, dao.SysComment.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.CommentId)
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Comment)
	if err != nil {
		return err
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Comment)
	if err != nil {
		return err
	}

	err = Media.RemoveRelated(tx, idList, shared.Comment)
	if err != nil {

		return err
	}

	//_, err = tx.Update(dao.SysComment.Table,
	//	g.Map{
	//		dao.SysComment.Columns.ParentId: 0,
	//	}, dao.SysComment.Columns.ParentId+" IN(?)", idList)
	//if err != nil {
	//	return err
	//}

	return nil
}
