package service

import (
	"fiber/app/dao"
	"github.com/gogf/gf/util/gconv"
)

var Config = new(configService)

type configService struct{}

func (s *configService) FindValue(key string) (string, error) {
	result, err := dao.SysConfig.Value(
		dao.SysConfig.Columns.ConfigValue,
		dao.SysConfig.Columns.ConfigKey,
		key)
	if err != nil {
		return "", err
	}
	return gconv.String(result), nil
}
