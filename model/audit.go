package model

import "time"

type AuditColumns struct {
	CreatedAt time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
