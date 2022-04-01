package valid

import (
	"fiber/app/system/index/shared"
	"fiber/library/redis"
	"github.com/gogf/gf/util/gconv"
)

var Account = new(accountService)

type accountService struct {
}

// CheckSign 检查用户是否签到
func (s *accountService) CheckSign(userId int64) bool {
	var redisCom redis.Com
	redisCom.Filed = gconv.String(userId)
	redisCom.Key = shared.UserSignToday
	res := redisCom.CheckHashField()
	if res {
		return true
	}
	return false
}
