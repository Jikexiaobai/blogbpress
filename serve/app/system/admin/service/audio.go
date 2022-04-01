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

var Audio = new(audioService)

type audioService struct{}

// SelectList 查询列表
func (s *audioService) SelectList(req *dto.AudioQuery) (int, []*result.AudioList, response.ResponseCode) {

	model := dao.SysAudio.SysAudioDao.Order(dao.SysAudio.Columns.UpdateTime + " desc")
	//model = model.Where(dao.SysAudio.Columns.Status, shared.REVIEWED)
	if req.Status == 4 {
		model = model.Where(dao.SysAudio.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysAudio.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysAudio.Columns.Status, req.Status)
	}
	if req.CateId != 0 {
		model = model.Where(dao.SysAudio.Columns.CateId, req.CateId)
	}
	if req.Title != "" {
		model = model.Where(dao.SysAudio.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysAudio.Columns.AudioId,
		dao.SysAudio.Columns.Title,
		dao.SysAudio.Columns.Cover,
		dao.SysAudio.Columns.Status,
		dao.SysAudio.Columns.CreateTime,
		dao.SysAudio.Columns.UserId,
		dao.SysAudio.Columns.CateId,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.AudioList
	for _, i := range list {
		var info *result.AudioList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.FAILD
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
func (s *audioService) Create(req *dto.AudioCreate) (code response.ResponseCode) {
	// 加入锁限制

	var entity model.SysAudio
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

	audio, err := tx.Insert(dao.SysAudio.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := audio.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, rid, shared.Audio)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 管理媒体库
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Link)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Audio)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *audioService) EditInfo(id int64) (*result.AudioEditInfo, response.ResponseCode) {
	var editInfo *result.AudioEditInfo

	err := dao.SysAudio.
		Where(dao.SysAudio.Columns.AudioId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	// 获取标签
	tagList, err := Tag.SelectRelatedList(id, shared.Audio)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	editInfo.TagList = tagList
	return editInfo, response.SUCCESS

}

// Edit 编辑
func (s *audioService) Edit(req *dto.AudioEdit) (code response.ResponseCode) {

	entity := g.Map{
		dao.SysAudio.Columns.Link:        req.Link,
		dao.SysAudio.Columns.Cover:       req.Cover,
		dao.SysAudio.Columns.Title:       req.Title,
		dao.SysAudio.Columns.CateId:      req.CateId,
		dao.SysAudio.Columns.Description: req.Description,
		dao.SysAudio.Columns.Price:       req.Price,
		dao.SysAudio.Columns.DownUrl:     req.DownUrl,
		dao.SysAudio.Columns.Attribute:   req.Attribute,
		dao.SysAudio.Columns.Purpose:     req.Purpose,
		dao.SysAudio.Columns.UpdateTime:  gtime.Now(),
	}

	entity[dao.SysAudio.Columns.HasDown] = req.HasDown
	entity[dao.SysAudio.Columns.DownMode] = req.DownMode
	entity[dao.SysAudio.Columns.UserId] = req.UserId
	entity[dao.SysAudio.Columns.Hots] = req.Hots
	entity[dao.SysAudio.Columns.Views] = req.Views
	entity[dao.SysAudio.Columns.Favorites] = req.Favorites
	entity[dao.SysAudio.Columns.Likes] = req.Likes
	entity[dao.SysAudio.Columns.Status] = shared.StatusReviewed

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
	_, err = tx.Update(dao.SysAudio.Table, entity, dao.SysAudio.Columns.AudioId, req.AudioId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, []int64{req.AudioId}, shared.Audio)
	if err != nil {

		return response.DELETE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, req.AudioId, shared.Audio)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.AudioId}, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}
	//pathList, err := utils.GetSrcLink(req.Content)
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Link)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.AudioId, shared.Audio)
		if err != nil {
			return response.FILE_SAVE_ERROR
		}
	}

	return response.SUCCESS
}

// Review 更新状态
func (s *audioService) Review(req *dto.Review) (code response.ResponseCode) {
	list, err := dao.SysAudio.Fields(
		dao.SysAudio.Columns.UserId,
		dao.SysAudio.Columns.Title,
		dao.SysAudio.Columns.AudioId,
		dao.SysAudio.Columns.Status).
		Where(dao.SysAudio.Columns.AudioId+" IN(?)", req.IdList).All()
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
	var tmpInfoList []*model.SysAudio
	for _, i := range list {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.AudioId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.FAILD
	}
	_, err = tx.Update(dao.SysAudio.Table,
		g.Map{
			dao.SysAudio.Columns.Status: req.Status,
			dao.SysAudio.Columns.Remark: req.Remark,
		}, dao.SysAudio.Columns.AudioId+" IN(?)", tmpIds)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.AudioId
		notice.DetailModule = shared.Audio
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的音频《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的音频《" + i.Title + "》未通过审核，原因：" + req.Remark
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
func (s *audioService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysAudio.Fields(
		dao.SysAudio.Columns.UserId,
		dao.SysAudio.Columns.AudioId,
		dao.SysAudio.Columns.Title).
		Where(dao.SysAudio.Columns.AudioId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}

	_, err = tx.Update(dao.SysAudio.Table, g.Map{
		dao.SysAudio.Columns.DeleteTime: gtime.Now(),
	}, dao.SysAudio.Columns.AudioId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_READ_ERROR
	}

	if req.Remark != "" {
		var noticeList []model.SysNotice
		for _, info := range list {
			var notice model.SysNotice
			notice.Type = shared.NoticeSystem
			notice.SystemType = shared.NoticeSysTemDeleteContent
			notice.DetailId = info.AudioId
			notice.DetailModule = shared.Audio
			notice.Status = shared.NoticeStatusReview
			notice.Content = "您发布的音频《" + info.Title + "》已被删除，原因：" + req.Remark
			notice.Receiver = info.UserId
			notice.CreateTime = gtime.Now()
			noticeList = append(noticeList, notice)
		}

		// 写入通知
		_, err = tx.Insert(dao.SysNotice.Table, noticeList)
		if err != nil {
			return response.DB_READ_ERROR
		}
	}

	return response.SUCCESS
}

// Reduction 还原
func (s *audioService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysAudio.Update(g.Map{
		dao.SysAudio.Columns.DeleteTime: nil,
	}, dao.SysAudio.Columns.AudioId+" IN(?)", idList)
	if err != nil {
		return response.DB_READ_ERROR
	}
	return response.SUCCESS
}

// Remove 删除
func (s *audioService) Remove(idList []int64) (code response.ResponseCode) {

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
	_, err = tx.Delete(dao.SysAudio.Table,
		dao.SysAudio.Columns.AudioId+" IN(?)", idList)

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Audio)
	if err != nil {

		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// RemoveByUser 删除用户时关联删除用户的音频
func (s *audioService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysAudio.Fields(
		dao.SysAudio.Columns.UserId,
		dao.SysAudio.Columns.AudioId,
		dao.SysAudio.Columns.Title).
		Where(dao.SysAudio.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysAudio.Table,
		dao.SysAudio.Columns.UserId+" IN(?)", Ids)

	var idList []int64

	for _, info := range list {
		idList = append(idList, info.AudioId)
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Audio)
	if err != nil {
		return err
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Audio)
	if err != nil {
		return err
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Audio)
	if err != nil {
		return err
	}

	err = Media.RemoveRelated(tx, idList, shared.Audio)
	if err != nil {

		return err
	}

	return nil
}
