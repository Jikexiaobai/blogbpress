package service

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/shared"
	"fiber/app/tools/response"
	upload_lib "fiber/library/upload"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"io/ioutil"
	"strconv"
	"strings"
)

var Upload = new(uploadService)

type uploadService struct {
}

func (s *uploadService) UploadFile(userId int64, files ghttp.UploadFiles) ([]string, response.ResponseCode) {

	// 获取文件保存的存储引擎

	FileSetting, err := Config.FindValue(shared.FileSetting)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	FileSettingJson := gjson.New(FileSetting)
	path := gconv.String(FileSettingJson.Get("path"))
	engine := gconv.Int(FileSettingJson.Get("engine"))

	switch engine {
	case shared.Local:
		BaseSetting, err := Config.FindValue("BaseSetting")
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		BaseSettingJson := gjson.New(BaseSetting)
		url := gconv.String(BaseSettingJson.Get("url"))

		var mediaEngine upload_lib.Engine
		mediaEngine = &upload_lib.LocalEngine{Path: path, Domain: url}
		mediaObject, err := mediaEngine.Upload(userId, files)
		if err != nil {
			return nil, response.FILE_SAVE_ERROR
		}
		err = Media.Create(mediaObject)
		if err != nil {
			return nil, response.DB_SAVE_ERROR
		}
		var mediaResult []string
		for _, i := range mediaObject {
			mediaResult = append(mediaResult, i.Link)
		}
		return mediaResult, response.SUCCESS
	case shared.AlyOss:
		AlyOssOption, err := Config.FindValue(shared.AlyOssOption)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		AlyOssOptionJson := gjson.New(AlyOssOption)
		endpoint := gconv.String(AlyOssOptionJson.Get("endpoint"))
		accessKeyId := gconv.String(AlyOssOptionJson.Get("accessKeyId"))
		accessKeySecret := gconv.String(AlyOssOptionJson.Get("accessKeySecret"))
		bucketName := gconv.String(AlyOssOptionJson.Get("bucketName"))

		var mediaEngine upload_lib.Engine
		mediaEngine = &upload_lib.AlyEngine{
			Path:            path,
			Endpoint:        endpoint,
			AccessKeyId:     accessKeyId,
			AccessKeySecret: accessKeySecret,
			BucketName:      bucketName,
		}
		mediaObject, err := mediaEngine.Upload(userId, files)
		if err != nil {
			return nil, response.FILE_SAVE_ERROR
		}
		err = Media.Create(mediaObject)
		if err != nil {
			return nil, response.DB_SAVE_ERROR
		}
		var mediaResult []string
		for _, i := range mediaObject {
			mediaResult = append(mediaResult, i.Link)
		}
		return mediaResult, response.SUCCESS
	case shared.QiNiuOss:
		QiNiuOssOption, err := Config.FindValue(shared.QiNiuOssOption)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		QiNiuOssOptionJson := gjson.New(QiNiuOssOption)
		endpoint := gconv.String(QiNiuOssOptionJson.Get("endpoint"))
		accessKeyId := gconv.String(QiNiuOssOptionJson.Get("accessKeyId"))
		accessKeySecret := gconv.String(QiNiuOssOptionJson.Get("accessKeySecret"))
		bucketName := gconv.String(QiNiuOssOptionJson.Get("bucketName"))
		address := gconv.String(QiNiuOssOptionJson.Get("address"))

		var mediaEngine upload_lib.Engine
		mediaEngine = &upload_lib.QnyEngine{
			Path:            path,
			Endpoint:        endpoint,
			AccessKeyId:     accessKeyId,
			AccessKeySecret: accessKeySecret,
			BucketName:      bucketName,
			Address:         address,
		}
		mediaObject, err := mediaEngine.Upload(userId, files)
		if err != nil {
			return nil, response.FILE_SAVE_ERROR
		}
		err = Media.Create(mediaObject)
		if err != nil {
			return nil, response.DB_SAVE_ERROR
		}
		var mediaResult []string
		for _, i := range mediaObject {
			mediaResult = append(mediaResult, i.Link)
		}
		return mediaResult, response.SUCCESS
	}
	return nil, response.FILE_SAVE_ERROR
}

// UploadChunk 上传切片
func (s *uploadService) UploadChunk(req *dto.UploadChunkFrom) response.ResponseCode {
	folderPath := "./public/chunk/" + req.Identifier
	fileName := folderPath + "/" + req.ChunkNumber + ".temp" //文件完全路径名
	req.File.Filename = fileName

	_, err := req.File.Save(folderPath, false)
	if err != nil {
		return response.FILE_SAVE_ERROR
	}
	return response.SUCCESS
}

// GetUploadChunkFiles 检查切片是否存在
func (s *uploadService) GetUploadChunkFiles(identifier string) ([]int, response.ResponseCode) {
	// 判断是否存在这个文件夹
	folderPath := "./public/chunk/" + identifier
	if gfile.Exists(folderPath) {
		//	如果存在
		list, _ := gfile.ScanDir(folderPath, "*", true)
		var fileNames []int
		for _, v := range list {
			tmpNames := gstr.Split(gfile.Basename(v), ".")
			fileNames = append(fileNames, gconv.Int(tmpNames[0]))
		}
		return fileNames, response.SUCCESS
	}

	return []int{}, response.SUCCESS
}

// MergeChunk 合并切片
func (s *uploadService) MergeChunk(req *dto.UploadChunkMergeFrom) ([]string, response.ResponseCode) {
	tmpFolderPath := "./public/chunk/" + req.Identifier
	if !gfile.Exists(tmpFolderPath) {
		return nil, response.FILE_SAVE_ERROR
	}

	tmpFolderFiles, err := ioutil.ReadDir(tmpFolderPath)
	if err != nil {
		return nil, response.FILE_SAVE_ERROR
	}
	FileSetting, err := Config.FindValue(shared.FileSetting)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}

	FileSettingJson := gjson.New(FileSetting)
	path := gconv.String(FileSettingJson.Get("path"))
	filePath := "./public/" + path + "/" + gtime.Date() + "/"
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	filePath = filePath + name + gfile.Ext(req.FileName)

	fileHandle, err := gfile.Create(filePath)
	if err != nil {

		return nil, response.FILE_SAVE_ERROR
	}
	defer fileHandle.Close()
	for index, _ := range tmpFolderFiles {
		tmpFileName := gconv.String(index+1) + ".temp"
		fileBuffer, err := ioutil.ReadFile(tmpFolderPath + "/" + tmpFileName)
		if err != nil {
			return nil, response.FILE_SAVE_ERROR
		}

		_, err = fileHandle.Write(fileBuffer)
		if err != nil {
			return nil, response.FILE_SAVE_ERROR
		}
	}

	// 删除目录
	err = gfile.Remove(tmpFolderPath)
	if err != nil {
		return nil, response.FILE_SAVE_ERROR
	}

	BaseSetting, err := Config.FindValue(shared.BaseSetting)
	if err != nil {
		return nil, response.DB_READ_ERROR
	}
	BaseSettingJson := gjson.New(BaseSetting)
	url := gconv.String(BaseSettingJson.Get("url"))
	var mediaObject []upload_lib.MediaObject
	path = filePath[1:]
	mediaType := gfile.Ext(req.FileName)
	mediaType = mediaType[1:]
	media := upload_lib.MediaObject{
		UserId:     req.UserId,
		Link:       url + "/public" + path[7:],
		Path:       path,
		Name:       name + gfile.Ext(req.FileName),
		OrName:     req.FileName,
		Size:       gconv.String(req.Size),
		UploadKey:  1,
		Status:     1,
		MediaType:  gstr.ToUpper(mediaType),
		Ext:        gfile.Ext(req.FileName),
		CreateTime: gtime.Now(),
	}
	mediaObject = append(mediaObject, media)
	err = Media.Create(mediaObject)
	if err != nil {
		return nil, response.FILE_SAVE_ERROR
	}
	var mediaResult []string
	for _, i := range mediaObject {
		mediaResult = append(mediaResult, i.Link)
	}
	return mediaResult, response.SUCCESS
}
