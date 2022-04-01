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

var Audio = new(audioService)

type audioService struct{}

// SelectList 查询文章列表
func (s *audioService) SelectList(req *dto.AudioQuery) (int, []*result.AudioList, response.ResponseCode) {

	model := dao.SysAudio.SysAudioDao.Where(dao.SysAudio.Columns.UserId, req.UserId).
		Where(dao.SysAudio.Columns.DeleteTime, nil).Order(dao.SysAudio.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysAudio.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysAudio.Columns.AudioId,
		dao.SysAudio.Columns.Title,
		dao.SysAudio.Columns.Cover,
		dao.SysAudio.Columns.Description,
		dao.SysAudio.Columns.Status,
		dao.SysAudio.Columns.CreateTime,
	)
	var res []*result.AudioList
	err = model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return total, res, response.SUCCESS
}

// Create 创建音频内容
func (s *audioService) Create(req *dto.AudioCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.AudioCreateCount+gconv.String(req.UserId),
		shared.AudioCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysAudio
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
		err = Tag.AddTags(tx, req.Tags, rid, shared.Audio)
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

	// 写入到es

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *audioService) EditInfo(userId, id int64) (*result.AudioEditInfo, response.ResponseCode) {
	var editInfo *result.AudioEditInfo

	err := dao.SysAudio.
		Where(dao.SysAudio.Columns.AudioId, id).
		Where(dao.SysAudio.Columns.UserId, userId).
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

// Edit 编辑内容
func (s *audioService) Edit(req *dto.AudioEdit) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.AudioEditCount+gconv.String(req.UserId),
		shared.AudioEditLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}
	var entity = g.Map{
		dao.SysAudio.Columns.Status:      shared.StatusReview,
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
	_, err = tx.Update(dao.SysAudio.Table, entity, dao.SysAudio.Columns.AudioId, req.AudioId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, req.AudioId, shared.Audio)
	if err != nil {
		return response.UPDATE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, req.AudioId, shared.Audio)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.AudioId, shared.Audio)
	if err != nil {
		return response.UPDATE_FAILED
	}
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Link)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.AudioId, shared.Audio)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	return response.SUCCESS
}

// Remove 删除音频
func (s *audioService) Remove(userId, id int64) (code response.ResponseCode) {

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

	_, err = tx.Model(dao.SysAudio.Table).
		Where(dao.SysAudio.Columns.AudioId, id).
		Delete(dao.SysAudio.Columns.UserId, userId)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, id, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, id, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Audio)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// SelectFilterList 查询列表
func (s *audioService) SelectFilterList(req *dto.QueryParam) (int, []*result.AudioListInfo, error) {

	model := dao.SysAudio.SysAudioDao.Where(dao.SysAudio.Columns.DeleteTime, nil)
	model = model.Where(dao.SysAudio.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysAudio.Columns.Title+" like ?", "%"+req.Title+"%")
	}
	if req.CateId != 0 {
		model = model.Where(dao.SysAudio.Columns.CateId, req.CateId)
	}
	if req.IsDown != 0 {
		model = model.Where(dao.SysAudio.Columns.HasDown, req.IsDown)
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
			if i.Module == shared.Audio {
				relatedIds = append(relatedIds, i.RelatedId)
			}
		}
		model = model.Where(dao.SysAudio.Columns.AudioId+" IN(?)", relatedIds)
	}

	if req.IsBuy && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysOrder.
			Where(dao.SysOrder.Columns.UserId, req.UserId).
			Where(dao.SysOrder.Columns.DetailModule, shared.Audio).
			Where(dao.SysOrder.Columns.Status, shared.StatusReviewed).
			Where(dao.SysOrder.Columns.OrderType, shared.OrderTypeThree).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.DetailId)
		}
		model = model.Where(dao.SysAudio.Columns.AudioId+" IN(?)", ids)
	}

	if req.IsFavorite && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserFavorite.
			Where(dao.SysUserFavorite.Columns.UserId, req.UserId).
			Where(dao.SysUserFavorite.Columns.Module, shared.Audio).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.FavoriteId)
		}
		model = model.Where(dao.SysAudio.Columns.AudioId+" IN(?)", ids)
	}

	if req.UserId != 0 && !req.IsFavorite && !req.IsBuy {
		model = model.Where(dao.SysAudio.Columns.UserId, req.UserId)
	}

	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysAudio.Columns.CreateTime, "desc")
	case shared.ModeHot:
		model = model.Order(dao.SysAudio.Columns.Hots, "desc")
	default:
		model = model.Order(dao.SysAudio.Columns.UpdateTime, "desc")
	}
	total, err := model.Count()
	if err != nil {
		return 0, nil, err
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, err
	}
	var res []*result.AudioListInfo
	for _, i := range list {
		var info *result.AudioListInfo
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return 0, nil, err
		}
		info.Id = i.AudioId
		info.Module = shared.Audio
		res = append(res, info)
	}
	return total, res, nil
}

// SelectByHomeList 查询首页列表
func (s *audioService) SelectByHomeList(ids string) ([]*result.AudioListInfo, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysAudio.Where(dao.SysAudio.Columns.AudioId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.AudioListInfo
	for _, i := range list {
		var info *result.AudioListInfo
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return nil, err
		}
		info.Id = i.AudioId
		info.Module = shared.Audio
		res = append(res, info)
	}
	return res, nil
}

// SelectInfo 查询文章信息
func (s *audioService) SelectInfo(userId, id int64) (*result.AudioInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysAudio.Update(g.Map{
		dao.SysAudio.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysAudio.Columns.AudioId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Audio + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.AudioInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if rs.DownMode == shared.DownModePay && userId != 0 {
			g.Dump(rs)
			rs.IsDown = Order.CheckIsPay(userId, rs.AudioId, shared.Audio, shared.OrderTypeThree)
		}

		if rs.DownMode == shared.DownModeComment {
			isComment := Comment.checkCommentStatus(rs.AudioId, userId, shared.Audio)
			rs.IsDown = isComment
		}

		if rs.DownMode == shared.DownModeCommon || (rs.DownMode == shared.DownModeLogin && userId != 0) || userId == rs.UserInfo.UserId {
			rs.IsDown = true
		}

		if !rs.IsDown {
			rs.DownUrl = ""
		}

		if userId != 0 {
			rs.IsFavorite = User.CheckUserFavorite(userId, id, shared.Audio)
			rs.IsLike = User.CheckUserLike(userId, id, shared.Audio)
		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysAudio.
		Where(dao.SysAudio.Columns.AudioId, id).
		Where(dao.SysAudio.Columns.DeleteTime, nil).
		Where(dao.SysAudio.Columns.Status, shared.StatusReviewed).
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
func (s *audioService) info(userId int64, info *model.SysAudio) (*result.AudioInfo, error) {
	var res result.AudioInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	if info.DownMode == shared.DownModePay && userId != 0 {
		res.IsDown = Order.CheckIsPay(userId, info.AudioId, shared.Audio, shared.OrderTypeThree)
	}

	if info.DownMode == shared.DownModeComment {
		isComment := Comment.checkCommentStatus(info.AudioId, userId, shared.Audio)
		res.IsDown = isComment
	}

	if info.DownMode == shared.DownModeCommon || (info.DownMode == shared.DownModeLogin && userId != 0) || userId == info.UserId {
		res.IsDown = true
	}

	if !res.IsDown {
		res.DownUrl = ""
	}

	//groupList, err := Group.SelectGroupRelatedList(info.AudioId, shared.Audio)
	//if err != nil {
	//	return nil, err
	//}
	//res.GroupList = groupList

	cateInfo, err := Category.SelectInfo(info.CateId, shared.Audio)
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
	tagList, err := Tag.SelectRelatedList(info.AudioId, shared.Audio)
	if err != nil {
		return nil, err
	}
	res.TagList = tagList

	// 获取评论数量
	comments, err := dao.SysComment.
		Where(dao.SysComment.Columns.RelatedId, info.AudioId).
		Where(dao.SysComment.Columns.Module, shared.Audio).Count()
	if err != nil {
		return nil, err
	}
	res.Comments = gconv.Int64(comments)

	if userId != 0 {
		if User.CheckUserFavorite(userId, info.AudioId, shared.Audio) {

			res.IsFavorite = true
		}

		if User.CheckUserLike(userId, info.AudioId, shared.Audio) {
			res.IsLike = true
		}

	}
	return &res, nil
}

// Like 点赞
func (s *audioService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.AudioLikeCount+gconv.String(userId)+gconv.String(id),
		shared.AudioLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysAudio.
		Fields(dao.SysAudio.Columns.UserId, dao.SysAudio.Columns.Title).
		Where(dao.SysAudio.Columns.AudioId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Audio)
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
	redisCom.Key = shared.Audio + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Audio
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
		_, err = tx.Update(dao.SysAudio.Table, g.Map{
			dao.SysAudio.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysAudio.Columns.AudioId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.AudioInfo
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
		if !Notice.CheckHasNoticeLike(shared.Audio, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Audio
			notice.Content = "点赞了你发布的《" + info.Title + "》音频"
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
			Where(dao.SysUserLike.Columns.Module, shared.Audio).
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
		_, err = tx.Update(dao.SysAudio.Table, g.Map{
			dao.SysAudio.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysAudio.Columns.AudioId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.AudioInfo
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
func (s *audioService) Favorite(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.AudioFavoriteCount+gconv.String(userId)+gconv.String(id),
		shared.AudioFavoriteLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, id).
		Count(dao.SysUserFavorite.Columns.Module, shared.Audio)
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
	redisCom.Key = shared.Audio + shared.InfoById + gconv.String(id)
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
		entity.Module = shared.Audio
		_, err := tx.Insert(dao.SysUserFavorite.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysAudio.Table, g.Map{
			dao.SysAudio.Columns.Favorites: gdb.Raw("favorites+1"),
		}, dao.SysAudio.Columns.AudioId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.AudioInfo
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
			Where(dao.SysUserFavorite.Columns.Module, shared.Audio).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysAudio.Table, g.Map{
			dao.SysAudio.Columns.Favorites: gdb.Raw("favorites-1"),
		}, dao.SysAudio.Columns.AudioId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.AudioInfo
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
