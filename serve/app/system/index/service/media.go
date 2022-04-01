package service

import (
	"fiber/app/dao"
	"fiber/app/system/index/shared"
	upload_lib "fiber/library/upload"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Media = new(mediaService)

type mediaService struct {
}

func (s *mediaService) SelectMediaPath(url string) string {
	path, err := dao.SysMedia.Value(dao.SysMedia.Columns.Path,
		dao.SysMedia.Columns.Link,
		url)
	if err != nil {
		return ""
	}
	return gconv.String(path)
}

func (s *mediaService) Create(object []upload_lib.MediaObject) error {
	_, err := dao.SysMedia.Insert(object)
	if err != nil {
		return err
	}
	return nil
}

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
			dao.SysMedia.Columns.Status: shared.StatusReviewed,
		}, dao.SysMedia.Columns.MediaId+" IN(?)", tmpId)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (s *mediaService) RemoveRelated(tx *gdb.TX, relatedId int64, module string) error {

	if relatedId != 0 && module != "" {
		mediaList, err := dao.SysMediaRelated.
			Fields(dao.SysMediaRelated.Columns.MediaId).
			Where(dao.SysMediaRelated.Columns.RelatedId, relatedId).
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
			Delete(dao.SysMediaRelated.Columns.RelatedId,
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
		Where(dao.SysMediaRelated.Columns.RelatedId, relatedId).
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
			dao.SysMedia.Columns.Status: shared.StatusReview,
		}, dao.SysMedia.Columns.MediaId+" IN(?)", tmpId)
		if err != nil {
			return err
		}
	}

	return nil
}
