package upload_lib

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// AlyEngine 阿里云oss存储引擎
type AlyEngine struct {
	Path            string
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

// Upload 阿里云的上传
func (e *AlyEngine) Upload(Id int64, files ghttp.UploadFiles) ([]MediaObject, error) {
	client, err := oss.New(e.Endpoint, e.AccessKeyId, e.AccessKeySecret)
	if err != nil {
		return nil, gerror.New("文件上传错误")
	}

	// 获取存储空间。
	bucket, err := client.Bucket(e.BucketName)
	if err != nil {
		return nil, gerror.New("文件上传错误")
	}
	// fileNames
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

		// 上传文件。
		err = bucket.PutObject(path, file)
		if err != nil {
			return nil, gerror.New("文件上传错误")
		}
		mediaType := gfile.Ext(item.FileHeader.Filename)
		mediaType = mediaType[1:]
		media := MediaObject{
			UserId:     Id,
			Link:       "http://" + e.BucketName + "." + e.Endpoint + "/" + path,
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

// Remove 阿里云删除文件
func (e *AlyEngine) Remove() error {
	client, err := oss.New(e.Endpoint, e.AccessKeyId, e.AccessKeySecret)
	if err != nil {
		return gerror.New("连接错误")
	}
	// 获取存储空间。
	bucket, err := client.Bucket(e.BucketName)
	if err != nil {
		return gerror.New("连接错误")
	}
	// 删除文件。
	err = bucket.DeleteObject(e.Path)
	if err != nil {
		return gerror.New("连接错误")
	}
	return nil
}

// GetToken 阿里云获取上传签名
func (e *AlyEngine) GetToken() error {
	client, err := sts.NewClientWithAccessKey("cn-shenzhen", e.AccessKeyId, e.AccessKeySecret)
	if err != nil {
		return gerror.New("文件上传错误")
	}

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见《API参考》。
	request.RoleArn = "acs:ram::1986051699684916:role/ramosstest"
	request.RoleSessionName = "oss-session"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	return nil
}
