// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysComment is the golang structure for table sys_comment.
type SysComment struct {
    CommentId  int64       `orm:"comment_id,primary" json:"commentId"`  //                         
    ReplyId    int64       `orm:"reply_id"           json:"replyId"`    //                         
    UserId     int64       `orm:"user_id"            json:"userId"`     // 发布用户                
    ParentId   int64       `orm:"parent_id"          json:"parentId"`   //                         
    TopId      int64       `orm:"top_id"             json:"topId"`      //                         
    RelatedId  int64       `orm:"related_id"         json:"relatedId"`  //                         
    Module     string      `orm:"module"             json:"module"`     // 模块                    
    Content    string      `orm:"content"            json:"content"`    // 内容                    
    Type       int         `orm:"type"               json:"type"`       // 类型 1图片 2视频 3文字  
    Files      string      `orm:"files"              json:"files"`      // 文件链接                
    Likes      int64       `orm:"likes"              json:"likes"`      // 点赞数                  
    Unlikes    int64       `orm:"unlikes"            json:"unlikes"`    // 点踩                    
    Remark     string      `orm:"remark"             json:"remark"`     //                         
    Status     int         `orm:"status"             json:"status"`     // 2已审核，1未审核        
    CreateTime *gtime.Time `orm:"create_time"        json:"createTime"` //                         
    UpdateTime *gtime.Time `orm:"update_time"        json:"updateTime"` //                         
    DeleteTime *gtime.Time `orm:"delete_time"        json:"deleteTime"` //                         
}