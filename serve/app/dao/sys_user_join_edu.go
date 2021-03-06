// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysUserJoinEduDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysUserJoinEduDao struct {
	*internal.SysUserJoinEduDao
}

var (
	// SysUserJoinEdu is globally public accessible object for table sys_user_join_edu operations.
	SysUserJoinEdu = &sysUserJoinEduDao{
		internal.SysUserJoinEdu,
	}
)

// Fill with you ideas below.