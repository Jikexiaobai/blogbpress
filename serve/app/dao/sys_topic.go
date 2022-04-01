// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysTopicDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysTopicDao struct {
	*internal.SysTopicDao
}

var (
	// SysTopic is globally public accessible object for table sys_topic operations.
	SysTopic = &sysTopicDao{
		internal.SysTopic,
	}
)

// Fill with you ideas below.