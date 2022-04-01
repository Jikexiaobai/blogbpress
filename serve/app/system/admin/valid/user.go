package valid

import "fiber/app/dao"

var User = new(userValid)

type userValid struct {
}

// CheckEmailUniqueAll 检查邮箱是否唯一
func (s *userValid) CheckEmailUniqueAll(email string, id ...interface{}) bool {
	if id != nil {
		rs, err := dao.SysUser.
			Where(dao.SysUser.Columns.UserId+" !=", id).
			Where(dao.SysUser.Columns.Email, email).Count()
		if err != nil || rs > 0 {
			return true
		}
		return false
	}

	rs, err := dao.SysUser.
		Where(dao.SysUser.Columns.Email, email).Count()
	if err != nil || rs > 0 {
		return true
	}
	return false
}

// CheckPhoneUniqueAll 检查手机号是否唯一
func (s *userValid) CheckPhoneUniqueAll(phone string, id ...interface{}) bool {
	if id != nil {
		rs, err := dao.SysUser.
			Where(dao.SysUser.Columns.UserId+" !=", id).
			Where(dao.SysUser.Columns.Phone, phone).Count()
		if err != nil || rs > 0 {
			return true
		}
		return false
	}

	rs, err := dao.SysUser.
		Where(dao.SysUser.Columns.Phone, phone).Count()
	if err != nil || rs > 0 {
		return true
	}
	return false
}
