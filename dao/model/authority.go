package model

import (
	"time"
)

//权限
type Authority struct {
	ID      uint64 `gorm:"primary_key"`
	Name    string
	Disable bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

//权限条目 例如 客户权限 下的可创建权限
type AuthorityItem struct {
	ID uint64 `gorm:"primary_key"`

	Name        string
	Remark      string //备注
	AuthorityID uint64
	Disable     bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type MemberAuthority struct {
	ID              uint64 `gorm:"primary_key"`
	AuthorityItemID uint64
	MemberID        uint64

	CreatedAt time.Time
}
