// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysConfigDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysConfigDao struct {
	*internal.SysConfigDao
}

var (
	// SysConfig is globally public accessible object for table sys_config operations.
	SysConfig = &sysConfigDao{
		internal.SysConfig,
	}
)

// Fill with you ideas below.