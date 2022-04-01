package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	"fiber/library/redis"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Video = new(videoService)

type videoService struct {
}

// SelectList 查询视频列表
func (s *videoService) SelectList(req *dto.VideoQuery) (int, []*result.VideoList, response.ResponseCode) {

	model := dao.SysVideo.SysVideoDao.
		Where(dao.SysVideo.Columns.UserId, req.UserId).
		Where(dao.SysVideo.Columns.DeleteTime, nil).
		Order(dao.SysVideo.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysVideo.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysVideo.Columns.VideoId,
		dao.SysVideo.Columns.Title,
		dao.SysVideo.Columns.Cover,
		dao.SysVideo.Columns.Description,
		dao.SysVideo.Columns.Status,
		dao.SysVideo.Columns.CreateTime,
	)
	var res []*result.VideoList
	err = model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return total, res, response.SUCCESS
}

// SelectFilterList 查询过滤视频列表
func (s *videoService) SelectFilterList(req *dto.QueryParam) (int, []*result.VideoListInfo, error) {
	model := dao.SysVideo.SysVideoDao.Where(dao.SysVideo.Columns.DeleteTime, nil)
	model = model.Where(dao.SysVideo.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysVideo.Columns.Title+" like ?", "%"+req.Title+"%")
	}
	if req.CateId != 0 {
		model = model.Where(dao.SysVideo.Columns.CateId, req.CateId)
	}
	if req.IsDown != 0 {
		model = model.Where(dao.SysVideo.Columns.HasDown, req.IsDown)
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
			if i.Module == shared.Video {
				relatedIds = append(relatedIds, i.RelatedId)
			}
		}
		model = model.Where(dao.SysVideo.Columns.VideoId+" IN(?)", relatedIds)
	}

	if req.UserId != 0 && !req.IsFavorite && !req.IsBuy {
		model = model.Where(dao.SysVideo.Columns.UserId, req.UserId)
	}

	if req.IsBuy && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysOrder.
			Where(dao.SysOrder.Columns.UserId, req.UserId).
			Where(dao.SysOrder.Columns.Status, shared.StatusReviewed).
			Where(dao.SysOrder.Columns.DetailModule, shared.Video).
			Where(dao.SysOrder.Columns.OrderType, shared.OrderTypeThree).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.DetailId)
		}
		model = model.Where(dao.SysVideo.Columns.VideoId+" IN(?)", ids)
	}

	if req.IsFavorite && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserFavorite.
			Where(dao.SysUserFavorite.Columns.UserId, req.UserId).
			Where(dao.SysUserFavorite.Columns.Module, shared.Video).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.FavoriteId)
		}
		model = model.Where(dao.SysVideo.Columns.VideoId+" IN(?)", ids)
	}

	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysVideo.Columns.CreateTime, "desc")
	case shared.ModeHot:
		model = model.Order(dao.SysVideo.Columns.Hots, "desc")
	default:
		model = model.Order(dao.SysVideo.Columns.UpdateTime, "desc")
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
	var res []*result.VideoListInfo

	for _, i := range list {
		var info *result.VideoListInfo
		rs, err := s.info(0, i)
		if err != nil {
			return 0, nil, nil
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return 0, nil, nil
		}
		info.Id = i.VideoId
		info.Module = shared.Video
		res = append(res, info)
	}
	return total, res, nil
}

// SelectByHomeList 查询首页列表
func (s *videoService) SelectByHomeList(ids string) ([]*result.VideoListInfo, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysVideo.Where(dao.SysVideo.Columns.VideoId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.VideoListInfo
	for _, i := range list {
		var info *result.VideoListInfo
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return nil, err
		}
		info.Id = i.VideoId
		info.Module = shared.Video
		res = append(res, info)

	}
	return res, nil
}

// Create 创建视频
func (s *videoService) Create(req *dto.VideoCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.VideoCreateCount+gconv.String(req.UserId),
		shared.VideoCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysVideo
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReview

	tx, err := g.DB().Begin()
	if err != nil {
		return response.ADD_FAILED
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
		err = Tag.AddTags(tx, req.Tags, rid, shared.Video)
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

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *videoService) EditInfo(userId, id int64) (*result.VideoEditInfo, response.ResponseCode) {
	var editInfo *result.VideoEditInfo

	err := dao.SysVideo.
		Where(dao.SysVideo.Columns.VideoId, id).
		Where(dao.SysVideo.Columns.UserId, userId).
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
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.VideoEditCount+gconv.String(req.UserId),
		shared.VideoEditLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity = g.Map{
		dao.SysVideo.Columns.Status:      shared.StatusReview,
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

	tx, err := g.DB().Begin()
	if err != nil {
		return response.UPDATE_FAILED
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
	err = Tag.RemoveRelated(tx, req.VideoId, shared.Video)
	if err != nil {
		return response.UPDATE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, req.VideoId, shared.Video)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.VideoId, shared.Video)
	if err != nil {
		return response.UPDATE_FAILED
	}
	//pathList, err := utils.GetSrcLink(req.Content)
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Link)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.VideoId, shared.Video)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	return response.SUCCESS
}

// Remove 删除
func (s *videoService) Remove(userId, id int64) (code response.ResponseCode) {

	tx, err := g.DB().Begin()
	if err != nil {
		return response.DELETE_FAILED
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysVideo.Table).Where(dao.SysVideo.Columns.UserId, userId).
		Delete(dao.SysVideo.Columns.VideoId, id)
	// 删除 关联圈子
	//err = Group.RemoveRelatedList(id, shared.Video)
	//if err != nil {
	//	return response.DELETE_FAILED
	//}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, id, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, id, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Video)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// SelectInfo 查询信息
func (s *videoService) SelectInfo(userId, id int64) (*result.VideoInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysVideo.Update(g.Map{
		dao.SysVideo.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysVideo.Columns.VideoId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Video + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.VideoInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if rs.DownMode == shared.DownModePay && userId != 0 {
			rs.IsDown = Order.CheckIsPay(userId, rs.VideoId, shared.Video, shared.OrderTypeThree)
		}

		if rs.DownMode == shared.DownModeComment {
			isComment := Comment.checkCommentStatus(rs.VideoId, userId, shared.Video)
			rs.IsDown = isComment
		}

		if rs.DownMode == shared.DownModeCommon || (rs.DownMode == shared.DownModeLogin && userId != 0) || userId == rs.UserInfo.UserId {
			rs.IsDown = true
		}

		if !rs.IsDown {
			rs.DownUrl = ""
		}

		if userId != 0 {
			rs.IsFavorite = User.CheckUserFavorite(userId, id, shared.Video)
			rs.IsLike = User.CheckUserLike(userId, id, shared.Video)
		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysVideo.
		Where(dao.SysVideo.Columns.VideoId, id).
		Where(dao.SysVideo.Columns.DeleteTime, nil).
		Where(dao.SysVideo.Columns.Status, shared.StatusReviewed).
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

func (s *videoService) info(userId int64, info *model.SysVideo) (*result.VideoInfo, error) {
	var res result.VideoInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	if info.DownMode == shared.DownModePay && userId != 0 {
		res.IsDown = Order.CheckIsPay(userId, info.VideoId, shared.Video, shared.OrderTypeThree)
	}

	if info.DownMode == shared.DownModeComment {
		isComment := Comment.checkCommentStatus(info.VideoId, userId, shared.Video)
		res.IsDown = isComment
	}

	if info.DownMode == shared.DownModeCommon || (info.DownMode == shared.DownModeLogin && userId != 0) || userId == info.UserId {
		res.IsDown = true
	}

	if !res.IsDown {
		res.DownUrl = ""
	}

	cateInfo, err := Category.SelectInfo(info.CateId, shared.Video)
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
	tagList, err := Tag.SelectRelatedList(info.VideoId, shared.Video)
	if err != nil {
		return nil, err
	}
	res.TagList = tagList

	// 获取评论数量
	comments, err := dao.SysComment.
		Where(dao.SysComment.Columns.RelatedId, info.VideoId).
		Where(dao.SysComment.Columns.Module, shared.Video).Count()
	if err != nil {
		return nil, err
	}
	res.Comments = gconv.Int64(comments)

	if userId != 0 {
		res.IsFavorite = User.CheckUserFavorite(userId, info.VideoId, shared.Video)

		res.IsLike = User.CheckUserLike(userId, info.VideoId, shared.Video)

	}
	return &res, nil
}

// Like 点赞
func (s *videoService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.VideoLikeCount+gconv.String(userId)+gconv.String(id),
		shared.VideoLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysVideo.
		Fields(dao.SysVideo.Columns.UserId, dao.SysVideo.Columns.Title).
		Where(dao.SysVideo.Columns.VideoId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Video)
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
	redisCom.Key = shared.Video + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Video
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
		_, err = tx.Update(dao.SysVideo.Table, g.Map{
			dao.SysVideo.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysVideo.Columns.VideoId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.VideoInfo
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
		if !Notice.CheckHasNoticeLike(shared.Video, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Video
			notice.Content = "点赞了你发布的《" + info.Title + "》视频"
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
			Where(dao.SysUserLike.Columns.Module, shared.Video).
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
		_, err = tx.Update(dao.SysVideo.Table, g.Map{
			dao.SysVideo.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysVideo.Columns.VideoId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.VideoInfo
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
func (s *videoService) Favorite(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.VideoFavoriteCount+gconv.String(userId)+gconv.String(id),
		shared.VideoFavoriteLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, id).
		Count(dao.SysUserFavorite.Columns.Module, shared.Video)
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
	redisCom.Key = shared.Video + shared.InfoById + gconv.String(id)
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
		entity.Module = shared.Video
		_, err := tx.Insert(dao.SysUserFavorite.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysVideo.Table, g.Map{
			dao.SysVideo.Columns.Favorites: gdb.Raw("favorites+1"),
		}, dao.SysVideo.Columns.VideoId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.VideoInfo
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
			Where(dao.SysUserFavorite.Columns.Module, shared.Video).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysVideo.Table, g.Map{
			dao.SysVideo.Columns.Favorites: gdb.Raw("favorites-1"),
		}, dao.SysVideo.Columns.VideoId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.VideoInfo
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
