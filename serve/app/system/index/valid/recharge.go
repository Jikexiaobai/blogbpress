package valid

import (
	"fiber/app/dao"
	"github.com/gogf/gf/util/gconv"
)

var Recharge = new(rechargeValid)

type rechargeValid struct {
}

// CheckPay 检查订单是否支付
func (s *rechargeValid) CheckPay(userId int64, code string) bool {
	res, err := dao.SysRecharge.
		Where(dao.SysRecharge.Columns.UserId, userId).
		Value(dao.SysRecharge.Columns.Status,
			dao.SysRecharge.Columns.Code, code)
	if err == nil && gconv.Int(res) == 2 {
		return true
	}
	return false
}

// CheckRechargeModeOneOrTwo 检查订单是否为支付宝或微信
func (s *rechargeValid) CheckRechargeModeOneOrTwo(userId int64, code string) bool {
	res, err := dao.SysRecharge.
		Where(dao.SysRecharge.Columns.UserId, userId).
		Value(dao.SysRecharge.Columns.Type,
			dao.SysRecharge.Columns.Code, code)
	if err == nil && (gconv.Int(res) == 1 || gconv.Int(res) == 2) {
		return true
	}
	return false
}
