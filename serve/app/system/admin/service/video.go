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

var Video = new(videoService)

type videoService struct {
}

// SelectList 查询
func (s *videoService) SelectList(req *dto.VideoQuery) (int, []*result.VideoList, response.ResponseCode) {
	model := dao.SysVideo.SysVideoDao.Order(dao.SysVideo.Columns.UpdateTime + " desc")

	if req.Status == 4 {
		model = model.Where(dao.SysVideo.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysVideo.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysVideo.Columns.Status, req.Status)
	}
	if req.CateId != 0 {
		model = model.Where(dao.SysVideo.Columns.CateId, req.CateId)
	}
	if req.Title != "" {
		model = model.Where(dao.SysVideo.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysVideo.Columns.UserId,
		dao.SysVideo.Columns.VideoId,
		dao.SysVideo.Columns.Title,
		dao.SysVideo.Columns.Cover,
		dao.SysVideo.Columns.Status,
		dao.SysVideo.Columns.CreateTime,
		dao.SysVideo.Columns.CateId,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.VideoList
	for _, i := range list {
		var info *result.VideoList
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

// Create 创建
func (s *videoService) Create(req *dto.VideoCreate) (code response.ResponseCode) {

	var entity model.SysVideo
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

	rs, err := tx.Insert(dao.SysVideo.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, rid, shared.Video)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 管理媒体库
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Link)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Video)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 写入到es

	return response.SUCCESS
}

// EditInfo 获取信息
func (s *videoService) EditInfo(id int64) (*result.VideoEditInfo, response.ResponseCode) {
	var editInfo *result.VideoEditInfo

	err := dao.SysVideo.
		Where(dao.SysVideo.Columns.VideoId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}

	// 获取标签
	tagList, err := Tag.SelectRelatedList(id, shared.Video)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	editInfo.TagList = tagList

	return editInfo, response.SUCCESS

}

// Edit 编辑
func (s *videoService) Edit(req *dto.VideoEdit) (code response.ResponseCode) {

	entity := g.Map{
		dao.SysVideo.Columns.Link:        req.Link,
		dao.SysVideo.Columns.Cover:       req.Cover,
		dao.SysVideo.Columns.Title:       req.Title,
		dao.SysVideo.Columns.CateId:      req.CateId,
		dao.SysVideo.Columns.Description: req.Description,
		dao.SysVideo.Columns.Price:       req.Price,
		dao.SysVideo.Columns.DownUrl:     req.DownUrl,
		dao.SysVideo.Columns.Attribute:   req.Attribute,
		dao.SysVideo.Columns.Purpose:     req.Purpose,
		dao.SysVideo.Columns.UpdateTime:  gtime.Now(),
	}

	entity[dao.SysVideo.Columns.HasDown] = req.HasDown
	entity[dao.SysVideo.Columns.DownMode] = req.DownMode
	entity[dao.SysVideo.Columns.UserId] = req.UserId
	entity[dao.SysVideo.Columns.Hots] = req.Hots
	entity[dao.SysVideo.Columns.Views] = req.Views
	entity[dao.SysVideo.Columns.Favorites] = req.Favorites
	entity[dao.SysVideo.Columns.Likes] = req.Likes
	entity[dao.SysVideo.Columns.Status] = shared.StatusReviewed

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

	_, err = tx.Update(dao.SysVideo.Table, entity, dao.SysVideo.Columns.VideoId, req.VideoId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, []int64{req.VideoId}, shared.Video)
	if err != nil {

		return response.DB_SAVE_ERROR
	}
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, req.VideoId, shared.Video)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.VideoId}, shared.Video)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	//pathList, err := utils.GetSrcLink(req.Content)
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Link)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.VideoId, shared.Video)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// Review 更新状态
func (s *videoService) Review(req *dto.Review) (code response.ResponseCode) {
	list, err := dao.SysVideo.Fields(
		dao.SysVideo.Columns.UserId,
		dao.SysVideo.Columns.Title,
		dao.SysVideo.Columns.VideoId,
		dao.SysVideo.Columns.Status).
		Where(dao.SysVideo.Columns.VideoId+" IN(?)", req.IdList).All()
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
	var tmpInfoList []*model.SysVideo
	for _, i := range list {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.VideoId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysVideo.Table,
		g.Map{
			dao.SysVideo.Columns.Status: req.Status,
			dao.SysVideo.Columns.Remark: req.Remark,
		}, dao.SysVideo.Columns.VideoId+" IN(?)", tmpIds)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.VideoId
		notice.DetailModule = shared.Video
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的视频《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的视频《" + i.Title + "》未通过审核，原因：" + req.Remark
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
func (s *videoService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysVideo.Fields(
		dao.SysVideo.Columns.UserId,
		dao.SysVideo.Columns.VideoId,
		dao.SysVideo.Columns.Title).
		Where(dao.SysVideo.Columns.VideoId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysVideo.Table, g.Map{
		dao.SysVideo.Columns.DeleteTime: gtime.Now(),
	}, dao.SysVideo.Columns.VideoId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}

	var noticeList []model.SysNotice
	for _, info := range list {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.SystemType = shared.NoticeSysTemDeleteContent
		notice.DetailId = info.VideoId
		notice.DetailModule = shared.Video
		notice.Status = shared.NoticeStatusReview
		notice.Content = "您发布的视频《" + info.Title + "》已被删除，原因：" + req.Remark
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
func (s *videoService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysVideo.Update(g.Map{
		dao.SysVideo.Columns.DeleteTime: nil,
	}, dao.SysVideo.Columns.VideoId+" IN(?)", idList)
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	return response.DB_SAVE_ERROR
}

// Remove 删除
func (s *videoService) Remove(idList []int64) (code response.ResponseCode) {

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

	_, err = tx.Delete(dao.SysVideo.Table,
		dao.SysVideo.Columns.VideoId+" IN(?)", idList)

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Video)
	if err != nil {

		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *videoService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysVideo.Fields(
		dao.SysVideo.Columns.UserId,
		dao.SysVideo.Columns.VideoId,
		dao.SysVideo.Columns.Title).
		Where(dao.SysVideo.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysVideo.Table,
		dao.SysVideo.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	var idList []int64
	for _, info := range list {
		idList = append(idList, info.VideoId)
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Video)
	if err != nil {
		return err
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Video)
	if err != nil {
		return err
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Video)
	if err != nil {
		return err
	}

	err = Media.RemoveRelated(tx, idList, shared.Video)
	if err != nil {

		return err
	}

	return nil
}
