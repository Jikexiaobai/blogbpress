// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysGroup is the golang structure for table sys_group.
type SysGroup struct {
    GroupId     int64       `orm:"group_id,primary" json:"groupId"`     //                                           
    UserId      int64       `orm:"user_id"          json:"userId"`      // 创建人                                    
    CateId      int64       `orm:"cate_id"          json:"cateId"`      // 小组分类                                  
    Title       string      `orm:"title"            json:"title"`       // 小组名称                                  
    Cover       string      `orm:"cover"            json:"cover"`       // 封面                                      
    JoinMode    int         `orm:"join_mode"        json:"joinMode"`    // 小组类型1 公共小组，2付费小组，3专属小组  
    Price       float64     `orm:"price"            json:"price"`       // 费用                                      
    SecretKey   string      `orm:"secret_key"       json:"secretKey"`   // 加入角色                                  
    Icon        string      `orm:"icon"             json:"icon"`        // 小组图标                                  
    Joins       int64       `orm:"joins"            json:"joins"`       //                                           
    Hots        int64       `orm:"hots"             json:"hots"`        //                                           
    Contents    int64       `orm:"contents"         json:"contents"`    //                                           
    Views       int64       `orm:"views"            json:"views"`       //                                           
    Description string      `orm:"description"      json:"description"` // 小组描述                                  
    Remark      string      `orm:"remark"           json:"remark"`      // 备注                                      
    Status      int         `orm:"status"           json:"status"`      // 状态 0 全部  1待审   2通过  3拒绝         
    CreateTime  *gtime.Time `orm:"create_time"      json:"createTime"`  //                                           
    UpdateTime  *gtime.Time `orm:"update_time"      json:"updateTime"`  // 跟新时间                                  
    DeleteTime  *gtime.Time `orm:"delete_time"      json:"deleteTime"`  //                                           
}