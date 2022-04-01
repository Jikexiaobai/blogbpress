package valid

import (
	"fiber/app/dao"
	"fiber/app/system/admin/shared"
)

var Topic = new(topicValid)

type topicValid struct {
}

// CheckModule 检查模块
func (s *topicValid) CheckModule(module string) bool {
	switch module {
	case shared.Article:
		return true
	case shared.Audio:
		return true
	case shared.Video:
		return true
	case shared.Resource:
		return true
	case shared.Edu:
		return true
	default:
		return false
	}
}

// CheckIsMyCreate 是否我发布的
func (s *topicValid) CheckIsMyCreate(userId, id int64) bool {
	count, err := dao.SysTopic.Where(dao.SysTopic.Columns.TopicId, id).
		Where(dao.SysTopic.Columns.UserId, userId).Count()
	if err != nil || count == 0 {
		return false
	}

	return true
}
