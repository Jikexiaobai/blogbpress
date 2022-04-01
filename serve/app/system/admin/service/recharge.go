package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
)

var Recharge = new(rechargeService)

type rechargeService struct {
}

// SelectList 查询列表
func (s *rechargeService) SelectList(req *dto.RechargeQuery) (int, []*result.RechargeList, response.ResponseCode) {
	model := dao.SysRecharge.SysRechargeDao.Order(dao.SysRecharge.Columns.UpdateTime + " desc")

	if req.Status != 0 {
		model = model.Where(dao.SysRecharge.Columns.Status, req.Status)
	}
	if req.Mode != 0 {
		model = model.Where(dao.SysRecharge.Columns.Mode, req.Mode)
	}
	if req.Code != "" {
		model = model.Where(dao.SysRecharge.Columns.Code, req.Code)
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
	var res []*result.RechargeList
	for _, i := range list {
		var info *result.RechargeList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)
		if i.Mode == shared.RechargeModeThree {
			money, err := dao.SysCard.Value(dao.SysCard.Columns.Money, dao.SysCard.Columns.SecretKey, i.CardKey)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			info.Money = gconv.Float64(money)
		}
		res = append(res, info)
	}

	return total, res, response.SUCCESS
}

// Review 更新状态
func (s *rechargeService) Review(req *dto.RechargeReview) (code response.ResponseCode) {
	info, err := dao.SysRecharge.Fields(
		dao.SysRecharge.Columns.UserId,
		dao.SysRecharge.Columns.Mode,
		dao.SysRecharge.Columns.Money,
		dao.SysRecharge.Columns.CardKey,
		dao.SysRecharge.Columns.RechargeId).
		Where(dao.SysRecharge.Columns.Code, req.Code).One()
	if err != nil {
		return response.UPDATE_FAILED
	}

	entity := g.Map{
		dao.SysRecharge.Columns.Status: req.Status,
		dao.SysRecharge.Columns.Remark: req.Remark,
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
	_, err = tx.Update(dao.SysRecharge.Table,
		entity, dao.SysRecharge.Columns.Code, req.Code)
	if err != nil {
		return response.UPDATE_FAILED
	}
	// 设置通知
	var notice model.SysNotice
	notice.Type = shared.NoticeSystem
	notice.SystemType = shared.NoticeUserRecharge
	notice.DetailId = info.RechargeId
	notice.Status = shared.NoticeStatusReview
	notice.Receiver = info.UserId
	notice.CreateTime = gtime.Now()
	if req.Status == shared.StatusReviewed {
		notice.Content = "您发起的充值订单《" + req.Code + "》已通过审核" + req.Remark
		if info.Mode == shared.RechargeModeThree {
			money, err := dao.SysCard.Value(dao.SysCard.Columns.Money, dao.SysCard.Columns.SecretKey, info.CardKey)
			if err != nil {
				return response.DB_READ_ERROR
			}
			// 获取充值用户的余额
			userBalance, err := dao.SysUser.Value(dao.SysUser.Columns.Balance, dao.SysUser.Columns.UserId, info.UserId)
			if err != nil {
				return response.DB_READ_ERROR
			}
			balance := decimal.NewFromFloat(gconv.Float64(userBalance)).Add(decimal.NewFromFloat(gconv.Float64(money)))
			_, err = tx.Update(dao.SysUser.Table, g.Map{
				dao.SysUser.Columns.Balance: gconv.Float64(balance),
			}, dao.SysUser.Columns.UserId, info.UserId)
			if err != nil {
				return response.DB_READ_ERROR
			}

			_, err = tx.Update(dao.SysCard.Table, g.Map{
				dao.SysCard.Columns.Status: shared.StatusReviewed,
			}, dao.SysCard.Columns.SecretKey, info.CardKey)
			if err != nil {
				return response.DB_READ_ERROR
			}
		}

		if info.Mode == shared.RechargeModeFour {
			// 获取充值用户的余额
			userBalance, err := dao.SysUser.Value(dao.SysUser.Columns.Balance, dao.SysUser.Columns.UserId, info.UserId)
			if err != nil {
				return response.DB_READ_ERROR
			}
			balance := decimal.NewFromFloat(gconv.Float64(userBalance)).Add(decimal.NewFromFloat(info.Money))
			_, err = tx.Update(dao.SysUser.Table, g.Map{
				dao.SysUser.Columns.Balance: gconv.Float64(balance),
			}, dao.SysUser.Columns.UserId, info.UserId)
			if err != nil {
				return response.DB_READ_ERROR
			}
		}
	}

	if req.Status == shared.StatusRefuse {
		notice.Content = "您发起的充值订单《" + req.Code + "》未通过审核原因：" + req.Remark
	}
	_, err = tx.Insert(dao.SysNotice.Table, notice)
	if err != nil {
		return response.DB_TX_ERROR
	}

	return response.SUCCESS
}

// Remove 删除
func (s *rechargeService) Remove(idList []int64) response.ResponseCode {
	_, err := dao.SysRecharge.Delete(dao.SysRecharge.Columns.RechargeId+" IN(?)", idList)
	if err != nil {
		return response.DELETE_FAILED
	}
	return response.SUCCESS
}

// RemoveByUser 删除
func (s *rechargeService) RemoveByUser(Ids []int64) error {
	_, err := dao.SysRecharge.Delete(dao.SysRecharge.Columns.UserId+" IN(?)", Ids)
	if err != nil {
		return err
	}
	return nil
}
