// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysMangerDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysMangerDao struct {
	*internal.SysMangerDao
}

var (
	// SysManger is globally public accessible object for table sys_manger operations.
	SysManger = &sysMangerDao{
		internal.SysManger,
	}
)

// Fill with you ideas below.