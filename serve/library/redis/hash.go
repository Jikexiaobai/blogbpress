package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (e *Com) GetHashField() int64 {
	res, _ := g.Redis().Do("HGET", e.Key, e.Filed)
	return gconv.Int64(gconv.String(res))
}

func (e *Com) GetHashFieldString() string {
	res, _ := g.Redis().Do("HGET", e.Key, e.Filed)
	return gconv.String(res)
}

func (e *Com) GetHashAllField() []string {
	res, _ := g.Redis().Do("HKEYS", e.Key)
	return gconv.Strings(res)
}

func (e *Com) SetHashField() error {
	_, err := g.Redis().Do("HSET", e.Key, e.Filed, e.Data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) INCHashField() (int64, error) {
	res, err := g.Redis().Do("HINCRBY", e.Key, e.Filed, e.Data)
	if err != nil {
		return gconv.Int64(gconv.String(res)), err
	}
	return gconv.Int64(gconv.String(res)), nil
}

func (e *Com) CheckHashField() bool {
	hasKey, err := g.Redis().Do("HEXISTS", e.Key, e.Filed)
	if gconv.Int(gconv.String(hasKey)) != 0 && err == nil {
		return true
	}
	return false
}

func (e *Com) DeleteHashField() error {
	_, err := g.Redis().Do("HDEL", e.Key, e.Filed)
	if err != nil {
		return err
	}
	return nil
}
