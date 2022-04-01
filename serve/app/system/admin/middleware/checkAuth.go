package middleware

import (
	"fiber/app/dao"
	"fiber/app/system/admin/service"
	"fiber/app/tools/response"
	"github.com/gogf/gf/net/ghttp"
)

func CheckAuth(r *ghttp.Request) {
	url := r.URL.Path[13:]

	//获取用户token
	tokenUserId, err := service.Auth.GetTokenId(r)
	if err != nil {
		response.Error(r).
			SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage(response.CodeMsg(response.ACCESS_TOKEN_TIMEOUT)).Send()
	}

	// 获取权限id
	authorityId, err := dao.SysAuthority.Value(dao.SysAuthority.Columns.AuthorityId, dao.SysAuthority.Columns.Perms, url)
	if err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	}
	//获取用户角色Id
	list, err := dao.SysMangerRole.Fields(dao.SysMangerRole.Columns.RoleId).Where(dao.SysMangerRole.Columns.UserId, tokenUserId).All()
	if err != nil {
		response.Error(r).
			SetCode(response.DB_READ_ERROR).
			SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	}
	var roleIds []int64
	for _, i := range list {
		roleIds = append(roleIds, i.RoleId)
	}

	res, err := dao.SysRoleAuthority.
		Where(dao.SysRoleAuthority.Columns.RoleId+" IN(?)", roleIds).
		Where(dao.SysRoleAuthority.Columns.AuthorityId, authorityId).
		All()
	if err != nil || len(res) == 0 {
		response.Error(r).
			SetCode(response.AUTH_ERROR).
			SetMessage(response.CodeMsg(response.AUTH_ERROR)).Send()
	}
	r.Middleware.Next()
}
