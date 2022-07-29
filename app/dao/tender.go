// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"toubiaogo/app/dao/internal"
)

// tenderDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type tenderDao struct {
	internal.TenderDao
}

var (
	// Tender is globally public accessible object for table tender operations.
	Tender = tenderDao{
		internal.Tender,
	}
)

// Fill with you ideas below.
