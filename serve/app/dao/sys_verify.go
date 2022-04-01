// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysVerifyDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysVerifyDao struct {
	*internal.SysVerifyDao
}

var (
	// SysVerify is globally public accessible object for table sys_verify operations.
	SysVerify = &sysVerifyDao{
		internal.SysVerify,
	}
)

// Fill with you ideas below.