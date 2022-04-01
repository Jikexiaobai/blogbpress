package dto

import "github.com/gogf/gf/net/ghttp"

// QueryParam 参数
//type MediaQueryParam struct {
//	Ext   []string `p:"ext" v:"required#请设置需要查询的文件后缀"`
//	Page  int      `p:"page"  v:"required#请设置页数"`
//	Limit int      `p:"limit"  v:"max:100#最大获取值为100"`
//}

type BigUploadDto struct {
	Hash     string            // 标题
	FileName string            //
	File     *ghttp.UploadFile //
}

type FileInfoDto struct {
	FileName string `p:"fileName"  v:"required#请设置文件名称"`
	Hash     string `p:"hash"  v:"required#请设置hash值"`
}
