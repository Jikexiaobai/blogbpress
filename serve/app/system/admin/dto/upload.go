package dto

import "github.com/gogf/gf/net/ghttp"

type UploadChunkFrom struct {
	Identifier  string `p:"identifier"  v:"required#请上传文件唯一标识"`
	ChunkNumber string `p:"chunkNumber"  v:"required#请设置切片索引"`
	TotalChunks string `p:"totalChunks"  v:"required#请设置切片总数"`
	File        *ghttp.UploadFile
}

type UploadChunkMergeFrom struct {
	UserId     int64
	Identifier string `p:"identifier"  v:"required#请上传文件唯一标识"`
	FileName   string `p:"fileName"  v:"required#请设置文件名"`
	Size       int64  `p:"size"  v:"required#请设置文件大小"`
}
