package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (e *Com) GetString() (interface{}, error) {
	res, err := g.Redis().Do("GET", e.Key)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Com) SetStringEX() error {
	_, err := g.Redis().Do("SETEX", e.Key, e.Time, e.Data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) INCString() (int64, error) {
	res, err := g.Redis().Do("INCR", e.Key)
	if err != nil {
		return gconv.Int64(gconv.String(res)), err
	}
	return gconv.Int64(gconv.String(res)), nil
}

func (e *Com) DELString() error {
	_, err := g.Redis().Do("DEL", e.Key)
	if err != nil {
		return err
	}
	return nil
}
