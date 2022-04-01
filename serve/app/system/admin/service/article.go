package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/regex"
	"fiber/app/tools/response"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Article = new(articleService)

type articleService struct {
}

// SelectList 查询列表
func (s *articleService) SelectList(req *dto.ArticleQuery) (int, []*result.ArticleList, response.ResponseCode) {

	model := dao.SysArticle.SysArticleDao.Order(dao.SysArticle.Columns.UpdateTime + " desc")

	if req.Status == 4 {
		model = model.Where(dao.SysArticle.Columns.DeleteTime+" !=", "")
	}
	if req.Status != 4 {
		model = model.Where(dao.SysArticle.Columns.DeleteTime, nil)
	}

	if req.Status != 0 && req.Status != 4 {
		model = model.Where(dao.SysArticle.Columns.Status, req.Status)
	}
	if req.Title != "" {
		model = model.Where(dao.SysArticle.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.CateId != 0 {
		model = model.Where(dao.SysArticle.Columns.CateId, req.CateId)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysArticle.Columns.ArticleId,
		dao.SysArticle.Columns.Title,
		dao.SysArticle.Columns.Cover,
		dao.SysArticle.Columns.Status,
		dao.SysArticle.Columns.UserId,
		dao.SysArticle.Columns.CateId,
		dao.SysArticle.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.ArticleList

	for _, i := range list {
		var info *result.ArticleList
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
func (s *articleService) Create(req *dto.ArticleCreate) (code response.ResponseCode) {

	var entity model.SysArticle
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
	res, err := tx.Insert(dao.SysArticle.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := res.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, rid, shared.Article)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 管理媒体库
	pathList, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Article)
		if err != nil {
			return response.ADD_FAILED
		}
	}
	return response.SUCCESS
}

// EditInfo 编辑信息
func (s *articleService) EditInfo(id int64) (*result.ArticleEditInfo, response.ResponseCode) {
	var editInfo *result.ArticleEditInfo

	err := dao.SysArticle.
		Where(dao.SysArticle.Columns.ArticleId, id).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	// 获取标签
	tagList, err := Tag.SelectRelatedList(id, shared.Article)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	editInfo.TagList = tagList

	return editInfo, response.SUCCESS
}

// Edit 编辑
func (s *articleService) Edit(req *dto.ArticleEdit) (code response.ResponseCode) {

	entity := make(map[string]interface{})
	entity[dao.SysArticle.Columns.UserId] = req.UserId
	entity[dao.SysArticle.Columns.Hots] = req.Hots
	entity[dao.SysArticle.Columns.Views] = req.Views
	entity[dao.SysArticle.Columns.Favorites] = req.Favorites
	entity[dao.SysArticle.Columns.Likes] = req.Likes
	entity[dao.SysArticle.Columns.Status] = shared.StatusReviewed
	entity[dao.SysArticle.Columns.Cover] = req.Cover
	entity[dao.SysArticle.Columns.Title] = req.Title
	entity[dao.SysArticle.Columns.Content] = req.Content
	entity[dao.SysArticle.Columns.CateId] = req.CateId
	entity[dao.SysArticle.Columns.Description] = req.Description
	entity[dao.SysArticle.Columns.UpdateTime] = gtime.Now()

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
	_, err = tx.Update(dao.SysArticle.Table, entity, dao.SysArticle.Columns.ArticleId, req.ArticleId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, []int64{req.ArticleId}, shared.Article)
	if err != nil {

		return response.UPDATE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddRelated(tx, req.Tags, req.ArticleId, shared.Article)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, []int64{req.ArticleId}, shared.Article)
	if err != nil {

		return response.UPDATE_FAILED
	}
	pathList, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.ArticleId, shared.Article)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	return response.SUCCESS
}

// Recover 软删除
func (s *articleService) Recover(req *dto.Remove) (code response.ResponseCode) {

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

	list, err := dao.SysArticle.Fields(
		dao.SysArticle.Columns.UserId,
		dao.SysArticle.Columns.ArticleId,
		dao.SysArticle.Columns.Title).
		Where(dao.SysArticle.Columns.ArticleId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_TX_ERROR
	}

	_, err = tx.Update(dao.SysArticle.Table, g.Map{
		dao.SysArticle.Columns.DeleteTime: gtime.Now(),
	},
		dao.SysArticle.Columns.ArticleId+" IN(?)", req.IdList)
	if err != nil {
		return response.DB_TX_ERROR
	}

	if req.Remark != "" {
		var noticeList []model.SysNotice
		for _, info := range list {
			var notice model.SysNotice
			notice.Type = shared.NoticeSystem
			notice.SystemType = shared.NoticeSysTemDeleteContent
			notice.DetailId = info.ArticleId
			notice.DetailModule = shared.Article
			notice.Status = shared.NoticeStatusReview
			notice.Content = "你发布的文章《" + info.Title + "》已被删除，原因：" + req.Remark
			notice.Receiver = info.UserId
			notice.CreateTime = gtime.Now()
			noticeList = append(noticeList, notice)
		}
		// 写入通知
		if len(noticeList) > 0 {
			_, err = tx.Insert(dao.SysNotice.Table, noticeList)
			if err != nil {
				return response.DB_TX_ERROR
			}
		}
	}
	return response.SUCCESS
}

// Reduction 还原
func (s *articleService) Reduction(idList []int64) response.ResponseCode {
	_, err := dao.SysArticle.Update(g.Map{
		dao.SysArticle.Columns.DeleteTime: nil,
	}, dao.SysArticle.Columns.ArticleId+" IN(?)", idList)
	if err != nil {
		return response.UPDATE_FAILED
	}
	return response.SUCCESS
}

// Review 审核
func (s *articleService) Review(req *dto.Review) (code response.ResponseCode) {
	info, err := dao.SysArticle.Fields(
		dao.SysArticle.Columns.UserId,
		dao.SysArticle.Columns.ArticleId,
		dao.SysArticle.Columns.Title,
		dao.SysArticle.Columns.Status).
		Where(dao.SysArticle.Columns.ArticleId+" IN(?)", req.IdList).All()
	if err != nil {
		return response.DB_READ_ERROR
	}
	entity := g.Map{
		dao.SysArticle.Columns.Status: req.Status,
		dao.SysArticle.Columns.Remark: req.Remark,
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
	var tmpInfoList []*model.SysArticle
	for _, i := range info {
		if i.Status == 1 {
			tmpIds = append(tmpIds, i.ArticleId)
			tmpInfoList = append(tmpInfoList, i)
		}
	}

	if len(tmpIds) < 1 {
		return response.UPDATE_FAILED
	}
	_, err = tx.Update(dao.SysArticle.Table,
		entity, dao.SysArticle.Columns.ArticleId+" IN(?)", tmpIds)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 设置通知
	var noticeList []model.SysNotice
	for _, i := range tmpInfoList {
		var notice model.SysNotice
		notice.Type = shared.NoticeSystem
		notice.DetailId = i.ArticleId
		notice.DetailModule = shared.Article
		notice.Status = shared.NoticeStatusReview
		notice.Receiver = i.UserId
		notice.CreateTime = gtime.Now()
		notice.SystemType = shared.NoticeSysTemReview
		if req.Status == shared.StatusReviewed {
			notice.Content = "您发布的文章《" + i.Title + "》已通过审核" + req.Remark
			err = Integral.SetUserContentIntegral(tx, i.UserId)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}
		if req.Status == shared.StatusRefuse {
			notice.Content = "您发布的文章《" + i.Title + "》未通过审核，原因：" + req.Remark
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
func (s *articleService) Remove(idList []int64) (code response.ResponseCode) {

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

	_, err = tx.Delete(dao.SysArticle.Table,
		dao.SysArticle.Columns.ArticleId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, idList, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除用户时关联删除用户的文章
func (s *articleService) RemoveByUser(tx *gdb.TX, Ids []int64) error {

	list, err := dao.SysArticle.Fields(
		dao.SysArticle.Columns.UserId,
		dao.SysArticle.Columns.ArticleId,
		dao.SysArticle.Columns.Title).
		Where(dao.SysArticle.Columns.UserId+" IN(?)", Ids).All()
	if err != nil {
		return err
	}

	_, err = tx.Delete(dao.SysArticle.Table, dao.SysArticle.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}

	// 删除redis 内容
	var idList []int64
	for _, info := range list {
		idList = append(idList, info.ArticleId)
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, idList, shared.Article)
	if err != nil {
		return err
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, idList, shared.Article)
	if err != nil {
		return err
	}

	err = User.RemoveUserFavorite(tx, idList, shared.Article)
	if err != nil {
		return err
	}

	err = Media.RemoveRelated(tx, idList, shared.Article)
	if err != nil {

		return err
	}

	return nil
}
