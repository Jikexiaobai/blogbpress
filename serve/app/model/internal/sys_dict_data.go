// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
    DictCode   int64       `orm:"dict_code,primary" json:"dictCode"`   // 字典编码                  
    CreateBy   int64       `orm:"create_by"         json:"createBy"`   //                           
    DictSort   int         `orm:"dict_sort"         json:"dictSort"`   // 字典排序                  
    DictLabel  string      `orm:"dict_label"        json:"dictLabel"`  // 字典标签                  
    DictValue  string      `orm:"dict_value"        json:"dictValue"`  // 字典键值                  
    DictType   string      `orm:"dict_type"         json:"dictType"`   // 字典类型                  
    CssClass   string      `orm:"css_class"         json:"cssClass"`   // 样式属性（其他样式扩展）  
    ListClass  string      `orm:"list_class"        json:"listClass"`  // 表格回显样式              
    IsDefault  int         `orm:"is_default"        json:"isDefault"`  // 是否默认 1是 2否          
    Status     int         `orm:"status"            json:"status"`     // 状态（0正常 1停用）       
    CreateTime *gtime.Time `orm:"create_time"       json:"createTime"` // 创建时间                  
    UpdateTime *gtime.Time `orm:"update_time"       json:"updateTime"` // 更新时间                  
    Remark     string      `orm:"remark"            json:"remark"`     // 备注                      
}