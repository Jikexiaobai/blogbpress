package valid

import (
	"fiber/app/dao"
)

var Authority = new(authorityValid)

type authorityValid struct {
}

func (s *authorityValid) CheckAuthorityUnique(perms string) bool {
	rs, err := dao.SysAuthority.Where(dao.SysAuthority.Columns.Perms, perms).Count()
	if err != nil || rs > 0 {
		return true
	}

	return false
}

func (s *authorityValid) CheckAuthorityUniqueExceptYourself(id int64, perms string) bool {
	rs, err := dao.SysAuthority.Where(dao.SysAuthority.Columns.AuthorityId+" !=", id).Where(dao.SysAuthority.Columns.Perms, perms).Count()
	if err != nil || rs > 0 {
		return true
	}
	return false
}

//func (s *authorityValid) CheckUserAuthority(userId int64, url string) bool {
//
//	var redisCom redis.Com
//	redisCom.Key = shared.UserAuthority + gconv.String(userId)
//	authorityListObj, err := redisCom.GetString()
//	if err != nil {
//		return false
//	}
//
//	if authorityListObj != nil {
//		//authorityListStr :=
//		authorityList := gstr.Split(gconv.String(authorityListObj), ",")
//		for _, authority := range authorityList {
//			if authority == url {
//				return true
//			}
//		}
//	}
//
//	userAdmin, err := dao.SysUser.Value(dao.SysUser.Columns.Admin, dao.SysUser.Columns.UserId, userId)
//	if err != nil {
//		return false
//	}
//
//	authorityList, err := s.SelectUserRoleAuthorityList(gconv.Int64(userAdmin))
//	if err != nil || authorityList == nil {
//		return false
//	}
//
//	strAuthorityList := gstr.Join(authorityList, ",")
//	//存入缓存
//	redisCom.Time = "600"
//	redisCom.Data = strAuthorityList
//	err = redisCom.SetStringEX()
//	if err != nil {
//		return false
//	}
//
//	for _, authority := range authorityList {
//		if authority == url {
//			return true
//		}
//	}
//
//	return false
//}
