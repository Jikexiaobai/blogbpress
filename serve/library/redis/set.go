package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (e *Com) ADDSet() error {
	_, err := g.Redis().Do("SADD", e.Key, e.Data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) REMSet() error {
	_, err := g.Redis().Do("SREM", e.Key, e.Data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) GetAllSet() []string {
	res, _ := g.Redis().Do("SMEMBERS", e.Key)
	return gconv.Strings(res)
}

func (e *Com) GetRandSet(count int) []string {
	res, _ := g.Redis().Do("SRANDMEMBER", e.Key, count)
	return gconv.Strings(res)
}
