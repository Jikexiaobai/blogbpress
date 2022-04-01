package valid

import (
	"fiber/app/dao"
	"fiber/app/system/admin/shared"
	"github.com/gogf/gf/util/gconv"
)

var Recharge = new(rechargeValid)

type rechargeValid struct {
}

// CheckMode 检查类型
func (s *rechargeValid) CheckMode(code string) bool {
	mode, err := dao.SysRecharge.
		Where(dao.SysRecharge.Columns.Status, shared.StatusReview).
		Value(dao.SysRecharge.Columns.Mode, dao.SysRecharge.Columns.Code, code)
	if err == nil && (gconv.Int(mode) == 3 || gconv.Int(mode) == 4) {
		return true
	}
	return false
}
