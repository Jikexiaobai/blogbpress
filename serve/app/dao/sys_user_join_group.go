// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysUserJoinGroupDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysUserJoinGroupDao struct {
	*internal.SysUserJoinGroupDao
}

var (
	// SysUserJoinGroup is globally public accessible object for table sys_user_join_group operations.
	SysUserJoinGroup = &sysUserJoinGroupDao{
		internal.SysUserJoinGroup,
	}
)

// Fill with you ideas below.