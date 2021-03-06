// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysMedia is the golang structure for table sys_media.
type SysMedia struct {
    MediaId    int64       `orm:"media_id,primary" json:"mediaId"`    //                                     
    UserId     uint64      `orm:"user_id"          json:"userId"`     // 上传的用户                          
    Link       string      `orm:"link"             json:"link"`       // 文件链接                            
    Path       string      `orm:"path"             json:"path"`       // 存放路径                            
    Name       string      `orm:"name"             json:"name"`       // 文件名称                            
    OrName     string      `orm:"or_name"          json:"orName"`     // 原始文件名称                        
    Size       string      `orm:"size"             json:"size"`       // 文件大小                            
    UploadKey  int         `orm:"upload_key"       json:"uploadKey"`  // 上传方式 1 为本地上传， 2为oss上传  
    Ext        string      `orm:"ext"              json:"ext"`        // 文件后缀                            
    Status     int         `orm:"status"           json:"status"`     //                                     
    CreateTime *gtime.Time `orm:"create_time"      json:"createTime"` // 创建时间                            
    DeleteTime *gtime.Time `orm:"delete_time"      json:"deleteTime"` //                                     
    MediaType  string      `orm:"media_type"       json:"mediaType"`  // 文件类型                            
    Remark     string      `orm:"remark"           json:"remark"`     // 备注                                
}