// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysAnswer is the golang structure for table sys_answer.
type SysAnswer struct {
    AnswerId   int64       `orm:"answer_id,primary" json:"answerId"`   //                                              
    UserId     int64       `orm:"user_id"           json:"userId"`     //                                              
    TopicId    int64       `orm:"topic_id"          json:"topicId"`    //                                              
    Doc        string      `orm:"doc"               json:"doc"`        // 文档下载                                     
    Content    string      `orm:"content"           json:"content"`    //                                              
    IsAdoption int         `orm:"is_adoption"       json:"isAdoption"` // 是否被采纳，1未采纳，2已采纳                 
    Price      float64     `orm:"price"             json:"price"`      //                                              
    Likes      int64       `orm:"likes"             json:"likes"`      //                                              
    Views      int64       `orm:"views"             json:"views"`      //                                              
    Remark     string      `orm:"remark"            json:"remark"`     //                                              
    Status     int         `orm:"status"            json:"status"`     // 状态：0全部, 1待审核，2已发布，3拒绝，4草稿  
    CreateTime *gtime.Time `orm:"create_time"       json:"createTime"` //                                              
    UpdateTime *gtime.Time `orm:"update_time"       json:"updateTime"` //                                              
    DeleteTime *gtime.Time `orm:"delete_time"       json:"deleteTime"` //                                              
}