package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (e *Com) ZIncr() error {
	_, err := g.Redis().Do("ZINCRBY", e.Key, 2, e.Data)
	if err != nil {
		return err
	}
	return nil
}

// ZRange
func (e *Com) ZRange() []string {
	res, _ := g.Redis().Do("ZRANGE", e.Key, "0", "5")
	return gconv.Strings(res)
}
