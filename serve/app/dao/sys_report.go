// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"fiber/app/dao/internal"
)

// sysReportDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysReportDao struct {
	*internal.SysReportDao
}

var (
	// SysReport is globally public accessible object for table sys_report operations.
	SysReport = &sysReportDao{
		internal.SysReport,
	}
)

// Fill with you ideas below.