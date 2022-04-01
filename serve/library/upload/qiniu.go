package upload_lib

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"golang.org/x/net/context"
)

// 七牛云存储引擎
type QnyEngine struct {
	Path            string
	Endpoint        string
	Address         string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

// 七牛云的上传
func (e *QnyEngine) Upload(Id int64, files ghttp.UploadFiles) ([]MediaObject, error) {
	// 鉴权
	mac := qbox.NewMac(e.AccessKeyId, e.AccessKeySecret)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   e.BucketName,
		Expires: 7200,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{}
	switch e.Endpoint {
	case "huadong":
		cfg.Zone = &storage.ZoneHuadong //指定上传的区域
	case "huabei":
		cfg.Zone = &storage.ZoneHuabei //指定上传的区域
	case "huanan":
		cfg.Zone = &storage.ZoneHuanan //指定上传的区域
	case "beimei":
		cfg.Zone = &storage.ZoneBeimei //指定上传的区域
	default:
		cfg.Zone = &storage.ZoneXinjiapo //指定上传的区域
	}
	cfg.UseHTTPS = false      // 是否使用https域名
	cfg.UseCdnDomains = false //是否使用CDN上传加速

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	var mediaObject []MediaObject
	for _, item := range files {
		// 打开文件
		file, err := item.Open()
		if err != nil {
			return nil, gerror.New("文件上传错误")
		}
		defer file.Close() // 文件关闭
		// 定义路径
		path := e.Path + "/" + gtime.Date() + "/"
		uuid, _ := gmd5.Encrypt(gtime.TimestampNanoStr())
		fileSuffix := gfile.Ext(item.FileHeader.Filename) //获取文件名
		// 重组路径
		name := uuid + fileSuffix
		path = path + name

		// 上传文件
		err = formUploader.Put(context.Background(), &ret, upToken, path, file, item.Size, nil)
		if err != nil {
			return nil, gerror.New("文件上传错误")
		}
		mediaType := gfile.Ext(item.FileHeader.Filename)
		mediaType = mediaType[1:]
		media := MediaObject{
			UserId:     Id,
			Link:       e.Address + "/" + ret.Key,
			Path:       path,
			Name:       name,
			OrName:     item.Filename,
			Size:       gconv.String(item.Size),
			UploadKey:  2,
			Status:     1,
			MediaType:  gstr.ToUpper(mediaType),
			Ext:        gfile.Ext(item.FileHeader.Filename),
			CreateTime: gtime.Now(),
		}
		mediaObject = append(mediaObject, media)
	}

	return mediaObject, nil
}

//七牛云删除文件
func (e *QnyEngine) Remove() error {
	// 鉴权
	mac := qbox.NewMac(e.AccessKeyId, e.AccessKeySecret)

	// 上传Config对象
	cfg := storage.Config{}
	switch e.Endpoint {
	case "huadong":
		cfg.Zone = &storage.ZoneHuadong //指定上传的区域
	case "huabei":
		cfg.Zone = &storage.ZoneHuabei //指定上传的区域
	case "huanan":
		cfg.Zone = &storage.ZoneHuanan //指定上传的区域
	case "beimei":
		cfg.Zone = &storage.ZoneBeimei //指定上传的区域
	default:
		cfg.Zone = &storage.ZoneXinjiapo //指定上传的区域
	}
	cfg.UseHTTPS = false // 是否使用https域名
	//cfg.UseCdnDomains = false //是否使用CDN上传加速

	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(e.BucketName, e.Path)
	if err != nil {
		return gerror.New("文件删除失败")
	}
	return nil
}
