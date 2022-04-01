package upload_lib

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// 本地存储引擎
type LocalEngine struct {
	Path   string
	Domain string
}

// Upload 本地上传
func (e *LocalEngine) Upload(Id int64, files ghttp.UploadFiles) ([]MediaObject, error) {

	path := "./public/" + e.Path + "/" + gtime.Date() + "/"

	names, err := files.Save(path, true)
	if err != nil {
		return nil, gerror.New("文件上传错误")
	}
	var mediaObject []MediaObject

	for i, item := range files {
		path = path[1:] + names[i]
		mediaType := gfile.Ext(item.FileHeader.Filename)
		mediaType = mediaType[1:]
		media := MediaObject{
			UserId:     Id,
			Link:       e.Domain + "/public" + path[7:],
			Path:       path,
			Name:       names[i],
			OrName:     item.Filename,
			Size:       gconv.String(item.Size),
			UploadKey:  1,
			Status:     1,
			MediaType:  gstr.ToUpper(mediaType),
			Ext:        gfile.Ext(item.FileHeader.Filename),
			CreateTime: gtime.Now(),
		}
		mediaObject = append(mediaObject, media)
	}
	return mediaObject, nil
}

// Remove 本地删除
func (e *LocalEngine) Remove() error {
	// 判断文件是否存在
	if !gfile.Exists(e.Path) {
		return nil
	}

	//存在就删除文件
	err := gfile.Remove(e.Path)
	if err != nil {
		return gerror.New("文件不存在")
	}
	return nil
}
