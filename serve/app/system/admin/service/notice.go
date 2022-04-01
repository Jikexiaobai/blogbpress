package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Notice = new(noticeService)

type noticeService struct {
}

// 创建
func (s *noticeService) Create(req *dto.NoticeCreate) error {
	var entity *model.SysNotice
	err := gconv.Struct(req, &entity)
	if err != nil {
		return err
	}
	entity.Status = 1
	entity.CreateTime = gtime.Now()
	_, err = dao.SysNotice.Save(entity)
	if err != nil {
		return err
	}
	return nil
}
