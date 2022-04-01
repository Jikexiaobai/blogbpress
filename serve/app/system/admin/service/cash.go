package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
)

var Cash = new(cashService)

type cashService struct {
}

// SelectList 查询列表
func (s *cashService) SelectList(req *dto.CashQuery) (int, []*result.CashList, response.ResponseCode) {
	model := dao.SysCash.SysCashDao.Order(dao.SysCash.Columns.UpdateTime + " desc")
	if req.Status != 0 {
		model = model.Where(dao.SysCash.Columns.Status, req.Status)
	}

	if req.Code != "" {
		model = model.Where(dao.SysCash.Columns.Code, req.Code)
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
	var res []*result.CashList
	for _, i := range list {
		var info *result.CashList
		err = gconv.Struct(i, &info)
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Review 更新状态
func (s *cashService) Review(req *dto.CashReview) (code response.ResponseCode) {
	info, err := dao.SysCash.Fields(
		dao.SysCash.Columns.UserId,
		dao.SysCash.Columns.Money,
		dao.SysCash.Columns.CashId).
		Where(dao.SysCash.Columns.Code, req.Code).One()
	if err != nil {
		return response.UPDATE_FAILED
	}
	entity := g.Map{
		dao.SysCash.Columns.Status:     req.Status,
		dao.SysCash.Columns.Remark:     req.Remark,
		dao.SysCash.Columns.ReceiptNum: req.ReceiptNum,
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
	_, err = tx.Update(dao.SysCash.Table,
		entity, dao.SysCash.Columns.Code, req.Code)
	if err != nil {
		return response.UPDATE_FAILED
	}

	// 设置通知
	var notice model.SysNotice
	notice.Type = shared.NoticeSystem
	notice.SystemType = shared.NoticeUserCash
	notice.DetailId = info.CashId
	notice.Status = shared.NoticeStatusReview
	notice.Receiver = info.UserId
	notice.CreateTime = gtime.Now()

	if req.Status == shared.StatusReviewed {
		notice.Content = "您申请的提现订单《" + req.Code + "》已通过审核" + req.Remark
		// 获取充值用户的余额
		userBalance, err := dao.SysUser.Value(dao.SysUser.Columns.Balance, dao.SysUser.Columns.UserId, info.UserId)
		if err != nil {
			return response.DB_READ_ERROR
		}

		balance := decimal.NewFromFloat(gconv.Float64(userBalance)).Sub(decimal.NewFromFloat(info.Money))
		_, err = tx.Update(dao.SysUser.Table, g.Map{
			dao.SysUser.Columns.Balance: gconv.Float64(balance),
		}, dao.SysUser.Columns.UserId, info.UserId)
		if err != nil {
			return response.DB_READ_ERROR
		}
	}
	if req.Status == shared.StatusRefuse {
		notice.Content = "您申请的提现订单《" + req.Code + "》未通过审核原因：" + req.Remark
	}
	_, err = tx.Insert(dao.SysNotice.Table, notice)
	if err != nil {
		return response.DB_TX_ERROR
	}
	return response.SUCCESS
}

// Remove 删除
func (s *cashService) Remove(idList []int64) response.ResponseCode {
	_, err := dao.SysCash.Delete(dao.SysCash.Columns.CashId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *cashService) RemoveByUser(tx *gdb.TX, Ids []int64) error {
	_, err := tx.Delete(dao.SysCash.Table, dao.SysCash.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}
	return nil
}
