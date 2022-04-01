package valid

import (
	"fiber/app/system/admin/shared"
	"fiber/app/system/index/service"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

var Media = new(mediaService)

type mediaService struct {
}

// CheckMediaSize 检查文件大小
func (s *mediaService) CheckMediaSize(size int64) bool {
	config, err := service.Config.FindValue(shared.FileSetting)
	if err != nil {
		return true
	}

	ConfigValue, _ := gjson.DecodeToJson([]byte(config))

	mediaSize := ConfigValue.GetInt64("fileSize")

	if mediaSize*1024*1024 < size {
		return true
	}
	return false
}

// CheckMediaType 检查文件类型
func (s *mediaService) CheckMediaType(mediaType string) bool {
	FileSetting, err := service.Config.FindValue(shared.FileSetting)
	if err != nil {
		return false
	}
	FileSettingJson := gjson.New(FileSetting)
	imageType := gconv.Strings(FileSettingJson.Get("imageType"))
	audioType := gconv.Strings(FileSettingJson.Get("audioType"))
	videoType := gconv.Strings(FileSettingJson.Get("videoType"))

	fileType := make([]string, 0)
	fileType = append(fileType, imageType...)
	fileType = append(fileType, audioType...)
	fileType = append(fileType, videoType...)
	for _, i := range fileType {

		if i == mediaType {
			return true
		}
	}
	return false
}
