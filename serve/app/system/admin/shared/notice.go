package shared

const (
	NoticeSystem  = 1 //系统通知
	NoticeComment = 2 //评论通知
	NoticeAnswer  = 3 //回答通知
	NoticeLike    = 4 //获赞通知
	NoticeFollow  = 5 //收到关注通知
)

const (
	NoticeSysTemRegister      = 1 // 用户注册通知
	NoticeSysTemDeleteContent = 2 // 内容删除通知
	NoticeUserTips            = 3 // 用户打赏通知
	NoticeUserBuyContent      = 4 // 用户购买内容通知
	NoticeUserJoin            = 5 // 用户加入通知
	NoticeSysTemReview        = 6 // 内容审核通知
	NoticeUserCash            = 7 // 用户提现通知
	NoticeUserRecharge        = 8 // 用户充值通知
	NoticeSysTemUserReview    = 9 // 用户审核通知
)

// 状态
const (
	NoticeStatusReview   = 1 // 未读
	NoticeStatusReviewed = 2 // 已读
)
