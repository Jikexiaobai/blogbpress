package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (e *Com) LPush() error {
	_, err := g.Redis().Do("LPUSH", e.Key, e.Data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Com) LRange(page, limit int) ([]string, error) {
	res, err := g.Redis().Do("LRANGE", e.Key, page, limit)
	if err != nil {
		return nil, err
	}
	return gconv.Strings(res), nil
}

func (e *Com) LLength() (int, error) {
	res, err := g.Redis().Do("LLEN", e.Key)
	if err != nil {
		return 0, err
	}
	return gconv.Int(res), nil
}
