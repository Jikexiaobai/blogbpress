package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	"fiber/library/redis"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var System = new(systemService)

type systemService struct{}

// Save 保存配置
func (s *systemService) Save(req *dto.SystemCreate) (code response.ResponseCode) {
	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}
	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	cg, err := dao.SysConfig.FindOne(dao.SysConfig.Columns.ConfigKey, req.ConfigKey)
	if err != nil {
		return response.DB_READ_ERROR
	}
	if cg != nil {
		cg.ConfigKey = req.ConfigKey
		cg.ConfigName = req.ConfigName
		cg.ConfigValue = req.ConfigValue
		cg.Remark = req.Remark

		cg.UpdateTime = gtime.Now()
		_, err = tx.Update(dao.SysConfig.Table, cg, dao.SysConfig.Columns.ConfigId, cg.ConfigId)
		if err != nil {
			return response.UPDATE_FAILED
		}

		switch req.ConfigKey {
		case shared.BaseSetting:
			err = Media.RemoveRelated(tx, nil, shared.BaseSetting)
			if err != nil {
				return response.DELETE_FAILED
			}
			base := gjson.New(req.ConfigValue)
			logo := gconv.String(base.Get("logo"))
			icon := gconv.String(base.Get("icon"))
			pathList := make([]string, 0)
			pathList = append(pathList, logo)
			pathList = append(pathList, icon)
			if len(pathList) > 0 {
				err = Media.AddRelated(tx, pathList, 0, shared.BaseSetting)
				if err != nil {
					return response.FILE_SAVE_ERROR
				}
			}
		case shared.UserSetting:
			err = Media.RemoveRelated(tx, nil, shared.UserSetting)
			if err != nil {
				return response.DELETE_FAILED
			}
			base := gjson.New(req.ConfigValue)
			defaultCover := gconv.String(base.Get("defaultCover"))
			defaultAvatar := gconv.String(base.Get("defaultAvatar"))
			pathList := make([]string, 0)
			pathList = append(pathList, defaultCover)
			pathList = append(pathList, defaultAvatar)
			if len(pathList) > 0 {
				err = Media.AddRelated(tx, pathList, 0, shared.UserSetting)
				if err != nil {
					return response.FILE_SAVE_ERROR
				}
			}
		}

	} else {
		var config model.SysConfig
		err := gconv.Struct(req, &config)
		if err != nil {
			return response.FAILD
		}

		config.CreateTime = gtime.Now()
		config.UpdateTime = gtime.Now()

		_, err = dao.SysConfig.Insert(config)
		if err != nil {
			return response.ADD_FAILED
		}
		switch req.ConfigKey {
		case shared.BaseSetting:
			base := gjson.New(req.ConfigValue)
			logo := gconv.String(base.Get("logo"))
			icon := gconv.String(base.Get("icon"))
			pathList := make([]string, 0)
			pathList = append(pathList, logo)
			pathList = append(pathList, icon)
			if len(pathList) > 0 {
				err = Media.AddRelated(tx, pathList, 0, shared.BaseSetting)
				if err != nil {
					return response.FILE_SAVE_ERROR
				}
			}
		case shared.UserSetting:
			base := gjson.New(req.ConfigValue)
			defaultCover := gconv.String(base.Get("defaultCover"))
			defaultAvatar := gconv.String(base.Get("defaultAvatar"))
			pathList := make([]string, 0)
			pathList = append(pathList, defaultCover)
			pathList = append(pathList, defaultAvatar)
			if len(pathList) > 0 {
				err = Media.AddRelated(tx, pathList, 0, shared.UserSetting)
				if err != nil {
					return response.FILE_SAVE_ERROR
				}
			}
		}
	}
	return response.SUCCESS
}

// FindValue 查询配置
func (s *systemService) FindValue(key string) (string, error) {
	result, err := dao.SysConfig.Value(
		dao.SysConfig.Columns.ConfigValue,
		dao.SysConfig.Columns.ConfigKey,
		key)
	if err != nil {
		return "", err
	}
	return gconv.String(result), nil
}

// Value 查询配置
func (s *systemService) Value(key string) (string, response.ResponseCode) {
	result, err := s.FindValue(key)
	if err != nil {
		return "", response.NOT_FOUND
	}
	return gconv.String(result), response.SUCCESS
}

// Info 系统信息
func (s *systemService) Info() (*result.SystemInfo, error) {

	var redisCom redis.Com
	redisCom.Key = shared.AdminSystemInfo
	systemObj, err := redisCom.GetString()
	if err != nil {
		return nil, err
	}
	if systemObj != nil {
		var systemInfo *result.SystemInfo
		err := gconv.Struct(systemObj, &systemInfo)
		if err != nil {
			return nil, err
		}
		return systemInfo, nil
	}

	list, err := dao.SysConfig.
		Where(dao.SysConfig.Columns.ConfigKey+" IN(?)",
			[]string{shared.BaseSetting,
				shared.FileSetting}).All()
	if err != nil {
		return nil, err
	}

	var base result.Base
	var file result.File
	for _, i := range list {
		if i.ConfigKey == shared.BaseSetting {
			err = gconv.Struct(i.ConfigValue, &base)
			if err != nil {
				return nil, err
			}
		}
		if i.ConfigKey == shared.FileSetting {
			err = gconv.Struct(i.ConfigValue, &file)
			if err != nil {
				return nil, err
			}
		}
	}
	info := result.SystemInfo{File: &file, Base: &base}

	// 写入缓存
	redisCom.Time = "1200"
	redisCom.Data = info
	err = redisCom.SetStringEX()
	if err != nil {
		return nil, err
	}
	return &info, nil
}
