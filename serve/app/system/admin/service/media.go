package service

import (
	"fiber/app/dao"
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/result"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	upload_lib "fiber/library/upload"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Media = new(mediaService)

type mediaService struct {
}

// SelectList 查询列表
func (s *mediaService) SelectList(req *dto.MediaQuery) (int, []*result.MediaList, response.ResponseCode) {
	model := dao.SysMedia.SysMediaDao.Where(dao.SysAudio.Columns.DeleteTime, nil)
	if req.OrName != "" {
		model = model.Where(dao.SysMedia.Columns.OrName+" like ?", "%"+req.OrName+"%")
	}

	// 状态
	if req.Status != 0 {
		model = model.Where(dao.SysAudio.Columns.Status, req.Status)
	}
	// 未使用
	model = model.Order(dao.SysMedia.Columns.CreateTime, "desc")
	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	model = model.Page(req.Page, req.Limit)
	list, err := model.Fields(
		dao.SysMedia.Columns.MediaId,
		dao.SysMedia.Columns.OrName,
		dao.SysMedia.Columns.Size,
		dao.SysMedia.Columns.Ext,
		dao.SysMedia.Columns.Path,
		dao.SysMedia.Columns.Link,
		dao.SysMedia.Columns.UploadKey,
		dao.SysMedia.Columns.Status,
		dao.SysMedia.Columns.UserId,
		dao.SysMedia.Columns.CreateTime,
	).All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.MediaList
	for _, i := range list {
		var info *result.MediaList
		err = gconv.Struct(i, &info)
		if err != nil {
			return 0, nil, response.FAILD
		}

		nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, i.UserId)
		if err != nil {
			return 0, nil, response.DB_READ_ERROR
		}
		info.NickName = gconv.String(nickName)
		res = append(res, info)
	}
	return total, res, response.SUCCESS
}

// Create 创建记录
func (s *mediaService) Create(object []upload_lib.MediaObject) error {
	_, err := dao.SysMedia.Insert(object)
	if err != nil {
		return err
	}
	return nil
}

// AddRelated 添加关联
func (s *mediaService) AddRelated(tx *gdb.TX, paths []string, relateId int64, module string) error {
	mediaList, err := dao.SysMedia.
		Fields(dao.SysMedia.Columns.MediaId).
		Where(dao.SysMedia.Columns.Link+" IN(?)", paths).
		All()
	if err != nil {
		return err
	}
	if len(mediaList) > 0 {
		var tmpMediaList []map[string]interface{}
		var tmpId []int64
		for _, i := range mediaList {
			tmpMedia := make(map[string]interface{})
			tmpMedia[dao.SysMediaRelated.Columns.Module] = module
			tmpMedia[dao.SysMediaRelated.Columns.RelatedId] = relateId
			tmpMedia[dao.SysMediaRelated.Columns.MediaId] = i.MediaId
			tmpMediaList = append(tmpMediaList, tmpMedia)
			tmpId = append(tmpId, i.MediaId)
		}

		_, err = tx.Save(dao.SysMediaRelated.Table, tmpMediaList)
		if err != nil {
			return err
		}

		_, err = tx.Update(dao.SysMedia.Table, g.Map{
			dao.SysMedia.Columns.Status: 2,
		}, dao.SysMedia.Columns.MediaId+" IN(?)", tmpId)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

// RemoveRelated 删除关联
func (s *mediaService) RemoveRelated(tx *gdb.TX, relatedId []int64, module string) error {

	if relatedId != nil && module != "" {
		mediaList, err := dao.SysMediaRelated.
			Fields(dao.SysMediaRelated.Columns.MediaId).
			Where(dao.SysMediaRelated.Columns.RelatedId+" IN(?)", relatedId).
			Where(dao.SysMediaRelated.Columns.Module, module).
			All()
		if err != nil {
			return err
		}
		var tmpId []int64
		for _, i := range mediaList {
			tmpId = append(tmpId, i.MediaId)
		}
		_, err = tx.Model(dao.SysMediaRelated.Table).
			Where(dao.SysMediaRelated.Columns.Module, module).
			Delete(dao.SysMediaRelated.Columns.RelatedId+" IN(?)",
				relatedId)
		if err != nil {
			return err
		}

		tmpList, err := dao.SysMediaRelated.
			Where(dao.SysMediaRelated.Columns.MediaId+" IN(?)", tmpId).
			All()
		if err != nil {
			return err
		}

		if tmpList == nil {
			_, err = tx.Update(dao.SysMedia.Table, g.Map{
				dao.SysMedia.Columns.Status: 1,
			}, dao.SysMedia.Columns.MediaId+" IN(?)", tmpId)
			if err != nil {
				return err
			}
		}

		return nil
	}

	mediaList, err := dao.SysMediaRelated.
		Fields(dao.SysMediaRelated.Columns.MediaId).
		Where(dao.SysMediaRelated.Columns.RelatedId+" IN(?)", relatedId).
		All()
	if err != nil {
		return err
	}
	var tmpId []int64
	for _, i := range mediaList {
		tmpId = append(tmpId, i.MediaId)
	}
	_, err = tx.Delete(dao.SysMediaRelated.Table,
		dao.SysMediaRelated.Columns.Module, module)
	if err != nil {
		return err
	}

	tmpList, err := dao.SysMediaRelated.
		Where(dao.SysMediaRelated.Columns.MediaId+" IN(?)", tmpId).
		All()
	if err != nil {
		return err
	}

	if tmpList == nil {
		_, err = tx.Update(dao.SysMedia.Table, g.Map{
			dao.SysMedia.Columns.Status: 1,
		}, dao.SysMedia.Columns.MediaId+" IN(?)", tmpId)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove 删除媒体信息
func (s *mediaService) Remove(ids []int64) (code response.ResponseCode) {
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
	mediaList, err := dao.SysMedia.
		Where(dao.SysMedia.Columns.MediaId+" IN(?)", ids).
		All()

	if err != nil {
		return response.DB_READ_ERROR
	}
	for _, i := range mediaList {
		switch i.UploadKey {
		case shared.Local:
			//	本地引擎删除文件
			var mediaEngine upload_lib.Engine
			path := "." + i.Path
			mediaEngine = &upload_lib.LocalEngine{Path: path}
			err := mediaEngine.Remove()
			if err != nil {
				return response.DELETE_FAILED
			}
			// 表记录删除
			_, err = tx.Delete(dao.SysMedia.Table, dao.SysMedia.Columns.MediaId+" IN(?)", ids)
			if err != nil {
				return response.DELETE_FAILED
			}
			_, err = tx.Delete(dao.SysMediaRelated.Table, dao.SysMediaRelated.Columns.MediaId+" IN(?)", ids)
			if err != nil {
				return response.DELETE_FAILED
			}
		case shared.AlyOss:
			//阿里云引擎删除文件
			AlyOssOption, err := System.FindValue(shared.AlyOssOption)
			if err != nil {
				return response.DB_READ_ERROR
			}
			AlyOssOptionJson := gjson.New(AlyOssOption)
			endpoint := gconv.String(AlyOssOptionJson.Get("endpoint"))
			accessKeyId := gconv.String(AlyOssOptionJson.Get("accessKeyId"))
			accessKeySecret := gconv.String(AlyOssOptionJson.Get("accessKeySecret"))
			bucketName := gconv.String(AlyOssOptionJson.Get("bucketName"))
			//
			var mediaEngine upload_lib.Engine
			mediaEngine = &upload_lib.AlyEngine{
				Path:            i.Path,
				Endpoint:        endpoint,
				AccessKeyId:     accessKeyId,
				AccessKeySecret: accessKeySecret,
				BucketName:      bucketName,
			}
			err = mediaEngine.Remove()
			if err != nil {
				return response.DELETE_FAILED
			}
			// 表记录删除
			_, err = tx.Delete(dao.SysMedia.Table, dao.SysMedia.Columns.MediaId+" IN(?)", ids)
			if err != nil {
				return response.DELETE_FAILED
			}
			_, err = tx.Delete(dao.SysMediaRelated.Table, dao.SysMediaRelated.Columns.MediaId+" IN(?)", ids)
			if err != nil {
				return response.DELETE_FAILED
			}
		case shared.QiNiuOss:
			QiNiuOssOption, err := System.FindValue(shared.QiNiuOssOption)
			if err != nil {
				return response.DB_READ_ERROR
			}
			QiNiuOssOptionJson := gjson.New(QiNiuOssOption)
			endpoint := gconv.String(QiNiuOssOptionJson.Get("endpoint"))
			accessKeyId := gconv.String(QiNiuOssOptionJson.Get("accessKeyId"))
			accessKeySecret := gconv.String(QiNiuOssOptionJson.Get("accessKeySecret"))
			bucketName := gconv.String(QiNiuOssOptionJson.Get("bucketName"))
			address := gconv.String(QiNiuOssOptionJson.Get("address"))

			var mediaEngine upload_lib.Engine
			mediaEngine = &upload_lib.QnyEngine{
				Path:            i.Path,
				Endpoint:        endpoint,
				AccessKeyId:     accessKeyId,
				AccessKeySecret: accessKeySecret,
				BucketName:      bucketName,
				Address:         address,
			}
			err = mediaEngine.Remove()
			if err != nil {
				return response.DELETE_FAILED
			}
			_, err = tx.Delete(dao.SysMedia.Table, dao.SysMedia.Columns.MediaId+" IN(?)", ids)
			if err != nil {
				return response.DELETE_FAILED
			}
			_, err = tx.Delete(dao.SysMediaRelated.Table, dao.SysMediaRelated.Columns.MediaId+" IN(?)", ids)
			if err != nil {
				return response.DELETE_FAILED
			}
		}
	}

	return response.SUCCESS
}
