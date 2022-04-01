// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysOrderDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysOrderDao struct {
	*internal.SysOrderDao
}

var (
	// SysOrder is globally public accessible object for table sys_order operations.
	SysOrder = &sysOrderDao{
		internal.SysOrder,
	}
)

// Fill with you ideas below.