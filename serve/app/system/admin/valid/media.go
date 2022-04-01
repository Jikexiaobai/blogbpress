package valid

import (
	"fiber/app/dao"
	"fiber/app/system/admin/shared"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

var Media = new(mediaValid)

type mediaValid struct {
}

// CheckMediaSize 检查文件大小
func (s *mediaValid) CheckMediaSize(size int64) bool {
	//System, err := System.FindValue("FileSetting")
	res, err := dao.SysConfig.Value(dao.SysConfig.Columns.ConfigValue, dao.SysConfig.Columns.ConfigKey, shared.FileSetting)
	if err != nil {
		return true
	}

	SystemValue, _ := gjson.DecodeToJson([]byte(gconv.String(res)))

	mediaSize := SystemValue.GetInt64("fileSize")

	if mediaSize*1024*1024 < size {
		return true
	}
	return false
}

// CheckMediaType 检查文件类型
func (s *mediaValid) CheckMediaType(mediaType string) bool {
	FileSetting, err := dao.SysConfig.Value(dao.SysConfig.Columns.ConfigValue, dao.SysConfig.Columns.ConfigKey, shared.FileSetting)
	if err != nil {
		return true
	}
	FileSettingJson := gjson.New(FileSetting)
	imageType := gconv.Strings(FileSettingJson.Get("imageType"))
	audioType := gconv.Strings(FileSettingJson.Get("audioType"))
	videoType := gconv.Strings(FileSettingJson.Get("videoType"))
	otherType := gconv.Strings(FileSettingJson.Get("otherType"))

	fileType := make([]string, 0)
	fileType = append(fileType, imageType...)
	fileType = append(fileType, audioType...)
	fileType = append(fileType, videoType...)
	fileType = append(fileType, otherType...)
	for _, i := range fileType {

		if i == mediaType {
			return true
		}
	}
	return false
}
