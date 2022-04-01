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

var Question = new(questionService)

type questionService struct {
}

// SelectList 获取列表
func (s *questionService) SelectList(req *dto.QuestionQuery) (int, []*result.QuestionList, response.ResponseCode) {

	model := dao.SysQuestion.SysQuestionDao.
		Where(dao.SysQuestion.Columns.DeleteTime, nil).
		Where(dao.SysQuestion.Columns.Status, shared.StatusReviewed)
	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysQuestion.Columns.Hots + " desc")
	case shared.ModeHot:
		model = model.Order(dao.SysQuestion.Columns.CreateTime + " desc")
	case shared.ModeFavorite:
		model = model.Order(dao.SysQuestion.Columns.Favorites + " desc")
	case shared.ModeView:
		model = model.Order(dao.SysQuestion.Columns.Views + " desc")
	default:
		model = model.Order(dao.SysResource.Columns.UpdateTime, "desc")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.QuestionList
	for _, i := range list {
		var info *result.QuestionList
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		err = gconv.Struct(rs, &info)
		res = append(res, info)
	}

	return total, res, response.SUCCESS
}

// Create 创建问题
func (s *questionService) Create(req *dto.QuestionCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.QuestionCreateCount+gconv.String(req.UserId),
		shared.QuestionCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysQuestion
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

	insertRes, err := tx.Insert(dao.SysQuestion.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := insertRes.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 管理媒体库
	pathList, err := regex.GetSrcLink(req.Content)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Question)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	//if entity.Status == shared.REVIEWED {
	//	err = Integral.SetUserCommentIntegral(userId)
	//	if err != nil {
	//		return nil, response.DB_SAVE_ERROR
	//	}
	//}

	return response.SUCCESS
}

// Remove 删除问题
func (s *questionService) Remove(userId, id int64) (code response.ResponseCode) {
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

	_, err = tx.Model(dao.SysQuestion.Table).
		Where(dao.SysQuestion.Columns.UserId, userId).
		Delete(dao.SysQuestion.Columns.QuestionId, id)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联圈子
	//err = Group(id, shared.Question)
	//if err != nil {
	//	return response.DELETE_FAILED
	//}

	err = Media.RemoveRelated(tx, id, shared.Question)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// SelectFilterList 获取过滤列表
func (s *questionService) SelectFilterList(req *dto.QueryParam) (int, []*result.QuestionListInfo, error) {

	model := dao.SysQuestion.SysQuestionDao.
		Where(dao.SysQuestion.Columns.DeleteTime, nil)
	model = model.Where(dao.SysQuestion.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysQuestion.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.GroupId != 0 {
		model = model.Where(dao.SysQuestion.Columns.GroupId, req.GroupId)
	}

	if req.UserId != 0 && !req.IsFavorite {
		model = model.Where(dao.SysQuestion.Columns.UserId, req.UserId)
	}

	if req.IsFavorite && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserFavorite.
			Where(dao.SysUserFavorite.Columns.UserId, req.UserId).
			Where(dao.SysUserFavorite.Columns.Module, shared.Question).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.FavoriteId)
		}
		model = model.Where(dao.SysQuestion.Columns.QuestionId+" IN(?)", ids)
	}

	switch req.Mode {
	// 最新
	case shared.ModeHot:
		model = model.Order(dao.SysQuestion.Columns.Hots + " desc")
	case shared.ModeNew:
		model = model.Order(dao.SysQuestion.Columns.CreateTime + " desc")
	case shared.ModeFollow:
		if req.UserId != 0 {
			//	获取用户关注的
			followIds, err := dao.SysUserFollow.Fields(dao.SysUserFollow.Columns.FollowId).
				Where(dao.SysUserFollow.Columns.UserId, req.UserId).
				All()
			if err != nil {
				return 0, nil, err
			}
			var followIdList []int64
			for _, i := range followIds {
				followIdList = append(followIdList, i.FollowId)
			}

			model = model.Where(dao.SysQuestion.Columns.UserId+" IN(?)", followIdList).Or(dao.SysQuestion.Columns.UserId, req.UserId)
			model = model.Order(dao.SysQuestion.Columns.CreateTime + " desc")
		}
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

	var res []*result.QuestionListInfo

	for _, i := range list {
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}

		var contentInfo *result.QuestionListInfo
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.Id = i.QuestionId
		contentInfo.Module = shared.Question
		res = append(res, contentInfo)
	}

	return total, res, nil
}

// SelectByHomeList 查询首页列表
func (s *questionService) SelectByHomeList(ids string) ([]*result.QuestionListInfo, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysQuestion.Where(dao.SysQuestion.Columns.QuestionId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.QuestionListInfo
	for _, i := range list {
		var contentInfo *result.QuestionListInfo
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		err = gconv.Struct(rs, &contentInfo)
		contentInfo.Id = i.QuestionId
		contentInfo.Module = shared.Question
		res = append(res, contentInfo)
	}

	return res, nil
}

// SelectInfo 查询文章信息
func (s *questionService) SelectInfo(userId, id int64) (*result.QuestionInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysQuestion.Update(g.Map{
		dao.SysQuestion.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysQuestion.Columns.QuestionId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Question + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.QuestionInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if userId != 0 {

			rs.IsFavorite = User.CheckUserFavorite(userId, id, shared.Question)

			rs.IsLike = User.CheckUserLike(userId, id, shared.Question)

		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysQuestion.
		Where(dao.SysQuestion.Columns.QuestionId, id).
		Where(dao.SysQuestion.Columns.DeleteTime, nil).
		Where(dao.SysQuestion.Columns.Status, shared.StatusReviewed).
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

func (s *questionService) info(userId int64, info *model.SysQuestion) (*result.QuestionInfo, error) {
	var res result.QuestionInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	userInfo, err := User.SelectInfo(userId, info.UserId)
	if err != nil {
		return nil, err
	}
	res.UserInfo = userInfo

	group, err := Group.SelectRelatedGroup(info.GroupId)
	if err != nil {
		return nil, err
	}
	res.GroupInfo = group

	// 获取回答数
	answersCount, err := dao.SysAnswer.
		Where(dao.SysAnswer.Columns.TopicId, info.QuestionId).
		Where(dao.SysAnswer.Columns.DeleteTime, nil).
		Count()
	if err != nil {
		return nil, err
	}
	res.Answers = gconv.Int64(answersCount)
	if userId != 0 {

		if User.CheckUserFavorite(userId, info.QuestionId, shared.Question) {
			res.IsFavorite = true
		}

		if User.CheckUserLike(userId, info.QuestionId, shared.Question) {
			res.IsLike = true
		}

	}
	return &res, nil
}

// Like 点赞
func (s *questionService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.QuestionLikeCount+gconv.String(userId)+gconv.String(id),
		shared.QuestionLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysQuestion.
		Fields(dao.SysQuestion.Columns.UserId, dao.SysQuestion.Columns.Title).
		Where(dao.SysQuestion.Columns.QuestionId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Question)
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
	redisCom.Key = shared.Question + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Question
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
		_, err = tx.Update(dao.SysQuestion.Table, g.Map{
			dao.SysQuestion.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysQuestion.Columns.QuestionId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.QuestionInfo
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
		if !Notice.CheckHasNoticeLike(shared.Question, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Question
			notice.Content = "点赞了你发布的《" + info.Title + "》问题"
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
			Where(dao.SysUserLike.Columns.Module, shared.Question).
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
		_, err = tx.Update(dao.SysQuestion.Table, g.Map{
			dao.SysQuestion.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysQuestion.Columns.QuestionId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.QuestionInfo
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
func (s *questionService) Favorite(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.QuestionFavoriteCount+gconv.String(userId)+gconv.String(id),
		shared.QuestionFavoriteLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, id).
		Count(dao.SysUserFavorite.Columns.Module, shared.Question)
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
	redisCom.Key = shared.Question + shared.InfoById + gconv.String(id)
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
		entity.Module = shared.Question
		_, err := tx.Insert(dao.SysUserFavorite.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysQuestion.Table, g.Map{
			dao.SysQuestion.Columns.Favorites: gdb.Raw("favorites+1"),
		}, dao.SysQuestion.Columns.QuestionId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.QuestionInfo
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
			Where(dao.SysUserFavorite.Columns.Module, shared.Question).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysQuestion.Table, g.Map{
			dao.SysQuestion.Columns.Favorites: gdb.Raw("favorites-1"),
		}, dao.SysQuestion.Columns.QuestionId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.QuestionInfo
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
