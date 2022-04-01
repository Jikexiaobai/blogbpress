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
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Topic = new(topicService)

type topicService struct {
}

// SelectList 获取话题列表
func (s *topicService) SelectList(req *dto.TopicQuery) (int, []*result.TopicList, response.ResponseCode) {
	model := dao.SysTopic.SysTopicDao.
		Where(dao.SysTopic.Columns.DeleteTime, nil).
		Where(dao.SysTopic.Columns.Status, shared.StatusReviewed).
		Order(dao.SysTopic.Columns.UpdateTime + " desc")
	if req.Type != 0 {
		model = model.Where(dao.SysTopic.Columns.Type, req.Type)
	}
	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysTopic.Columns.CreateTime + " desc")
	case shared.ModeHot:
		model = model.Order(dao.SysTopic.Columns.Hots + " desc")
	case shared.ModeFollow:
		if req.UserId != 0 {
			//	获取用户关注的
			followIds, err := dao.SysUserFollow.Fields(dao.SysUserFollow.Columns.FollowId).
				Where(dao.SysUserFollow.Columns.UserId, req.UserId).
				All()
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			var followIdList []int64
			for _, i := range followIds {
				followIdList = append(followIdList, i.FollowId)
			}

			model = model.Where(dao.SysTopic.Columns.UserId+" IN(?)", followIdList).Or(dao.SysTopic.Columns.UserId, req.UserId)
			model = model.Order(dao.SysTopic.Columns.CreateTime + " desc")
		}
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

	var res []*result.TopicList
	for _, i := range list {
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		var info *result.TopicList
		err = gconv.Struct(rs, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}

		res = append(res, info)
	}
	return total, res, response.SUCCESS

}

// SelectTop 获取置顶
func (s *topicService) SelectTop() ([]*result.TopicTopList, response.ResponseCode) {
	var list []*result.TopicTopList
	err := dao.SysTopic.
		Where(dao.SysTopic.Columns.DeleteTime, nil).
		Where(dao.SysTopic.Columns.Status, 2).
		Where(dao.SysTopic.Columns.IsTop, 2).Structs(&list)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	return list, response.SUCCESS
}

// Create 创建话题
func (s *topicService) Create(req *dto.TopicCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.TopicCreateCount+gconv.String(req.UserId),
		shared.TopicCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysTopic
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}

	entity.UserId = req.UserId
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReviewed
	entity.IsTop = shared.StatusReview
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

	rs, err := tx.Insert(dao.SysTopic.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	files := gjson.New(req.Files)
	filesList := gconv.Strings(files.Value())
	pathList := gconv.Strings(filesList)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Topic)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	return response.SUCCESS
}

// SelectFilterList 获取话题列表
func (s *topicService) SelectFilterList(req *dto.QueryParam) (int, []*result.TopicListInfo, error) {
	model := dao.SysTopic.SysTopicDao.Where(dao.SysTopic.Columns.DeleteTime, nil)
	model = model.Where(dao.SysTopic.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysTopic.Columns.Title+" like ?", "%"+req.Title+"%")
	}
	if req.Type != 0 {
		model = model.Where(dao.SysTopic.Columns.Type, req.Type)
	}
	if req.IsTop != 0 {
		model = model.Where(dao.SysTopic.Columns.IsTop, req.IsTop)
	}
	if req.GroupId != 0 {
		model = model.Where(dao.SysTopic.Columns.GroupId, req.GroupId)
	}

	if req.UserId != 0 && req.Mode != shared.ModeFollow {
		model = model.Where(dao.SysTopic.Columns.UserId, req.UserId)
	}
	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysTopic.Columns.CreateTime + " desc")
	case shared.ModeHot:
		model = model.Order(dao.SysTopic.Columns.Hots + " desc")
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

			model = model.Where(dao.SysTopic.Columns.UserId+" IN(?)", followIdList).
				Or(dao.SysQuestion.Columns.UserId, req.UserId).Order(dao.SysTopic.Columns.CreateTime + " desc")

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

	var res []*result.TopicListInfo
	for _, i := range list {
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}

		var contentInfo *result.TopicListInfo
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.Id = i.TopicId
		contentInfo.Module = shared.Topic
		res = append(res, contentInfo)
	}
	return total, res, nil

}

// SelectByHomeList 查询首页列表
func (s *topicService) SelectByHomeList(ids string) ([]*result.TopicListInfo, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysTopic.Where(dao.SysTopic.Columns.TopicId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.TopicListInfo
	for _, i := range list {
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		var contentInfo *result.TopicListInfo
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return nil, err
		}

		contentInfo.Id = i.TopicId
		contentInfo.Module = shared.Topic
		res = append(res, contentInfo)
	}
	return res, nil
}

// Remove 删除话题
func (s *topicService) Remove(userId, id int64) (code response.ResponseCode) {

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

	_, err = tx.Model(dao.SysTopic.Table).
		Where(dao.SysTopic.Columns.UserId,
			userId).Delete(dao.SysTopic.Columns.TopicId, id)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除 关联圈子

	//err = Group.RemoveRelatedList(id, shared.Topic)
	//if err != nil {
	//	return response.DELETE_FAILED
	//}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Topic)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, id, shared.Topic)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// SelectInfo 查询文章信息
func (s *topicService) SelectInfo(userId, id int64) (*result.TopicInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysTopic.Update(g.Map{
		dao.SysTopic.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysTopic.Columns.TopicId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}
	var redisCom redis.Com
	redisCom.Key = shared.Topic + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.TopicInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if userId != 0 {
			rs.IsLike = User.CheckUserLike(userId, id, shared.Topic)
		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysTopic.
		Where(dao.SysTopic.Columns.TopicId, id).
		Where(dao.SysTopic.Columns.DeleteTime, nil).
		Where(dao.SysTopic.Columns.Status, shared.StatusReviewed).
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

// Like 点赞
func (s *topicService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.TopicLikeCount+gconv.String(userId)+gconv.String(id),
		shared.TopicLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysTopic.
		Fields(dao.SysTopic.Columns.UserId, dao.SysTopic.Columns.Title).
		Where(dao.SysTopic.Columns.TopicId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Topic)
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
	redisCom.Key = shared.Topic + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Topic
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
		_, err = tx.Update(dao.SysTopic.Table, g.Map{
			dao.SysTopic.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysTopic.Columns.TopicId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.TopicInfo
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
		if !Notice.CheckHasNoticeLike(shared.Topic, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Topic
			notice.Content = "点赞了你发布的《" + info.Title + "》帖子"
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
			Where(dao.SysUserLike.Columns.Module, shared.Topic).
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
		_, err = tx.Update(dao.SysTopic.Table, g.Map{
			dao.SysTopic.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysTopic.Columns.TopicId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.TopicInfo
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

// 转换内容
func (s *topicService) info(userId int64, info *model.SysTopic) (*result.TopicInfo, error) {
	var res result.TopicInfo
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

	// 获取评论数量
	comments, err := dao.SysComment.
		Where(dao.SysComment.Columns.RelatedId, info.TopicId).
		Where(dao.SysComment.Columns.Module, shared.Topic).Count()
	if err != nil {
		return nil, err
	}
	res.Comments = gconv.Int64(comments)

	if info.Type == 2 {
		var relatedInfo result.RelatedInfo
		switch info.Module {
		case shared.Article:
			tmpInfo, err := dao.SysArticle.
				Fields(dao.SysArticle.Columns.Title, dao.SysArticle.Columns.Cover).
				Where(dao.SysArticle.Columns.ArticleId, info.RelatedId).
				One()
			if err != nil {
				return nil, err
			}
			relatedInfo.Cover = tmpInfo.Cover
			relatedInfo.Title = tmpInfo.Title
			relatedInfo.Id = info.RelatedId
			relatedInfo.Module = info.Module
		case shared.Audio:
			tmpInfo, err := dao.SysAudio.
				Fields(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.Cover).
				Where(dao.SysAudio.Columns.AudioId, info.RelatedId).
				One()
			if err != nil {
				return nil, err
			}
			relatedInfo.Cover = tmpInfo.Cover
			relatedInfo.Title = tmpInfo.Title
			relatedInfo.Id = info.RelatedId
			relatedInfo.Module = info.Module
		case shared.Video:
			tmpInfo, err := dao.SysVideo.
				Fields(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.Cover).
				Where(dao.SysVideo.Columns.VideoId, info.RelatedId).
				One()
			if err != nil {
				return nil, err
			}
			relatedInfo.Cover = tmpInfo.Cover
			relatedInfo.Title = tmpInfo.Title
			relatedInfo.Id = info.RelatedId
			relatedInfo.Module = info.Module
		case shared.Resource:
			tmpInfo, err := dao.SysResource.
				Fields(dao.SysResource.Columns.Title, dao.SysResource.Columns.Cover).
				Where(dao.SysResource.Columns.ResourceId, info.RelatedId).
				One()
			if err != nil {
				return nil, err
			}
			relatedInfo.Cover = tmpInfo.Cover
			relatedInfo.Title = tmpInfo.Title
			relatedInfo.Id = info.RelatedId
			relatedInfo.Module = info.Module
		case shared.Edu:
			tmpInfo, err := dao.SysEdu.
				Fields(dao.SysEdu.Columns.Title, dao.SysEdu.Columns.Cover).
				Where(dao.SysEdu.Columns.EduId, info.RelatedId).
				One()
			if err != nil {
				return nil, err
			}
			relatedInfo.Cover = tmpInfo.Cover
			relatedInfo.Title = tmpInfo.Title
			relatedInfo.Id = info.RelatedId
			relatedInfo.Module = "course"
		}
		res.RelatedInfo = relatedInfo
	}

	if userId != 0 {
		res.IsLike = User.CheckUserLike(userId, info.TopicId, shared.Topic)
	}
	return &res, nil
}
