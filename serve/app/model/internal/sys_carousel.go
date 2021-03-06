// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysCarousel is the golang structure for table sys_carousel.
type SysCarousel struct {
    CarouselId int64       `orm:"carousel_id,primary" json:"carouselId"` //                               
    Mode       int         `orm:"mode"                json:"mode"`       // 类型1(投稿内容)，2(其他内容)  
    RelatedId  int64       `orm:"related_id"          json:"relatedId"`  //                               
    Module     string      `orm:"module"              json:"module"`     //                               
    Link       string      `orm:"link"                json:"link"`       // 链接                          
    Cover      string      `orm:"cover"               json:"cover"`      // 封面地址                      
    Type       int         `orm:"type"                json:"type"`       // 类型: 1(web),2(app)           
    CreateTime *gtime.Time `orm:"create_time"         json:"createTime"` //                               
    UpdateTime *gtime.Time `orm:"update_time"         json:"updateTime"` //                               
    DeleteTime *gtime.Time `orm:"delete_time"         json:"deleteTime"` //                               
}