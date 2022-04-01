package lock_utils

import (
	"fiber/library/redis"
	"github.com/gogf/gf/util/gconv"
)

func SetCount(key, lockKey string, expired int64, count int) (int, error) {
	Times := 0
	var redisCom redis.Com
	redisCom.Key = key
	TimeObj, err := redisCom.GetString()
	//TimeObj, err := g.Redis().Do("GET", key)
	if err != nil {
		return Times, err
	}
	if TimeObj != nil {
		TimeIncr, err := redisCom.INCString()
		Times = gconv.Int(TimeIncr)
		if err != nil {
			return Times, err
		}
	} else {
		redisCom.Data = Times
		redisCom.Time = expired
		err := redisCom.SetStringEX()
		if err != nil {
			return Times, err
		}
	}

	if Times >= count {
		setLock(lockKey, expired)
	}
	return Times, nil
}

func setLock(key string, expired int64) {
	var redisCom redis.Com
	redisCom.Key = key
	redisCom.Data = "true"
	redisCom.Time = expired
	_ = redisCom.SetStringEX()
}

// CheckLock 检查账号是否锁定
func CheckLock(key string) bool {
	result := false
	var redisCom redis.Com
	redisCom.Key = key
	lock, err := redisCom.GetString()
	if err != nil {
		return false
	}
	if lock != nil {
		lockContent := gconv.String(lock)
		if lockContent == "true" {
			result = true
		}
	}
	return result
}
