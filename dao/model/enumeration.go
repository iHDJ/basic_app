package model

import "time"

type Enumeration struct {
	ID       uint64 `gorm:"primary_key"`
	Name     string
	Remark   string
	IsSystem bool
	Opened   bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (enum Enumeration) Table() string {
	return "enumerations"
}

type EnumerationLabel struct {
	ID            uint64 `gorm:"primary_key"`
	Key           string
	Value         string
	EnumerationID uint64

	CreatedAt time.Time
	UpdatedAt time.Time
}
