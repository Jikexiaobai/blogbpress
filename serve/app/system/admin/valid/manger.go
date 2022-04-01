package valid

import "fiber/app/dao"

var Manger = new(mangerValid)

type mangerValid struct {
}

// CheckEmailUniqueAll 检查邮箱是否唯一
func (s *mangerValid) CheckEmailUniqueAll(email string, id ...interface{}) bool {
	if id != nil {
		rs, err := dao.SysManger.
			Where(dao.SysManger.Columns.UserId+" !=", id).
			Where(dao.SysManger.Columns.Email, email).Count()
		if err != nil || rs > 0 {
			return true
		}
		return false
	}

	rs, err := dao.SysManger.
		Where(dao.SysManger.Columns.Email, email).Count()
	if err != nil || rs > 0 {
		return true
	}
	return false
}

// CheckPhoneUniqueAll 检查手机号是否唯一
func (s *mangerValid) CheckPhoneUniqueAll(phone string, id ...interface{}) bool {
	if id != nil {
		rs, err := dao.SysManger.
			Where(dao.SysManger.Columns.UserId+" !=", id).
			Where(dao.SysManger.Columns.Phone, phone).Count()
		if err != nil || rs > 0 {
			return true
		}
		return false
	}

	rs, err := dao.SysManger.
		Where(dao.SysManger.Columns.Phone, phone).Count()
	if err != nil || rs > 0 {
		return true
	}
	return false
}
