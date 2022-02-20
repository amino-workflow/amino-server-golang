package model

import "time"

type AuditColumns struct {
	CreatedAt time.Time `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
