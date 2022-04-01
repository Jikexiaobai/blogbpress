package service

import (
	"fiber/app/system/admin/dto"
	"fiber/app/system/admin/shared"
	"fiber/app/tools/response"
	upload_lib "fiber/library/upload"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gogf/gf/os/gfile"

	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Upload = new(uploadService)

type uploadService struct {
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
	FileSetting, err := System.FindValue(shared.FileSetting)
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

	BaseSetting, err := System.FindValue(shared.BaseSetting)
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
