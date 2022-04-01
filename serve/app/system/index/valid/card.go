package valid

import (
	"fiber/app/dao"
	"github.com/gogf/gf/util/gconv"
)

var Card = new(cardValid)

type cardValid struct {
}

// CheckIsUse 检查订单是否支付
func (s *cardValid) CheckIsUse(secretKey string) bool {
	res, err := dao.SysCard.Value(dao.SysCard.Columns.Status,
		dao.SysCard.Columns.SecretKey, secretKey)
	if err == nil && gconv.Int(res) == 2 {
		return true
	}
	return false
}
