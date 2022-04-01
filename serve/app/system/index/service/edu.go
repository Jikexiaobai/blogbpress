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
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Edu = new(eduService)

type eduService struct {
}

// SelectList 查询列表
func (s *eduService) SelectList(req *dto.EduQuery) (int, []*result.EduList, response.ResponseCode) {

	model := dao.SysEdu.SysEduDao.Where(dao.SysEdu.Columns.UserId, req.UserId).
		Where(dao.SysEdu.Columns.DeleteTime, nil).
		Order(dao.SysEdu.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysEdu.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysEdu.Columns.EduId,
		dao.SysEdu.Columns.Title,
		dao.SysEdu.Columns.Cover,
		dao.SysEdu.Columns.Description,
		dao.SysEdu.Columns.Status,
		dao.SysEdu.Columns.CreateTime,
	)
	var res []*result.EduList
	err = model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return total, res, response.SUCCESS
}

// Create 创建课程
func (s *eduService) Create(req *dto.EduCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.EduCreateCount+gconv.String(req.UserId),
		shared.EduCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysEdu
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

	rs, err := tx.Insert(dao.SysEdu.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}

	// 增加 关联标签
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, rid, shared.Edu)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	// 管理媒体库
	pathList, err := regex.GetStringLink(req.Section)
	contentPath, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, contentPath...)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Edu)
		if err != nil {
			return response.ADD_FAILED
		}
	}
	// 写入到es

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *eduService) EditInfo(userId, id int64) (*result.EduEditInfo, response.ResponseCode) {

	var editInfo *result.EduEditInfo

	err := dao.SysEdu.
		Where(dao.SysEdu.Columns.EduId, id).
		Where(dao.SysEdu.Columns.UserId, userId).
		Struct(&editInfo)
	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}
	// 获取标签
	tagList, err := Tag.SelectRelatedList(id, shared.Edu)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	editInfo.TagList = tagList

	return editInfo, response.SUCCESS

}

// Edit 编辑课程
func (s *eduService) Edit(req *dto.EduEdit) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.EduEditCount+gconv.String(req.UserId),
		shared.EduEditLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity = g.Map{
		dao.SysEdu.Columns.Status:      shared.StatusReview,
		dao.SysEdu.Columns.Cover:       req.Cover,
		dao.SysEdu.Columns.Title:       req.Title,
		dao.SysEdu.Columns.Max:         req.Max,
		dao.SysEdu.Columns.Content:     req.Content,
		dao.SysEdu.Columns.Section:     req.Section,
		dao.SysEdu.Columns.CateId:      req.CateId,
		dao.SysEdu.Columns.Description: req.Description,
		dao.SysEdu.Columns.Price:       req.Price,
		dao.SysEdu.Columns.UpdateTime:  gtime.Now(),
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
	_, err = tx.Update(dao.SysEdu.Table, entity, dao.SysEdu.Columns.EduId, req.EduId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 增加 关联标签
	err = Tag.RemoveRelated(tx, req.EduId, shared.Edu)
	if err != nil {
		return response.UPDATE_FAILED
	}
	if len(req.Tags) > 0 {
		err = Tag.AddTags(tx, req.Tags, req.EduId, shared.Edu)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.EduId, shared.Edu)
	if err != nil {
		return response.UPDATE_FAILED
	}
	pathList, err := regex.GetStringLink(req.Section)
	contentPath, err := regex.GetSrcLink(req.Content)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, contentPath...)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.EduId, shared.Edu)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}
	return response.SUCCESS
}

// Remove 删除
func (s *eduService) Remove(userId, id int64) (code response.ResponseCode) {

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

	_, err = tx.Model(dao.SysEdu.Table).Where(dao.SysEdu.Columns.UserId,
		userId).Delete(dao.SysEdu.Columns.EduId, id)
	if err != nil {
		return response.DELETE_FAILED
	}

	//// 删除 关联圈子
	//err = Group.RemoveRelatedList(id, shared.Edu)
	//if err != nil {
	//	return response.DELETE_FAILED
	//}

	// 删除 关联标签
	err = Tag.RemoveRelated(tx, id, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	// 删除用户收藏和点赞
	err = User.RemoveUserLike(tx, id, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = User.RemoveUserFavorite(tx, id, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Edu)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// SelectFilterList 查询过滤列表
func (s *eduService) SelectFilterList(req *dto.QueryParam) (int, []*result.EduListInfo, error) {
	model := dao.SysEdu.SysEduDao.Where(dao.SysEdu.Columns.DeleteTime, nil)
	model = model.Where(dao.SysEdu.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysEdu.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.CateId != 0 {
		model = model.Where(dao.SysEdu.Columns.CateId, req.CateId)
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
			if i.Module == shared.Edu {
				relatedIds = append(relatedIds, i.RelatedId)
			}
		}
		model = model.Where(dao.SysEdu.Columns.EduId+" IN(?)", relatedIds)
	}

	if req.IsJoin && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserJoinEdu.
			Where(dao.SysUserJoinEdu.Columns.UserId, req.UserId).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.EduId)
		}
		model = model.Where(dao.SysEdu.Columns.EduId+" IN(?)", ids)
	}

	if req.IsFavorite && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserFavorite.
			Where(dao.SysUserFavorite.Columns.UserId, req.UserId).
			Where(dao.SysUserFavorite.Columns.Module, shared.Edu).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.FavoriteId)
		}
		model = model.Where(dao.SysEdu.Columns.EduId+" IN(?)", ids)
	}

	if req.UserId != 0 && !req.IsJoin && !req.IsFavorite {
		model = model.Where(dao.SysEdu.Columns.UserId, req.UserId)
	}

	if req.Type != 0 {
		model = model.Where(dao.SysEdu.Columns.Type, req.Type)
	}

	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysEdu.Columns.CreateTime, "desc")
	case shared.ModeHot:
		model = model.Order(dao.SysEdu.Columns.Hots, "desc")
	default:
		model = model.Order(dao.SysEdu.Columns.UpdateTime, "desc")
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
	var res []*result.EduListInfo
	for _, i := range list {
		var contentInfo *result.EduListInfo
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.Id = i.EduId
		contentInfo.Module = "course"
		res = append(res, contentInfo)
	}

	return total, res, nil
}

// SelectByHomeList 查询首页列表
func (s *eduService) SelectByHomeList(ids string) ([]*result.EduListInfo, error) {
	idList := gstr.Split(ids, ",")
	list, err := dao.SysEdu.Where(dao.SysEdu.Columns.EduId+" IN(?)", idList).All()
	if err != nil {
		return nil, err
	}
	var res []*result.EduListInfo
	for _, i := range list {
		var contentInfo *result.EduListInfo
		rs, err := s.info(0, i)
		if err != nil {
			return nil, err
		}
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return nil, err
		}
		contentInfo.Id = i.EduId
		contentInfo.Module = "course"
		res = append(res, contentInfo)
	}
	return res, nil
}

func (s *eduService) info(userId int64, info *model.SysEdu) (*result.EduInfo, error) {
	var res result.EduInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	cateInfo, err := Category.SelectInfo(info.CateId, shared.Edu)
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
	tagList, err := Tag.SelectRelatedList(info.EduId, shared.Edu)
	if err != nil {
		return nil, err
	}
	res.TagList = tagList

	// 获取评论数量
	comments, err := dao.SysComment.
		Where(dao.SysComment.Columns.RelatedId, info.EduId).
		Where(dao.SysComment.Columns.Module, shared.Edu).Count()
	if err != nil {
		return nil, err
	}
	res.Comments = gconv.Int64(comments)

	if userId != 0 {
		// 获取是否支付
		res.IsPay = Order.CheckIsPay(userId, info.EduId, shared.Edu, shared.OrderTypeSix)

		// 检查是否报名
		res.IsJoin = !s.CheckIsJoin(userId, info.EduId)

		res.IsFavorite = User.CheckUserFavorite(userId, info.EduId, shared.Edu)

		res.IsLike = User.CheckUserLike(userId, info.EduId, shared.Edu)
	}

	if !res.IsJoin {

		var links []*result.SectionInfo
		j, err := gjson.DecodeToJson(res.Section)
		if err != nil {
			return nil, err
		}
		sections := j.Value()
		err = gconv.Structs(sections, &links)
		if err != nil {
			return nil, err
		}
		for _, i := range links {
			for _, child := range i.Children {
				if !child.IsWatch {
					child.Link = ""
				}
			}
		}
		jLinks, err := gjson.Encode(links)
		if err != nil {
			return nil, err
		}
		res.Section = gconv.String(jLinks)
	}

	return &res, nil
}

// SelectInfo 查询文章信息
func (s *eduService) SelectInfo(userId, id int64) (*result.EduInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysEdu.Update(g.Map{
		dao.SysEdu.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysEdu.Columns.EduId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Edu + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.EduInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if userId != 0 {
			// 获取是否支付
			rs.IsPay = Order.CheckIsPay(userId, rs.EduId, shared.Edu, shared.OrderTypeSix)

			// 检查是否报名
			rs.IsJoin = !s.CheckIsJoin(userId, rs.EduId)

			rs.IsFavorite = User.CheckUserFavorite(userId, rs.EduId, shared.Edu)

			rs.IsLike = User.CheckUserLike(userId, rs.EduId, shared.Edu)
		}
		if !rs.IsJoin {

			var links []*result.SectionInfo
			j, err := gjson.DecodeToJson(rs.Section)
			if err != nil {
				return nil, response.FAILD
			}
			sections := j.Value()
			err = gconv.Structs(sections, &links)
			if err != nil {
				return nil, response.FAILD
			}
			for _, i := range links {
				for _, child := range i.Children {
					if !child.IsWatch {
						child.Link = ""
					}
				}
			}
			jLinks, err := gjson.Encode(links)
			if err != nil {
				return nil, response.FAILD
			}
			rs.Section = gconv.String(jLinks)
		}

		return rs, response.SUCCESS
	}

	info, err := dao.SysEdu.
		Where(dao.SysEdu.Columns.EduId, id).
		Where(dao.SysEdu.Columns.DeleteTime, nil).
		Where(dao.SysEdu.Columns.Status, shared.StatusReviewed).
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

func (s *eduService) Join(req *dto.EduJoinCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.JoinEduCount+gconv.String(req.UserId)+gconv.String(req.EduId),
		shared.JoinEduLock+gconv.String(req.UserId)+gconv.String(req.EduId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	var entity model.SysUserJoinEdu
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}
	entity.CreateTime = gtime.Now()

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

	_, err = tx.Insert(dao.SysUserJoinEdu.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}

	_, err = tx.Model(dao.SysEdu.Table).Update(g.Map{
		dao.SysEdu.Columns.Joins: gdb.Raw(dao.SysEdu.Columns.Joins + "+1"),
	}, dao.SysEdu.Columns.EduId, req.EduId)
	if err != nil {
		return response.ADD_FAILED
	}

	var redisCom redis.Com
	//	设置用户活跃度
	redisCom.Key = shared.UserHot
	redisCom.Data = req.UserId
	err = redisCom.ADDSet()
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}
	return response.SUCCESS
}

// Like 点赞
func (s *eduService) Like(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.EduLikeCount+gconv.String(userId)+gconv.String(id),
		shared.EduLikeLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	// 获取作者id
	info, err := dao.SysEdu.
		Fields(dao.SysEdu.Columns.UserId, dao.SysEdu.Columns.Title).
		Where(dao.SysEdu.Columns.EduId, id).One()
	if err != nil {
		return response.DB_READ_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserLike.
		Where(dao.SysUserLike.Columns.UserId, userId).
		Where(dao.SysUserLike.Columns.RelatedId, id).
		Count(dao.SysUserLike.Columns.Module, shared.Edu)
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
	redisCom.Key = shared.Edu + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	if count == 0 {
		// 写入点赞数据库
		var entity model.SysUserLike
		entity.UserId = userId
		entity.RelatedId = id
		entity.Module = shared.Edu
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
		_, err = tx.Update(dao.SysEdu.Table, g.Map{
			dao.SysEdu.Columns.Likes: gdb.Raw("likes+1"),
		}, dao.SysEdu.Columns.EduId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.EduInfo
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
		if !Notice.CheckHasNoticeLike(shared.Edu, id) {
			// 通知被点赞用户
			var notice model.SysNotice
			notice.Type = shared.NoticeLike
			notice.FromUserId = userId
			notice.DetailId = id
			notice.DetailModule = shared.Edu
			notice.Content = "点赞了你发布的《" + info.Title + "》课程"
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
			Where(dao.SysUserLike.Columns.Module, shared.Edu).
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
		_, err = tx.Update(dao.SysEdu.Table, g.Map{
			dao.SysEdu.Columns.Likes: gdb.Raw("likes-1"),
		}, dao.SysEdu.Columns.EduId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.EduInfo
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
func (s *eduService) Favorite(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.EduFavoriteCount+gconv.String(userId)+gconv.String(id),
		shared.EduFavoriteLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}

	//判断是否点赞
	count, err := dao.SysUserFavorite.
		Where(dao.SysUserFavorite.Columns.UserId, userId).
		Where(dao.SysUserFavorite.Columns.FavoriteId, id).
		Count(dao.SysUserFavorite.Columns.Module, shared.Edu)
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
	redisCom.Key = shared.Edu + shared.InfoById + gconv.String(id)
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
		entity.Module = shared.Edu
		_, err := tx.Insert(dao.SysUserFavorite.Table, entity)
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysEdu.Table, g.Map{
			dao.SysEdu.Columns.Favorites: gdb.Raw("favorites+1"),
		}, dao.SysEdu.Columns.EduId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.EduInfo
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
			Where(dao.SysUserFavorite.Columns.Module, shared.Edu).
			Delete()
		if err != nil {
			return response.DELETE_FAILED
		}

		// 更新文章点赞
		_, err = tx.Update(dao.SysEdu.Table, g.Map{
			dao.SysEdu.Columns.Favorites: gdb.Raw("favorites-1"),
		}, dao.SysEdu.Columns.EduId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		// 修改缓存的点赞数据
		if infoObj != nil {
			var rs *result.EduInfo
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

func (s *eduService) SelectUserJointList(req *dto.EduJoinQueryParam) (int, []*result.EduUserJoinListInfo, error) {

	model := dao.SysUserJoinEdu.SysUserJoinEduDao.
		Where(dao.SysUserJoinEdu.Columns.EduId, req.EduId)
	if req.Name != "" {
		model = model.Where(dao.SysUserJoinEdu.Columns.Name, req.Name)
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
	var res []*result.EduUserJoinListInfo
	for _, i := range list {
		var contentInfo *result.EduUserJoinListInfo
		err = gconv.Struct(i, &contentInfo)
		if err != nil {
			return 0, nil, err
		}
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.NickName = gconv.String(nickName)
		res = append(res, contentInfo)
	}

	return total, res, nil
}

// CheckJoin 检查是否可以报名
func (s *eduService) CheckJoin(id int64) bool {

	info, err := dao.SysEdu.Fields(
		dao.SysEdu.Columns.Joins,
		dao.SysEdu.Columns.Max,
		dao.SysEdu.Columns.Type,
	).Where(dao.SysEdu.Columns.EduId, id).One()
	if err != nil {
		return false
	}
	if info.Max == 0 && info.Type == 2 {
		return true
	}
	if info.Joins <= gconv.Int64(info.Max) {
		return true
	}

	return false
}

// CheckIsJoin 检查是否报名了
func (s *eduService) CheckIsJoin(userId, id int64) bool {
	count, err := dao.SysUserJoinEdu.Where(dao.SysEdu.Columns.UserId, userId).
		Where(dao.SysEdu.Columns.EduId, id).Count()
	if err != nil || count > 0 {
		return false
	}
	return true
}

// CheckIsPay 检查是否需要支付
func (s *eduService) CheckIsPay(userId, id int64) bool {
	info, err := dao.SysEdu.Fields(
		dao.SysEdu.Columns.JoinMode,
	).
		Where(dao.SysEdu.Columns.UserId, userId).
		Where(dao.SysEdu.Columns.EduId, id).One()
	if err != nil {
		return false
	}
	if info.JoinMode == 1 {
		return Order.CheckIsPay(userId, id, shared.Edu, shared.OrderTypeSix)
	}

	return true
}
