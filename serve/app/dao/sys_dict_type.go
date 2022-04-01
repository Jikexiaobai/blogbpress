// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysDictTypeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysDictTypeDao struct {
	*internal.SysDictTypeDao
}

var (
	// SysDictType is globally public accessible object for table sys_dict_type operations.
	SysDictType = &sysDictTypeDao{
		internal.SysDictType,
	}
)

// Fill with you ideas below.