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

var Answer = new(answerService)

type answerService struct {
}

// info 详细信息数据转换
func (s *answerService) info(userId int64, info *model.SysAnswer) (*result.AnswerInfo, error) {
	var res result.AnswerInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	userInfo, err := User.SelectInfo(userId, info.UserId)
	if err != nil {
		return nil, err
	}
	res.UserInfo = userInfo

	if userId != 0 {
		if User.CheckUserLike(userId, info.AnswerId, shared.Answer) {
			res.IsLike = true
		}
		if info.IsAdoption == 1 {
			res.IsAdoption = false
		}
		if info.IsAdoption == 2 {
			res.IsAdoption = true
		}
	}

	return &res, nil
}

// SelectList 获取列表
func (s *answerService) SelectList(userId int64, req *dto.AnswerQuery) (int, []*result.AnswerInfo, response.ResponseCode) {
	model := dao.SysAnswer.SysAnswerDao.
		Where(dao.SysAnswer.Columns.DeleteTime, nil).
		Where(dao.SysAnswer.Columns.Status, shared.StatusReviewed)

	model = model.Where(dao.SysAnswer.Columns.TopicId, req.TopicId)

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Order(dao.SysComment.Columns.UpdateTime + " desc")
	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.AnswerInfo
	for _, i := range list {
		info, err := s.info(userId, i)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Create 创建
func (s *answerService) Create(req *dto.AnswerCreate) (answerRes *result.AnswerInfo, code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.AnswerCreateCount+gconv.String(req.UserId),
		shared.AnswerCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	var entity model.SysAnswer
	err = gconv.Struct(req, &entity)
	if err != nil {
		return nil, response.INVALID
	}
	entity.UserId = req.UserId
	entity.IsAdoption = 1
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReviewed

	// 开启事务
	tx, err := g.DB().Begin()
	if err != nil {
		return nil, response.DB_TX_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	res, err := tx.Insert(dao.SysAnswer.Table, entity)
	if err != nil {
		_ = tx.Rollback()
		return nil, response.ADD_FAILED
	}
	rid, err := res.LastInsertId()
	if err != nil || rid <= 0 {
		return nil, response.ADD_FAILED
	}

	files := gjson.New(req.Files)
	filesList := gconv.Strings(files.Value())
	pathList := gconv.Strings(filesList)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Answer)
		if err != nil {
			return nil, response.ADD_FAILED
		}
	}

	var redisCom redis.Com
	//	设置用户活跃
	redisCom.Key = shared.UserHot
	redisCom.Data = req.UserId
	err = redisCom.ADDSet()
	if err != nil {
		return nil, response.CACHE_SAVE_ERROR
	}

	//通知问题作者
	var notice model.SysNotice

	// 通知问题作者
	notice.Type = shared.NoticeAnswer
	topic, err := dao.SysTopic.
		Fields(dao.SysTopic.Columns.UserId, dao.SysTopic.Columns.Title).
		Where(dao.SysTopic.Columns.TopicId, req.TopicId).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	notice.Receiver = topic.UserId
	notice.Content = req.Content
	notice.DetailId = rid
	notice.DetailModule = shared.Answer
	notice.Status = shared.NoticeStatusReview
	notice.CreateTime = gtime.Now()
	_, err = tx.Insert(dao.SysNotice.Table, notice)
	if err != nil {
		return nil, response.DB_SAVE_ERROR
	}

	err = Integral.SetUserAnswerIntegral(redisCom, tx, req.UserId)
	if err != nil {
		return nil, response.DB_SAVE_ERROR
	}

	err = gconv.Struct(entity, &answerRes)
	if err != nil {
		return nil, response.INVALID
	}
	userInfo, err := User.SelectInfo(req.UserId, req.UserId)
	if err != nil {
		return nil, response.ADD_FAILED
	}
	answerRes.UserInfo = userInfo
	answerRes.AnswerId = rid

	return answerRes, response.SUCCESS
}

// Like 点赞
func (s *answerService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.AnswerLikeCount+gconv.String(userId)+gconv.String(id),
		shared.AnswerLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysAnswer.
		Fields(dao.SysAnswer.Columns.UserId, dao.SysAnswer.Columns.Content).
		Where(dao.SysAnswer.Columns.AnswerId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Answer)
	if err != nil {
		return response.DB_READ_ERROR
	}

	// 开启事务
	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var redisCom redis.Com
	redisCom.Key = shared.Answer + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Answer
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
		_, err = tx.Update(dao.SysAnswer.Table, g.Map{
			dao.SysAnswer.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysAnswer.Columns.AnswerId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.AnswerInfo
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
		if !Notice.CheckHasNoticeLike(shared.Answer, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Answer
			notice.Content = "点赞了你发布的《" + info.Content + "》答案"
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
			Where(dao.SysUserLike.Columns.Module, shared.Answer).
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
		_, err = tx.Update(dao.SysAnswer.Table, g.Map{
			dao.SysAnswer.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysAnswer.Columns.AnswerId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.AnswerInfo
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

// Adoption 采纳
func (s *answerService) Adoption(userId, topicId, answerId int64) (code response.ResponseCode) {

	_, err := dao.SysAnswer.
		Where(dao.SysAnswer.Columns.AnswerId, answerId).
		Where(dao.SysAnswer.Columns.TopicId, topicId).Update(g.Map{
		dao.SysAnswer.Columns.IsAdoption: 2,
	})
	if err != nil {
		return response.DB_SAVE_ERROR
	}
	var redisCom redis.Com
	redisCom.Key = shared.UserHot
	redisCom.Data = userId
	err = redisCom.ADDSet()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	return response.SUCCESS
}

// Remove 删除
func (s *answerService) Remove(userId, id int64) (code response.ResponseCode) {
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

	_, err = tx.Model(dao.SysAnswer.Table).
		Where(dao.SysAnswer.Columns.UserId, userId).
		Delete(dao.SysAnswer.Columns.AnswerId, id)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Answer)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Answer)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}
