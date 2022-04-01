package service

import (
	"fiber/app/dao"
	"fiber/app/model"
)

var Dict = new(dictService)

type dictService struct {
}

func (s *dictService) SelectDictDataByType(dictType string) ([]*model.SysDictData, error) {
	result, err := dao.SysDictData.Where(dao.SysDictData.Columns.DictType, dictType).All()
	if err != nil {
		return nil, err
	}
	return result, nil
}
