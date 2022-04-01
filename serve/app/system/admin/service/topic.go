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

var Topic = new(topicService)

type topicService struct {
}

// SelectList 查询列表
func (s *topicService) SelectList(req *dto.TopicQuery) (int, []*result.TopicList, response.ResponseCode) {
	model := dao.SysTopic.SysTopicDao.Order(dao.SysTopic.Columns.UpdateTime + " desc")
	if req.Status == 4 {
		model = model.Where(dao.SysTopic.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysTopic.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysTopic.Columns.Status, req.Status)
	}
	if req.Title != "" {
		model = model.Where(dao.SysTopic.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysTopic.Columns.UserId,
		dao.SysTopic.Columns.TopicId,
		dao.SysTopic.Columns.IsTop,
		dao.SysTopic.Columns.Title,
		dao.SysTopic.Columns.Status,
		dao.SysTopic.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	var res []*result.TopicList

	for _, i := range list {
		var info *result.TopicList
		err := gconv.Struct(i, &info)
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
func (s *topicService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysTopic.Fields(
		dao.SysTopic.Columns.UserId,
		dao.SysTopic.Columns.TopicId,
		dao.SysTopic.Columns.Title,
		dao.SysTopic.Columns.Status).
		Where(dao.SysTopic.Columns.TopicId+" IN(?)", req.IdList).All()
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
	var tmpInfoList []*model.SysTopic
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.TopicId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysTopic.Table,
		g.Map{
			dao.SysTopic.Columns.Status: req.Status,
			dao.SysTopic.Columns.Remark: req.Remark,
		}, dao.SysTopic.Columns.TopicId+" IN(?)", tmpIds)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.TopicId
		notice.DetailModule = shared.Article
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的帖子《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的帖子《" + i.Title + "》未通过审核，原因：" + req.Remark
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

// SetTop 更新状态
func (s *topicService) SetTop(req *dto.TopicTop) response.ResponseCode {
	var dto = g.Map{
		dao.SysTopic.Columns.IsTop: req.IsTop,
	}
	_, err := dao.SysTopic.Update(dto, dao.SysTopic.Columns.TopicId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	return response.SUCCESS
}

// Recover 软删除
func (s *topicService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysTopic.Fields(
		dao.SysTopic.Columns.UserId,
		dao.SysTopic.Columns.TopicId,
		dao.SysTopic.Columns.Title).
		Where(dao.SysTopic.Columns.TopicId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysTopic.Table, g.Map{
		dao.SysTopic.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysTopic.Columns.TopicId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	var noticeList []model.SysNotice
	for _, info := range list {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.SystemType = shared.NoticeSysTemDeleteContent
		notice.DetailId = info.TopicId
		notice.DetailModule = shared.Topic
		notice.Status = shared.NoticeStatusReview
		notice.Content = "您发布的帖子《" + info.Title + "》已被删除，原因：" + req.Remark
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
func (s *topicService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysTopic.Update(g.Map{
		dao.SysTopic.Columns.DeleteTime: nil,
	}, dao.SysTopic.Columns.TopicId+" IN(?)", idList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	return response.SUCCESS
}

// Remove 删除
func (s *topicService) Remove(idList []int64) (code response.ResponseCode) {

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

	_, err = tx.Delete(dao.SysTopic.Table,
		dao.SysTopic.Columns.TopicId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Topic)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, idList, shared.Topic)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

func (s *topicService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysTopic.Fields(
		dao.SysTopic.Columns.UserId,
		dao.SysTopic.Columns.TopicId,
		dao.SysTopic.Columns.Title).
		Where(dao.SysTopic.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}
	_, err = tx.Delete(dao.SysTopic.Table, dao.SysTopic.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.TopicId)
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Topic)
	if err != nil {
		return err
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, idList, shared.Topic)
	if err != nil {
		return err
	}

	return nil
}
