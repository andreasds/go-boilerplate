package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type Foo struct {
	ID           uuid.UUID   `db:"id"`
	ReceiptName  string      `db:"receipt_name"`
	ReceiptNotes null.String `db:"receipt_notes"`
	TotalPrice   float64     `db:"total_price"`
	IsActive     bool        `db:"is_active"`
	CreatedTime  time.Time   `db:"created_time"`
	CreatedBy    uuid.UUID   `db:"created_by"`
	UpdatedTime  null.Time   `db:"updated_time"`
	UpdatedBy    null.String `db:"updated_by"`
	RemovedTime  null.Time   `db:"removed_time"`
	RemovedBy    null.String `db:"removed_by"`
}
