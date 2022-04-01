package result

import "github.com/gogf/gf/os/gtime"

type ReportList struct {
	ReportId    int64       `json:"id"`                                 //
	NickName    string      `json:"nickName"`                           //
	RelatedId   int64       `json:"relatedId"`                          //
	Module      string      `json:"module"`                             //
	Title       string      `json:"title"`                              //
	Type        int         `json:"type"`                               // 举报类型 1广告垃圾，2违规内容，3恶意灌水，4重复发帖
	Description string      `json:"description"`                        // 描述
	Status      int         `json:"status"`                             // 状态，1 已处理，2 未处理
	CreateTime  *gtime.Time `orm:"create_time"       json:"createTime"` //
}
