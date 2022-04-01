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
	"github.com/gogf/gf/util/gconv"
)

var Group = new(groupService)

type groupService struct{}

// SelectList 查询圈子列表
func (s *groupService) SelectList(req *dto.GroupQuery) (int, []*result.GroupList, response.ResponseCode) {

	model := dao.SysGroup.SysGroupDao.Where(dao.SysGroup.Columns.UserId, req.UserId).
		Where(dao.SysGroup.Columns.DeleteTime, nil).
		Order(dao.SysGroup.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysGroup.Columns.Status, req.Status)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	model = model.Fields(
		dao.SysGroup.Columns.GroupId,
		dao.SysGroup.Columns.Title,
		dao.SysGroup.Columns.Cover,
		dao.SysGroup.Columns.Description,
		dao.SysGroup.Columns.Status,
		dao.SysGroup.Columns.CreateTime,
	)
	var res []*result.GroupList
	err = model.Structs(&res)
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	return total, res, response.SUCCESS
}

// SelectFilterList 获取过滤列表
func (s *groupService) SelectFilterList(req *dto.QueryParam) (int, []*result.GroupListInfo, error) {

	model := dao.SysGroup.SysGroupDao.
		Where(dao.SysGroup.Columns.DeleteTime, nil)
	model = model.Where(dao.SysGroup.Columns.Status, shared.StatusReviewed)
	if req.Title != "" {
		model = model.Where(dao.SysGroup.Columns.Title+" like ?", "%"+req.Title+"%")
	}

	if req.UserId != 0 && !req.IsJoin {
		model = model.Where(dao.SysGroup.Columns.UserId, req.UserId)
	}

	// 获取我加入的圈子
	if req.IsJoin && req.UserId != 0 {
		var ids []int64
		relateIds, err := dao.SysUserJoinGroup.
			Where(dao.SysUserJoinGroup.Columns.UserId, req.UserId).
			All()
		if err != nil {
			return 0, nil, err
		}
		for _, i := range relateIds {
			ids = append(ids, i.GroupId)
		}
		model = model.Where(dao.SysGroup.Columns.GroupId+" IN(?)", ids).
			Or(dao.SysGroup.Columns.UserId, req.UserId)
	}

	switch req.Mode {
	case shared.ModeNew:
		model = model.Order(dao.SysGroup.Columns.CreateTime, "desc")
	case shared.ModeHot:
		model = model.Order(dao.SysGroup.Columns.Hots, "desc")
	default:
		model = model.Order(dao.SysGroup.Columns.UpdateTime, "desc")
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, err
	}
	if req.Limit != 0 {
		model = model.Page(req.Page, req.Limit)
	}
	model = model.Fields(
		dao.SysGroup.Columns.GroupId,
		dao.SysGroup.Columns.UserId,
		dao.SysGroup.Columns.Title,
		dao.SysGroup.Columns.Cover,
		dao.SysGroup.Columns.Description,
		dao.SysGroup.Columns.Joins,
		dao.SysGroup.Columns.Hots,
		dao.SysGroup.Columns.Views,
		dao.SysGroup.Columns.Contents,
		dao.SysGroup.Columns.Status,
		dao.SysGroup.Columns.CreateTime,
		dao.SysGroup.Columns.CateId,
	)
	list, err := model.All()
	if err != nil {
		return 0, nil, err
	}
	var res []*result.GroupListInfo
	for _, i := range list {
		var contentInfo *result.GroupListInfo
		rs, err := s.info(req.UserId, i)
		if err != nil {
			return 0, nil, err
		}
		err = gconv.Struct(rs, &contentInfo)
		if err != nil {
			return 0, nil, err
		}
		contentInfo.Id = i.GroupId
		contentInfo.Module = shared.Group
		res = append(res, contentInfo)
	}

	return total, res, nil
}

// Create 创建小组
func (s *groupService) Create(req *dto.GroupCreate) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.GroupCreateCount+gconv.String(req.UserId),
		shared.GroupCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
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

	var entity model.SysGroup
	err = gconv.Struct(req, &entity)
	if err != nil {
		return response.INVALID
	}

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = shared.StatusReview

	group, err := tx.Insert(dao.SysGroup.Table, entity)
	if err != nil {
		return response.ADD_FAILED
	}
	rid, err := group.LastInsertId()
	if err != nil || rid <= 0 {
		return response.ADD_FAILED
	}
	// 管理媒体库
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, rid, shared.Group)
		if err != nil {
			return response.ADD_FAILED
		}
	}

	return response.SUCCESS
}

// EditInfo 获取编辑信息
func (s *groupService) EditInfo(userId, id int64) (*result.GroupEditInfo, response.ResponseCode) {
	var editInfo *result.GroupEditInfo
	err := dao.SysGroup.
		Where(dao.SysGroup.Columns.GroupId, id).
		Where(dao.SysGroup.Columns.UserId, userId).
		Struct(&editInfo)

	if editInfo == nil || err != nil {
		return nil, response.NOT_FOUND
	}

	return editInfo, response.SUCCESS

}

// Edit 编辑修改圈子
func (s *groupService) Edit(req *dto.GroupEdit) (code response.ResponseCode) {

	// 加入锁限制
	_, err := lock_utils.SetCount(shared.GroupEditCount+gconv.String(req.UserId),
		shared.GroupEditLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
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

	var entity = g.Map{
		dao.SysGroup.Columns.Status:      shared.StatusReview,
		dao.SysGroup.Columns.Cover:       req.Cover,
		dao.SysGroup.Columns.Title:       req.Title,
		dao.SysGroup.Columns.CateId:      req.CateId,
		dao.SysGroup.Columns.Description: req.Description,
		dao.SysGroup.Columns.Price:       req.Price,
		dao.SysGroup.Columns.JoinMode:    req.JoinMode,
		dao.SysGroup.Columns.SecretKey:   req.SecretKey,
		dao.SysGroup.Columns.Icon:        req.Icon,
		dao.SysGroup.Columns.UpdateTime:  gtime.Now(),
	}

	_, err = dao.SysGroup.Update(entity, dao.SysGroup.Columns.GroupId, req.GroupId)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 删除媒体
	err = Media.RemoveRelated(tx, req.GroupId, shared.Group)
	if err != nil {
		return response.UPDATE_FAILED
	}
	//pathList, err := utils.GetSrcLink(req.Content)
	pathList := make([]string, 0)
	pathList = append(pathList, req.Cover)
	pathList = append(pathList, req.Icon)
	if len(pathList) > 0 {
		err = Media.AddRelated(tx, pathList, req.GroupId, shared.Group)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}
	return response.SUCCESS
}

// Remove 删除圈子
func (s *groupService) Remove(userId, id int64) (code response.ResponseCode) {

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

	_, err = tx.Model(dao.SysGroup.Table).
		Where(dao.SysGroup.Columns.UserId, userId).
		Delete(dao.SysGroup.Columns.GroupId, id)
	if err != nil {
		return response.DELETE_FAILED
	}

	err = Media.RemoveRelated(tx, id, shared.Group)
	if err != nil {
		return response.DELETE_FAILED
	}

	return response.SUCCESS
}

// info 圈子信息转换
func (s *groupService) info(userId int64, info *model.SysGroup) (*result.GroupInfo, error) {
	var res *result.GroupInfo
	err := gconv.Struct(info, &res)
	if err != nil {
		return nil, err
	}

	cateInfo, err := Category.SelectInfo(info.CateId, shared.Group)
	if err != nil {
		return nil, err
	}
	res.CateInfo = cateInfo

	userInfo, err := User.SelectInfo(userId, info.UserId)
	if err != nil {
		return nil, err
	}
	res.UserInfo = userInfo

	// 设置查看数量 -----------------
	if userId != 0 {
		res.IsJoin = User.CheckUserGroup(userId, info.GroupId)
	}
	return res, nil
}

// SelectInfo 查询圈子信息
func (s *groupService) SelectInfo(userId, id int64) (*result.GroupInfo, response.ResponseCode) {
	// 修改阅读数
	_, err := dao.SysGroup.Update(g.Map{
		dao.SysGroup.Columns.Views: gdb.Raw("views+1"),
	}, dao.SysGroup.Columns.GroupId, id)
	if err != nil {
		return nil, response.UPDATE_FAILED
	}

	var redisCom redis.Com
	redisCom.Key = shared.Group + shared.InfoById + gconv.String(id)
	// 获取缓存
	infoObj, err := redisCom.GetString()
	if err != nil {
		return nil, response.CACHE_READ_ERROR
	}
	if infoObj != nil {
		var rs *result.GroupInfo
		err := gconv.Struct(infoObj, &rs)
		if err != nil {
			return nil, response.CACHE_READ_ERROR
		}

		if userId != 0 {
			rs.IsJoin = s.CheckIsJoin(userId, id)
		}
		return rs, response.SUCCESS
	}

	info, err := dao.SysGroup.
		Where(dao.SysGroup.Columns.GroupId, id).
		Where(dao.SysGroup.Columns.DeleteTime, nil).
		Where(dao.SysGroup.Columns.Status, shared.StatusReviewed).
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

// Join 加入圈子
func (s *groupService) Join(userId, id int64) (code response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.JoinGroupCount+gconv.String(userId)+gconv.String(id),
		shared.JoinGroupLock+gconv.String(userId)+gconv.String(id), 60, 5)
	if err != nil {
		return response.CACHE_SAVE_ERROR
	}
	count, err := dao.SysUserJoinGroup.
		Where(dao.SysUserJoinGroup.Columns.UserId, userId).
		Where(dao.SysUserJoinGroup.Columns.GroupId, id).
		Count()
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
	redisCom.Key = shared.Group + shared.InfoById + gconv.String(id)
	infoObj, err := redisCom.GetString()
	if err != nil {
		return response.CACHE_READ_ERROR
	}
	// 获取用户是否加入
	if count == 0 {
		// 写入据库
		_, err := tx.Save(dao.SysUserJoinGroup.Table, g.Map{
			dao.SysUserJoinGroup.Columns.GroupId: id,
			dao.SysUserJoinGroup.Columns.UserId:  userId,
		})
		if err != nil {
			return response.ADD_FAILED
		}

		// 更新
		_, err = tx.Update(dao.SysGroup.Table, g.Map{
			dao.SysGroup.Columns.Joins: gdb.Raw("joins+1"),
		}, dao.SysGroup.Columns.GroupId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}
		// 修改缓存的数据
		if infoObj != nil {
			var rs *result.GroupInfo
			err := gconv.Struct(infoObj, &rs)
			if err != nil {
				return response.CACHE_READ_ERROR
			}
			rs.Joins += 1
			redisCom.Time = 600
			redisCom.Data = rs
			err = redisCom.SetStringEX()
			if err != nil {
				return response.CACHE_SAVE_ERROR
			}
		}

	} else {

		_, err := tx.Model(dao.SysUserJoinGroup.Table).
			Where(dao.SysUserJoinGroup.Columns.UserId, userId).
			Where(dao.SysUserJoinGroup.Columns.GroupId, id).
			Delete()

		// 更新
		_, err = tx.Update(dao.SysGroup.Table, g.Map{
			dao.SysGroup.Columns.Joins: gdb.Raw("joins-1"),
		}, dao.SysGroup.Columns.GroupId, id)
		if err != nil {
			return response.UPDATE_FAILED
		}

		if infoObj != nil {
			var rs *result.GroupInfo
			err := gconv.Struct(infoObj, &rs)
			if err != nil {
				return response.CACHE_READ_ERROR
			}
			rs.Joins -= 1
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

// CheckCanJoin 检查是否能加入圈子
func (s *groupService) CheckCanJoin(userId, id int64, secretKey string) bool {
	info, err := dao.SysGroup.
		Fields(
			dao.SysGroup.Columns.JoinMode,
			dao.SysGroup.Columns.SecretKey,
		).
		Where(dao.SysGroup.Columns.GroupId, id).
		One()

	if err != nil || info == nil {
		return false
	}

	// 免费加入
	if info.JoinMode == 1 {
		return true
	}
	// 付费加入
	if info.JoinMode == 2 {
		return Order.CheckIsPay(userId, id, shared.Group, shared.OrderTypeFive)
	}
	//密钥加入
	if info.JoinMode == 3 {
		return secretKey == info.SecretKey
	}

	return false
}

// CheckIsJoin 检查是否加入了
func (s *groupService) CheckIsJoin(userId, id int64) bool {
	count, err := dao.SysUserJoinGroup.
		Where(dao.SysUserJoinGroup.Columns.UserId, userId).
		Where(dao.SysUserJoinGroup.Columns.GroupId, id).Count()
	if err == nil && count > 0 {
		return true
	}

	return false
}

// MyJoinGroupList 我加入的小组
func (s *groupService) MyJoinGroupList(userId int64) ([]*result.GroupJoin, error) {

	// 获取我加入小组的信息
	var group1 []*result.GroupJoin
	var group2 []*result.GroupJoin
	err := dao.SysGroup.Where(dao.SysGroup.Columns.UserId, userId).
		Where(dao.SysGroup.Columns.Status, 2).
		Fields(dao.SysGroup.Columns.GroupId, dao.SysGroup.Columns.Title).
		Structs(&group1)
	if err != nil {
		return nil, err
	}

	err = dao.SysGroup.As("a").
		LeftJoin("sys_user_join_group as b", "b.group_id = a.group_id").
		Where("b.user_id", userId).And("a.status", 2).
		Fields("a.group_id", "a.title").
		Structs(&group2)
	if err != nil {
		return nil, err
	}
	group1 = append(group1, group2...)

	return group1, nil
}

// SelectRelatedGroup 获取关联圈子
func (s *groupService) SelectRelatedGroup(id int64) (*result.GroupJoin, error) {
	var group *result.GroupJoin
	err := dao.SysGroup.Where(dao.SysGroup.Columns.GroupId, id).Struct(&group)
	if err != nil {
		return nil, err
	}
	return group, nil
}
