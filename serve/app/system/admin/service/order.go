package service

import (
	"fiber/app/dao"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

var Order = new(orderService)

type orderService struct {
}

// SelectList 查询列表
func (s *orderService) SelectList(req *dto.OrderQuery) (int, []*result.OrderList, response.ResponseCode) {
	model := dao.SysOrder.SysOrderDao.Order(dao.SysOrder.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysOrder.Columns.Status, req.Status)
	}
	if req.OrderType != 0 {
		model = model.Where(dao.SysOrder.Columns.OrderType, req.OrderType)
	}
	if req.OrderNum != "" {
		model = model.Where(dao.SysOrder.Columns.OrderNum, req.OrderNum)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysOrder.Columns.OrderId,
		dao.SysOrder.Columns.UserId,
		dao.SysOrder.Columns.OrderNum,
		dao.SysOrder.Columns.OrderType,
		dao.SysOrder.Columns.OrderMoney,
		dao.SysOrder.Columns.PaymentMoney,
		dao.SysOrder.Columns.PayMethod,
		dao.SysOrder.Columns.Status,
		dao.SysOrder.Columns.CreateTime,
		dao.SysOrder.Columns.PayTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.OrderList
	for _, i := range list {
		var info *result.OrderList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
			dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)
		res = append(res, info)
	}

	return total, res, response.SUCCESS
}

// SelectInfo 查询详细信息
func (s *orderService) SelectInfo(id int64) (*result.OrderInfo, response.ResponseCode) {

	order, err := dao.SysOrder.Where(dao.SysOrder.Columns.OrderId, id).One()
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	var info *result.OrderInfo
	err = gconv.Struct(order, &info)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
		dao.SysUser.Columns.UserId, order.UserId)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	info.NickName = gconv.String(nickName)

	authorName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
		dao.SysUser.Columns.UserId, order.AuthorId)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	info.AuthorName = gconv.String(authorName)
	switch order.OrderType {
	case shared.OrderTypeFour:
		title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title,
			dao.SysTopic.Columns.TopicId, order.DetailId)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		info.Title = gconv.String(title)
	case shared.OrderTypeThree:
		switch order.DetailModule {
		case shared.Resource:
			title, err := dao.SysResource.Value(dao.SysResource.Columns.Title,
				dao.SysResource.Columns.ResourceId, order.DetailId)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			info.Title = gconv.String(title)
		case shared.Audio:
			title, err := dao.SysAudio.Value(dao.SysAudio.Columns.Title,
				dao.SysAudio.Columns.AudioId, order.DetailId)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			info.Title = gconv.String(title)
		case shared.Video:
			title, err := dao.SysVideo.Value(dao.SysVideo.Columns.Title,
				dao.SysVideo.Columns.VideoId, order.DetailId)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			info.Title = gconv.String(title)
		}
	case shared.OrderTypeFive:
		title, err := dao.SysGroup.Value(dao.SysGroup.Columns.Title,
			dao.SysGroup.Columns.GroupId, order.DetailId)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		info.Title = gconv.String(title)
	case shared.OrderTypeSix:
		title, err := dao.SysEdu.Value(dao.SysEdu.Columns.Title,
			dao.SysEdu.Columns.EduId, order.DetailId)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		info.Title = gconv.String(title)
	case shared.OrderTypeSeven:
		questionId, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.TopicId,
			dao.SysAnswer.Columns.AnswerId, order.DetailId)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		title, err := dao.SysQuestion.Value(dao.SysQuestion.Columns.Title,
			dao.SysQuestion.Columns.QuestionId, questionId)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		info.Title = gconv.String(title)
	}
	return info, response.SUCCESS
}

// Remove 删除
func (s *orderService) Remove(idList []int64) response.ResponseCode {
	_, err := dao.SysOrder.Delete(dao.SysOrder.Columns.OrderId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *orderService) RemoveByUser(tx *gdb.TX, Ids []int64) error {
	_, err := tx.Delete(dao.SysOrder.Table, dao.SysOrder.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}
	return nil
}
