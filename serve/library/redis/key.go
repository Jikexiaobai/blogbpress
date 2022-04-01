package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (e *Com) Delete() error {
	_, err := g.Redis().Do("DEL", e.Key)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) SetExpire() error {
	_, err := g.Redis().Do("EXPIRE", e.Key, e.Time)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) CheckExists() bool {
	res, err := g.Redis().Do("EXISTS", e.Key)
	if err != nil || gconv.Int(res) == 0 {
		return false
	}
	return true
}
