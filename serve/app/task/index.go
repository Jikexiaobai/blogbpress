package task

import (
	"fiber/app/system/index/shared"
	"fiber/library/redis"
	"github.com/gogf/gf/os/gcron"
)

func Init() {
	//	用户签到 每天定时删除key
	loadSign()
	deleteRedisKey()
	//定时文字数据写入
	//loadArticle()
}

func deleteRedisKey() {
	_, _ = gcron.Add("0 0 2 * * *", func() {
		var redisCom redis.Com
		redisCom.Key = shared.UserHot
		isHaveHotUser := redisCom.CheckExists()
		if isHaveHotUser {
			_ = redisCom.Delete()
		}
		redisCom.Key = shared.Search
		isHaveSearch := redisCom.CheckExists()
		if isHaveSearch {
			_ = redisCom.Delete()
		}
	})
}
