package valid

import (
	"fiber/app/dao"
	"fiber/app/system/admin/shared"
)

var Cash = new(cashValid)

type cashValid struct {
}

// CheckIsCash 检查提现订单是否已经提现
func (s *cashValid) CheckIsCash(code string) bool {
	count, err := dao.SysCash.Where(dao.SysCash.Columns.Code, code).
		Where(dao.SysCash.Columns.Status, shared.StatusReviewed).Count()
	if err == nil && count > 0 {
		return true
	}
	return false
}
