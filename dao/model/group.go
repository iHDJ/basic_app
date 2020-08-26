package model

import "time"

const (
	DefaultManagerGroupID = 1
	DefaultUserGroupID    = 2
)

var (
	ManagerGroup Group
	UserGroup    Group
)

type Group struct {
	ID      uint64 `gorm:"primary_key"`
	Name    string
	Route   string //指的是从root到该组需要经过的父级id路径, 以#符号分隔
	GroupID uint64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
