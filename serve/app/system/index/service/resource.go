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

var Resource = new(resourceService)

type resourceService struct {
}

// SelectList 查询列表
func (s *resourceService) SelectList(req *dto.ResourceQuery) (int, []*result.ResourceList, response.ResponseCode) {

	model := dao.SysResource.SysResourceDao.Where(dao.SysResource.Columns.UserId, req.UserId).
		Where(dao.SysResource.Columns.DeleteTime, nil).
		Order(dao.SysResource.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysResource.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysResource.Columns.ResourceId,
		dao.SysResource.Columns.Title,
		dao.SysResource.Columns.Cover,
		dao.SysResource.Columns.Description,
		dao.SysResource.Columns.Status,
		dao.SysResource.Columns.CreateTime,
	)
	var res []*result.ResourceList
	err = model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return total, res, response.SUCCESS
}

// Create 创建资源
func (s *resourceService) Create(req *dto.ResourceCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ResourceCreateCount+gconv.String(req.UserId),
		shared.ResourceCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysResource
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
	rs, err := tx.Insert(dao.SysResource.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	//// 增加 关联圈子
	//if len(req.GroupId) > 0 {
	//	err = Group.AddGroupRelated(req.GroupId, rid, shared.Resource)
	//	if err != nil {
	//		return response.ADD_FAILED
	//	}
	//}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, rid, shared.Resource)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 管理媒体库
	pathList, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Resource)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 写入到es

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *resourceService) EditInfo(userId, id int64) (*result.ResourceEditInfo, response.ResponseCode) {
	var editInfo *result.ResourceEditInfo
	err := dao.SysResource.
		Where(dao.SysResource.Columns.ResourceId, id).
		Where(dao.SysResource.Columns.UserId, userId).
		Struct(&editInfo)

	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}

	// 获取标签
	tagList, err := Tag.SelectRelatedList(id, shared.Resource)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	editInfo.TagList = tagList

	//groupList, err := Group.SelectGroupRelatedList(id, shared.Resource)
	//if err != nil {
	//	return nil, response.DB_READ_ERROR
	//}
	//editInfo.GroupList = groupList

	return editInfo, response.SUCCESS

}

// Edit 编辑修改资源
func (s *resourceService) Edit(req *dto.ResourceEdit) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ResourceEditCount+gconv.String(req.UserId),
		shared.ResourceEditLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity = g.Map{
		dao.SysResource.Columns.Status:      shared.StatusReview,
		dao.SysResource.Columns.Cover:       req.Cover,
		dao.SysResource.Columns.Title:       req.Title,
		dao.SysResource.Columns.Content:     req.Content,
		dao.SysResource.Columns.CateId:      req.CateId,
		dao.SysResource.Columns.Description: req.Description,
		dao.SysResource.Columns.Price:       req.Price,
		dao.SysResource.Columns.DownUrl:     req.DownUrl,
		dao.SysResource.Columns.Attribute:   req.Attribute,
		dao.SysResource.Columns.Purpose:     req.Purpose,
		dao.SysResource.Columns.UpdateTime:  gtime.Now(),
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
	_, err = tx.Update(dao.SysResource.Table, entity, dao.SysResource.Columns.ResourceId, req.ResourceId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, req.ResourceId, shared.Resource)
	if err != nil {
		return response.UPDATE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, req.ResourceId, shared.Resource)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.ResourceId, shared.Resource)
	if err != nil {
		return response.UPDATE_FAILED
	}
	pathList, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.ResourceId, shared.Resource)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}
	return response.SUCCESS
}

// Remove 删除资源
func (s *resourceService) Remove(userId, id int64) (code response.ResponseCode) {

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

	_, err = tx.Model(dao.SysResource.Table).
		Where(dao.SysResource.Columns.UserId, userId).
		Delete(dao.SysResource.Columns.ResourceId, id)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联圈子
	//err = Group.RemoveRelatedList(0, shared.Resource)
	//if err != nil {
	//	return response.DELETE_FAILED
	//}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, id, shared.Resource)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Resource)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, id, shared.Resource)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, id, shared.Resource)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// SelectFilterList 查询过滤列表
func (s *resourceService) SelectFilterList(req *dto.QueryParam) (int, []*result.ResourceListInfo, error) {

	model := dao.SysResource.SysResourceDao.Where(dao.SysResource.Columns.DeleteTime, nil)
	model = model.Where(dao.SysResource.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysResource.Columns.Title+" like ?", "%"+req.Title+"%")
	}
	if req.CateId != 0 {
		model = model.Where(dao.SysResource.Columns.CateId, req.CateId)
	}

	if req.IsDown != 0 {
		model = model.Where(dao.SysResource.Columns.HasDown, req.IsDown)
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
			if i.Module == shared.Resource {
				relatedIds = append(relatedIds, i.RelatedId)
			}
		}
		model = model.Where(dao.SysResource.Columns.ResourceId+" IN(?)", relatedIds)
	}

	if req.UserId != 0 && !req.IsFavorite && !req.IsBuy {
		model = model.Where(dao.SysResource.Columns.UserId, req.UserId)
	}

	if req.IsFavorite && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserFavorite.
			Where(dao.SysUserFavorite.Columns.UserId, req.UserId).
			Where(dao.SysUserFavorite.Columns.Module, shared.Resource).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.FavoriteId)
		}
		model = model.Where(dao.SysResource.Columns.ResourceId+" IN(?)", ids)
	}

	if req.IsBuy && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysOrder.
			Where(dao.SysOrder.Columns.UserId, req.UserId).
			Where(dao.SysOrder.Columns.Status, shared.StatusReview).
			Where(dao.SysOrder.Columns.DetailModule, shared.Resource).
			Where(dao.SysOrder.Columns.OrderType, shared.OrderTypeThree).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.DetailId)
		}
		model = model.Where(dao.SysResource.Columns.ResourceId+" IN(?)", ids)
	}
	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysResource.Columns.CreateTime, "desc")
	case shared.ModeHot:
		model = model.Order(dao.SysResource.Columns.Hots, "desc")
	default:
		model = model.Order(dao.SysResource.Columns.UpdateTime, "desc")
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
	var res []*result.ResourceListInfo

	for _, i := range list {
		var info *result.ResourceListInfo
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return 0, nil, err
		}
		info.Id = i.ResourceId
		info.Module = shared.Resource
		res = append(res, info)
	}
	return total, res, nil
}

// SelectByHomeList 查询首页列表
func (s *resourceService) SelectByHomeList(ids string) ([]*result.ResourceListInfo, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysResource.Where(dao.SysResource.Columns.ResourceId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.ResourceListInfo
	for _, i := range list {

		var info *result.ResourceListInfo
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		err = gconv.Struct(rs, &info)
		if err != nil {
			return nil, err
		}
		info.Id = i.ResourceId
		info.Module = shared.Resource
		res = append(res, info)
	}
	return res, nil
}

// SelectInfo 查询文章信息
func (s *resourceService) SelectInfo(userId, id int64) (*result.ResourceInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysResource.Update(g.Map{
		dao.SysResource.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysResource.Columns.ResourceId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Resource + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.ResourceInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if rs.DownMode == shared.DownModePay && userId != 0 {
			rs.IsDown = Order.CheckIsPay(userId, rs.ResourceId, shared.Resource, shared.OrderTypeThree)
		}

		if rs.DownMode == shared.DownModeComment {
			isComment := Comment.checkCommentStatus(rs.ResourceId, userId, shared.Resource)
			rs.IsDown = isComment
		}

		if rs.DownMode == shared.DownModeCommon || (rs.DownMode == shared.DownModeLogin && userId != 0) || userId == rs.UserInfo.UserId {
			rs.IsDown = true
		}

		if !rs.IsDown {
			rs.DownUrl = ""
		}

		if userId != 0 {
			rs.IsFavorite = User.CheckUserFavorite(userId, id, shared.Resource)
			rs.IsLike = User.CheckUserLike(userId, id, shared.Resource)
		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysResource.
		Where(dao.SysResource.Columns.ResourceId, id).
		Where(dao.SysResource.Columns.DeleteTime, nil).
		Where(dao.SysResource.Columns.Status, shared.StatusReviewed).
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

// Info 详细信息数据转换
func (s *resourceService) info(userId int64, info *model.SysResource) (*result.ResourceInfo, error) {
	var res result.ResourceInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	if info.DownMode == shared.DownModePay && userId != 0 {
		res.IsDown = Order.CheckIsPay(userId, info.ResourceId, shared.Resource, shared.OrderTypeThree)
	}

	if info.DownMode == shared.DownModeComment {
		isComment := Comment.checkCommentStatus(info.ResourceId, userId, shared.Resource)
		res.IsDown = isComment
	}

	if info.DownMode == shared.DownModeCommon || (info.DownMode == shared.DownModeLogin && userId != 0) || userId == info.UserId {
		res.IsDown = true
	}

	if !res.IsDown {
		res.DownUrl = ""
	}

	cateInfo, err := Category.SelectInfo(info.CateId, shared.Resource)
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
	tagList, err := Tag.SelectRelatedList(info.ResourceId, shared.Resource)
	if err != nil {
		return nil, err
	}
	res.TagList = tagList

	// 获取评论数量
	comments, err := dao.SysComment.
		Where(dao.SysComment.Columns.RelatedId, info.ResourceId).
		Where(dao.SysComment.Columns.Module, shared.Resource).Count()
	if err != nil {
		return nil, err
	}
	res.Comments = gconv.Int64(comments)

	if userId != 0 {
		res.IsFavorite = User.CheckUserFavorite(userId, info.ResourceId, shared.Resource)
		res.IsLike = User.CheckUserLike(userId, info.ResourceId, shared.Resource)

	}
	return &res, nil
}

// Like 点赞
func (s *resourceService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ResourceLikeCount+gconv.String(userId)+gconv.String(id),
		shared.ResourceLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysResource.
		Fields(dao.SysResource.Columns.UserId, dao.SysResource.Columns.Title).
		Where(dao.SysResource.Columns.ResourceId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Resource)
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
	redisCom.Key = shared.Resource + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Resource
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
		_, err = tx.Update(dao.SysResource.Table, g.Map{
			dao.SysResource.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysResource.Columns.ResourceId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.ResourceInfo
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
		if !Notice.CheckHasNoticeLike(shared.Resource, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Resource
			notice.Content = "点赞了你发布的《" + info.Title + "》资源"
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
			Where(dao.SysUserLike.Columns.Module, shared.Resource).
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
		_, err = tx.Update(dao.SysResource.Table, g.Map{
			dao.SysResource.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysResource.Columns.ResourceId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.ResourceInfo
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
func (s *resourceService) Favorite(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.ResourceFavoriteCount+gconv.String(userId)+gconv.String(id),
		shared.ResourceFavoriteLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, id).
		Count(dao.SysUserFavorite.Columns.Module, shared.Resource)
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
	redisCom.Key = shared.Resource + shared.InfoById + gconv.String(id)
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
		entity.Module = shared.Resource
		_, err := tx.Insert(dao.SysUserFavorite.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysResource.Table, g.Map{
			dao.SysResource.Columns.Favorites: gdb.Raw("favorites+1"),
		}, dao.SysResource.Columns.ResourceId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.ResourceInfo
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
			Where(dao.SysUserFavorite.Columns.Module, shared.Resource).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysResource.Table, g.Map{
			dao.SysResource.Columns.Favorites: gdb.Raw("favorites-1"),
		}, dao.SysResource.Columns.ResourceId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.ResourceInfo
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
