package model

import "time"

type Member struct {
	ID uint64 `gorm:"primary_key"`

	UserID         uint64
	GroupID        uint64
	OrganizationID uint64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
