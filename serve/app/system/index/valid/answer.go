package valid

import (
	"fiber/app/dao"
	"github.com/gogf/gf/util/gconv"
)

var Answer = new(answerValid)

type answerValid struct {
}

// CheckIsAdoption 是否已经采纳
func (s *answerValid) CheckIsAdoption(id int64) bool {
	adoption, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.IsAdoption, dao.SysAnswer.Columns.AnswerId, id)
	if err != nil || gconv.Int(adoption) == 2 {
		return true
	}
	return false
}
