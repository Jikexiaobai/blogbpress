// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysManger is the golang structure for table sys_manger.
type SysManger struct {
    UserId     int64       `orm:"user_id,primary" json:"userId"`     // 用户ID                     
    NickName   string      `orm:"nick_name"       json:"nickName"`   // 用户昵称                   
    Email      string      `orm:"email"           json:"email"`      // 用户邮箱                   
    Phone      string      `orm:"phone"           json:"phone"`      // 手机号码                   
    Sex        int         `orm:"sex"             json:"sex"`        // 用户性别（1男 2女 3未知）  
    Avatar     string      `orm:"avatar"          json:"avatar"`     // 头像地址                   
    Password   string      `orm:"password"        json:"password"`   // 密码                       
    Salt       string      `orm:"salt"            json:"salt"`       // 密码盐                     
    Status     int         `orm:"status"          json:"status"`     // 帐号状态（1停用,2正常）    
    LoginIp    string      `orm:"login_ip"        json:"loginIp"`    // 最后登陆IP                 
    LoginTime  *gtime.Time `orm:"login_time"      json:"loginTime"`  // 最后登陆时间               
    CreateTime *gtime.Time `orm:"create_time"     json:"createTime"` // 创建时间                   
    UpdateTime *gtime.Time `orm:"update_time"     json:"updateTime"` // 更新时间                   
    Remark     string      `orm:"remark"          json:"remark"`     // 备注                       
    DeleteTime *gtime.Time `orm:"delete_time"     json:"deleteTime"` // 软删除                     
}