package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/regex"
	"fiber/app/tools/response"
	"fiber/library/redis"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Article = new(articleService)

type articleService struct {
}

// SelectList 查询文章列表
func (s *articleService) SelectList(req *dto.ArticleQuery) (int, []*result.ArticleList, response.ResponseCode) {

	model := dao.SysArticle.SysArticleDao.Where(dao.SysArticle.Columns.UserId, req.UserId).
		Where(dao.SysArticle.Columns.DeleteTime, nil).
		Order(dao.SysArticle.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysArticle.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysArticle.Columns.ArticleId,
		dao.SysArticle.Columns.Title,
		dao.SysArticle.Columns.Cover,
		dao.SysArticle.Columns.Description,
		dao.SysArticle.Columns.Status,
		dao.SysArticle.Columns.CreateTime,
	)
	var res []*result.ArticleList
	err = model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return total, res, response.SUCCESS
}

// SelectFilterList 查询过滤文章列表
func (s *articleService) SelectFilterList(req *dto.QueryParam) (int, []*result.ArticleFilterList, error) {

	model := dao.SysArticle.SysArticleDao.
		Where(dao.SysArticle.Columns.DeleteTime, nil)

	model = model.Where(dao.SysArticle.Columns.Status, shared.StatusReviewed)

	if req.Title != "" {
		model = model.Where(dao.SysArticle.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.CateId != 0 {
		model = model.Where(dao.SysArticle.Columns.CateId, req.CateId)
	}

	if req.UserId != 0 && !req.IsFavorite {
		model = model.Where(dao.SysArticle.Columns.UserId, req.UserId)
	}

	if req.IsFavorite && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserFavorite.
			Where(dao.SysUserFavorite.Columns.UserId, req.UserId).
			Where(dao.SysUserFavorite.Columns.Module, shared.Article).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.FavoriteId)
		}
		model = model.Where(dao.SysArticle.Columns.ArticleId+" IN(?)", ids)
	}

	if req.TagId != 0 {
		//	获取标签关联的id
		tagRelatedList, err := dao.SysTagRelated.
			Where(dao.SysTagRelated.Columns.TagId, req.TagId).
			All()
		if err != nil {
			return 0, nil, err
		}

		var relatedIds []int64
		for _, i := range tagRelatedList {
			if i.Module == shared.Article {
				relatedIds = append(relatedIds, i.RelatedId)
			}
		}
		model = model.Where(dao.SysArticle.Columns.ArticleId+" IN(?)", relatedIds)
	}

	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysArticle.Columns.CreateTime, "desc")
	case shared.ModeHot:
		model = model.Order(dao.SysArticle.Columns.Hots, "desc")
	default:
		model = model.Order(dao.SysArticle.Columns.UpdateTime, "desc")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, err
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysArticle.Columns.ArticleId,
		dao.SysArticle.Columns.UserId,
		dao.SysArticle.Columns.Title,
		dao.SysArticle.Columns.Cover,
		dao.SysArticle.Columns.Description,
		dao.SysArticle.Columns.Likes,
		dao.SysArticle.Columns.Hots,
		dao.SysArticle.Columns.Views,
		dao.SysArticle.Columns.Favorites,
		dao.SysArticle.Columns.Status,
		dao.SysArticle.Columns.CreateTime,
		dao.SysArticle.Columns.CateId,
	)
	list, err := model.All()

	if err != nil {
		return 0, nil, nil
	}
	var res []*result.ArticleFilterList
	for _, i := range list {
		var info *result.ArticleFilterList
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return 0, nil, err
		}
		info.Id = i.ArticleId
		info.Module = shared.Article
		res = append(res, info)
	}

	return total, res, nil
}

// SelectByHomeList 查询首页列表
func (s *articleService) SelectByHomeList(ids string) ([]*result.ArticleFilterList, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysArticle.Where(dao.SysArticle.Columns.ArticleId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.ArticleFilterList
	for _, i := range list {
		var info *result.ArticleFilterList
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return nil, err
		}
		info.Id = i.ArticleId
		info.Module = shared.Article
		res = append(res, info)
	}
	return res, nil
}

// Create 创建文章
func (s *articleService) Create(req *dto.ArticleCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ArticleCreateCount+gconv.String(req.UserId),
		shared.ArticleCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}
	var entity model.SysArticle
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReview

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
	res, err := tx.Insert(dao.SysArticle.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := res.LastInsertId()
	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 增加 关联圈子
	//if len(req.GroupId) > 0 {
	//	err = Group.AddGroupRelated(req.GroupId, rid, shared.Article)
	//	if err != nil {
	//		return response.ADD_FAILED
	//	}
	//}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, rid, shared.Article)
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

// EditInfo 取文章编辑信息
func (s *articleService) EditInfo(userId int64, id int64) (*result.ArticleEditInfo, response.ResponseCode) {
	var editInfo *result.ArticleEditInfo
	err := dao.SysArticle.
		Where(dao.SysArticle.Columns.ArticleId, id).
		Where(dao.SysArticle.Columns.UserId, userId).
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

	//groupList, err := Group.SelectGroupRelatedList(id, shared.Article)
	//if err != nil {
	//	return nil, response.DB_READ_ERROR
	//}
	//editInfo.GroupList = groupList

	return editInfo, response.SUCCESS
}

// Edit 编辑修改文章
func (s *articleService) Edit(req *dto.ArticleEdit) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ArticleEditCount+gconv.String(req.UserId),
		shared.ArticleEditLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity = g.Map{
		dao.SysArticle.Columns.Status:      shared.StatusReview,
		dao.SysArticle.Columns.Cover:       req.Cover,
		dao.SysArticle.Columns.Title:       req.Title,
		dao.SysArticle.Columns.Content:     req.Content,
		dao.SysArticle.Columns.CateId:      req.CateId,
		dao.SysArticle.Columns.Description: req.Description,
		dao.SysArticle.Columns.UpdateTime:  gtime.Now(),
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_READ_ERROR
	}
	defer func() {
		if err != nil {
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
	err = Tag.RemoveRelated(tx, req.ArticleId, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, req.ArticleId, shared.Article)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.ArticleId, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}
	pathList, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.ArticleId, shared.Article)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	return response.SUCCESS
}

// Remove 删除文章
func (s *articleService) Remove(userId, id int64) (code response.ResponseCode) {

	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_READ_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Model(dao.SysArticle.Table).
		Where(dao.SysArticle.Columns.ArticleId, id).
		Delete(dao.SysArticle.Columns.UserId, userId)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, id, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, id, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Article)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// SelectInfo 查询文章信息
func (s *articleService) SelectInfo(userId, id int64) (*result.ArticleInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysArticle.Update(g.Map{
		dao.SysArticle.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysArticle.Columns.ArticleId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Article + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.ArticleInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if userId != 0 {
			rs.IsFavorite = User.CheckUserFavorite(userId, id, shared.Article)
			rs.IsLike = User.CheckUserLike(userId, id, shared.Article)
		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysArticle.
		Where(dao.SysArticle.Columns.ArticleId, id).
		Where(dao.SysArticle.Columns.DeleteTime, nil).
		Where(dao.SysArticle.Columns.Status, shared.StatusReviewed).
		One()
	if err != nil || info == nil {
		return nil, response.NOT_FOUND
	}

	rs, err := s.info(userId, info)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	redisCom.Time = 600
	redisCom.Data = rs
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	return rs, response.SUCCESS
}

// info 详细信息数据转换
func (s *articleService) info(userId int64, info *model.SysArticle) (*result.ArticleInfo, error) {
	var res result.ArticleInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	//groupList, err := Group.SelectGroupRelatedList(info.ArticleId, shared.Article)
	//if err != nil {
	//	return nil, err
	//}
	//res.GroupList = groupList

	cateInfo, err := Category.SelectInfo(info.CateId, shared.Article)
	if err != nil {
		return nil, err
	}
	res.CateInfo = cateInfo

	userInfo, err := User.SelectInfo(userId, info.UserId)
	if err != nil {
		return nil, err
	}
	res.UserInfo = userInfo

	// 获取标签
	tagList, err := Tag.SelectRelatedList(info.ArticleId, shared.Article)
	if err != nil {
		return nil, err
	}
	res.TagList = tagList

	// 获取评论数量
	comments, err := dao.SysComment.
		Where(dao.SysComment.Columns.RelatedId, info.ArticleId).
		Where(dao.SysComment.Columns.Module, shared.Article).Count()
	if err != nil {
		return nil, err
	}
	res.Comments = gconv.Int64(comments)

	// 设置查看数量 -----------------
	if userId != 0 {
		if User.CheckUserFavorite(userId, info.ArticleId, shared.Article) {
			res.IsFavorite = true
		}

		if User.CheckUserLike(userId, info.ArticleId, shared.Article) {
			res.IsLike = true
		}
	}
	return &res, nil
}

// Like 点赞
func (s *articleService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ArticleLikeCount+gconv.String(userId)+gconv.String(id),
		shared.ArticleLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysArticle.
		Fields(dao.SysArticle.Columns.UserId, dao.SysArticle.Columns.Title).
		Where(dao.SysArticle.Columns.ArticleId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Article)
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

	var redisCom redis.Com
	redisCom.Key = shared.Article + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Article
		_, err := tx.Insert(dao.SysUserLike.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新 作者获赞数
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysUser.Columns.UserId, info.UserId)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysArticle.Table, g.Map{
			dao.SysArticle.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysArticle.Columns.ArticleId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.ArticleInfo
			err := gconv.Struct(infoObj, &rs)
			if err != nil {
				return response.CACHE_READ_ERROR
			}
			rs.Likes += 1
			redisCom.Time = 600
			redisCom.Data = rs
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}

		// 检查是否点赞通知已经存在
		if !Notice.CheckHasNoticeLike(shared.Article, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Article
			notice.Content = "点赞了你发布的《" + info.Title + "》文章"
			notice.Status = shared.NoticeStatusReview
			notice.Receiver = gconv.Int64(info.UserId)
			notice.CreateTime = gtime.Now()
			_, err = tx.Insert(dao.SysNotice.Table, notice)
			if err != nil {
				return response.DB_SAVE_ERROR
			}
		}

		//	设置点赞积分
		err = Integral.SetUserLikeAndFavoriteIntegral(redisCom, tx, userId)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	} else {
		_, err := tx.Model(dao.SysUserLike.Table).
			Where(dao.SysUserLike.Columns.UserId, userId).
			Where(dao.SysUserLike.Columns.RelatedId, id).
			Where(dao.SysUserLike.Columns.Module, shared.Article).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新 作者获赞数
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysUser.Columns.UserId, info.UserId)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysArticle.Table, g.Map{
			dao.SysArticle.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysArticle.Columns.ArticleId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.ArticleInfo
			err := gconv.Struct(infoObj, &rs)
			if err != nil {
				return response.CACHE_READ_ERROR
			}
			rs.Likes -= 1
			redisCom.Time = 600
			redisCom.Data = rs
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}
	}

	//	设置用户活跃度
	redisCom.Key = shared.UserHot
	redisCom.Data = userId
	err = redisCom.ADDSet()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	return response.SUCCESS
}

// Favorite 收藏
func (s *articleService) Favorite(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ArticleFavoriteCount+gconv.String(userId)+gconv.String(id),
		shared.ArticleFavoriteLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, id).
		Count(dao.SysUserFavorite.Columns.Module, shared.Article)
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

	var redisCom redis.Com
	redisCom.Key = shared.Article + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	// 判断是否收藏
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserFavorite
		entity.UserId = userId
		entity.FavoriteId = id
		entity.Module = shared.Article
		_, err := tx.Insert(dao.SysUserFavorite.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysArticle.Table, g.Map{
			dao.SysArticle.Columns.Favorites: gdb.Raw("favorites+1"),
		}, dao.SysArticle.Columns.ArticleId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.ArticleInfo
			err := gconv.Struct(infoObj, &rs)
			if err != nil {
				return response.CACHE_READ_ERROR
			}
			rs.Favorites += 1
			redisCom.Time = 600
			redisCom.Data = rs
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}

		//	设置点赞积分
		err = Integral.SetUserLikeAndFavoriteIntegral(redisCom, tx, userId)
		if err != nil {
			return response.DB_SAVE_ERROR
		}
	} else {
		_, err := tx.Model(dao.SysUserFavorite.Table).
			Where(dao.SysUserFavorite.Columns.UserId, userId).
			Where(dao.SysUserFavorite.Columns.FavoriteId, id).
			Where(dao.SysUserFavorite.Columns.Module, shared.Article).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysArticle.Table, g.Map{
			dao.SysArticle.Columns.Favorites: gdb.Raw("favorites-1"),
		}, dao.SysArticle.Columns.ArticleId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.ArticleInfo
			err := gconv.Struct(infoObj, &rs)
			if err != nil {
				return response.CACHE_READ_ERROR
			}
			rs.Favorites -= 1
			redisCom.Time = 600
			redisCom.Data = rs
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}
	}

	//	设置用户活跃度
	redisCom.Key = shared.UserHot
	redisCom.Data = userId
	err = redisCom.ADDSet()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	return response.SUCCESS
}
