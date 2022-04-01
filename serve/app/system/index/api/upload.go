package api

import (
	"fiber/app/system/index/dto"
	"fiber/app/system/index/service"
	"fiber/app/system/index/valid"
	"fiber/app/tools/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gvalid"
)

var Upload = new(uploadApi)

type uploadApi struct{}

func (c *uploadApi) LoadRouter(group *ghttp.RouterGroup) {
	//group.POST("/upload", c.upload)
	group.Group("/upload", func(group *ghttp.RouterGroup) {
		group.GET("/chunk", c.getUploadChunk)
		group.POST("/chunk", c.postUploadChunk)
		group.POST("/mergeChunk", c.postUploadChunkMerge)
	})

}

func (c *uploadApi) upload(r *ghttp.Request) {
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}

	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, "upload") {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}

	files := r.GetUploadFiles("file")
	if files == nil {
		response.Error(r).SetCode(response.INVALID).Send()
	}

	// 检查文件大小
	for _, file := range files {
		if valid.Media.CheckMediaSize(file.Size) {
			response.Error(r).SetCode(response.INVALID).SetMessage("文件太大了").Send()
		}
		fileSuffix := gfile.Ext(file.FileHeader.Filename) //获取文件名
		//检查文件类型
		if !valid.Media.CheckMediaType(fileSuffix) {
			response.Error(r).SetCode(response.INVALID).SetMessage("文件格式不正确").Send()
		}
	}

	if result, code := service.Upload.UploadFile(tokenUserId, files); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("link", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *uploadApi) getUploadChunk(r *ghttp.Request) {
	//	接受参数
	rule := "required"
	msg := map[string]string{
		"required": "请设置文件夹唯一标识",
	}
	identifier := r.GetQueryString("identifier")
	if err := gvalid.Check(identifier, rule, msg); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}
	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, "upload") {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}
	if result, code := service.Upload.GetUploadChunkFiles(identifier); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("result", result)
		response.Success(r).SetData(data).Send()
	}
}

func (c *uploadApi) postUploadChunk(r *ghttp.Request) {
	//	接受参数
	var req *dto.UploadChunkFrom
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}
	file := r.GetUploadFile("file")
	if file == nil {
		response.Error(r).SetCode(response.INVALID).Send()
	}
	req.File = file

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}
	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, "upload") {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}
	//	检查是否存在文件
	if code := service.Upload.UploadChunk(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		if req.ChunkNumber == req.TotalChunks {
			data.Set("result", "needMerge")
		} else {
			data.Set("result", "分片上传完成")
		}
		response.Success(r).SetData(data).Send()
	}
}

func (c *uploadApi) postUploadChunkMerge(r *ghttp.Request) {
	//	接受参数
	var req *dto.UploadChunkMergeFrom
	if err := r.Parse(&req); err != nil {
		response.Error(r).
			SetCode(response.PARAM_INVALID).
			SetMessage(err.Error()).
			Send()
	}

	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).SetCode(response.INVALID).SetMessage(err.Error()).Send()
	}
	req.UserId = tokenUserId
	//检查是否有权限操作
	if !service.Grade.CheckHasCommon(tokenUserId, "upload") {
		response.Error(r).SetCode(response.INVALID).SetMessage("当前等级无权操作").Send()
	}
	if result, code := service.Upload.MergeChunk(req); code != response.SUCCESS {
		response.Error(r).SetCode(code).SetMessage(response.CodeMsg(code)).Send()
	} else {
		data := gmap.New(true)
		data.Set("link", result)
		response.Success(r).SetData(data).Send()
	}
}
