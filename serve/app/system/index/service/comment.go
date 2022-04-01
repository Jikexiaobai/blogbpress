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
	"github.com/gogf/gf/util/gconv"
)

var Comment = new(commentService)

type commentService struct {
}

// Info 详细信息数据转换
func (s *commentService) info(userId int64, info *model.SysComment) (*result.CommentInfo, error) {
	var res *result.CommentInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	userInfo, err := User.SelectInfo(userId, info.UserId)
	if err != nil {
		return nil, err
	}
	res.UserInfo = userInfo

	switch info.Module {
	case shared.Article:
		title, err := dao.SysArticle.Value(dao.SysArticle.Columns.Title, dao.SysArticle.Columns.ArticleId, info.RelatedId)
		if err != nil {
			return nil, err
		}
		res.Title = gconv.String(title)
	case shared.Audio:
		title, err := dao.SysAudio.Value(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.AudioId, info.RelatedId)
		if err != nil {
			return nil, err
		}
		res.Title = gconv.String(title)
	case shared.Edu:
		title, err := dao.SysEdu.Value(dao.SysEdu.Columns.Title, dao.SysEdu.Columns.EduId, info.RelatedId)
		if err != nil {
			return nil, err
		}
		res.Title = gconv.String(title)
	case shared.Resource:
		title, err := dao.SysResource.Value(dao.SysResource.Columns.Title, dao.SysResource.Columns.ResourceId, info.RelatedId)
		if err != nil {
			return nil, err
		}
		res.Title = gconv.String(title)
	case shared.Video:
		title, err := dao.SysVideo.Value(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.VideoId, info.RelatedId)
		if err != nil {
			return nil, err
		}
		res.Title = gconv.String(title)
	case shared.Topic:
		title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title, dao.SysTopic.Columns.TopicId, info.RelatedId)
		if err != nil {
			return nil, err
		}
		res.Title = gconv.String(title)
	}

	if userId != 0 {
		res.IsLike = User.CheckUserLike(userId, info.CommentId, shared.Comment)
	}
	return res, nil
}

// Create 创建 评论
func (s *commentService) Create(userId int64, req *dto.CommentCreate) (res *result.CommentInfo, code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.CommentCreateCount+gconv.String(userId),
		shared.CommentCreateLock+gconv.String(userId), 60, 5)
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	var entity *model.SysComment
	err = gconv.Struct(req, &entity)
	if err != nil {
		return nil, response.INVALID
	}

	entity.UserId = userId
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReviewed

	tx, err := g.DB().Begin()
	if err != nil {
		return nil, response.DB_TX_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	insertRes, err := tx.Insert(dao.SysComment.Table, entity)
	if err != nil {
		return nil, response.ADD_FAILED
	}
	rid, err := insertRes.LastInsertId()

	files := gjson.New(req.Files)
	filesList := gconv.Strings(files.Value())
	pathList := gconv.Strings(filesList)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Comment)
		if err != nil {
			return nil, response.ADD_FAILED
		}
	}

	var redisCom redis.Com
	//	设置用户活跃
	redisCom.Key = shared.UserHot
	redisCom.Data = userId
	err = redisCom.ADDSet()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	if entity.Status == shared.StatusReviewed {
		//通知问题作者
		var notice model.SysNotice
		notice.FromUserId = userId
		notice.Type = shared.NoticeComment
		notice.DetailModule = shared.Comment
		notice.Status = shared.NoticeStatusReview
		notice.CreateTime = gtime.Now()
		if req.TopId == 0 && req.ParentId == 0 {
			switch req.Module {
			case shared.Article:
				info, err := dao.SysArticle.
					Fields(dao.SysArticle.Columns.UserId, dao.SysArticle.Columns.Title).
					Where(dao.SysArticle.Columns.ArticleId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}
				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			case shared.Answer:
				info, err := dao.SysAnswer.
					Fields(dao.SysAnswer.Columns.UserId, dao.SysAnswer.Columns.Content).
					Where(dao.SysAnswer.Columns.AnswerId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}

				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			case shared.Audio:
				info, err := dao.SysAudio.
					Fields(dao.SysAudio.Columns.UserId, dao.SysAudio.Columns.Title).
					Where(dao.SysAudio.Columns.AudioId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}

				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			case shared.Video:
				info, err := dao.SysVideo.
					Fields(dao.SysVideo.Columns.UserId, dao.SysVideo.Columns.Title).
					Where(dao.SysVideo.Columns.VideoId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}

				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			case shared.Resource:
				info, err := dao.SysResource.
					Fields(dao.SysResource.Columns.UserId, dao.SysResource.Columns.Title).
					Where(dao.SysResource.Columns.ResourceId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}

				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			case shared.Topic:
				info, err := dao.SysTopic.
					Fields(dao.SysTopic.Columns.UserId, dao.SysTopic.Columns.Title).
					Where(dao.SysTopic.Columns.TopicId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}

				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			case shared.Edu:
				info, err := dao.SysEdu.
					Fields(dao.SysEdu.Columns.UserId, dao.SysEdu.Columns.Title).
					Where(dao.SysEdu.Columns.EduId, req.RelatedId).One()
				if err != nil {
					return nil, response.DB_READ_ERROR
				}

				notice.Receiver = gconv.Int64(info.UserId)
				notice.Content = req.Content
			}
			_, err = tx.Insert(dao.SysNotice.Table, notice)
			if err != nil {
				return nil, response.DB_SAVE_ERROR
			}
		}

		// 通知回复
		if req.TopId != 0 {
			notice.DetailId = req.ParentId
			notice.DetailModule = shared.Comment
			notice.Receiver = req.ReplyId
			notice.Content = req.Content
			_, err = tx.Insert(dao.SysNotice.Table, notice)
			if err != nil {
				return nil, response.DB_SAVE_ERROR
			}
		}

		err = Integral.SetUserCommentIntegral(redisCom, tx, userId)
		if err != nil {
			return nil, response.DB_SAVE_ERROR
		}
	}

	err = gconv.Struct(entity, &res)
	res.CommentId = rid
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	userInfo, err := User.SelectInfo(userId, userId)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	res.UserInfo = userInfo

	return res, response.SUCCESS
}

// SelectList 查询对应模块评论列表
func (s *commentService) SelectList(userId int64, req *dto.CommentQuery) (int, []*result.CommentInfo, error) {
	model := dao.SysComment.SysCommentDao.
		Where(dao.SysComment.Columns.DeleteTime, nil)
	model = model.Where(dao.SysComment.Columns.Status, shared.StatusReviewed)

	if req.RelatedId != 0 && req.Module != "" {
		model = model.Where(dao.SysComment.Columns.RelatedId, req.RelatedId)
		model = model.Where(dao.SysComment.Columns.Module, req.Module)
	}

	if req.Module != "" {
		model = model.Where(dao.SysComment.Columns.Module, req.Module)
	}

	if req.Content != "" {
		model = model.Where(dao.SysComment.Columns.Content+" like ?", "%"+req.Content+"%")
	}

	if req.RelatedId == 0 && userId != 0 {
		model = model.Where(dao.SysComment.Columns.UserId, userId)
	}

	model = model.Order(dao.SysComment.Columns.UpdateTime + " desc")
	// 只获取最上层的评论
	model = model.Where(dao.SysComment.Columns.TopId, 0)
	total, err := model.Count()
	if err != nil {
		return 0, nil, err
	}
	model = model.Page(req.Page, req.Limit)

	list, err := model.All()
	if err != nil {
		return 0, nil, err
	}

	// 获取顶级评论下的前5条回复
	var res []*result.CommentInfo
	for _, i := range list {
		// 获取所有的评论回复列表下前5条内容

		info, err := s.info(userId, i)
		if err != nil {
			return 0, nil, err
		}
		res = append(res, info)
	}
	children, err := s.getReply(userId, list)
	if err != nil {
		return 0, nil, err
	}
	res = append(res, children...)

	return total, res, nil
}

//getReply 查询回复内容
func (s *commentService) getReply(userId int64, list []*model.SysComment) ([]*result.CommentInfo, error) {
	var res []*result.CommentInfo
	if len(list) > 0 {
		for _, i := range list {
			children, err := dao.SysComment.
				Where(dao.SysComment.Columns.TopId, i.CommentId).
				Where(dao.SysComment.Columns.Status, shared.StatusReviewed).
				Page(0, 5).
				All()
			if err != nil {
				return nil, err
			}
			if len(children) > 0 {
				for _, c := range children {
					info, err := s.info(userId, c)
					if err != nil {
						return nil, err
					}
					res = append(res, info)
				}
			}
		}
	}
	return res, nil
}

// Like 点赞
func (s *commentService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.CommentLikeCount+gconv.String(userId)+gconv.String(id),
		shared.CommentLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysComment.
		Fields(dao.SysComment.Columns.UserId, dao.SysComment.Columns.Content).
		Where(dao.SysComment.Columns.CommentId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Comment)
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
	redisCom.Key = shared.Comment + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Comment
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
		_, err = tx.Update(dao.SysComment.Table, g.Map{
			dao.SysComment.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysComment.Columns.CommentId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.CommentInfo
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
		if !Notice.CheckHasNoticeLike(shared.Comment, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Comment
			notice.Content = "点赞了你发布的《" + info.Content + "》评论"
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
			Where(dao.SysUserLike.Columns.Module, shared.Comment).
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
		_, err = tx.Update(dao.SysComment.Table, g.Map{
			dao.SysComment.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysComment.Columns.CommentId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.CommentInfo
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

// Remove 删除评论
func (s *commentService) Remove(userId, id int64) (code response.ResponseCode) {

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
	_, err = tx.Model(dao.SysComment.Table).
		Where(dao.SysComment.Columns.UserId, userId).
		Delete(dao.SysComment.Columns.CommentId, id)
	if err != nil {
		return response.DELETE_FAILED
	}
	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Comment)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, id, shared.Comment)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Comment)
	if err != nil {

		return response.DELETE_FAILED
	}

	//_, err = tx.Update(dao.SysComment.Table,
	//	g.Map{
	//		dao.SysComment.Columns.ParentId: 0,
	//	}, dao.SysComment.Columns.ParentId, id)
	//if err != nil {
	//	return response.DELETE_FAILED
	//}

	return response.SUCCESS
}

// 检查评论状态
func (s *commentService) checkCommentStatus(relatedId, userId int64, module string) bool {
	res, err := dao.SysComment.
		Where(dao.SysComment.Columns.UserId, userId).
		Where(dao.SysComment.Columns.RelatedId, relatedId).
		Where(dao.SysComment.Columns.Module, module).
		All()
	if err != nil {
		return false
	}
	for _, i := range res {
		if i.Status == 2 {
			return true
		}
	}
	return false
}
