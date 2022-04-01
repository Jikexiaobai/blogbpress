package dto

import "github.com/gogf/gf/net/ghttp"

// 用户报名查询参数
type MediaQuery struct {
	OrName string `p:"orName"`
	Status int    `p:"Status"`
	Page   int    `p:"page" v:"required#请设置页数"`
	Limit  int    `p:"limit" v:"between:1,100#参数只允许1到100"`
}
type BigUploadDto struct {
	Hash     string            // 标题
	FileName string            //
	File     *ghttp.UploadFile //
}

type FileInfoDto struct {
	FileName string `p:"fileName"  v:"required#请设置文件名称"`
	Hash     string `p:"hash"  v:"required#请设置hash值"`
}
